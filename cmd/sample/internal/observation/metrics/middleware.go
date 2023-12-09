package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
)

// Middleware returns the middleware to record the metrics.
func (m *Handler) Middleware(log logger.Logger) api.Middleware {
	// This is the actual middleware function to be executed.
	mw := func(handler api.Handler) api.Handler {
		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			log.Info(
				"metrics middleware called",
				"method", r.Method,
				"uri", r.RequestURI,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)

			start := time.Now()

			// Run the next handler and catch any propagated error.
			err := handler(ctx, w, r)
			if err != nil {
				return err
			}

			m.AddRequest(r)
			m.AddLatency(r, start)

			log.Info(
				"metrics middleware ended",
				"method", r.Method,
				"uri", r.RequestURI,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)

			return nil
		}

		return h
	}

	return mw
}
