// Package config provides the configuration for the application.
//
// It also provides method to read and unmarshal the config in the config object.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the configuration for the application.
type Config struct {
	// LaunchDarkly holds the configuration for the launch darkly.
	LaunchDarkly LaunchDarklyConfig `mapstructure:"launch_darkly"`

	// DB holds the configuration for the MySQL.
	DB DBConfig `mapstructure:"db"`

	// Jaeger holds the configuration for the Jaeger.
	Jaeger JaegerConfig `mapstructure:"jaeger"`
}

// LaunchDarklyConfig holds the configuration for the launch darkly.
type LaunchDarklyConfig struct {
	// SecretKey holds the secret key for the launch darkly.
	SecretKey string `mapstructure:"secret_key"`
}

// MySQLConfig holds the configuration for the MySQL.
type DBConfig struct {
	// Name holds the name of the database.
	Name string `mapstructure:"name"`

	// User holds the user for the database.
	User string `mapstructure:"user"`

	// Password holds the password for the database.
	Password string `mapstructure:"password"`

	// Host holds the host for the database.
	Host string `mapstructure:"host"`

	// Port holds the port for the database.
	Port int `mapstructure:"port"`
}

// JaegerConfig holds the configuration for the Jaeger.
type JaegerConfig struct {
	// Host holds the host for the Jaeger.
	Host string `mapstructure:"host"`

	// Path holds the path for the Jaeger.
	Path string `mapstructure:"path"`
}

// Read reads the configuration from the config file and environment variables.
func Read() (*Config, error) {
	c := new(Config)

	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return c, nil
}
