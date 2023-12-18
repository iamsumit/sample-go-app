package activity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/iamsumit/sample-go-app/pkg/api"
	httppkg "github.com/iamsumit/sample-go-app/pkg/http"
	"github.com/iamsumit/sample-go-app/pkg/util/app"
)

// Client holds the activity client.
type Client struct {
	host   string
	client *http.Client
}

// New returns a new activity client.
func New(host string) *Client {
	return &Client{
		host:   host,
		client: &http.Client{},
	}
}

// Activity holds the activity information.
type Activity struct {
	AppName   string `json:"app_name"`
	Entity    string `json:"entity"`
	Operation string `json:"operation"`
}

// Record sends the activity to the activity service.
func (c *Client) Record(ctx context.Context, entity string, operation string) error {
	act := Activity{
		AppName:   app.Name(),
		Entity:    entity,
		Operation: operation,
	}

	body, err := json.Marshal(act)
	if err != nil {
		return err
	}

	hc := httppkg.NewClient()
	res, err := hc.Post(ctx, url.URL{
		Scheme: "http",
		Host:   c.host,
		Path:   "/v1/activity",
	}, bytes.NewBuffer(body))

	if res.StatusCode != http.StatusCreated {
		return api.NewError(
			fmt.Errorf("error while recording activity, status code: %d", res.StatusCode),
			res.StatusCode,
			nil,
		)
	}

	return nil
}
