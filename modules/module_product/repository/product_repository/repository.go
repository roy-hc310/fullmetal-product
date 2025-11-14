package product_repository

import (
	"github.com/roy-hc310/fullmetal-product/pkg/infrastructure/database"
)

type ProductRepository struct {
	PostgresInfra *database.PostgresInfra
}

func NewProductRepository(postgresInfra *database.PostgresInfra) *ProductRepository {

	return &ProductRepository{
		PostgresInfra: postgresInfra,
	}
}
