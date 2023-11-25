package metrics

import (
	"github.com/iamsumit/sample-go-app/pkg/metrics/common"
	"github.com/iamsumit/sample-go-app/pkg/metrics/internal/otel"
)

type ProviderType int

const (
	UnknownProvider ProviderType = iota
	Otel            ProviderType = iota + 1
)

type Exporter int

const (
	UnknownExporter Exporter = iota
	Prometheus      Exporter = iota + 1
)

// We can use enum code generator here as well.
func (pt Exporter) String() string {
	switch pt {
	case Prometheus:
		return "prometheus"
	default:
		return "unknown"
	}
}

type Config struct {
	Name     string
	Type     ProviderType
	Exporter Exporter
}

// New returns a new metrics provider based on the given config.
func New(config *Config) (common.Provider, error) {
	switch config.Type {
	case Otel:
		return otel.New(config.Name, config.Exporter.String())
	default:
		return nil, nil
	}
}
