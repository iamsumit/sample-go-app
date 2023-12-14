package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	errpkg "github.com/iamsumit/sample-go-app/pkg/error"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

//--------------------------------------------------------------------------
// Error Definitions
//--------------------------------------------------------------------------

var (
	// ErrDecode is used when there is an error decoding the request body.
	ErrDecode = func(err error) *errpkg.Error {
		e := NewError(
			err,
			http.StatusBadRequest,
			nil,
			errpkg.WithMessage("error decoding request body"),
		)

		return e
	}
)

// Decode decodes a JSON request body into the provided type.
func Decode(ctx context.Context, r *http.Request, log logger.Logger, d interface{}) error {
	// Start the tracing here.
	_, span := tracer.Global("api.Decode").Start(ctx, "api.Decode")
	defer span.End()

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		log.Error(
			"decoding error while reading the body",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)

		return ErrDecode(err)
	}
	defer r.Body.Close()

	// Decode JSON into a map[string]interface{}.
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		log.Error(
			"decoding error while unmarshalling the body",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)

		return ErrDecode(err)
	}

	// Use mapstructure to decode the map into the struct.
	err = mapstructure.Decode(data, d)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		log.Error(
			"decoding error while decoding the body into the struct",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)

		return ErrDecode(err)
	}

	span.SetAttributes(
		attribute.String("body", string(body)),
	)

	return nil
}
