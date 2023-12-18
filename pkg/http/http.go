package http

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// Client is a HTTP client.
type Client struct {
	client *http.Client
}

// NewClient returns a new Client
func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

// Post sends a POST request to the given URL with the given data.
func (c *Client) Post(ctx context.Context, u url.URL, body io.Reader) (*http.Response, error) {
	// Start the tracing here.
	ctx, span := tracer.Global("http.Post").Start(ctx, u.Path)
	defer span.End()

	// This will pass the current context to the activity.
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), body)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	span.SetAttributes(
		attribute.String("url", u.String()),
		attribute.Int("status_code", res.StatusCode),
	)

	return res, nil
}
