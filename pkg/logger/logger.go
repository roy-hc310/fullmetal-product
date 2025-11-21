package logger

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

var Log zerolog.Logger

// InitLogger initializes the global logger with the specified environment and log level
func InitLogger(environment, level string) {
	// Set log level
	logLevel := parseLogLevel(level)
	zerolog.SetGlobalLevel(logLevel)

	// Configure output format based on environment
	var output io.Writer
	if strings.ToLower(environment) == "production" {
		// Production: JSON format
		output = os.Stdout
		Log = zerolog.New(output).With().Timestamp().Logger()
	} else {
		// Development: Pretty format with colors
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		Log = zerolog.New(output).With().Timestamp().Caller().Logger()
	}
}

// parseLogLevel converts string log level to zerolog.Level
func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn", "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

// FromContext extracts logger from context with trace ID if available
func FromContext(ctx context.Context) zerolog.Logger {
	logger := Log

	// Extract trace ID from OpenTelemetry context
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		traceID := span.SpanContext().TraceID().String()
		spanID := span.SpanContext().SpanID().String()
		logger = logger.With().
			Str("trace_id", traceID).
			Str("span_id", spanID).
			Logger()
	}

	return logger
}

// WithTraceID adds trace ID to the logger
func WithTraceID(ctx context.Context, logger zerolog.Logger) zerolog.Logger {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		return logger.With().
			Str("trace_id", span.SpanContext().TraceID().String()).
			Str("span_id", span.SpanContext().SpanID().String()).
			Logger()
	}
	return logger
}

// Info creates an info level log event
func Info(ctx context.Context) *zerolog.Event {
	logger := FromContext(ctx)
	return logger.Info()
}

// Debug creates a debug level log event
func Debug(ctx context.Context) *zerolog.Event {
	logger := FromContext(ctx)
	return logger.Debug()
}

// Warn creates a warn level log event
func Warn(ctx context.Context) *zerolog.Event {
	logger := FromContext(ctx)
	return logger.Warn()
}

// Error creates an error level log event
func Error(ctx context.Context) *zerolog.Event {
	logger := FromContext(ctx)
	return logger.Error()
}

// Fatal creates a fatal level log event
func Fatal(ctx context.Context) *zerolog.Event {
	logger := FromContext(ctx)
	return logger.Fatal()
}
