package gokit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// endpoint wraps the endpoint function with tracing.
func (h *Handler) endpoint(r *http.Request, e endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// Start the tracing here.
		_, span := tracer.Global("apigokit.Endpoint").Start(ctx, "apigokit.Endpoint")
		defer span.End()

		span.SetAttributes(
			attribute.String("path", r.URL.Path),
			attribute.String("method", r.Method),
			attribute.String("request_data", fmt.Sprintf("%+v", request)),
		)

		res, err := e(ctx, request)
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}

		return res, nil
	}
}
