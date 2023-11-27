package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// Handler is the GCLOUD PubSub implementation of the pubsub.Service Interface
type Handler struct {
	client *pubsub.Client
}

// Topic contains the topic client handler.
type Topic struct {
	client *pubsub.Topic
	pbsb   *Handler
}

// Subscription contains the subscription client handler.
type Subscription struct {
	client             *pubsub.Subscription
	cancelSubscription context.CancelFunc
}

// Message is the message that we will publish to pubsub
type Message struct {
	// Attributes are the pubsub message attributes.
	Attributes map[string]string

	// Data is the message of the pubsub message
	Data json.RawMessage
}

// New creates a pubsub using GCE PubSub backend.
func New(ctx context.Context, projectID string) (*Handler, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	ps := Handler{
		client: client,
	}

	return &ps, nil
}

// CreateTopic creates a topic if it don't exist yet.
func (h *Handler) CreateTopic(ctx context.Context, topicName string) (*Topic, error) {
	// Create a topic.
	topic := h.client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error checking if topic exists: %v", err)
	}

	if !exists {
		topic, err = h.client.CreateTopic(ctx, topicName)
		if err != nil {
			return nil, fmt.Errorf("Error creating topic: %v", err)
		}

		fmt.Printf("Topic %s created\n", topicName)
	} else {
		fmt.Printf("Topic %s already exists\n", topicName)
	}

	return &Topic{
		client: topic,
		pbsb:   h,
	}, nil
}

// CreateSubscription creates a subscription if it don't exist yet.
func (t *Topic) CreateSubscription(ctx context.Context, subscriptionName string) (*Subscription, error) {
	subscription := t.pbsb.client.Subscription(subscriptionName)
	found, err := subscription.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to check if subscription [%s] exists: %v", subscriptionName, err)
	}

	if !found {
		subscription, err = t.pbsb.client.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
			Topic: t.client,
		})

		if err != nil {
			return nil, err
		}
	}

	return &Subscription{
		client: subscription,
	}, nil
}

// Close closes the subscription.
func (s *Subscription) Close() {
	if s.cancelSubscription != nil {
		s.cancelSubscription()
	}
}
