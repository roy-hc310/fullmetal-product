package product_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
)

func (r *ProductRepository) GetListProduct(ctx context.Context, params *dto_v1.GetListProductRequest) (res []*dto_v1.GetListProductResponse, pagination *dto_v1.PaginationResponse, err error) {

	db := r.PostgresInfra.DBRead.WithContext(ctx).Model(&entity.Product{})

	if params.ShopID != nil {
		db = db.Where("shop_id = ?", params.ShopID)
	}

	if params.Limit == 0 {
		params.Limit = constant.DefaultLimit
	}

	if params.SortBy == "" {
		params.SortBy = constant.DefaultSortBy
	}

	if params.Order == "" {
		params.Order = constant.DefaultOrder
	}

	if params.Cursor != "" {
		if strings.ToLower(params.Order) == "asc" {
			db = db.Where(fmt.Sprintf("%s > ?", params.SortBy), params.Cursor)
		} else {
			db = db.Where(fmt.Sprintf("%s < ?", params.SortBy), params.Cursor)
		}
	}

	db = db.Order(params.SortBy + " " + params.Order).Limit(params.Limit + 1)

	products := []entity.Product{}
	err = db.Find(&products).Error
	if err != nil {
		return nil, nil, err
	}

	hasNext := false
	if len(products) > int(params.Limit) {
		hasNext = true
		products = products[:params.Limit]
	}

	var nextCursor, prevCursor string
	if len(products) > 0 {
		first := products[0]
		last := products[len(products)-1]
		if strings.ToLower(params.Order) == "asc" {
			nextCursor = last.ID.String()
			prevCursor = first.ID.String()
		} else {
			nextCursor = first.ID.String()
			prevCursor = last.ID.String()
		}
	}

	res = []*dto_v1.GetListProductResponse{}
	err = copier.Copy(&res, &products)
	if err != nil {
		return nil, nil, err
	}

	pagination = &dto_v1.PaginationResponse{
		Limit:      int64(params.Limit),
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		HasNext:    hasNext,
	}

	return res, pagination, nil
}
