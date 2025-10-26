package rpc_v1

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/server"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/service"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/gen/kitex/rpc_product"
	"github.com/roy-hc310/fullmetal-product/pkg/gen/kitex/rpc_product/productservice"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductRPC struct {
	ProductService service.ProductServiceInterface
}

func NewProductRPC(env *config.Env, svr *server.Server, productService service.ProductServiceInterface) *ProductRPC {
	productRPC := &ProductRPC{
		ProductService: productService,
	}
	err := productservice.RegisterService(*svr, productRPC)
	if err != nil {
		fmt.Printf("Failed to register RPC service: %v\n", err)
	}
	return productRPC
}

func (p *ProductRPC) CreateProduct(ctx context.Context, req *rpc_product.CreateProductRequest) (res *rpc_product.CreateProductResponse, err error) {
	res = &rpc_product.CreateProductResponse{}

	product := &dto_v1.CreateProductRequest{}
	if err := copier.Copy(product, req); err != nil {
		return nil, err
	}

	result, _, _ := p.ProductService.CreateProduct(ctx, product)
	res.Id = result
	return res, nil
}

func (p *ProductRPC) DeleteProduct(ctx context.Context, req *rpc_product.DeleteProductRequest) (res *rpc_product.DeleteProductResponse, err error) {
	res = &rpc_product.DeleteProductResponse{}

	result, _, err := p.ProductService.DeleteProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res.Id = result
	return res, nil
}

func (p *ProductRPC) GetDetailProduct(ctx context.Context, req *rpc_product.GetDetailProductRequest) (res *rpc_product.GetDetailProductResponse, err error) {
	res = &rpc_product.GetDetailProductResponse{}

	result, _, err := p.ProductService.GetDetailProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(res, result); err != nil {
		return nil, err
	}
	res.CreatedAt = timestamppb.New(*result.CreatedAt)
	res.UpdatedAt = timestamppb.New(*result.UpdatedAt)
	res.BrandId = result.BrandID
	res.CategoryId = result.CategoryID
	res.ShopId = result.ShopID
	res.Name = result.Name
	res.Sku = result.Sku

	return res, nil
}
func (p *ProductRPC) GetListProduct(ctx context.Context, req *rpc_product.GetListProductRequest) (res *rpc_product.GetListProductResponse, err error) {
	return
}
func (p *ProductRPC) UpdateProduct(ctx context.Context, req *rpc_product.UpdateProductRequest) (res *rpc_product.UpdateProductResponse, err error) {
	return
}
