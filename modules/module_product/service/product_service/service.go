package product_service

import (
	"github.com/roy-hc310/fullmetal-product/modules/module_product/repository"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/repository/product_repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepositoryInteface
}

func NewProductService(productRepository *product_repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}
