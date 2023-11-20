package cmd

import (
	"fmt"

	pbsbMsg "github.com/iamsumit/sample-go-app/pkg/pubsub/definitions/message"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read the messages from topic",
	Run:   ReadMessage,
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func ReadMessage(cmd *cobra.Command, args []string) {
	messageChannel := make(chan []byte)
	message := pbsbMsg.Message{}

	go func() {
		pubsubClient.SubscriptionConfig.ReceiveMessages(messageChannel)
	}()

	for {
		select {
		case data := <-messageChannel:
			proto.Unmarshal(data, &message)
			fmt.Printf("Received message: %v\n", message)
		}
	}
}
