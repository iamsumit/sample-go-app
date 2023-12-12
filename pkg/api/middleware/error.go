package middleware

import (
	"context"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/api"
	errpkg "github.com/iamsumit/sample-go-app/pkg/error"
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
				return api.NewError(err, http.StatusInternalServerError, nil)
			}

			// Run the next handler and catch any propagated error.
			err = handler(ctx, w, r)
			if err != nil {
				log.Error("CLIENT ERROR", "error", err)

				api.SetIsError(ctx)
				var status int

				// Create a new error response.
				er := api.ErrorResponse{
					Error: err.Error(),
				}

				// Check if this is a normal or a wrapped error.
				switch err.(type) {
				case *errpkg.Error:
					// The errors provided by the error package related stuff.
					e := err.(*errpkg.Error)

					// Log the actual error message.
					//
					// This is what will be shown in the services' logs for
					// internal purpose only.
					//
					// The error message prefixed with "internal:",
					// will be updated for clients.
					log.Error(
						"ORIGINAL ERROR",
						"error_type", e.Type(),
						"error", e.OriginalError(),
						"status", e.StatusCode(),
						"attributes", e.Attributes(),
						"path", r.URL.Path,
						"method", r.Method,
					)

					// Update any attributes or status set in the error.
					er.Data = e.Attributes()
					status = e.StatusCode()
				default:
					// This is an unknown error. Log it and set the status code
					status = http.StatusInternalServerError
				}

				// Respond with the error back to the client
				if err := api.Respond(ctx, w, er, status); err != nil {
					return err
				}
			}

			return nil
		}

		return h
	}

	return m
}
