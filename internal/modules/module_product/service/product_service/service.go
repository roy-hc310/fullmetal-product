package product_service

import (
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/cache"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/event"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/telemetry"

	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/repository"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/repository/product_repository"
)

type ProductService struct {
	Repository repository.ProductRepositoryInteface
	Cache      cache.CacheInterface
	Event      event.EventProducerInterface
	Otel       telemetry.OtelInterface
}

func NewProductService(infra *infrastructure.Infrastructure, productRepository *product_repository.ProductRepository) *ProductService {
	return &ProductService{
		Repository: productRepository,
		Cache:      infra.Redis,
		Event:      infra.Kafka,
		Otel:       infra.Otel,
	}
}
