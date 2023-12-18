// Package router Sample API.
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
//	Host: localhost:8080
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
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/api/middleware"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/sample/internal/client/activity"
	userv1 "github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/v1"
	redoc "github.com/mvrilo/go-redoc"
	swgui "github.com/swaggest/swgui/v5cdn"
)

// Routes holds the additional routes information for the application.
type Routes struct {
	Path    string
	Handler api.Handler
}

// Config holds the configuration for the router.
type Config struct {
	Log      logger.Logger
	DB       *sql.DB
	Activity *activity.Client
}

// ConfigureRoutes configures the routes for the application.
func ConfigureRoutes(shutdown chan os.Signal, notracepath []string, routes []Routes, cfg Config, mw ...api.Middleware) http.Handler {
	// -------------------------------------------------------------------
	// Middlewares
	// -------------------------------------------------------------------
	mw = append(mw, middleware.Logger(cfg.Log))
	mw = append(mw, middleware.Errors(cfg.Log))

	// -------------------------------------------------------------------
	// API Handler
	// -------------------------------------------------------------------
	a := api.New(shutdown, notracepath, mw...)

	// Provides home endpoint.
	a.Handle(http.MethodGet, "/", func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		err := api.Respond(ctx, w, "Welcome to Sample Go App!", http.StatusOK)

		return err
	})

	// -------------------------------------------------------------------
	// Additional Routes
	// -------------------------------------------------------------------
	for _, r := range routes {
		a.Handle(http.MethodGet, r.Path, r.Handler)
	}

	// -------------------------------------------------------------------
	// V1 Routes
	// -------------------------------------------------------------------
	SetV1Routes(a, cfg)

	return a
}

// SetV1Routes returns the http handler for the v1 routes.
func SetV1Routes(a *api.API, cfg Config) {
	// -------------------------------------------------------------------
	// User Handler & Routes
	// -------------------------------------------------------------------
	userV1 := userv1.New(cfg.Log, cfg.DB, cfg.Activity)

	a.Handle(http.MethodGet, "/v1/user/{id}", userV1.ByID)
	a.Handle(http.MethodGet, "/v1/users", userV1.All)
	a.Handle(http.MethodPost, "/v1/user", userV1.CreateUser)

	ServeSWGUIDocsRoutes(a, "v1")
	ServeReDocsRoutes(a, "v1")
}

// ServeSWGUIDocsRoutes serves the swagger documentation routes usign swgui.
func ServeSWGUIDocsRoutes(a *api.API, version string) {
	// -------------------------------------------------------------------
	// Swagger JSON
	// -------------------------------------------------------------------
	vSwagPath := fmt.Sprintf("/%s/swagger.json", version)
	a.Handle(http.MethodGet, vSwagPath, func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// Construct the full path to the JSON file
		swagJSONPath := filepath.Join("./docs", "swagger", version, "/swagger.json")
		http.ServeFile(
			w,
			r,
			swagJSONPath,
		)

		return nil
	})

	// -------------------------------------------------------------------
	// Swagger UI
	// -------------------------------------------------------------------
	docPath := fmt.Sprintf("/%s/swgui-doc", version)
	a.Handle(http.MethodGet, docPath, func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		swgui.New(
			"Sample API",
			vSwagPath,
			docPath,
		).ServeHTTP(w, r)

		return nil
	})
}

// ServeReDocsRoutes serves the swagger documentation using redoc library.
func ServeReDocsRoutes(a *api.API, version string) {
	swagJSONPath := filepath.Join("./docs", "swagger", version, "/swagger.json")
	doc := redoc.Redoc{
		Title:       "Sample API",
		Description: "Sample V1 API",
		SpecFile:    swagJSONPath,
		SpecPath:    fmt.Sprintf("/%s/swagger.json", version),
		DocsPath:    fmt.Sprintf("/%s/redoc", version),
	}

	a.Handle(http.MethodGet, doc.DocsPath, func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		doc.Handler().ServeHTTP(w, r)
		return nil
	})
}
