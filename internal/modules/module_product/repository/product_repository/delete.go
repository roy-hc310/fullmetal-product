package product_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
	"gorm.io/gorm"
)

func (r *ProductRepository) DeleteProduct(ctx context.Context, id string) (res string, err error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		return res, err
	}

	result := r.PostgresInfra.DB.WithContext(ctx).Delete(&entity.Product{}, "id = ?", parseID)
	if result.Error != nil {
		return res, err
	}

	if result.RowsAffected == 0 {
		return res, gorm.ErrRecordNotFound
	}
	return id, nil
}
