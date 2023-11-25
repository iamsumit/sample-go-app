package otel

import (
	"errors"

	"github.com/iamsumit/sample-go-app/pkg/metrics/common"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

type Provider struct {
	provider *metric.MeterProvider
	meter    api.Meter
}

func New(name string, exporter string) (*Provider, error) {
	switch exporter {
	case "prometheus":
		promExporter, err := newPrometheusExporter()
		if err != nil {
			return nil, err
		}

		provider := metric.NewMeterProvider(metric.WithReader(promExporter))

		return &Provider{
			provider: provider,
			meter:    provider.Meter(name),
		}, nil
	}

	return nil, errors.New("unknown exporter")
}

func newPrometheusExporter() (metric.Reader, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	return exporter, nil
}

type Counter struct {
	counter api.Float64Counter
	labels  []string
}

// Creates a new counter of metrics with given name and labels.
func (p Provider) NewCounter(name string, description string, labels ...string) (common.Counter, error) {
	// This is the equivalent of prometheus.NewCounterVec
	counter, err := p.meter.Float64Counter(name, api.WithDescription(description))
	if err != nil {
		return nil, err
	}

	return &Counter{
		counter: counter,
		labels:  labels,
	}, nil
}
