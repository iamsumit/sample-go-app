package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
)

// Logger is a middleware that logs the incoming request information.
func Logger(log logger.Logger) api.Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler api.Handler) api.Handler {
		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			start := time.Now()
			log.Info(
				"request started",
				"method", r.Method,
				"uri", r.RequestURI,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
				"start_time", start,
			)

			v, err := api.GetContextValues(ctx)
			if err != nil {
				log.Error("api value missing from context", "error", err)
				return api.NewError(err, http.StatusInternalServerError, nil)
			}

			// Call the next handler
			err = handler(ctx, w, r)
			if err != nil {
				log.Info("request error",
					"error", err.Error(),
					"status_code", v.StatusCode,
					"method", r.Method,
					"uri", r.RequestURI,
					"remote_addr", r.RemoteAddr,
					"user_agent", r.UserAgent(),
				)
				return err
			}

			latency := time.Since(start)
			s := float64(latency.Microseconds()) / float64(1000000)
			log.Info("request completed",
				"duration_s", s,
				"status_code", v.StatusCode,
				"method", r.Method,
				"uri", r.RequestURI,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)
			return err
		}
		return h
	}
	return m
}
