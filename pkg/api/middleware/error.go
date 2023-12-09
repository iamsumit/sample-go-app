package middleware

import (
	"context"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/validator"
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
				return api.NewRequestError(err, http.StatusInternalServerError)
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
				case *api.RequestError:
					status = err.(*api.RequestError).Status
				case validator.Error:
					er.Data = validator.Attributes(err)
					status = http.StatusBadRequest
				default:
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
