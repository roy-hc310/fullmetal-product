package product_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
)

func (r *ProductRepository) GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, err error) {

	parseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	result, err := r.Read.GetDetailProduct(ctx, parseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return nil, err
	}

	res = &dto_v1.GetDetailProductResponse{}
	if err := copier.Copy(res, &result); err != nil {
		return nil, err
	}

	return res, nil
}
