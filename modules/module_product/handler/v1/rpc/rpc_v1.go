package rpc_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/server"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/modules/module_product/service"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/gen/kitex/rpc_product"
	"github.com/roy-hc310/fullmetal-product/pkg/gen/kitex/rpc_product/productservice"
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
	res = &rpc_product.CreateProductResponse{
		Data: &rpc_product.ProductID{},
		Meta: &rpc_product.ResponseMeta{},
	}

	product := &dto_v1.CreateProductRequest{}
	if err := copier.Copy(product, req); err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, err
	}

	result, traceID, err := p.ProductService.CreateProduct(ctx, product)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, err
	}

	res.Data.Id = result
	res.Meta.TraceId = traceID
	res.Meta.Success = true

	return res, nil
}

func (p *ProductRPC) DeleteProduct(ctx context.Context, req *rpc_product.DeleteProductRequest) (res *rpc_product.DeleteProductResponse, err error) {
	res = &rpc_product.DeleteProductResponse{
		Data: &rpc_product.ProductID{},
		Meta: &rpc_product.ResponseMeta{},
	}

	result, traceID, err := p.ProductService.DeleteProduct(ctx, req.Id)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, err
	}

	res.Data.Id = result
	res.Meta.TraceId = traceID
	res.Meta.Success = true

	return res, nil
}

func (p *ProductRPC) GetDetailProduct(ctx context.Context, req *rpc_product.GetDetailProductRequest) (res *rpc_product.GetDetailProductResponse, err error) {
	res = &rpc_product.GetDetailProductResponse{
		// Data: &rpc_product.Product{},
		Meta: &rpc_product.ResponseMeta{},
	}

	result, traceID, err := p.ProductService.GetDetailProduct(ctx, req.Id)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, nil
	}

	if result != nil {
		res.Data = &rpc_product.Product{}
		err = copier.Copy(&res.Data, result)
		if err != nil {
			res.Meta.Errors = append(res.Meta.Errors, err.Error())
			return res, nil
		}
	}

	res.Meta.TraceId = traceID
	res.Meta.Success = true

	return res, nil
}
func (p *ProductRPC) GetListProduct(ctx context.Context, req *rpc_product.GetListProductRequest) (res *rpc_product.GetListProductResponse, err error) {
	res = &rpc_product.GetListProductResponse{
		Data:       []*rpc_product.Product{},
		Meta:       &rpc_product.ResponseMeta{},
		Pagination: &rpc_product.PaginationResponse{},
	}

	query := &dto_v1.GetListProductRequest{}
	err = copier.Copy(&query, req)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, nil
	}

	result, pagination, traceID, err := p.ProductService.GetListProduct(ctx, query)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, nil
	}

	err = copier.CopyWithOption(&res.Data, &result, copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: "",
				Fn: func(src interface{}) (interface{}, error) {
					t := src.(time.Time)
					return t.Format(time.RFC3339), nil
				},
			},
		},
	})
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, nil
	}

	res.Meta.TraceId = traceID
	res.Meta.Success = true
	res.Pagination.NextCursor = pagination.NextCursor
	res.Pagination.PrevCursor = pagination.PrevCursor
	res.Pagination.Limit = pagination.Limit

	return res, nil
}

func (p *ProductRPC) UpdateProduct(ctx context.Context, req *rpc_product.UpdateProductRequest) (res *rpc_product.UpdateProductResponse, err error) {
	res = &rpc_product.UpdateProductResponse{
		Data: &rpc_product.ProductID{},
		Meta: &rpc_product.ResponseMeta{},
	}

	// if err := copier.Copy(&product, req); err != nil {
	// 	res.Meta.Errors = append(res.Meta.Errors, err.Error())
	// 	return res, nil
	// }

	jsonData, err := json.Marshal(req)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, nil
	}

	product := map[string]interface{}{}
	err = json.Unmarshal(jsonData, &product)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, nil
	}

	result, traceID, err := p.ProductService.UpdateProduct(ctx, product)
	if err != nil {
		res.Meta.Errors = append(res.Meta.Errors, err.Error())
		return res, nil
	}

	res.Data.Id = result
	res.Meta.TraceId = traceID
	res.Meta.Success = true

	return res, nil
}
