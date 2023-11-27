## Pubsub

The package is used to initiate the pub/sub handler which provides the following functions:

- To create a new pubsub handler which can be used to create topic.
- To create a new topic which can be used to create a new subscription and publish message.
- To create a subscription which can be used to receive the messages published in the topic.

### Example code:

```go
// Initiate a new pubsub handler.
pbsbHandler, _ := pubsub.New(ctx, "sample-project-id")

// Create any number topics by using this method on the handler.
topic, _ := pbsbHandler.CreateTopic(ctx, "sample-topic")
// Create any number of subscription by using this method.
subscription, _ := topic.CreateSubscription(ctx, "sample-subscription")

// Create a channel to receive the messages from subscription.
myChannel := make(chan []byte)
message := api.Message{}

go func() {
  // Start a goroutine to read the messages from the subscription.
  subscription.ReceiveMessages(myChannel)
}()

for {
  select {
  // Any message read in the goroutine above will be sent to this channel.
  case data := <-messageChannel:
    proto.Unmarshal(data, &message)
    fmt.Printf("Received message: %v\n", message)
  }
}

// Publish the message to the topic.
topic.Publish(ctx, pubsub.Message{})
```
