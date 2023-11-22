package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/iamsumit/sample-go-app/pkg/config"
	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/spf13/viper"
)

var (
	configuration config.Configuration
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
}

func main() {
	fmt.Println(configuration.Http.Port)

	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
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

	log.Println("Database connected!")

	defer func() {
		_ = sqlDB.Close()
	}()

	// -------------------------------------------------------------------
	// Launch Darkly
	// -------------------------------------------------------------------

	// ldClient = launchdarkly.NewClient(configuration.LaunchDarkly.SecretKey)

	// -------------------------------------------------------------------
	// Routing
	// -------------------------------------------------------------------
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/reload-config", reloadConfigHandler)
	http.HandleFunc("/user/", userHandler)

	// -------------------------------------------------------------------
	// Server
	// -------------------------------------------------------------------
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	return nil
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path by '/'
	parts := strings.Split(r.URL.Path, "/")

	// Get the last part of the URL path
	uid := parts[len(parts)-1]
	id, _ := strconv.Atoi(uid)

	fmt.Fprintf(w, "User: %v", id)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// isEnabled := ldClient.ReadFlag(enablePrintFlagKey)

	enabled := "true"
	// if !isEnabled {
	// 	enabled = "false"
	// }

	fmt.Fprintf(w, "Pubsub Name: %s, %s; EnvVar: %s; Flag Enabled: %s", configuration.PubSub.Name, viper.Get("pubsub.name"), configuration.Environment.Env, enabled)
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
