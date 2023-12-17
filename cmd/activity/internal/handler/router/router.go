// Package router Activity API.
//
// the purpose of this application is to provide basic routes
// to play with.
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//	Schemes: http, https
//	Host: localhost:8081
//	BasePath: /
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package router

import (
	"database/sql"
	"net/http"
	"os"

	activityv1 "github.com/iamsumit/sample-go-app/activity/internal/handler/activity/v1"
	"github.com/iamsumit/sample-go-app/activity/internal/handler/router/welcome"
	activitydb "github.com/iamsumit/sample-go-app/activity/internal/repository/activity/db"
	welcomedb "github.com/iamsumit/sample-go-app/activity/internal/repository/welcome/db"
	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/api/apigokit"
	"github.com/iamsumit/sample-go-app/pkg/api/middleware"
	"github.com/iamsumit/sample-go-app/pkg/logger"
)

// Routes holds the additional routes information for the application.
type Routes struct {
	Path    string
	Handler api.Handler
}

// Config holds the configuration for the router.
type Config struct {
	Log logger.Logger
	DB  *sql.DB
}

// ConfigureRoutes configures the routes for the application.
func ConfigureRoutes(shutdown chan os.Signal, notracepath []string, routes []Routes, cfg Config, mw ...api.Middleware) http.Handler {
	// -------------------------------------------------------------------
	// Middlewares
	// -------------------------------------------------------------------
	mw = append(mw, middleware.Logger(cfg.Log))

	// -------------------------------------------------------------------
	// API Handler
	// -------------------------------------------------------------------
	a := apigokit.New(shutdown, cfg.Log, notracepath, mw...)

	// -------------------------------------------------------------------
	// Routes
	// -------------------------------------------------------------------

	// Configure the welcome route.
	ConfigureWelcomeRoute(a, cfg)

	// Configure the documents routes.
	ConfigureDocsRoutes(a, cfg)

	// Configure the activity routes.
	ConfigureActivityRoutes(a, cfg)

	// -------------------------------------------------------------------
	// Additional Routes
	// -------------------------------------------------------------------
	for _, r := range routes {
		a.RoutingHandler().Handle(http.MethodGet, r.Path, r.Handler)
	}

	return a.RoutingHandler()
}

// ConfigureWelcomeRoute configures the welcome route.
func ConfigureWelcomeRoute(a *apigokit.Handler, cfg Config) {
	// -------------------------------------------------------------------
	// Service
	// -------------------------------------------------------------------

	// Creates a new router repository.
	wrepo := welcomedb.New()

	// Creates a new router service.
	wservice := welcome.NewService(wrepo)

	// -------------------------------------------------------------------
	// Endpoints
	// -------------------------------------------------------------------
	we := welcome.MakeEndpoints(wservice)

	// -------------------------------------------------------------------
	// Routes
	// -------------------------------------------------------------------
	a.Handle(http.MethodGet, "/", we.Welcome, welcome.Decode)
}

// ConfigureActivityRoute configures the activity routes.
func ConfigureActivityRoutes(a *apigokit.Handler, cfg Config) {
	// -------------------------------------------------------------------
	// Activity Handler & Routes
	// -------------------------------------------------------------------

	// Creates a new activity repository.
	aRepo := activitydb.New(cfg.DB)

	// Creates a new activity service.
	aService := activityv1.NewService(aRepo)

	//-------------------------------------------------------------------
	// Endpoints
	// -------------------------------------------------------------------

	// Creates a new activity endpoint.
	ae := activityv1.MakeEndpoints(aService)

	// Creates a new activity handler.
	a.Handle(
		http.MethodPost,
		"/v1/activity",
		ae.Create,
		activityv1.DecodeNewActivity,
	)
}
