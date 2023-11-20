package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/iamsumit/sample-go-app/pkg/config"
	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/iamsumit/sample-go-app/pkg/launchdarkly"
	"github.com/spf13/viper"
)

var (
	configuration  config.Configuration
	ldClient       *launchdarkly.LaunchDarklyClient
	databaseClient *db.DataBaseClient
)

const (
	enablePrintFlagKey = "enable-print-flag"
)

func init() {
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

	// ldClient = launchdarkly.NewClient(configuration.LaunchDarkly.SecretKey)

	databaseClient = db.NewClient()
}

func main() {
	fmt.Println(configuration.Http.Port)

	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/reload-config", reloadConfigHandler)
	http.HandleFunc("/user/", userHandler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path by '/'
	parts := strings.Split(r.URL.Path, "/")

	// Get the last part of the URL path
	uid := parts[len(parts)-1]
	id, _ := strconv.Atoi(uid)

	user := databaseClient.GetConfigByUID(id)
	fmt.Fprintf(w, "User: %v", user)
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
