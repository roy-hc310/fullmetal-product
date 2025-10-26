package product_repository

import (
	"github.com/roy-hc310/fullmetal-product/modules/module_product/entity"
	"github.com/roy-hc310/fullmetal-product/pkg/infrastructure/database"
)

type ProductRepository struct {
	Write *entity.Queries
	Read  *entity.Queries
}

func NewProductRepository(postgresInfra *database.PostgresInfra) *ProductRepository {

	write := entity.New(postgresInfra.DBWrite)
	read := entity.New(postgresInfra.DBRead)

	return &ProductRepository{
		Write: write,
		Read:  read,
	}
}
