package mocks

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock implementation of ProductRepositoryInterface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(ctx context.Context, data *entity.Product) (string, error) {
	args := m.Called(ctx, data)
	return args.String(0), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(ctx context.Context, id string) (string, error) {
	args := m.Called(ctx, id)
	return args.String(0), args.Error(1)
}

func (m *MockProductRepository) GetDetailProduct(ctx context.Context, id string) (*dto_v1.GetDetailProductResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto_v1.GetDetailProductResponse), args.Error(1)
}

func (m *MockProductRepository) GetListProduct(ctx context.Context, params *dto_v1.GetListProductRequest) ([]*dto_v1.GetListProductResponse, *dto_v1.PaginationResponse, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).([]*dto_v1.GetListProductResponse), args.Get(1).(*dto_v1.PaginationResponse), args.Error(2)
}

func (m *MockProductRepository) UpdateProduct(ctx context.Context, id string, data map[string]interface{}) (string, error) {
	args := m.Called(ctx, id, data)
	return args.String(0), args.Error(1)
}
