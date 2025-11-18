package product_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
)

func (r *ProductRepository) UpdateProduct(ctx context.Context, id string, data map[string]interface{}) (res string, err error) {

	productID, err := uuid.Parse(id)
	if err != nil {
		return res, err
	}

	data["updated_at"] = time.Now()

	err = r.PostgresInfra.DB.WithContext(ctx).Model(&entity.Product{}).Where("id = ?", productID).Updates(data).Error
	if err != nil {
		return res, err
	}

	return id, nil
}
