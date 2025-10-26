package module_product

import (
	"github.com/cloudwego/kitex/server"
	rpc_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/handler/v1/rpc"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/repository/product_repository"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/service/product_service"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/infrastructure"
)

type ModuleProduct struct {
	ProductRPC        *rpc_v1.ProductRPC
	ProductService    *product_service.ProductService
	ProductRepository *product_repository.ProductRepository
}

func NewModuleProduct(env *config.Env, svr *server.Server, infra *infrastructure.Infrastructure) *ModuleProduct {
	module := &ModuleProduct{}

	module.ProductRepository = product_repository.NewProductRepository(infra.Postgres)
	module.ProductService = product_service.NewProductService(module.ProductRepository)
	module.ProductRPC = rpc_v1.NewProductRPC(env, svr, module.ProductService)

	return module
}
