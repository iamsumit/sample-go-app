package gokit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// DecoderFunc returns the
type DecoderFunc func(context.Context, *http.Request) (request interface{}, err error)

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
//
//   - method: HTTP method
//   - path: HTTP path
//   - e: endpoint function
//   - d: request decoder function
func (h *Handler) Handle(method string, path string, e endpoint.Endpoint, d DecoderFunc, mw ...api.Middleware) {
	opts := []httptransport.ServerOption{
		// Custom error encoder
		httptransport.ServerErrorEncoder(
			h.encodeError,
		),
	}

	gokitHandler := api.Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		httptransport.NewServer(
			h.endpoint(r, e),
			h.decoder(r, d),
			h.encodeResponse(r),
			opts...,
		).ServeHTTP(w, r.WithContext(ctx))

		return nil
	})

	h.api.Handle(method, path, gokitHandler, mw...)
}

// decoder helps tracing the decoding.
func (h *Handler) decoder(r *http.Request, d DecoderFunc) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		// Start the tracing here.
		_, span := tracer.Global("apigokit.Decode").Start(ctx, "apigokit.Decode")
		defer span.End()

		// Decode the request.
		d, err := d(ctx, r)
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}

		span.SetAttributes(
			attribute.String("path", r.URL.Path),
			attribute.String("method", r.Method),
			attribute.String("data", fmt.Sprintf("%+v", d)),
		)
		return d, nil
	}
}

// encodeError is a transport/http.EncodeErrorFunc that encodes
func (h *Handler) encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	api.RespondWithError(ctx, w, h.log, err)
}

// encodeResponse is a transport/http.EncodeResponseFunc that encodes
func (h *Handler) encodeResponse(r *http.Request) func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		// Start the tracing here.
		_, span := tracer.Global("apigokit.encodeResponse").Start(ctx, "apigokit.encodeResponse")
		defer span.End()

		if e, ok := response.(error); ok {
			// Not a Go kit transport error, but a business-logic error.
			// Provide those as HTTP errors.
			span.SetStatus(codes.Error, e.Error())
			encodeError(ctx, e, w)
			return nil
		}

		status := http.StatusOK

		// Set the default status code for a success kind response.
		//
		// The status codes for error kind response are returned by the
		// handler function. In form of errpkg.Error.
		//
		// See: api.NewError function.
		switch r.Method {
		case http.MethodPost:
			status = http.StatusCreated
		case http.MethodDelete:
			status = http.StatusNoContent
		default:
			status = http.StatusOK
		}

		return api.Respond(ctx, w, response, status)
	}
}

// encodeError is a custom error encoder that encodes the error to the HTTP response.
func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	status := http.StatusBadRequest
	if err == nil {
		err = ErrInternal
		status = http.StatusInternalServerError
	}

	api.Respond(ctx, w, api.NewError(err, status, nil), status)
}
