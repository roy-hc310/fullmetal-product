package bootstrap

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/server"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/cache"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/database"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/event"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/telemetry"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
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
		fmt.Printf("Failed to initialize Redis: %v\n", err)
	}
	infrastructure.Redis = redis

	postgres, err := database.NewPostgresInfra(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize Postgres: %v", err)
		return nil, err
	}
	infrastructure.Postgres = postgres

	kafka, err := event.NewKafkaInfra()
	if err != nil {
		fmt.Printf("Failed to initialize Kafka: %v\n", err)
	}
	infrastructure.Kafka = kafka

	otel, err := telemetry.NewOtelInfra(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize OpenTelemetry: %v\n", err)
	}
	infrastructure.Otel = otel
	application.Infrastructure = infrastructure

	application.ModuleProduct = module_product.NewModuleProduct(env, svr, infrastructure)

	return application, nil
}

func (app *Application) Shutdown(ctx context.Context) {
	if app.Infrastructure.Redis != nil {
		err := app.Infrastructure.Redis.Shutdown()
		if err != nil {
			fmt.Printf("Failed to shutdown Redis: %v\n", err)
		}
	}

	if app.Infrastructure.Kafka != nil {
		err := app.Infrastructure.Kafka.Shutdown(ctx)
		if err != nil {
			fmt.Printf("Failed to shutdown Kafka: %v\n", err)
		}
	}

	if app.Infrastructure.Otel != nil {
		err := app.Infrastructure.Otel.Shutdown(ctx)
		if err != nil {
			fmt.Printf("Failed to shutdown OpenTelemetry: %v\n", err)
		}
	}
}
