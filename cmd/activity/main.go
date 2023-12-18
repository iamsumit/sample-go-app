// Package main provides the main entry point for the application.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/iamsumit/sample-go-app/activity/internal/config"
	"github.com/iamsumit/sample-go-app/activity/internal/handler/router"
	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	smetrics "github.com/iamsumit/sample-go-app/pkg/metrics"
	"github.com/iamsumit/sample-go-app/pkg/observation/metrics"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"github.com/iamsumit/sample-go-app/pkg/util/app"
	"github.com/spf13/viper"
)

func main() {
	// -------------------------------------------------------------------
	// Logger
	// -------------------------------------------------------------------
	log, err := logger.New(logger.WithSlogger(), logger.WithJSONFormat())
	if err != nil {
		fmt.Println("Error while creating logger", err)

		return
	}

	// -------------------------------------------------------------------
	// Server
	// -------------------------------------------------------------------
	if err := start(log); err != nil {
		log.Error("Error while starting the server", "error", err)

		return
	}
}

//nolint:funlen
func start(log logger.Logger) error {
	// -------------------------------------------------------------------
	// Configurations
	// -------------------------------------------------------------------
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// Read through environment variable.
	viper.SetEnvPrefix("ACTIVITY")

	// Read the configuration on load.
	configuration, err := config.Read()
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Metrics
	// -------------------------------------------------------------------
	mProvider, err := smetrics.New(&smetrics.Config{
		Name:     app.Name(),
		Type:     smetrics.Otel,
		Exporter: smetrics.Prometheus,
	})
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Database
	// -------------------------------------------------------------------
	sqlDB, err := db.New(db.Config{
		Type:     db.Postgres,
		Name:     configuration.DB.Name,
		User:     configuration.DB.User,
		Password: configuration.DB.Password,
		Port:     configuration.DB.Port,
		Host:     configuration.DB.Host,
	})
	if err != nil {
		return err
	}

	defer func() {
		_ = sqlDB.Close()
	}()

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Database connection information
	// -------------------------------------------------------------------
	log.Info(
		"Database connected!",
		"database", configuration.DB.Name,
		"host", configuration.DB.Host,
	)

	// -------------------------------------------------------------------
	// Tracer
	// -------------------------------------------------------------------

	// Once initiated, it will set the global tracer provider.
	//
	// tracer.Global("activity") can be used to get the global tracer instance.
	_, err = tracer.New(context.Background(), &tracer.Config{
		Name:        app.Name(),
		ServiceName: app.TraceName,
		Jaeger: tracer.JaegerConfig{
			Host: configuration.Jaeger.Host,
			Path: configuration.Jaeger.Path,
		},
	})
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Observation:- Metrics
	// -------------------------------------------------------------------

	// Exclude paths from the metrics and tracing.
	exPaths := []string{
		"/metrics",
	}

	mInt, err := metrics.New(
		app.Name(),
		metrics.WithMetricsProvider(mProvider),
		metrics.WithNoMetricsPath(exPaths),
	)
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Routing
	// -------------------------------------------------------------------

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	mw := make([]api.Middleware, 0)
	mw = append(mw, mInt.Middleware(log))

	routes := []router.Routes{
		{
			Path:    "/metrics",
			Handler: mInt.Handler(),
		},
	}

	handler := router.ConfigureRoutes(
		shutdown,
		exPaths,
		routes,
		router.Config{
			Log: log,
			DB:  sqlDB,
		},
		mw...,
	)

	// -------------------------------------------------------------------
	// Server
	// -------------------------------------------------------------------

	server := &http.Server{
		Addr:    ":8080", // Set your desired port
		Handler: handler,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info("startup.api", "status", "api router started", "host", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// -------------------------------------------------------------------
	// Signal handling
	// -------------------------------------------------------------------
	select {
	case err := <-serverErrors:
		fmt.Printf("server error: %v", err)
	case sig := <-shutdown:
		log.Info("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info("shutdown", "status", "shutdown completed", "signal", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("web.shutdownTimeout"))
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := server.Shutdown(ctx); err != nil {
			log.Error(err.Error())

			_ = server.Close()
		}
	}

	return nil
}
