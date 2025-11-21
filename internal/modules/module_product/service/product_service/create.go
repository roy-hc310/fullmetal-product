package product_service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
	"github.com/roy-hc310/fullmetal-product/pkg/utils"
)

func (s *ProductService) CreateProduct(ctx context.Context, data *dto_v1.CreateProductRequest) (res string, traceID string, err error) {
	ctx, span := s.Otel.Tracer().Start(ctx, "span_create_product")
	defer span.End()
	traceID = span.SpanContext().TraceID().String()

	product := &entity.Product{
		ID:     uuid.Must(uuid.NewV7()),
		UserID: utils.GetUserID(ctx),
	}
	err = copier.Copy(product, data)
	if err != nil {
		return res, traceID, err
	}

	err = product.Validation()
	if err != nil {
		return res, traceID, err
	}

	func() {
		_, dbSpan := s.Otel.Tracer().Start(ctx, "db")
		defer dbSpan.End()

		res, err = s.Repository.CreateProduct(ctx, product)
	}()

	if err != nil {
		return res, traceID, err
	}

	go func() {
		ctxBackground := context.Background()

		cacheKey := fmt.Sprintf("product:%s", product.ID.String())

		jsonData, err := utils.StructToJSON(product)
		if err != nil {
			logger.Error(ctxBackground).Err(err).
				Str("product_id", product.ID.String()).
				Msg("Failed to marshal product to JSON for caching")
			return
		}

		err = s.Cache.Set(ctxBackground, cacheKey, string(jsonData), 60*time.Minute)
		if err != nil {
			logger.Error(ctxBackground).Err(err).
				Str("product_id", product.ID.String()).
				Str("cache_key", cacheKey).
				Msg("Failed to cache product")
		}

		err = s.Event.Publish(ctxBackground, constant.ProductCreateTopic, product.ID.String(), jsonData)
		if err != nil {
			logger.Error(ctxBackground).Err(err).
				Str("product_id", product.ID.String()).
				Str("topic", constant.ProductCreateTopic).
				Msg("Failed to publish product create event")
		}
	}()

	return res, traceID, err
}
