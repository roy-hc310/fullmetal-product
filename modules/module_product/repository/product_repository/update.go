package product_repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
)

func (r *ProductRepository) UpdateProduct(ctx context.Context, data map[string]interface{}) (res string, err error) {

	product := entity.UpdateProductParams{}
	err = copier.Copy(&product, data)
	if err != nil {
		return "", err
	}

	err = r.Write.UpdateProduct(ctx, product)
	if err != nil {
		return "", err
	}

	str, ok := data["id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to convert id to string")
	}
	res = str

	return res, nil
}
