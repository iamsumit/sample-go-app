// Package metrics provides the metrics handler.
//
// It can be used to add metrics for the application.
// It provides Handler to be used to show metrics details.
package metrics

import (
	"net/http"
	"time"

	"github.com/iamsumit/sample-go-app/pkg/metrics/common"
	"github.com/iamsumit/sample-go-app/pkg/util/strings"
)

// Handler is the metrics handler.
type Handler struct {
	name     string
	provider common.Provider
	request  common.Counter
	latency  common.Counter
	ntm      []string
}

// New returns a new metrics handler.
func New(name string, opts ...Option) (*Handler, error) {
	m, err := newConfig(name, opts...)
	if err != nil {
		return nil, err
	}

	m.name = name

	err = m.counter()
	if err != nil {
		return nil, err
	}

	return m, nil
}

// counter registers the counter for metrics.
func (m *Handler) counter() error {
	var err error
	m.request, err = m.provider.NewCounter(m.name+"/request", "Number of requests", "path", "method")
	if err != nil {
		return err
	}

	m.latency, err = m.provider.NewCounter(m.name+"/latency", "Latency of requests", "path", "method")
	if err != nil {
		return err
	}

	return nil
}

// AddRequest adds the request count.
func (m *Handler) AddRequest(r *http.Request) {
	m.request.Record(r.Context(), 1, r.URL.Path, r.Method)
}

// AddLatency adds the latency of the request.
func (m *Handler) AddLatency(r *http.Request, startTime time.Time) {
	m.latency.Record(
		r.Context(),
		float64(time.Now().UnixMilli()-startTime.UnixMilli()),
		r.URL.Path,
		r.Method,
	)
}

// isNoMetricsPath checks if the given path is no metrics path.
func (m *Handler) isNoMetricsPath(r *http.Request) bool {
	if strings.Contain(m.ntm, r.URL.Path) {
		return true
	}

	return false
}
