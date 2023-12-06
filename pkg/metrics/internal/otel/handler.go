package otel

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler is to handle the metrics reports based on the exporter.
func (p Provider) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	switch p.exporter {
	case "prometheus":
		promhttp.Handler().ServeHTTP(w, r)
		return nil
	}

	return nil
}
