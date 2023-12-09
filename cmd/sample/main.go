package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/metrics"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"github.com/iamsumit/sample-go-app/sample/internal/config"
	"github.com/iamsumit/sample-go-app/sample/internal/handler/router"
	pMetricsInt "github.com/iamsumit/sample-go-app/sample/internal/observation/metrics"
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
	}
}

func start(log logger.Logger) error {
	// -------------------------------------------------------------------
	// Configurations
	// -------------------------------------------------------------------

	configuration := new(config.Configuration)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// Read through environment variable.
	viper.SetEnvPrefix("SAMPLE")
	viper.AutomaticEnv()

	// Read the configuration on load.
	config.ReadConfig(configuration)

	// -------------------------------------------------------------------
	// Metrics
	// -------------------------------------------------------------------
	mProvider, err := metrics.New(&metrics.Config{
		Name:     "sample",
		Type:     metrics.Otel,
		Exporter: metrics.Prometheus,
	})
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Tracer
	// -------------------------------------------------------------------
	_, err = tracer.New(context.Background(), &tracer.Config{
		Name:        "sample",
		ServiceName: "sample-go-app",
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
	mInt, err := pMetricsInt.New("sample", pMetricsInt.WithMetricsProvider(mProvider))
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Database
	// -------------------------------------------------------------------
	sqlDB, err := db.New(db.Config{
		Type:     db.MySQL,
		Name:     configuration.MySQL.Name,
		User:     configuration.MySQL.User,
		Password: configuration.MySQL.Password,
		Port:     configuration.MySQL.Port,
		Host:     configuration.MySQL.Host,
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
		"database", configuration.MySQL.Name,
		"host", configuration.MySQL.Host,
	)

	// -------------------------------------------------------------------
	// Routing
	// -------------------------------------------------------------------

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	mw := make([]api.Middleware, 0)
	mw = append(mw, mInt.Middleware(log))

	handler := router.ConfigureRoutes(shutdown, mInt.Handler(), router.Config{
		Log: log,
		DB:  sqlDB,
	}, mw...)

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
			_ = server.Close()
			log.Error(err.Error())
		}
	}

	return nil
}
