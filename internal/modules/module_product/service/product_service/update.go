package product_service

import (
	"context"
	"fmt"
	"log"

	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/roy-hc310/fullmetal-product/pkg/shared"
)

func (s *ProductService) UpdateProduct(ctx context.Context, id string, data map[string]interface{}) (res string, traceID string, err error) {

	res, err = s.Repository.UpdateProduct(ctx, id, data)
	if err != nil {
		return res, traceID, err
	}

	cacheKey := fmt.Sprintf("product:%s", id)
	err = s.Cache.Delete(ctx, cacheKey)
	if err != nil {
		log.Default().Println("UpdateProduct: ", err)
	}

	jsonData, err := shared.JSONToString(data)
	if err != nil {
		fmt.Printf("Failed to convert to JSON: %v\n", err)
	} else {
		err = s.Event.Publish(ctx, constant.ProductUpdateTopic, id, []byte(jsonData))
		if err != nil {
			log.Default().Println("UpdateProduct: ", err)
		}
	}

	return res, traceID, nil
}
