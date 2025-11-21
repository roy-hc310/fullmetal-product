package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/cloudwego/kitex/server"
	"github.com/roy-hc310/fullmetal-product/internal/bootstrap"
	middleware_rpc "github.com/roy-hc310/fullmetal-product/internal/middleware/rpc"
	"github.com/roy-hc310/fullmetal-product/internal/shared"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
)

func main() {
	err := config.LoadGlobalEnv(".")
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to load environment configuration")
		return
	}

	// Initialize logger with environment configuration
	environment := config.GlobalEnv.Environment
	if environment == "" {
		environment = "development"
	}
	logLevel := config.GlobalEnv.LogLevel
	if logLevel == "" {
		logLevel = "info"
	}
	logger.InitLogger(environment, logLevel)

	logger.Log.Info().
		Str("service", config.GlobalEnv.ServiceName).
		Str("environment", environment).
		Str("log_level", logLevel).
		Msg("Starting microservice")

	addr, err := net.ResolveTCPAddr("tcp", "localhost:"+config.GlobalEnv.RPCPort)
	if err != nil {
		logger.Log.Fatal().Err(err).Str("port", config.GlobalEnv.RPCPort).Msg("Failed to resolve TCP address")
		return
	}

	publicKey, err := shared.ParseRSAPublicKey(config.GlobalEnv.PublicKey)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to parse RSA public key")
		return
	}

	svr := server.NewServer(
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware_rpc.LoggingMiddleware()),
		server.WithMiddleware(middleware_rpc.TimeoutMiddleware()),
		server.WithMiddleware(middleware_rpc.JWTAuthMiddleware(publicKey)),
	)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	application, err := bootstrap.NewApplication(ctx, &config.GlobalEnv, &svr)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to bootstrap application")
		return
	}

	go func() {
		logger.Log.Info().Str("address", addr.String()).Msg("RPC server listening")
		err = svr.Run()
		if err != nil {
			logger.Log.Error().Err(err).Msg("Server stopped with error")
		}
	}()

	<-ctx.Done()
	logger.Log.Info().Msg("Shutdown signal received, gracefully stopping server")
	application.Shutdown(context.Background())
	logger.Log.Info().Msg("Server shutdown complete")
}
