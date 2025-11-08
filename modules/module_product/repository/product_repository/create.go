package product_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
)

func (r *ProductRepository) CreateProduct(ctx context.Context, data *dto_v1.CreateProductRequest) (res string, err error) {

	product := entity.CreateProductParams{}
	product.ID = uuid.Must(uuid.NewV7())
	err = copier.Copy(&product, data)
	if err != nil {
		return "", err
	}

	// product.BrandID = uuid.MustParse(uuid.NewV7())

	result, err := r.Write.CreateProduct(ctx, product)
	if err != nil {
		return "", err
	}
	res = result.String()
	return res, nil
}
