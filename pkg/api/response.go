package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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
	v, err := GetContextValues(ctx)
	if err != nil {
		return NewRequestError(err, http.StatusInternalServerError)
	}

	// Set the status code for the request logger middleware in the context.
	err = SetStatusCode(ctx, statusCode)
	if err != nil {
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
		return err
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
