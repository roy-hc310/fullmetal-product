package product_service

import (
	"context"
)

func (s *ProductService) DeleteProduct(ctx context.Context, id string) (res string, traceID string, err error) {

	res, err = s.Repository.DeleteProduct(ctx, id)
	return res, traceID, err
}
