package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

// AzureServiceBusConfig holds configuration for Azure Service Bus
type AzureServiceBusConfig struct {
	ConnectionString string
	TopicName        string
	SubscriptionName string
}

// ServiceBusClient struct have the project id and the client object
type ServiceBusClient struct {
	connectionString string
	client           *azservicebus.Client
}

// NewMicrosoftServiceBusClient creates new pub/sub client using the native library
func NewMicrosoftServiceBusClient(connectionString string) *ServiceBusClient {
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		fmt.Printf("failed to create Service Bus client: %v", err)
	}

	return &ServiceBusClient{
		connectionString: connectionString,
		client:           client,
	}
}

// Publisher is a Watermill publisher for Azure Service Bus
type Publisher struct {
	client *azservicebus.Client
	sender  *azservicebus.Sender
}

// NewPublisher creates a new Publisher
func NewPublisher(config AzureServiceBusConfig) (*Publisher, error) {
	client, err := azservicebus.NewClientFromConnectionString(config.ConnectionString, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating Service Bus client")
	}

	topicSender, err := client.NewSender(config.TopicName, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating topic sender")
	}

	return &Publisher{
		client: client,
		sender:  topicSender,
	}, nil
}

// Publish sends a message to Azure Service Bus
func (p *Publisher) Publish(topic string, messages ...*message.Message) error {
	for _, msg := range messages {
		// Create a new Azure Service Bus message
		sbMessage := &azservicebus.Message{
			Body: msg.Payload,
		}

		// Add metadata as user properties
		for k, v := range msg.Metadata {
			sbMessage.ApplicationProperties[k] = v
		}

		// Send the message
		err := p.sender.SendMessage(context.Background(), sbMessage, nil)
		if err != nil {
			return errors.Wrap(err, "sending message to Service Bus")
		}
	}
	return nil
}

// Close closes the publisher
func (p *Publisher) Close() error {
	if err := p.sender.Close(context.Background()); err != nil {
		return errors.Wrap(err, "closing topic sender")
	}
	return nil
}

// Subscriber is a Watermill subscriber for Azure Service Bus
type Subscriber struct {
	client    *azservicebus.Client
	receiver  *azservicebus.Receiver
}

// NewSubscriber creates a new Subscriber
func NewSubscriber(config AzureServiceBusConfig) (*Subscriber, error) {
	client, err := azservicebus.NewClientFromConnectionString(config.ConnectionString, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating Service Bus client")
	}

	subscriptionReceiver, err := client.NewReceiverForSubscription(config.TopicName, config.SubscriptionName, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating subscription receiver")
	}

	return &Subscriber{
		client:       client,
		receiver: subscriptionReceiver,
	}, nil
}

// Subscribe subscribes to messages from Azure Service Bus
func (s *Subscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	output := make(chan *message.Message)

	go func() {
		defer close(output)

		// Receive messages
		for {
			msg, err := s.receiver.ReceiveMessages(ctx, 1, nil)
			if err != nil {
				if err.Error() == "context canceled" {
					return
				}
				fmt.Printf("Error receiving message: %v", err)
				continue
			}

			// Create Watermill message from Azure Service Bus message
			wmMsg := message.NewMessage(msg[0].MessageID, msg[0].Body)

			// Add user properties to metadata
			for k, v := range msg[0].ApplicationProperties {
				wmMsg.Metadata.Set(k, v.(string))
			}

			// Send message to channel
			select {
			case output <- wmMsg:
			case <-ctx.Done():
				return
			}

			// Complete the message to acknowledge receipt
			if err := s.receiver.CompleteMessage(context.Background(), msg[0], nil); err != nil {
				fmt.Printf("Error completing message: %v", err)
			}
		}
	}()

	return output, nil
}

// Close closes the subscriber
func (s *Subscriber) Close() error {
	if err := s.receiver.Close(context.Background()); err != nil {
		return errors.Wrap(err, "closing subscription receiver")
	}
	return nil
}
