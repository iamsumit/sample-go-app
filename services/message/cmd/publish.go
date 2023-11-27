package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/golang/protobuf/proto"
	api "github.com/iamsumit/sample-go-app/message/api/message"
	pbsb "github.com/iamsumit/sample-go-app/pkg/pubsub"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish the message",
	Run:   PublishMessage,
}

func init() {
	rootCmd.AddCommand(publishCmd)
}

func PublishMessage(cmd *cobra.Command, args []string) {
	// A sample message for VCS type.
	message := api.Message{
		Message: api.MessageType_VCS,
		Vcs: &api.VCSMessage{
			Type: api.VCSType_GITHUB,
			Request: &api.PullRequest{
				Number: 1,
				Title:  "Sample Pull Request",
				Details: &api.PullRequest_Details{
					Type:         api.PullRequest_OPENED,
					SourceBranch: "source-branch",
					DestBranch:   "dest-branch",
				},
			},
		},
	}

	messageBytes, _ := proto.Marshal(&message)
	msgs, _ := pubsubClient.Publish(ctx, &pbsb.Message{
		Data: messageBytes,
	})

	fmt.Printf("Message Published: %v", msgs)
}
