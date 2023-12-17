package middleware

import (
	"context"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
)

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
func Errors(log logger.Logger) api.Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler api.Handler) api.Handler {
		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			_, err := api.GetContextValues(ctx)
			if err != nil {
				log.Error(
					"api value missing from context",
					"error", err,
					"from", "api - error middleware",
				)
				return api.NewError(err, http.StatusInternalServerError, nil)
			}

			// Run the next handler and catch any propagated error.
			err = handler(ctx, w, r)
			if err != nil {
				// Respond with the error back to the client
				if err := api.RespondWithError(ctx, w, log, err); err != nil {
					return err
				}

			}

			return nil
		}

		return h
	}

	return m
}
