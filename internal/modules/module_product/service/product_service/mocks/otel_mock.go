package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

// MockOtel is a mock implementation of OtelInterface
type MockOtel struct {
	mock.Mock
	tracerProvider trace.TracerProvider
}

// NewMockOtel creates a new MockOtel with a noop tracer provider
func NewMockOtel() *MockOtel {
	return &MockOtel{
		tracerProvider: noop.NewTracerProvider(),
	}
}

type MockTracerProvider struct {
	noop.TracerProvider
}

type MockSpan struct {
	mock.Mock
}

func (m *MockOtel) Tracer() trace.Tracer {
	// Use the noop tracer which properly implements the trace.Tracer interface
	return &mockTracer{}
}

func (m *MockOtel) Shutdown(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// mockTracer wraps noop.Tracer but returns our MockSpan
type mockTracer struct {
	noop.Tracer
}

func (m *mockTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	ms := &mockSpan{}
	// Create a valid trace ID for testing
	traceID, _ := trace.TraceIDFromHex("00000000000000000000000000000001")
	spanID, _ := trace.SpanIDFromHex("0000000000000001")

	ms.On("End").Return()
	ms.On("SpanContext").Return(trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceID,
		SpanID:  spanID,
	}))

	return ctx, ms
}

// MockSpan embeds noop.Span to implement trace.Span interface
type mockSpan struct {
	noop.Span
	mock.Mock
}

func (m *mockSpan) End(options ...trace.SpanEndOption) {
	m.Called()
}

func (m *mockSpan) SpanContext() trace.SpanContext {
	args := m.Called()
	return args.Get(0).(trace.SpanContext)
}
