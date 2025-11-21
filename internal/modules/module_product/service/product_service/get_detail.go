package product_service

import (
	"context"
	"fmt"
	"time"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
	"github.com/roy-hc310/fullmetal-product/pkg/utils"
)

func (s *ProductService) GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, traceID string, err error) {
	ctx, span := s.Otel.Tracer().Start(ctx, "span_get_detail_product")
	defer span.End()
	traceID = span.SpanContext().TraceID().String()

	cacheKey := fmt.Sprintf("product:%s", id)

	logger.Info(ctx).
		Str("product_id", id).
		Msg("Fetching product detail")

	// 1. Try cache
	val, cacheErr := s.Cache.Get(ctx, cacheKey)
	if cacheErr == nil && val != "" {
		// Cache hit → unmarshal JSON
		var cached dto_v1.GetDetailProductResponse
		jsonErr := utils.JSONToStruct([]byte(val), &cached)
		if jsonErr == nil {
			logger.Debug(ctx).
				Str("product_id", id).
				Str("cache_key", cacheKey).
				Msg("Redis cache hit")
			return &cached, traceID, nil
		}
		// If unmarshal fails → fallback to DB
		logger.Warn(ctx).Err(jsonErr).
			Str("product_id", id).
			Msg("Failed to unmarshal cached product, fetching from DB")
	} else if cacheErr != nil {
		logger.Debug(ctx).Err(cacheErr).
			Str("product_id", id).
			Msg("Redis cache miss")
	}

	// 2. Cache miss → fetch from DB
	func() {
		_, dbSpan := s.Otel.Tracer().Start(ctx, "db")
		defer dbSpan.End()

		res, err = s.Repository.GetDetailProduct(ctx, id)
	}()

	if err != nil {
		logger.Error(ctx).Err(err).
			Str("product_id", id).
			Msg("Failed to fetch product from database")
		return nil, traceID, err
	}

	// 3. Write result to cache
	if res != nil {
		jsonData, marshalErr := utils.StructToJSON(res)
		if marshalErr == nil {
			if cacheErr := s.Cache.Set(ctx, cacheKey, string(jsonData), time.Hour); cacheErr != nil {
				logger.Error(ctx).Err(cacheErr).
					Str("product_id", id).
					Str("cache_key", cacheKey).
					Msg("Failed to cache product after DB fetch")
			} else {
				logger.Debug(ctx).
					Str("product_id", id).
					Msg("Product cached successfully")
			}
		}
	}

	return res, traceID, nil
}
