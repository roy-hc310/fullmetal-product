package middleware_rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
)

// LoggingMiddleware logs RPC requests and responses with timing information
func LoggingMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			start := time.Now()

			// Extract method name from context if available
			method := "unknown"

			// Log incoming request
			logger.Info(ctx).
				Str("method", method).
				Msg("RPC request received")

			// Call the next handler
			err = next(ctx, req, resp)

			// Calculate duration
			duration := time.Since(start)

			// Log response with status
			if err != nil {
				logger.Error(ctx).
					Err(err).
					Str("method", method).
					Dur("duration_ms", duration).
					Msg("RPC request failed")
			} else {
				logger.Info(ctx).
					Str("method", method).
					Dur("duration_ms", duration).
					Msg("RPC request completed")
			}

			return err
		}
	}
}
