package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	LaunchDarkly LaunchDarklyConfig `mapstructure:"launch_darkly"`
	MySQL        MySQLConfig
	Environment  EnvironmentConfig
	Jaeger       JaegerConfig
}

type LaunchDarklyConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

type MySQLConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

type EnvironmentConfig struct {
	Env      string
	OtherEnv string
	EmptyEnv string
}

type JaegerConfig struct {
	Host string
	Path string
}

func ReadConfig(configuration *Configuration) {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Enable VIPER to read Environment Variables
	viper.BindEnv("env", "ENV")
	viper.BindEnv("other_env", "OTHER_ENV")

	// Read environment variables
	configuration.Environment = EnvironmentConfig{
		Env:      viper.GetString("env"),
		OtherEnv: viper.GetString("other_env"),
	}
}
