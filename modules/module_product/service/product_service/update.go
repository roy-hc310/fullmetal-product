package product_service

import (
	"context"
)

func (s *ProductService) UpdateProduct(ctx context.Context, id string, data *interface{}) (res string, traceID string, err error) {
	return res, traceID, nil
}
