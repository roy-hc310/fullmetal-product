package product_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *ProductRepository) DeleteProduct(ctx context.Context, id string) (res string, err error) {
	var parseID pgtype.UUID
	parseID.Scan(id)
	err = r.Write.DeleteProduct(ctx, parseID)
	if err != nil {
		return "", err
	}
	return id, nil
}
