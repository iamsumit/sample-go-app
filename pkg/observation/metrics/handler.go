package metrics

import "github.com/iamsumit/sample-go-app/pkg/api"

// Handler returns the handler for the /metrics endpoint.
func (m *Handler) Handler() api.Handler {
	return m.provider.Handler
}
