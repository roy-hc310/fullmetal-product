package middleware_rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
)

func TimeoutMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			if _, ok := ctx.Deadline(); !ok {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, time.Duration(config.GlobalEnv.ContextTimeout)*time.Second)
				defer cancel()
			}
			return next(ctx, req, resp)
		}
	}
}
