package product_service

import (
	"context"
	"fmt"

	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
	"github.com/roy-hc310/fullmetal-product/pkg/utils"
)

func (s *ProductService) UpdateProduct(ctx context.Context, id string, data map[string]interface{}) (res string, traceID string, err error) {
	ctx, span := s.Otel.Tracer().Start(ctx, "span_update_product")
	defer span.End()
	traceID = span.SpanContext().TraceID().String()

	logger.Info(ctx).
		Str("product_id", id).
		Msg("Updating product")

	// Update in database with separate span
	func() {
		_, dbSpan := s.Otel.Tracer().Start(ctx, "db")
		defer dbSpan.End()

		res, err = s.Repository.UpdateProduct(ctx, id, data)
	}()

	if err != nil {
		logger.Error(ctx).Err(err).
			Str("product_id", id).
			Msg("Failed to update product in database")
		return res, traceID, err
	}

	logger.Info(ctx).
		Str("product_id", id).
		Msg("Product updated successfully")

	// Background: invalidate cache and publish event
	go func() {
		ctxBackground := context.Background()

		// Invalidate cache
		cacheKey := fmt.Sprintf("product:%s", id)
		err = s.Cache.Delete(ctxBackground, cacheKey)
		if err != nil {
			logger.Error(ctxBackground).Err(err).
				Str("product_id", id).
				Str("cache_key", cacheKey).
				Msg("Failed to delete product from cache")
		}

		// Publish update event
		dataByte, err := utils.StructToJSON(data)
		if err != nil {
			logger.Error(ctxBackground).Err(err).
				Str("product_id", id).
				Msg("Failed to convert update data to JSON")
		} else {
			err = s.Event.Publish(ctxBackground, constant.ProductUpdateTopic, id, dataByte)
			if err != nil {
				logger.Error(ctxBackground).Err(err).
					Str("product_id", id).
					Str("topic", constant.ProductUpdateTopic).
					Msg("Failed to publish product update event")
			}
		}
	}()

	return res, traceID, nil
}
