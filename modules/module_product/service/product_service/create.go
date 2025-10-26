package product_service

import (
	"context"

	"github.com/google/uuid"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
)

func (s *ProductService) CreateProduct(ctx context.Context, data *dto_v1.CreateProductRequest) (res string, traceID string, err error) {
	data.ProductID = uuid.NewString()
	data.BrandID = uuid.NewString()
	data.CategoryID = uuid.NewString()
	data.ShopID = uuid.NewString()
	res, err = s.ProductRepository.CreateProduct(ctx, data)
	return res, traceID, err
}
