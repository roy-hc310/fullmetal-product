package product_service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/service/product_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestUpdateProduct_Success tests successful product update
func TestUpdateProduct_Success(t *testing.T) {
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
	updateData := map[string]interface{}{
		"name":        "Updated Product",
		"description": "Updated Description",
	}

	mockRepo.On("UpdateProduct", mock.Anything, productID, mock.MatchedBy(func(data map[string]interface{}) bool {
		return data["name"] == "Updated Product"
	})).Return(productID, nil)

	// Mock background operations
	mockCache.On("Delete", mock.Anything, mock.Anything).Return(nil)
	mockEvent.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Act
	res, traceID, err := service.UpdateProduct(ctx, productID, updateData)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, productID, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}

// TestUpdateProduct_RepositoryError tests update with repository error
func TestUpdateProduct_RepositoryError(t *testing.T) {
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
	updateData := map[string]interface{}{
		"name": "Updated Product",
	}

	expectedError := errors.New("update failed")
	mockRepo.On("UpdateProduct", mock.Anything, productID, mock.Anything).Return("", expectedError)

	// Act
	res, traceID, err := service.UpdateProduct(ctx, productID, updateData)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}

// TestUpdateProduct_EmptyData tests update with empty data
func TestUpdateProduct_EmptyData(t *testing.T) {
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
	emptyData := map[string]interface{}{}

	// Even with empty data, repository should be called
	mockRepo.On("UpdateProduct", mock.Anything, productID, mock.Anything).Return(productID, nil)

	// Mock background operations
	mockCache.On("Delete", mock.Anything, mock.Anything).Return(nil)
	mockEvent.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Act
	res, traceID, err := service.UpdateProduct(ctx, productID, emptyData)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, productID, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}
