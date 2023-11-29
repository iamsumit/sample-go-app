package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/metrics"
	"github.com/iamsumit/sample-go-app/pkg/metrics/common"
	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"github.com/iamsumit/sample-go-app/sample/internal/config"
	"github.com/iamsumit/sample-go-app/sample/internal/handler/user"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

var (
	configuration  config.Configuration
	RequestCounter common.Counter
	LatencyCounter common.Counter
)

const (
	enablePrintFlagKey = "enable-print-flag"
)

func init() {
	// -------------------------------------------------------------------
	// Configurations
	// -------------------------------------------------------------------

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	// Watch the config file for changes.
	viper.WatchConfig()

	// Read the configuration on load.
	ReadConfig(&configuration)

	// Read the configuration on every time config changes.
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		// ReadConfig(&configuration)
	})

	// -------------------------------------------------------------------
	// Metrics
	// -------------------------------------------------------------------
	otelMetrics, err := metrics.New(&metrics.Config{
		Name:     "sample",
		Type:     metrics.Otel,
		Exporter: metrics.Prometheus,
	})
	if err != nil {
		panic(err)
	}

	RequestCounter, err = otelMetrics.NewCounter("sample_request_count", "Number of requests", "path", "method")
	if err != nil {
		panic(err)
	}

	LatencyCounter, err = otelMetrics.NewCounter("sample_latency", "Latency of each request", "path", "method")
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	// -------------------------------------------------------------------
	// Logger
	// -------------------------------------------------------------------
	log, err := logger.New(logger.WithSlogger(), logger.WithJSONFormat())
	if err != nil {
		panic(err)
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

	// -------------------------------------------------------------------
	// Database connection information
	// -------------------------------------------------------------------
	log.Info(
		"Database connected!",
		"database", configuration.MySQL.Name,
		"host", configuration.MySQL.Host,
	)

	// -------------------------------------------------------------------
	// Launch Darkly
	// -------------------------------------------------------------------

	// ldClient = launchdarkly.NewClient(configuration.LaunchDarkly.SecretKey)

	// -------------------------------------------------------------------
	// Tracer
	// -------------------------------------------------------------------
	otelTracer, err := tracer.New(context.Background(), &tracer.Config{
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
	// Routing Group Handler
	// -------------------------------------------------------------------
	userHandler, err := user.New(context.Background(), log, otelTracer)
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Routing
	// -------------------------------------------------------------------
	http.HandleFunc("/", helloWorldHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/reload-config", reloadConfigHandler)
	http.HandleFunc("/user/", userHandler.GetUser)

	// -------------------------------------------------------------------
	// Server
	// -------------------------------------------------------------------
	log.Info(
		"Server is gettig started!",
		"host", "localhost",
		"port", "8080",
	)

	// -------------------------------------------------------------------
	// Start the server
	// -------------------------------------------------------------------

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	RequestCounter.Record(r.Context(), 1, r.URL.Path, r.Method)
	startTime := time.Now().UnixMilli()

	// isEnabled := ldClient.ReadFlag(enablePrintFlagKey)
	enabled := "true"
	// if !isEnabled {
	// 	enabled = "false"
	// }

	fmt.Fprintf(w, "EnvVar: %s; Flag Enabled: %s", configuration.Environment.Env, enabled)
	time.Sleep(1)
	LatencyCounter.Record(r.Context(), float64(time.Now().UnixMilli()-startTime), r.URL.Path, r.Method)
}

func reloadConfigHandler(w http.ResponseWriter, r *http.Request) {
	ReadConfig(&configuration)

	fmt.Fprintf(w, "Config Reloaded.")
}

func ReadConfig(configuration *config.Configuration) {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Enable VIPER to read Environment Variables
	viper.BindEnv("env", "TEST_ENV")
	viper.BindEnv("other_env", "TEST_OTHER_ENV")

	// Read environment variables
	configuration.Environment = config.EnvironmentConfig{
		Env:      viper.GetString("env"),
		OtherEnv: viper.GetString("other_env"),
	}
}
