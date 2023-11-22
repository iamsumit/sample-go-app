package cmd

import (
	"context"
	"fmt"
	"os"

	pbsb "github.com/iamsumit/sample-go-app/pkg/pubsub"
	"github.com/spf13/cobra"
)

var (
	ctx          context.Context
	pubsubClient *pbsb.PubSub
	projectID    = "sample-go-app"
	topic        = "sample-topic"
	subscription = "sample-subscription"
)

func init() {
	ctx = context.Background()
	pubsubClient, _ = pbsb.New(ctx, projectID, topic, subscription)

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
