package product_repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *ProductRepository) DeleteProduct(ctx context.Context, id string) (res string, err error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	err = r.Write.DeleteProduct(ctx, parseID)
	if err != nil {
		return "", err
	}
	return id, nil
}
