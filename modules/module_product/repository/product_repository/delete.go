package product_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
)

func (r *ProductRepository) DeleteProduct(ctx context.Context, id string) (res string, err error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		return res, err
	}

	err = r.PostgresInfra.DBWrite.WithContext(ctx).Delete(&entity.Product{}, "id = ?", parseID).Error
	if err != nil {
		return res, err
	}
	return id, nil
}
