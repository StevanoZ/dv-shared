package pubsub_client

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	shrd_service "github.com/StevanoZ/dv-shared/service"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
)

type GooglePubSub interface {
	CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error)
	CreateSubscription(ctx context.Context, id string, cfg pubsub.SubscriptionConfig) (*pubsub.Subscription, error)
	Topic(id string) *pubsub.Topic
	Subscription(id string) *pubsub.Subscription
	Close() error
}

type PubSubClientImpl struct {
	config *shrd_utils.BaseConfig
	pubSub GooglePubSub
}

func NewGooglePubSub(config *shrd_utils.BaseConfig) (c *pubsub.Client, err error) {
	ctx := context.Background()
	projectId := config.GCP_PROJECT_ID

	return pubsub.NewClient(ctx, projectId)
}

func NewPubSubClient(config *shrd_utils.BaseConfig, pubSub GooglePubSub) shrd_service.PubSubClient {
	return &PubSubClientImpl{config: config, pubSub: pubSub}
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

	tpc, err = p.pubSub.CreateTopic(ctx, topicName)
	tpc.EnableMessageOrdering = true

	return tpc, err
}

func (p *PubSubClientImpl) CreateSubscriptionIfNotExists(ctx context.Context, id string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	sub := p.pubSub.Subscription(id)

	ok, err := sub.Exists(ctx)

	if err != nil {
		return nil, err
	}

	if ok {
		return sub, nil
	}

	return p.pubSub.CreateSubscription(ctx, id, pubsub.SubscriptionConfig{
		Topic:                 topic,
		EnableMessageOrdering: true,
		AckDeadline:           20 * time.Second,
		DeadLetterPolicy: &pubsub.DeadLetterPolicy{
			DeadLetterTopic:     p.config.DLQ_TOPIC,
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
		topic.EnableMessageOrdering = true
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

		log.Println("publish message with ID: ", id)
	}

	return nil
}

func (p *PubSubClientImpl) PullMessages(ctx context.Context, id string, topic *pubsub.Topic, callback func(ctx context.Context, msg *pubsub.Message)) error {
	defer p.pubSub.Close()

	sub, err := p.CreateSubscriptionIfNotExists(ctx, id, topic)
	if err != nil {
		return err
	}

	return sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Println("received message with ID: ", msg.ID)

		callback(ctx, msg)
		msg.Ack()
	})
}

func (p *PubSubClientImpl) Close() error {
	return p.pubSub.Close()
}
