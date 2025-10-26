package infrastructure

import (
	"github.com/roy-hc310/fullmetal-product/pkg/infrastructure/cache"
	"github.com/roy-hc310/fullmetal-product/pkg/infrastructure/database"
	"github.com/roy-hc310/fullmetal-product/pkg/infrastructure/event"
	"github.com/roy-hc310/fullmetal-product/pkg/infrastructure/telemetry"
)

type Infrastructure struct {
	Redis    *cache.RedisInfra
	Postgres *database.PostgresInfra
	Kafka    *event.KafkaInfra
	Otel     *telemetry.OtelInfra
}
