package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/golang/protobuf/proto"
	pbsb "github.com/iamsumit/sample-go-app/pkg/pubsub"
	pbsbMsg "github.com/iamsumit/sample-go-app/pkg/pubsub/definitions/message"
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
	message := pbsbMsg.Message{
		Message: pbsbMsg.MessageType_VCS,
		Vcs: &pbsbMsg.VCSMessage{
			Type: pbsbMsg.VCSType_GITHUB,
			Request: &pbsbMsg.PullRequest{
				Number: 1,
				Title:  "Sample Pull Request",
				Details: &pbsbMsg.PullRequest_Details{
					Type:         pbsbMsg.PullRequest_OPENED,
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
