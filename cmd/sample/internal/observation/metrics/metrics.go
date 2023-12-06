package metrics

import (
	"net/http"
	"time"

	"github.com/iamsumit/sample-go-app/pkg/metrics/common"
)

// Handler is the metrics handler.
type Handler struct {
	provider common.Provider
	request  common.Counter
	latency  common.Counter
}

// New returns a new metrics handler.
func New(name string, opts ...Option) (*Handler, error) {
	m, err := newConfig(name, opts...)
	if err != nil {
		return nil, err
	}

	m.counter()

	return m, nil
}

// counter registers the counter for metrics.
func (m *Handler) counter() error {
	var err error
	m.request, err = m.provider.NewCounter("sample_request", "Number of requests", "path", "method")
	if err != nil {
		return err
	}

	m.latency, err = m.provider.NewCounter("sample_latency", "Latency of requests", "path", "method")
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
