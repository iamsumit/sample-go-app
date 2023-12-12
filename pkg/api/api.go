package api

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/gorilla/mux"
)

// A Handler is a type that handles a http request within the framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type API struct {
	shutdown chan os.Signal
	mux      *mux.Router
	mw       []Middleware
}

func New(shutdown chan os.Signal, mw ...Middleware) *API {
	return &API{
		shutdown: shutdown,
		mux:      mux.NewRouter(),
		mw:       mw,
	}
}

// SignalShutdown is used to gracefully shut down the app when an integrity
// issue is identified.
func (a *API) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *API) Handle(method string, path string, handler Handler, mw ...Middleware) {
	// First wrap handler specific middleware around this handler
	handler = wrapMiddleware(mw, handler)

	// Add the api's general middleware to the handler chain.
	handler = wrapMiddleware(a.mw, handler)

	// Execute each specific request
	// The function to execute for each request.
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the context with the required values to
		// process the request.
		v := ContextValues{}
		ctx := context.WithValue(r.Context(), key, &v)

		if err := handler(ctx, w, r); err != nil {
			a.SignalShutdown()
		}
		return
	})

	a.mux.Handle(path, h).Methods(method)
}

// ServeHTTP implements the http.Handler interface. It's the entry point for
// all http traffic and allows the opentelemetry mux to run first to handle
// tracing. The opentelemetry mux then calls the application mux to handle
// application traffic.  See NewApi function above for implementation.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
