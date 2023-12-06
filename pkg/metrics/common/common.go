package common

import (
	"context"
	"net/http"
)

// Provider is the interface for every provider to use.
type Provider interface {
	NewCounter(name string, description string, labels ...string) (Counter, error)
	Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

// Counter is the interface for every counter of the metrics to use.
type Counter interface {
	Record(ctx context.Context, incr float64, labelValues ...string)
}
