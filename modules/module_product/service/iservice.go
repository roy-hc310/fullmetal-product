package service

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
)

type ProductServiceInterface interface {
	CreateProduct(ctx context.Context, data *dto_v1.CreateProductRequest) (res string, traceID string, err error)
	DeleteProduct(ctx context.Context, id string) (res string, traceID string, err error)
	GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, traceID string, err error)
	GetListProduct(ctx context.Context, data *dto_v1.GetListProductRequest) (res *dto_v1.GetListProductResponse, traceID string, err error)
	UpdateProduct(ctx context.Context, id string, data *interface{}) (res string, traceID string, err error)
}
