package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	errpkg "github.com/iamsumit/sample-go-app/pkg/error"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"go.opentelemetry.io/otel/codes"
)

//--------------------------------------------------------------------------
// Error response
//--------------------------------------------------------------------------

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	// in:body
	//
	//example: data is not in proper format
	Error string `json:"error"`
	// in:body
	//
	//example: {"field": "error message for this specific field"}
	Data map[string]interface{} `json:"data,omitempty"`
}

//--------------------------------------------------------------------------
// The response
//--------------------------------------------------------------------------

// Response is the form used for API responses for success in the API.
type Response struct {
	// Success
	//
	Success bool `json:"success"`
	// Status
	//
	Status int `json:"status"`
	// Timestamp
	//
	// example: 1234567
	Timestamp int64 `json:"timestamp"`
	// Data
	// in: body
	Data interface{} `json:"data,omitempty"`
	// Errors
	// in: body
	Errors interface{} `json:"errors,omitempty"`
}

// Respond returns json to client.
func Respond(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {
	// Start the tracing here.
	ctx, span := tracer.Global("api.Response").Start(ctx, "api.Response")
	defer span.End()

	v, err := GetContextValues(ctx)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return NewError(err, http.StatusInternalServerError, nil)
	}

	// Set the status code for the request logger middleware in the context.
	err = SetStatusCode(ctx, statusCode)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	r := Response{
		Status:    statusCode,
		Timestamp: time.Now().UTC().Unix(),
	}

	if !v.IsError {
		r.Success = true
		r.Data = data
	} else {
		r.Success = false
		r.Errors = data
	}

	// Convert the response to json
	jd, err := json.Marshal(r)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	if v.IsError {
		span.SetStatus(codes.Error, string(jd))
	} else {
		span.SetStatus(codes.Ok, "request completed successfully")
	}

	// set the content type now that we know there was no marshal error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Send the result back to the client
	if _, err := w.Write(jd); err != nil {
		return fmt.Errorf("write fail: %+v", jd)
	}

	return nil
}

// RespondWithError returns parsed error to client.
func RespondWithError(ctx context.Context, w http.ResponseWriter, log logger.Logger, err error) error {
	// Log the error.
	log.Error("CLIENT ERROR", "error", err)

	// Set the error in the context.
	SetIsError(ctx)

	status := http.StatusInternalServerError

	// Create a new error response.
	er := ErrorResponse{
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
		log.Error(
			"ORIGINAL ERROR",
			"error_type", e.Type(),
			"error", e.OriginalError(),
			"status", e.StatusCode(),
			"attributes", e.Attributes(),
		)

		// Update any attributes or status set in the error.
		er.Data = e.Attributes()
		status = e.StatusCode()
	default:
		// This is an unknown error. Log it and set the status code
		status = http.StatusInternalServerError
	}

	// Respond with the error back to the client
	if err := Respond(ctx, w, er, status); err != nil {
		return err
	}

	return nil
}
