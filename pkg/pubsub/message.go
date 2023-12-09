package pubsub

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

// Publish publishes a message to the pubsub topic.
func (t *Topic) Publish(ctx context.Context, msg *Message) ([]string, error) {
	m := &pubsub.Message{
		Attributes: msg.Attributes,
		Data:       msg.Data,
	}

	defer t.client.Stop()
	result := t.client.Publish(ctx, m)
	msgIDs, err := result.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("publish failed, %w", err)
	}

	return []string{msgIDs}, nil
}

// ReceiveMessages receives and processes messages from the subscription.
func (s *Subscription) ReceiveMessages(messageChannel chan<- []byte) {
	ctx := context.Background()

	err := s.client.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", msg.ID)
		msg.Ack()

		messageChannel <- msg.Data
	})
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
}
