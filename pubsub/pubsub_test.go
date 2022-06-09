package pubsub_client

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	shrd_helper "github.com/StevanoZ/dv-shared/helper"
	shrd_service "github.com/StevanoZ/dv-shared/service"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/stretchr/testify/assert"
)

const PROJECT = "TEST-PROJECT"
const TOPIC = "TOPIC"
const SUBSCRIPTION = "TESTING"
const DLQ = "DLQ"
const MESSAGE = "TEST MESSAGE"
const ORDER_KEY = "order-key"

var CONTEXT = context.Background()

func loadBaseConfig() *shrd_utils.BaseConfig {
	return shrd_utils.LoadBaseConfig("../app", "test")
}

func initPubSubClient(t *testing.T, gPubSub *pubsub.Client) shrd_service.PubSubClient {
	pubSubClient := NewPubSubClient(loadBaseConfig(), gPubSub)
	assert.NotNil(t, pubSubClient)
	return pubSubClient
}

func createTopicAndDLQ(t *testing.T, client shrd_service.PubSubClient) (topic *pubsub.Topic, dlqTopic *pubsub.Topic) {
	topic, err := client.CreateTopicIfNotExists(CONTEXT, TOPIC)
	assert.NoError(t, err)
	assert.Equal(t, topic.ID(), TOPIC)
	assert.Equal(t, true, topic.EnableMessageOrdering)

	dlqTopic, err = client.CreateTopicIfNotExists(CONTEXT, DLQ)
	assert.NoError(t, err)
	assert.Equal(t, dlqTopic.ID(), DLQ)

	return topic, dlqTopic
}

func publishTopic(t *testing.T, client shrd_service.PubSubClient) *pubsub.Topic {
	topic, _ := createTopicAndDLQ(t, client)
	err := client.PublishTopics(CONTEXT, []*pubsub.Topic{topic, topic}, MESSAGE, ORDER_KEY)
	assert.NoError(t, err)

	return topic
}

func TestNewGooglePubSub(t *testing.T) {
	// JUST LOAD FAKE CREDENTIALS
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "service-account.json")

	gPubSub, err := NewGooglePubSub(loadBaseConfig())
	assert.NoError(t, err)
	assert.NotNil(t, gPubSub)
}

func TestNewPubSubClient(t *testing.T) {
	gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
	defer close()

	client := initPubSubClient(t, gPubSub)
	assert.NotNil(t, client)
}

func TestCreateTopicIfNotExists(t *testing.T) {
	t.Run("Create topic if not exist", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)

		topic, err := client.CreateTopicIfNotExists(CONTEXT, TOPIC)
		assert.Equal(t, topic.ID(), TOPIC)
		assert.Equal(t, true, topic.EnableMessageOrdering)
		assert.NoError(t, err)
	})

	t.Run("Not create topic if already exists", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)
		topic, err := gPubSub.CreateTopic(CONTEXT, TOPIC)
		assert.NoError(t, err)
		assert.NotNil(t, topic)

		topic, err = client.CreateTopicIfNotExists(CONTEXT, TOPIC)
		assert.NoError(t, err)
		assert.Equal(t, false, topic.EnableMessageOrdering)
	})

	t.Run("Failed when creating topic", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		close()

		client := initPubSubClient(t, gPubSub)

		topic, err := client.CreateTopicIfNotExists(CONTEXT, TOPIC)
		assert.Error(t, err)
		assert.Nil(t, topic)
	})
}

func TestCreateSubscriptionIfNotExists(t *testing.T) {
	t.Run("Create subscription if not exist", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)
		topic, _ := createTopicAndDLQ(t, client)

		sub, err := client.CreateSubscriptionIfNotExists(CONTEXT, SUBSCRIPTION, topic)
		assert.NoError(t, err)
		assert.Equal(t, sub.ID(), SUBSCRIPTION)
	})

	t.Run("Not create subscription if already exist", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)

		topic, _ := createTopicAndDLQ(t, client)

		sub, err := client.CreateSubscriptionIfNotExists(CONTEXT, SUBSCRIPTION, topic)
		assert.NoError(t, err)
		assert.Equal(t, sub.ID(), SUBSCRIPTION)

		sub, err = client.CreateSubscriptionIfNotExists(CONTEXT, SUBSCRIPTION, topic)
		assert.NoError(t, err)
		assert.Equal(t, sub.ID(), SUBSCRIPTION)
	})

	t.Run("Failed when subscribe topic", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		close()

		client := initPubSubClient(t, gPubSub)

		sub, err := client.CreateSubscriptionIfNotExists(CONTEXT, SUBSCRIPTION, &pubsub.Topic{})
		assert.Error(t, err)
		assert.Nil(t, sub)
	})
}

