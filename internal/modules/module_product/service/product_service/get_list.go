package product_service

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
)

func (s *ProductService) GetListProduct(ctx context.Context, data *dto_v1.GetListProductRequest) (res []*dto_v1.GetListProductResponse, pagination *dto_v1.PaginationResponse, traceID string, err error) {
	res, pagination, err = s.ProductRepository.GetListProduct(ctx, data)
	if err != nil {
		return nil, nil, traceID, err
	}

	return res, pagination, traceID, nil
}
