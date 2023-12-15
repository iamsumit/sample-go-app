package test

import (
	"testing"

	"github.com/iamsumit/sample-go-app/sample/internal/config"
	"github.com/spf13/viper"
)

// New returns a new config for testing.
func New(t *testing.T, configDirPath string) (*config.Config, error) {
	viper.SetConfigName("config-test")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDirPath)

	// Read through environment variable.
	viper.SetEnvPrefix("SAMPLE")

	t.Log("Reading config...")

	c, err := config.Read()
	if err != nil {
		t.Error(err)
		return nil, err
	}

	t.Log("Read config successfully!")
	return c, nil
}
