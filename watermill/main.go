package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	config := AzureServiceBusConfig{
		ConnectionString: os.Getenv(("SERVICE_BUS_CONNECTION_STRING")),
		TopicName:        os.Getenv("SERVICE_BUS_TOPIC_NAME"),
		SubscriptionName: os.Getenv("SERVICE_BUS_SUBSCRIPTION_NAME"),
	}

	// Publisher
	publisher, err := NewPublisher(config)
	if err != nil {
		fmt.Printf("failed to create publisher: %v", err)
	}
	defer publisher.Close()

	// Publish a message
	msg := message.NewMessage("1", []byte("Hello, Azure Service Bus!"))
	err = publisher.Publish(config.TopicName, msg)
	if err != nil {
		fmt.Printf("failed to publish message: %v", err)
	}

	// Subscriber
	subscriber, err := NewSubscriber(config)
	if err != nil {
		fmt.Printf("failed to create subscriber: %v", err)
	}
	defer subscriber.Close()

	// Subscribe to messages
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	messages, err := subscriber.Subscribe(ctx, config.TopicName)
	if err != nil {
		fmt.Printf("failed to subscribe: %v", err)
	}

	for msg := range messages {
		fmt.Printf("Received message: %s", string(msg.Payload))
	}
}