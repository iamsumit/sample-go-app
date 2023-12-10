package middleware

import (
	"context"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/db"
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
				case *api.Error:
					// The errors provided by the api package related stuff.
					apiErr := err.(*api.Error)

					// Update any attributes or status set in the apiErr.
					er.Data = apiErr.Attributes
					status = apiErr.Status
				case *validator.Error:
					// The errors provided by the validate package related stuff.
					vErr := err.(*validator.Error)

					// Update any attributes or status set in the vErr.
					er.Data = vErr.Attributes
					status = vErr.Status
				case *db.Error:
					// The errors provided by the db package related stuff.
					dbErr := err.(*db.Error)

					// Log the actual error message.
					//
					// This is what will be shown in the services' logs for
					// internal purpose only.
					//
					// The error message prefixed with "internal:",
					// will be updated for clients.
					log.Error("DATABASE ERROR", "error", dbErr.Err.Error())

					// Update any attributes or status set in the dbErr.
					er.Data = dbErr.Attributes
					status = dbErr.Status

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
