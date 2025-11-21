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

// TestDeleteProduct_Success tests successful product deletion
func TestDeleteProduct_Success(t *testing.T) {
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

	// Mock repository to return success
	mockRepo.On("DeleteProduct", mock.Anything, productID).Return(productID, nil)

	// Mock background operations
	mockCache.On("Delete", mock.Anything, mock.Anything).Return(nil)
	mockEvent.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Act
	res, traceID, err := service.DeleteProduct(ctx, productID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, productID, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}

// TestDeleteProduct_RepositoryError tests deletion with repository error
func TestDeleteProduct_RepositoryError(t *testing.T) {
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
	expectedError := errors.New("product not found")

	mockRepo.On("DeleteProduct", mock.Anything, productID).Return("", expectedError)

	// Act
	res, traceID, err := service.DeleteProduct(ctx, productID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, res)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}
