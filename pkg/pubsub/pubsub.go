package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

// PubSub is the GCLOUD PubSub implementation of the pubsub.Service Interface
type PubSub struct {
	Client             *pubsub.Client
	PubSubConfig       PubSubConfig
	SubscriptionConfig SubscriptionConfig
}

// PubSubConfig for pubsub details
type PubSubConfig struct {
	ProjectID string
	TopicName string
}

type SubscriptionConfig struct {
	SubscriptionName   string
	Subscription       *pubsub.Subscription
	CancelSubscription context.CancelFunc
}

// Message is the message that we will publish to pubsub
type Message struct {
	// Attributes are the pubsub message attributes.
	Attributes map[string]string

	// Data is the message of the pubsub message
	Data json.RawMessage
}

// NewPubSub creates a pubsub using GCE PubSub backend
func NewPubSub(ctx context.Context, projectID string, topicName string, subscriptionName string) (*PubSub, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return &PubSub{}, err
	}

	ps := PubSub{
		Client: client,
		PubSubConfig: PubSubConfig{
			ProjectID: projectID,
			TopicName: topicName,
		},
		SubscriptionConfig: SubscriptionConfig{
			SubscriptionName:   subscriptionName,
			CancelSubscription: nil,
		},
	}

	return &ps, nil
}

func (c *PubSub) CreateTopicAndSubscription(ctx context.Context) error {
	// Create a topic
	topic := c.Client.Topic(c.PubSubConfig.TopicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return fmt.Errorf("Error checking if topic exists: %v", err)
	}

	if !exists {
		_, err := c.Client.CreateTopic(ctx, c.PubSubConfig.TopicName)
		if err != nil {
			return fmt.Errorf("Error creating topic: %v", err)
		}

		fmt.Printf("Topic %s created\n", c.PubSubConfig.TopicName)
	} else {
		fmt.Printf("Topic %s already exists\n", c.PubSubConfig.TopicName)
	}

	// Create the subscription if it doesn't exist yet
	subscription := c.Client.Subscription(c.SubscriptionConfig.SubscriptionName)
	found, err := subscription.Exists(ctx)
	if err != nil {
		return fmt.Errorf("unable to check if subscription [%s] exists: %v", c.SubscriptionConfig.SubscriptionName, err)
	}

	if !found {
		subscription, _ = c.Client.CreateSubscription(ctx, c.SubscriptionConfig.SubscriptionName, pubsub.SubscriptionConfig{
			Topic: topic,
		})
	}

	c.SubscriptionConfig.Subscription = subscription

	return nil
}

func (c *PubSub) Publish(ctx context.Context, msg *Message) ([]string, error) {
	m := &pubsub.Message{
		Attributes: msg.Attributes,
		Data:       msg.Data,
	}

	topic := c.Client.Topic(c.PubSubConfig.TopicName)
	defer topic.Stop()
	result := topic.Publish(ctx, m)
	msgIDs, err := result.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("publish failed, %w", err)
	}

	return []string{msgIDs}, nil
}

// ReceiveMessages receives and processes messages from the subscription.
func (s *SubscriptionConfig) ReceiveMessages(messageChannel chan<- []byte) {
	ctx := context.Background()

	err := s.Subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", msg.ID)
		msg.Ack()

		messageChannel <- msg.Data
	})

	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
}

// Close closes the subscription.
func (s *SubscriptionConfig) Close() {
	if s.CancelSubscription != nil {
		s.CancelSubscription()
	}
}
