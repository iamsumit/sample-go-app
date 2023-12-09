// Package config provides the configuration for the application.
//
// It also provides method to read and unmarshal the config in the config object.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration holds the configuration for the application.
type Configuration struct {
	LaunchDarkly LaunchDarklyConfig `mapstructure:"launch_darkly"`
	MySQL        MySQLConfig        `mapstructure:"mysql"`
	Jaeger       JaegerConfig       `mapstructure:"jaeger"`
	Env          string             `mapstructure:"env"`
	OtherEnv     string             `mapstructure:"other_env"`
}

// LaunchDarklyConfig holds the configuration for the launch darkly.
type LaunchDarklyConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

// MySQLConfig holds the configuration for the MySQL.
type MySQLConfig struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

// JaegerConfig holds the configuration for the Jaeger.
type JaegerConfig struct {
	Host string `mapstructure:"host"`
	Path string `mapstructure:"path"`
}

// ReadConfig reads the configuration from the config file and environment variables.
func ReadConfig() *Configuration {
	c := new(Configuration)

	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return c
}
