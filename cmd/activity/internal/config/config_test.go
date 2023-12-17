package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	// Read through environment variable.
	viper.SetEnvPrefix("ACTIVITY")

	t.Run("ReadConfig", func(t *testing.T) {
		// Read the configuration on load.
		c, err := Read()
		if err != nil {
			t.Errorf("failed to read config: %v", err)
		}

		if c.DB.Name != "activity_db" {
			t.Errorf("expected activity_db, got %s", c.DB.Name)
		}
	})
}
