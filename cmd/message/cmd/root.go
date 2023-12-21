package cmd

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/iamsumit/sample-go-app/message/internal/config"
	pbsb "github.com/iamsumit/sample-go-app/pkg/pubsub"
	"github.com/iamsumit/sample-go-app/pkg/slogger"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	ctx           context.Context
	err           error
	configuration config.Configuration
	pbsbClient    *pbsb.Handler
	topic         *pbsb.Topic
	subscription  *pbsb.Subscription
	log           *slog.Logger
)

func init() {

	//--------------------------------------------------------------------
	// Logger
	//--------------------------------------------------------------------
	log = slogger.WithFormat(os.Stdout, "tint")

	// -------------------------------------------------------------------
	// Configurations
	// -------------------------------------------------------------------

	configFilePath := "./config/config.yml"

	// Read the configuration on load.
	err = ReadConfig(configFilePath, &configuration)
	if err != nil {
		log.Error("Error reading config file", "ERROR", err)
		panic(err)
	}

	log.Info("configuration loaded", "topic", configuration.PubSub.Topic)

	ctx = context.Background()
	pbsbClient, _ = pbsb.New(ctx, configuration.PubSub.Project)

	topic, err = pbsbClient.CreateTopic(ctx, configuration.PubSub.Topic)
	if err != nil {
		log.Error("Error creating topic", "ERROR", err)
		panic(err)
	}

	subscription, err = topic.CreateSubscription(ctx, configuration.PubSub.Subscription)
	if err != nil {
		log.Error("Error creating subscription", "ERROR", err)
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "message",
	Short: "Message command used to publish and subscribe to the pub/sub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("working fine with args ", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error("Error executing command", "ERROR", err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ReadConfig(cfp string, configuration *config.Configuration) error {
	cfgFile, err := os.Open(cfp)
	if err != nil {
		log.Error("Error reading config file", "ERROR", err)
		return err
	}

	defer cfgFile.Close()

	// Read the content of the config file
	content, err := io.ReadAll(cfgFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &configuration)
	if err != nil {
		log.Error("Error unmarshalling config file", "ERROR", err)
		return err
	}

	return nil
}
