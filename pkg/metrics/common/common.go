package common

import "context"

type Provider interface {
	NewCounter(name string, description string, labels ...string) (Counter, error)
}

type Counter interface {
	Record(ctx context.Context, incr float64, labelValues ...string)
}
