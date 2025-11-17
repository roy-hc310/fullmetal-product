package product_service

import (
	"context"
)

func (s *ProductService) UpdateProduct(ctx context.Context, data map[string]interface{}) (res string, traceID string, err error) {

	res, err = s.ProductRepository.UpdateProduct(ctx, data)
	if err != nil {
		return res, traceID, err
	}

	return res, traceID, nil
}
