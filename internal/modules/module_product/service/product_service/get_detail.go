package product_service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
)

func (s *ProductService) GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, traceID string, err error) {
	cacheKey := fmt.Sprintf("product:%s", id)

	// 1. Try cache
	val, cacheErr := s.ProductCache.Get(ctx, cacheKey)
	if cacheErr == nil && val != "" {
		// Cache hit → unmarshal JSON
		var cached dto_v1.GetDetailProductResponse
		if jsonErr := json.Unmarshal([]byte(val), &cached); jsonErr == nil {
			return &cached, traceID, nil
		}
		// If unmarshal fails → fallback to DB
	}

	// 2. Cache miss → fetch from DB
	res, err = s.ProductRepository.GetDetailProduct(ctx, id)
	if err != nil {
		return nil, traceID, err
	}

	// 3. Write result to cache
	if res != nil {
		jsonData, marshalErr := json.Marshal(res)
		if marshalErr == nil {
			_ = s.ProductCache.Set(ctx, cacheKey, string(jsonData), time.Hour)
		}
	}

	return res, traceID, nil
}
