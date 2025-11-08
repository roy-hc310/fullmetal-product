package product_repository

import (
	"context"

	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
)

func (r *ProductRepository) GetListProduct(ctx context.Context, data *dto_v1.GetListProductRequest) (res []*dto_v1.GetListProductResponse, pagination *dto_v1.PaginationResponse, err error) {
	params := entity.GetListProductParams{}

	err = copier.Copy(&params, data)
	if err != nil {
		return nil, nil, err
	}

	result, err := r.Read.GetListProduct(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	paginationResult, err := r.Read.CountProduct(ctx)
	if err != nil {
		return nil, nil, err
	}

	res = []*dto_v1.GetListProductResponse{}
	if err := copier.Copy(&res, &result); err != nil {
		return nil, nil, err
	}

	pagination = &dto_v1.PaginationResponse{}
	if err := copier.Copy(&pagination, paginationResult); err != nil {
		return nil, nil, err
	}
	// get previous cursor and next cursorhjhjhjhjsadfjwefojasdf

	return res, pagination, nil
}
