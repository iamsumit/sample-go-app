package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/iamsumit/sample-go-app/pkg/metrics"
	"github.com/iamsumit/sample-go-app/pkg/metrics/common"
	"github.com/iamsumit/sample-go-app/sample/handler/config"
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
	log, err := logger.New(&logger.Config{
		LoggerType: logger.SLog,
		LogFormat:  logger.JSON,
	})
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------
	// Database
	// -------------------------------------------------------------------
	sqlDB, err := db.Handler(&db.Config{
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
	// Routing
	// -------------------------------------------------------------------
	http.HandleFunc("/", helloWorldHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/reload-config", reloadConfigHandler)
	http.HandleFunc("/user/", userHandler)

	// -------------------------------------------------------------------
	// Server
	// -------------------------------------------------------------------
	log.Info(
		"Server is gettig started!",
		"host", "localhost",
		"port", configuration.Http.Port,
	)

	// -------------------------------------------------------------------
	// Start the server
	// -------------------------------------------------------------------

	addr := fmt.Sprintf(":%d", configuration.Http.Port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}

	return nil
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	RequestCounter.Record(r.Context(), 1, r.URL.Path, r.Method)
	startTime := time.Now().UnixMilli()
	// Split the URL path by '/'
	parts := strings.Split(r.URL.Path, "/")

	// Get the last part of the URL path
	uid := parts[len(parts)-1]
	id, _ := strconv.Atoi(uid)

	fmt.Fprintf(w, "User: %v", id)
	LatencyCounter.Record(r.Context(), float64(time.Now().UnixMilli()-startTime), r.URL.Path, r.Method)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	RequestCounter.Record(r.Context(), 1, r.URL.Path, r.Method)
	startTime := time.Now().UnixMilli()

	// isEnabled := ldClient.ReadFlag(enablePrintFlagKey)
	enabled := "true"
	// if !isEnabled {
	// 	enabled = "false"
	// }

	fmt.Fprintf(w, "Pubsub Name: %s, %s; EnvVar: %s; Flag Enabled: %s", configuration.PubSub.Name, viper.Get("pubsub.name"), configuration.Environment.Env, enabled)
	time.Sleep(1)
	LatencyCounter.Record(r.Context(), float64(time.Now().UnixMilli()-startTime), r.URL.Path, r.Method)
}

func reloadConfigHandler(w http.ResponseWriter, r *http.Request) {
	ReadConfig(&configuration)

	fmt.Fprintf(w, "Config Reloaded. Pubsub Name: %s, %s", configuration.PubSub.Name, viper.Get("pubsub.name"))
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
