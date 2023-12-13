package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	os.Setenv("SAMPLE_ENV", "someEnv")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	// Read through environment variable.
	viper.SetEnvPrefix("SAMPLE")

	t.Run("ReadConfig", func(t *testing.T) {
		// Read the configuration on load.
		c, err := Read()
		if err != nil {
			t.Errorf("failed to read config: %v", err)
		}

		if c.MySQL.Name != "sample_db" {
			t.Errorf("expected sample_db, got %s", c.MySQL.Name)
		}

		if c.Env != "someEnv" {
			t.Errorf("expected someEnv, got %s", c.Env)
		}
	})
}
