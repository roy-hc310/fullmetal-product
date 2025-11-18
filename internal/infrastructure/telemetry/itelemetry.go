package telemetry

import "go.opentelemetry.io/otel/trace"

type OtelInterface interface {
	Tracer() trace.Tracer
}
