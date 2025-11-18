package telemetry

import (
	"context"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type OtelInfra struct {
	tracer         trace.Tracer
	TracerProvider *sdktrace.TracerProvider
}

func NewOtelInfra(ctx context.Context) (*OtelInfra, error) {

	// 1️⃣ Create the OTLP gRPC exporter that sends spans to the collector.
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(config.GlobalEnv.OtelGrpcExporter),
	)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Define resource attributes (metadata about this service).
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(config.GlobalEnv.ServiceName),
		// semconv.ServiceVersionKey.String("1.0.0"),
		// semconv.DeploymentEnvironmentKey.String("production"),
	)

	// 3️⃣ Create a TracerProvider using the exporter and resource.
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// 4️⃣ Set the global provider and propagators.
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	tracer := otel.Tracer("fullmetal-product-tracer")

	return &OtelInfra{
		tracer:         tracer,
		TracerProvider: tracerProvider,
	}, nil
}

func (o *OtelInfra) Shutdown(ctx context.Context) error {
	if o.TracerProvider == nil {
		return nil
	}
	return o.TracerProvider.Shutdown(ctx)
}

func (o *OtelInfra) Tracer() trace.Tracer {
	return o.tracer
}
