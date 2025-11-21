package bootstrap

import (
	"context"

	"github.com/cloudwego/kitex/server"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/cache"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/database"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/event"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/telemetry"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
)

type Application struct {
	Infrastructure *infrastructure.Infrastructure
	ModuleProduct  *module_product.ModuleProduct
}

func NewApplication(ctx context.Context, env *config.Env, svr *server.Server) (*Application, error) {
	application := &Application{}
	infrastructure := &infrastructure.Infrastructure{}

	redis, err := cache.NewRedisInfra()
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to initialize Redis")
	}
	infrastructure.Redis = redis

	postgres, err := database.NewPostgresInfra(ctx)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to initialize Postgres")
		return nil, err
	}
	infrastructure.Postgres = postgres

	kafka, err := event.NewKafkaInfra()
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to initialize Kafka")
	}
	infrastructure.Kafka = kafka

	otel, err := telemetry.NewOtelInfra(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to initialize OpenTelemetry")
	}
	infrastructure.Otel = otel
	application.Infrastructure = infrastructure

	application.ModuleProduct = module_product.NewModuleProduct(ctx, svr, infrastructure)

	return application, nil
}

func (app *Application) Shutdown(ctx context.Context) {
	if app.Infrastructure.Redis != nil {
		err := app.Infrastructure.Redis.Shutdown()
		if err != nil {
			logger.Log.Error().Err(err).Msg("Failed to shutdown Redis")
		}
	}

	if app.Infrastructure.Kafka != nil {
		err := app.Infrastructure.Kafka.Shutdown(ctx)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Failed to shutdown Kafka")
		}
	}

	if app.Infrastructure.Otel != nil {
		err := app.Infrastructure.Otel.Shutdown(ctx)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Failed to shutdown OpenTelemetry")
		}
	}
}
