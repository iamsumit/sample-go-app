package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/iamsumit/sample-go-app/message/internal/config"
	pbsb "github.com/iamsumit/sample-go-app/pkg/pubsub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ctx           context.Context
	pubsubClient  *pbsb.PubSub
	configuration config.Configuration
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

	ctx = context.Background()
	pubsubClient, _ = pbsb.New(ctx, configuration.PubSub.Project, configuration.PubSub.Topic, configuration.PubSub.Subscription)

	err := pubsubClient.CreateTopicAndSubscription(ctx)
	if err != nil {
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
}