func TestPublishTopics(t *testing.T) {
	t.Run("Success publish topic", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)
		topic, _ := createTopicAndDLQ(t, client)

		err := client.PublishTopics(CONTEXT, []*pubsub.Topic{topic, topic}, MESSAGE, ORDER_KEY)
		assert.NoError(t, err)
	})

	t.Run("Failed when publish topic", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)
		topic, _ := createTopicAndDLQ(t, client)

		err := client.PublishTopics(CONTEXT, []*pubsub.Topic{topic}, func() {}, ORDER_KEY)
		assert.Error(t, err)
	})

	t.Run("Failed when publish topic", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)
		topic, _ := createTopicAndDLQ(t, client)
		topic.PublishSettings.Timeout = 1 * time.Microsecond

		err := client.PublishTopics(CONTEXT, []*pubsub.Topic{topic}, MESSAGE, ORDER_KEY)
		assert.Error(t, err)
	})
}

func TestPullMessages(t *testing.T) {
	t.Run("Success pull message", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
		defer close()

		client := initPubSubClient(t, gPubSub)

		topic, _ := createTopicAndDLQ(t, client)

		go func() {
			deliveryAttempt := 3
			msgData, err := json.Marshal(MESSAGE)
			assert.NoError(t, err)

			topic.Publish(CONTEXT, &pubsub.Message{
				Data:            msgData,
				OrderingKey:     ORDER_KEY,
				DeliveryAttempt: &deliveryAttempt,
			})
		}()

		go func() {
			err := client.PullMessages(CONTEXT, SUBSCRIPTION, topic, func(ctx context.Context, msg *pubsub.Message) {
				var data string

				err := json.Unmarshal(msg.Data, &data)
				assert.NoError(t, err)
				assert.Equal(t, MESSAGE, data)
			})
			assert.NoError(t, err)
		}()

		time.Sleep(600 * time.Millisecond)
	})

	t.Run("Failed when pull message", func(t *testing.T) {
		gPubSub, close := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)

		client := initPubSubClient(t, gPubSub)
		topic := publishTopic(t, client)

		go func() {
			close()
			err := client.PullMessages(CONTEXT, SUBSCRIPTION, topic, func(ctx context.Context, msg *pubsub.Message) {
				assert.Equal(t, nil, msg)
			})
			assert.Error(t, err)
		}()

		time.Sleep(300 * time.Millisecond)
	})
}

func TestClose(t *testing.T) {
	gPubSub, _ := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
	client := initPubSubClient(t, gPubSub)
	err := client.Close()
	assert.NoError(t, err)
}

func TestCheckTopicAndPublish(t *testing.T) {
	ctx := context.Background()
	gPubSub, _ := shrd_helper.CreateFakeGooglePubSub(t, PROJECT)
	client := initPubSubClient(t, gPubSub)
	defer gPubSub.Close()

	topic, _ := createTopicAndDLQ(t, client)
	t.Run("Publish message", func(t *testing.T) {
		client.CheckTopicAndPublish(ctx, []string{topic.ID()}, ORDER_KEY, MESSAGE)
	})

	t.Run("Should not publish message", func(t *testing.T) {
		client.CheckTopicAndPublish(ctx, []string{}, ORDER_KEY, MESSAGE)
	})

	t.Run("Failed when publish message", func(t *testing.T) {
		ctxCancel, cancel := context.WithCancel(ctx)
		cancel()
		client.CheckTopicAndPublish(ctxCancel, []string{TOPIC}, ORDER_KEY, "publish-message")
	})
}
