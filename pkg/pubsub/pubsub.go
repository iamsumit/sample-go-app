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
	TopicConfig        TopicConfig
	SubscriptionConfig SubscriptionConfig
}

// TopicConfig for pubsub topic details
type TopicConfig struct {
	ProjectID string
	TopicName string
}

// SubscriptionConfig for pubsub subscription details
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

// New creates a pubsub using GCE PubSub backend
func New(ctx context.Context, projectID string, topicName string, subscriptionName string) (*PubSub, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return &PubSub{}, err
	}

	ps := PubSub{
		Client: client,
		TopicConfig: TopicConfig{
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

// CreateTopicAndSubscription creates a topic and subscription if they don't exist yet.
func (c *PubSub) CreateTopicAndSubscription(ctx context.Context) error {
	// Create a topic
	topic := c.Client.Topic(c.TopicConfig.TopicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return fmt.Errorf("Error checking if topic exists: %v", err)
	}

	if !exists {
		_, err := c.Client.CreateTopic(ctx, c.TopicConfig.TopicName)
		if err != nil {
			return fmt.Errorf("Error creating topic: %v", err)
		}

		fmt.Printf("Topic %s created\n", c.TopicConfig.TopicName)
	} else {
		fmt.Printf("Topic %s already exists\n", c.TopicConfig.TopicName)
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

// Publish publishes a message to the pubsub topic.
func (c *PubSub) Publish(ctx context.Context, msg *Message) ([]string, error) {
	m := &pubsub.Message{
		Attributes: msg.Attributes,
		Data:       msg.Data,
	}

	topic := c.Client.Topic(c.TopicConfig.TopicName)
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
