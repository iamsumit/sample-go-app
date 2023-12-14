package tracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

// JaegerConfig is the configuration for Jaeger endpoint.
type JaegerConfig struct {
	Host string
	Path string
}

// Config is to configure the tracer provider.
type Config struct {
	Name        string
	ServiceName string
	Jaeger      JaegerConfig
}

// Tracer contains the OpenTelemetry tracer.
type Tracer struct {
	tracer   trace.Tracer
	provider *sdktrace.TracerProvider
}

// New returns the tracer instance to be used.
func New(ctx context.Context, config *Config) (*Tracer, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceName(config.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(config.Jaeger.Host),
		otlptracehttp.WithURLPath(config.Jaeger.Path),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	return &Tracer{
		tracer: otel.Tracer(config.Name),
	}, nil
}

// Global returns the global tracer instance.
func Global(name string) *Tracer {
	t := &Tracer{
		tracer: otel.Tracer(name),
	}

	return t
}

// Start starts a new span with the given operation name.
func (t Tracer) Start(ctx context.Context, operationName string) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, operationName)
}
