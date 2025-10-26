package product_service

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
)

func (s *ProductService) GetListProduct(ctx context.Context, data *dto_v1.GetListProductRequest) (res *dto_v1.GetListProductResponse, traceID string, err error) {
	return res, traceID, nil
}
