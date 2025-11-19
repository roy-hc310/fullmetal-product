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
	"github.com/roy-hc310/fullmetal-product/pkg/shared"
)

func (s *ProductService) CreateProduct(ctx context.Context, data *dto_v1.CreateProductRequest) (res string, traceID string, err error) {
	ctx, span := s.Otel.Tracer().Start(ctx, "CreateProduct")
	defer span.End()
	traceID = span.SpanContext().TraceID().String()

	product := &entity.Product{
		ID: uuid.Must(uuid.NewV7()),
	}
	err = copier.Copy(&product, data)
	if err != nil {
		return res, traceID, err
	}

	_, dbSpan := s.Otel.Tracer().Start(ctx, "CreateProductDB")
	defer dbSpan.End()

	res, err = s.Repository.CreateProduct(ctx, product)
	if err != nil {
		return res, traceID, err
	}

	cacheKey := fmt.Sprintf("product:%s", product.ID.String())

	jsonData, err := shared.JSONToString(product)
	if err != nil {
		fmt.Printf("Failed to convert to JSON: %v\n", err)
	}

	s.Cache.Set(ctx, cacheKey, jsonData, 24*time.Hour)

	s.Event.Publish(ctx, constant.ProductCreateTopic, product.ID.String(), []byte(jsonData))
	return res, traceID, err
}
