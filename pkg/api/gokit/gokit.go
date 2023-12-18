package gokit

import (
	"os"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/logger"
)

type Handler struct {
	api *api.API
	log logger.Logger
}

// New creates an API struct with provided middleware.
//
//	shutdown: channel to signal shutdown
//	log: logger to use
//	notracepath: list of paths to not trace, example: /metrics
//	mw: list of middleware to execute on each request
func New(shutdown chan os.Signal, log logger.Logger, notracepath []string, mw ...api.Middleware) *Handler {
	a := api.New(shutdown, notracepath, mw...)

	return &Handler{
		api: a,
		log: log,
	}
}

// RoutingHandler returns the handler that implements the http.Handler interface.
func (h *Handler) RoutingHandler() *api.API {
	return h.api
}
