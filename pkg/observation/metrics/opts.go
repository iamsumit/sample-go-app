package metrics

import (
	"github.com/iamsumit/sample-go-app/pkg/metrics"
	"github.com/iamsumit/sample-go-app/pkg/metrics/common"
)

// Option is to configure the metrics handler.
type Option func(*Handler)

// newConfig returns a new metrics handler configured with given options.
func newConfig(name string, opts ...Option) (*Handler, error) {
	m := &Handler{}
	m.merge(opts...)

	if m.provider != nil {
		return m, nil
	}

	// Set default meter provider if not provided in option.
	mProvider, err := metrics.New(&metrics.Config{
		Name:     name,
		Type:     metrics.Otel,
		Exporter: metrics.Prometheus,
	})
	if err != nil {
		return nil, err
	}

	m.provider = mProvider

	return m, nil
}

// Merge the configurations by the config set in the options.
func (c *Handler) merge(opts ...Option) {
	if c == nil {
		return
	}

	for _, f := range opts {
		f(c)
	}
}

// WithMetricsProvider sets the metrics provider to the handler.
func WithMetricsProvider(metrics common.Provider) func(*Handler) {
	return func(m *Handler) {
		m.provider = metrics
	}
}

// WithNoMetricsPath sets the no metrics path to the handler.
//
// This will not record the metrics for the given paths.
// This is useful when you want to skip the metrics for health check endpoints.
func WithNoMetricsPath(ntm []string) func(*Handler) {
	return func(m *Handler) {
		m.ntm = ntm
	}
}
