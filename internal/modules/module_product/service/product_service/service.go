package product_service

import (
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/cache"

	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/repository"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/repository/product_repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepositoryInteface
	ProductCache      cache.CacheInterface
}

func NewProductService(infra *infrastructure.Infrastructure, productRepository *product_repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
		ProductCache:      infra.Redis,
	}
}
