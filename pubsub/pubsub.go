package pubsub_client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
)

const DLQ_TOPIC = "DLQ-Topic"

type GooglePubSub interface {
	CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error)
	CreateSubscription(ctx context.Context, id string, cfg pubsub.SubscriptionConfig) (*pubsub.Subscription, error)
	Topic(id string) *pubsub.Topic
	Subscription(id string) *pubsub.Subscription
	Close() error
}

type PubSubClient interface {
	CreateTopicIfNotExists(ctx context.Context, topicName string) (*pubsub.Topic, error)
	CreateSubscriptionIfNotExists(ctx context.Context, client GooglePubSub, id string, topic *pubsub.Topic) (*pubsub.Subscription, error)
	PublishTopics(ctx context.Context, topics []*pubsub.Topic, data any, orderingKey string) error
	PullMessages(ctx context.Context, id string, topic *pubsub.Topic, callback func(ctx context.Context, msg *pubsub.Message))
}

type PubSubClientImpl struct {
	pubSub GooglePubSub
}

func NewGooglePubSub(config *shrd_utils.BaseConfig) (c *pubsub.Client, err error) {
	ctx := context.Background()
	projectId := config.GCP_PROJECT_ID

	return pubsub.NewClient(ctx, projectId)
}

func NewPubSubClient(pubSub GooglePubSub) PubSubClient {
	return &PubSubClientImpl{pubSub: pubSub}
}

func (p *PubSubClientImpl) CreateTopicIfNotExists(ctx context.Context, topicName string) (*pubsub.Topic, error) {
	tpc := p.pubSub.Topic(topicName)

	ok, err := tpc.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if ok {
		return tpc, nil
	}

	return p.pubSub.CreateTopic(ctx, topicName)
}

func (p *PubSubClientImpl) CreateSubscriptionIfNotExists(ctx context.Context, client GooglePubSub, id string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	sub := client.Subscription(id)
	ok, err := sub.Exists(ctx)

	if err != nil {
		return nil, err
	}

	if ok {
		return sub, nil
	}

	return client.CreateSubscription(ctx, id, pubsub.SubscriptionConfig{
		Topic:                 topic,
		EnableMessageOrdering: true,
		AckDeadline:           20 * time.Second,
		DeadLetterPolicy: &pubsub.DeadLetterPolicy{
			DeadLetterTopic:     DLQ_TOPIC,
			MaxDeliveryAttempts: 5,
		},
	})
}

func (p *PubSubClientImpl) PublishTopics(ctx context.Context, topics []*pubsub.Topic, data any, orderingKey string) error {
	var results []*pubsub.PublishResult
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		res := topic.Publish(ctx, &pubsub.Message{
			Data:        message,
			OrderingKey: orderingKey,
		})
		results = append(results, res)
	}

	for _, result := range results {
		id, err := result.Get(ctx)
		if err != nil {
			return err
		}

		fmt.Println("publish message with ID: ", id)
	}

	return nil
}

func (p *PubSubClientImpl) PullMessages(ctx context.Context, id string, topic *pubsub.Topic, callback func(ctx context.Context, msg *pubsub.Message)) {
	defer p.pubSub.Close()
	sub, _ := p.CreateSubscriptionIfNotExists(ctx, p.pubSub, id, topic)
	sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		callback(ctx, msg)
		msg.Ack()
	})

}