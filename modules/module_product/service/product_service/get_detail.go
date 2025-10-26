package product_service

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
)

func (s *ProductService) GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, traceID string, err error) {
	res, err = s.ProductRepository.GetDetailProduct(ctx, id)
	if err != nil {
		return nil, traceID, err
	}
	return res, traceID, err
}
