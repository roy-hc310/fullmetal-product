package product_service

import (
	"context"
	"fmt"

	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
)

func (s *ProductService) DeleteProduct(ctx context.Context, id string) (res string, traceID string, err error) {
	ctx, span := s.Otel.Tracer().Start(ctx, "span_delete_product")
	defer span.End()
	traceID = span.SpanContext().TraceID().String()

	logger.Info(ctx).
		Str("product_id", id).
		Msg("Deleting product")

	// Delete from database with separate span
	func() {
		_, dbSpan := s.Otel.Tracer().Start(ctx, "db")
		defer dbSpan.End()

		res, err = s.Repository.DeleteProduct(ctx, id)
	}()

	if err != nil {
		logger.Error(ctx).Err(err).
			Str("product_id", id).
			Msg("Failed to delete product from database")
		return res, traceID, err
	}

	logger.Info(ctx).
		Str("product_id", id).
		Msg("Product deleted successfully")

	// Background: invalidate cache and publish event
	go func() {
		ctxBackground := context.Background()

		// Invalidate cache
		cacheKey := fmt.Sprintf("product:%s", id)
		if err := s.Cache.Delete(ctxBackground, cacheKey); err != nil {
			logger.Error(ctxBackground).Err(err).
				Str("product_id", id).
				Str("cache_key", cacheKey).
				Msg("Failed to invalidate cache after delete")
		}

		// Publish delete event
		if err := s.Event.Publish(ctxBackground, constant.ProductDeleteTopic, id, []byte(id)); err != nil {
			logger.Error(ctxBackground).Err(err).
				Str("product_id", id).
				Str("topic", constant.ProductDeleteTopic).
				Msg("Failed to publish product delete event")
		}
	}()

	return res, traceID, nil
}
