// Package router provides the router for the application.
//
// It can be used to add all the possible routes for the application.
package router

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/api/middleware"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	pUserV1 "github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/v1"
)

// Config holds the configuration for the router.
type Config struct {
	Log logger.Logger
	DB  *sql.DB
}

// ConfigureRoutes configures the routes for the application.
func ConfigureRoutes(shutdown chan os.Signal, mHandler api.Handler, cfg Config, mw ...api.Middleware) http.Handler {
	// -------------------------------------------------------------------
	// Middlewares
	// -------------------------------------------------------------------
	mw = append(mw, middleware.Logger(cfg.Log))
	mw = append(mw, middleware.Errors(cfg.Log))

	// -------------------------------------------------------------------
	// API Handler
	// -------------------------------------------------------------------
	a := api.New(shutdown, mw...)

	// Provides home endpoint.
	a.Handle(http.MethodGet, "/", func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		err := api.Respond(ctx, w, "Welcome to Sample Go App!", http.StatusOK)

		return err
	})

	// Provides metrics endpoint.
	a.Handle(http.MethodGet, "/metrics", mHandler)

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
	userV1 := pUserV1.New(cfg.Log, cfg.DB)

	a.Handle(http.MethodGet, "/v1/user/{id}", userV1.ByID)
	a.Handle(http.MethodPost, "/v1/user", userV1.CreateUser)
}
