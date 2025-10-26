package product_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
)

func (r *ProductRepository) GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, err error) {
	var parseID pgtype.UUID
	parseID.Scan(id)

	result, err := r.Read.GetDetailProduct(ctx, parseID)
	if err != nil {
		return nil, err
	}

	res = &dto_v1.GetDetailProductResponse{}
	if err := copier.Copy(res, result); err != nil {
		return nil, err
	}

	return res, nil
}
