package infrastructure

import (
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/cache"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/database"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/event"
	"github.com/roy-hc310/fullmetal-product/internal/infrastructure/telemetry"
)

type Infrastructure struct {
	Redis    *cache.RedisInfra
	Postgres *database.PostgresInfra
	Kafka    *event.KafkaInfra
	Otel     *telemetry.OtelInfra
}
