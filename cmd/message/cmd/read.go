package cmd

import (
	api "github.com/iamsumit/sample-go-app/message/api/message"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var (
	readCmd = &cobra.Command{
		Use:   "read",
		Short: "Read the messages from topic",
		Run:   ReadMessage,
	}
)

func init() {
	rootCmd.AddCommand(readCmd)
}

func ReadMessage(cmd *cobra.Command, args []string) {
	messageChannel := make(chan []byte)
	message := api.Message{}

	go func() {
		subscription.ReceiveMessages(messageChannel)
	}()

	for {
		select {
		case data := <-messageChannel:
			proto.Unmarshal(data, &message)
			log.Info(
				"Received message",
				"message", message.Message.String(),
				"vcs", message.Vcs.Type.String(),
				"request", message.Vcs.Request.GetTitle(),
			)
		}
	}
}
