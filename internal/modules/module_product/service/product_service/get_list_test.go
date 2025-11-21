package product_service

import (
	"context"
	"errors"
	"testing"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/service/product_service/mocks"
	"github.com/roy-hc310/fullmetal-product/pkg/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestGetListProduct_Success tests successful product list retrieval
func TestGetListProduct_Success(t *testing.T) {
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
	request := &dto_v1.GetListProductRequest{
		Query: base.Query{
			Limit:  10,
			Cursor: "",
			SortBy: "created_at",
			Order:  "desc",
		},
	}

	expectedProducts := []*dto_v1.GetListProductResponse{
		{Name: "Product 1", Sku: "SKU1"},
		{Name: "Product 2", Sku: "SKU2"},
	}

	expectedPagination := &dto_v1.PaginationResponse{
		Limit:      10,
		NextCursor: "cursor123",
		PrevCursor: "cursor456",
		HasNext:    true,
	}

	mockRepo.On("GetListProduct", mock.Anything, request).Return(expectedProducts, expectedPagination, nil)

	// Act
	res, pagination, traceID, err := service.GetListProduct(ctx, request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res, 2)
	assert.Equal(t, expectedProducts[0].Name, res[0].Name)
	assert.NotNil(t, pagination)
	assert.True(t, pagination.HasNext)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}

// TestGetListProduct_RepositoryError tests list retrieval with repository error
func TestGetListProduct_RepositoryError(t *testing.T) {
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
	request := &dto_v1.GetListProductRequest{
		Query: base.Query{
			Limit: 10,
		},
	}

	expectedError := errors.New("database query failed")
	mockRepo.On("GetListProduct", mock.Anything, request).Return(nil, nil, expectedError)

	// Act
	res, pagination, traceID, err := service.GetListProduct(ctx, request)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, res)
	assert.Nil(t, pagination)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}

// TestGetListProduct_EmptyResult tests empty list return
func TestGetListProduct_EmptyResult(t *testing.T) {
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
	request := &dto_v1.GetListProductRequest{
		Query: base.Query{
			Limit: 10,
		},
	}

	emptyProducts := []*dto_v1.GetListProductResponse{}
	emptyPagination := &dto_v1.PaginationResponse{
		Limit:   10,
		HasNext: false,
	}

	mockRepo.On("GetListProduct", mock.Anything, request).Return(emptyProducts, emptyPagination, nil)

	// Act
	res, pagination, traceID, err := service.GetListProduct(ctx, request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res, 0)
	assert.False(t, pagination.HasNext)
	assert.NotEmpty(t, traceID)
	mockRepo.AssertExpectations(t)
}
