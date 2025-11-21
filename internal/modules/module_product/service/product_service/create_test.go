package product_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/service/product_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateProduct_Success tests successful product creation
func TestCreateProduct_Success(t *testing.T) {
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
	request := &dto_v1.CreateProductRequest{
		Name:        "Test Product",
		Sku:         "SKU123",
		Description: "Test Description",
		IsActive:    true,
		Status:      1,
	}

	// Mock repository to return success
	mockRepo.On("CreateProduct", mock.Anything, mock.MatchedBy(func(p *entity.Product) bool {
		return p.Name == request.Name && *p.SKU == request.Sku
	})).Return(productID, nil)

	// Mock background operations (cache and event)
	mockCache.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockEvent.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Act
	res, traceID, err := service.CreateProduct(ctx, request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, productID, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)

	// Give background goroutine time to execute
	time.Sleep(100 * time.Millisecond)
}

// TestCreateProduct_ValidationError tests product creation with validation failure
func TestCreateProduct_ValidationError(t *testing.T) {
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
	request := &dto_v1.CreateProductRequest{
		Name: "", // Invalid: empty name
	}

	// Act
	res, traceID, err := service.CreateProduct(ctx, request)

	// Assert
	assert.Error(t, err)
	assert.Empty(t, res)
	assert.NotEmpty(t, traceID) // Trace ID should still be returned
}

// TestCreateProduct_RepositoryError tests product creation with repository error
func TestCreateProduct_RepositoryError(t *testing.T) {
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
	request := &dto_v1.CreateProductRequest{
		Name:        "Test Product",
		Sku:         "SKU123",
		Description: "Test Description",
		IsActive:    true,
		Status:      1,
	}

	expectedError := errors.New("database connection error")
	mockRepo.On("CreateProduct", mock.Anything, mock.Anything).Return("", expectedError)

	// Act
	res, traceID, err := service.CreateProduct(ctx, request)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}
