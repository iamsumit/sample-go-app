package api

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"github.com/iamsumit/sample-go-app/pkg/util/strings"
	"go.opentelemetry.io/otel/attribute"

	"go.opentelemetry.io/otel/trace"
)

// A Handler is a type that handles a http request within the framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// API is the handler for api package.
type API struct {
	shutdown chan os.Signal
	mux      *mux.Router
	mw       []Middleware
	ntp      []string
}

// New creates an API struct with provided middleware.
//
//	shutdown: channel to signal shutdown
//	notracepath: list of paths to not trace, example: /metrics
//	mw: list of middleware to execute on each request
func New(shutdown chan os.Signal, notracepath []string, mw ...Middleware) *API {
	return &API{
		shutdown: shutdown,
		mux:      mux.NewRouter(),
		mw:       mw,
		ntp:      notracepath,
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
		var span trace.Span
		spanCtx := r.Context()

		// Check if the path is in the notracepath list.
		isNoTrace := strings.Contain(a.ntp, r.URL.Path)
		if !isNoTrace {
			// Start the tracing here.
			spanCtx, span = tracer.Global("api").Start(spanCtx, r.URL.Path)
			defer span.End()
		}

		// Set the context with the required values to
		// process the request.
		v := ContextValues{}
		ctx := context.WithValue(spanCtx, key, &v)

		if err := handler(ctx, w, r); err != nil {
			a.SignalShutdown()
		}

		if isNoTrace {
			return
		}

		span.SetAttributes(
			attribute.String("method", r.Method),
			attribute.String("path", r.URL.Path),
			attribute.Int("status_code", v.StatusCode),
		)

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
