package product_service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/service/product_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestGetDetailProduct_CacheHit tests getting product from cache
func TestGetDetailProduct_CacheHit(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockProductRepository)
	mockCache := new(mocks.MockCache)
	mockEvent := new(mocks.MockEventProducer)
	mockOtel := new(mocks.MockOtel)

	service := &ProductService{
		Repository: mockRepo,
		Cache:      mockCache,
		Event:      mockEvent,
		Otel:       mockOtel,
	}

	ctx := context.Background()
	productID := uuid.New().String()
	cachedData := `{"id":"` + productID + `","name":"Test Product","sku":"SKU123"}`

	// Mock cache to return hit
	mockCache.On("Get", mock.Anything, "product:"+productID).Return(cachedData, nil)

	// Act
	res, traceID, err := service.GetDetailProduct(ctx, productID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, traceID)
	mockCache.AssertExpectations(t)
	// Repository should NOT be called on cache hit
	mockRepo.AssertNotCalled(t, "GetDetailProduct")
}

// TestGetDetailProduct_CacheMiss tests getting product from database
func TestGetDetailProduct_CacheMiss(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockProductRepository)
	mockCache := new(mocks.MockCache)
	mockEvent := new(mocks.MockEventProducer)
	mockOtel := new(mocks.MockOtel)

	service := &ProductService{
		Repository: mockRepo,
		Cache:      mockCache,
		Event:      mockEvent,
		Otel:       mockOtel,
	}

	ctx := context.Background()
	productID := uuid.New().String()

	expectedProduct := &dto_v1.GetDetailProductResponse{
		Name: "Test Product",
		Sku:  "SKU123",
	}

	// Mock cache miss
	mockCache.On("Get", mock.Anything, "product:"+productID).Return("", errors.New("cache miss"))
	// Mock repository success
	mockRepo.On("GetDetailProduct", mock.Anything, productID).Return(expectedProduct, nil)
	// Mock cache set
	mockCache.On("Set", mock.Anything, "product:"+productID, mock.Anything, mock.Anything).Return(nil)

	// Act
	res, traceID, err := service.GetDetailProduct(ctx, productID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expectedProduct.Name, res.Name)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

// TestGetDetailProduct_RepositoryError tests database error
func TestGetDetailProduct_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockProductRepository)
	mockCache := new(mocks.MockCache)
	mockEvent := new(mocks.MockEventProducer)
	mockOtel := new(mocks.MockOtel)

	service := &ProductService{
		Repository: mockRepo,
		Cache:      mockCache,
		Event:      mockEvent,
		Otel:       mockOtel,
	}

	ctx := context.Background()
	productID := uuid.New().String()
	expectedError := errors.New("database connection error")

	// Mock cache miss
	mockCache.On("Get", mock.Anything, "product:"+productID).Return("", errors.New("cache miss"))
	// Mock repository error
	mockRepo.On("GetDetailProduct", mock.Anything, productID).Return(nil, expectedError)

	// Act
	res, traceID, err := service.GetDetailProduct(ctx, productID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}
