package product_service

import (
	"context"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
)

func (s *ProductService) GetListProduct(ctx context.Context, data *dto_v1.GetListProductRequest) (res []*dto_v1.GetListProductResponse, pagination *dto_v1.PaginationResponse, traceID string, err error) {
	ctx, span := s.Otel.Tracer().Start(ctx, "span_get_list_product")
	defer span.End()
	traceID = span.SpanContext().TraceID().String()

	logger.Info(ctx).
		Int("limit", data.Limit).
		Str("cursor", data.Cursor).
		Msg("Fetching product list")

	// Fetch from database with separate span
	func() {
		_, dbSpan := s.Otel.Tracer().Start(ctx, "db")
		defer dbSpan.End()

		res, pagination, err = s.Repository.GetListProduct(ctx, data)
	}()

	if err != nil {
		logger.Error(ctx).Err(err).
			Int("limit", data.Limit).
			Msg("Failed to fetch product list from database")
		return nil, nil, traceID, err
	}

	logger.Info(ctx).
		Int("count", len(res)).
		Bool("has_next", pagination.HasNext).
		Msg("Product list fetched successfully")

	return res, pagination, traceID, nil
}
