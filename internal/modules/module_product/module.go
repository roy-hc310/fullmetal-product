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
	Consumer   *consumer_v1.ProductConsumer
	RPC        *rpc_v1.ProductRPC
	Service    *product_service.ProductService
	Repository *product_repository.ProductRepository
}

func NewModuleProduct(ctx context.Context, svr *server.Server, infra *infrastructure.Infrastructure) *ModuleProduct {
	module := &ModuleProduct{}

	module.Repository = product_repository.NewProductRepository(infra.Postgres)
	module.Service = product_service.NewProductService(infra, module.Repository)
	module.Consumer = consumer_v1.NewProductConsumer(module.Service)
	module.RPC = rpc_v1.NewProductRPC(svr, module.Service)
	go infra.Kafka.RegisterConsumer(ctx, module.Consumer)

	return module
}
