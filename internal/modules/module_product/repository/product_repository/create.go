package product_repository

import (
	"context"

	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
)

func (r *ProductRepository) CreateProduct(ctx context.Context, data *entity.Product) (res string, err error) {

	err = r.PostgresInfra.DB.WithContext(ctx).Create(data).Error
	if err != nil {
		return res, err
	}
	return data.ID.String(), nil
}

// If you have relationships (foreign keys, etc.), you can use:

// db.Create(&product).Association("Brand").Append(&brand)

// — but for now, plain inserts are fine.
