package product_repository

import (
	"context"

	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
)

func (r *ProductRepository) CreateProduct(ctx context.Context, data *entity.Product) (res string, err error) {

	// product := entity.Product{
	// 	ID: uuid.Must(uuid.NewV7()),
	// }
	// err = copier.Copy(&product, data)
	// if err != nil {
	// 	return "", err
	// }

	// product.BrandID = uuid.MustParse(uuid.NewV7())

	err = r.PostgresInfra.DBWrite.WithContext(ctx).Create(data).Error
	if err != nil {
		return res, err
	}
	return data.ID.String(), nil
}

// If you have relationships (foreign keys, etc.), you can use:

// db.Create(&product).Association("Brand").Append(&brand)

// — but for now, plain inserts are fine.
