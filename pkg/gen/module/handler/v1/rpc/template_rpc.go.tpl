package {{.Snake}}_rpc_v1

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/server"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/gen/kitex/rpc_{{.Snake}}"
	"github.com/roy-hc310/fullmetal-product/pkg/gen/kitex/rpc_{{.Snake}}/{{.Snake}}service"
)

type ProductRPC struct {
}

func NewProductRPC(env *config.Env, svr *server.Server) *ProductRPC {
	{{.Snake}}RPC := &ProductRPC{}
	err := {{.Snake}}service.RegisterService(*svr, {{.Snake}}RPC)
	if err != nil {
		fmt.Printf("Failed to register RPC service: %v\n", err)
	}
	return {{.Snake}}RPC
}

func (p *ProductRPC) CreateProduct(ctx context.Context, req *rpc_{{.Snake}}.CreateProductRequest) (res *rpc_{{.Snake}}.CreateProductResponse, err error) {
	return
}
func (p *ProductRPC) GetDetailProduct(ctx context.Context, req *rpc_{{.Snake}}.GetDetailProductRequest) (res *rpc_{{.Snake}}.GetDetailProductResponse, err error) {
	return
}
func (p *ProductRPC) GetListProduct(ctx context.Context, req *rpc_{{.Snake}}.GetListProductRequest) (res *rpc_{{.Snake}}.GetListProductResponse, err error) {
	return
}
func (p *ProductRPC) UpdateProduct(ctx context.Context, req *rpc_{{.Snake}}.UpdateProductRequest) (res *rpc_{{.Snake}}.UpdateProductResponse, err error) {
	return
}

func (p *ProductRPC) DeleteProduct(ctx context.Context, req *rpc_{{.Snake}}.DeleteProductRequest) (res *rpc_{{.Snake}}.DeleteProductResponse, err error) {
	return
}
