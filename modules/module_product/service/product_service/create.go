package product_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
)

func (s *ProductService) CreateProduct(ctx context.Context, data *dto_v1.CreateProductRequest) (res string, traceID string, err error) {

	data.BrandID = uuid.Must(uuid.NewV7()).String()
	data.CategoryID = uuid.Must(uuid.NewV7()).String()
	data.ShopID = uuid.Must(uuid.NewV7()).String()

	product := &entity.Product{
		ID: uuid.Must(uuid.NewV7()),
	}
	err = copier.Copy(&product, data)
	if err != nil {
		return res, traceID, err
	}

	res, err = s.ProductRepository.CreateProduct(ctx, product)
	return res, traceID, err
}
