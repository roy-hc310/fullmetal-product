package product_repository

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
)

func (r *ProductRepository) GetListProduct(ctx context.Context, data *dto_v1.GetListProductRequest) (res *dto_v1.GetListProductResponse, err error) {
	return res, nil
}
