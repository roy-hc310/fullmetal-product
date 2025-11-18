package repository

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
	// product_entity "github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
	// "github.com/roy-hc310/hannie-product/package/database"
)

type ProductRepositoryInteface interface {
	CreateProduct(ctx context.Context, data *entity.Product) (res string, err error)
	DeleteProduct(ctx context.Context, id string) (res string, err error)
	GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, err error)
	GetListProduct(ctx context.Context, data *dto_v1.GetListProductRequest) (res []*dto_v1.GetListProductResponse, pagination *dto_v1.PaginationResponse, err error)
	UpdateProduct(ctx context.Context, id string, data map[string]interface{}) (res string, err error)
}
