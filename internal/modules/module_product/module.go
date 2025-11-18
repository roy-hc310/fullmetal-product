package module_product

import (
	"context"

	"github.com/cloudwego/kitex/server"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure"
	consumer_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/handler/v1/consumer"
	rpc_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/handler/v1/rpc"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/repository/product_repository"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/service/product_service"
)

type ModuleProduct struct {
	ProductConsumer   *consumer_v1.ProductConsumer
	ProductRPC        *rpc_v1.ProductRPC
	ProductService    *product_service.ProductService
	ProductRepository *product_repository.ProductRepository
}

func NewModuleProduct(svr *server.Server, infra *infrastructure.Infrastructure) *ModuleProduct {
	module := &ModuleProduct{}

	module.ProductRepository = product_repository.NewProductRepository(infra.Postgres)
	module.ProductService = product_service.NewProductService(infra, module.ProductRepository)
	module.ProductConsumer = consumer_v1.NewProductConsumer(module.ProductService)
	module.ProductRPC = rpc_v1.NewProductRPC(svr, module.ProductService)
	go infra.Kafka.RegisterConsumer(context.Background(), module.ProductConsumer)

	return module
}
