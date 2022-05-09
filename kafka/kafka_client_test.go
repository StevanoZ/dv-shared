package kafka_client

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	mock_kafka "github.com/StevanoZ/dv-shared/kafka/mock"
	shrd_service "github.com/StevanoZ/dv-shared/service"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	TOPIC                  = "TESTING"
	TOPIC_KEY              = "test-topic"
	MESSAGE                = "Just for testing"
	error_produce_message  = errors.New("error produce message")
	error_consumer_message = errors.New("error consumer message")
)

func loadBaseConfig() *shrd_utils.BaseConfig {
	return shrd_utils.LoadBaseConfig("../app", "test")
}

func initKafkaClient(ctrl *gomock.Controller) (
	shrd_service.MessageBrokerClient, *mock_kafka.MockKafkaProducer, *mock_kafka.MockKafkaConsumer) {
	producer := mock_kafka.NewMockKafkaProducer(ctrl)
	consumer := mock_kafka.NewMockKafkaConsumer(ctrl)

	return NewKafkaClient(producer, consumer), producer, consumer
}
func TestNewProducer(t *testing.T) {
	config := loadBaseConfig()

	t.Run("Local Kafka", func(t *testing.T) {
		producer, err := NewKafkaProducer(config)
		assert.Nil(t, err)
		assert.NotNil(t, producer)
	})

	t.Run("Remote Kafka", func(t *testing.T) {
		config.IsRemoteBroker = true
		producer, err := NewKafkaProducer(config)
		assert.Nil(t, err)
		assert.NotNil(t, producer)
	})
}

func TestNewConsumer(t *testing.T) {
	config := loadBaseConfig()

	t.Run("Local Consumer", func(t *testing.T) {
		consumer, err := NewKafkaConsumer(config)
		assert.NoError(t, err)
		assert.NotNil(t, consumer)
	})

	t.Run("Remote Consumer", func(t *testing.T) {
		config.IsRemoteBroker = true
		consumer, err := NewKafkaConsumer(config)
		assert.NoError(t, err)
		assert.NotNil(t, consumer)
	})
}

func TestNewKafkaClient(t *testing.T) {
	config := loadBaseConfig()
	producer, err := NewKafkaProducer(config)
	assert.NoError(t, err)
	assert.NotNil(t, producer)

	consumer, err := NewKafkaConsumer(config)
	assert.NoError(t, err)
	assert.NotNil(t, consumer)

	kafkaClient := NewKafkaClient(producer, consumer)
	assert.NotNil(t, kafkaClient)
}

func TestProducerSendEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kafkaClient, producer, _ := initKafkaClient(ctrl)

	t.Run("Success Produce Message", func(t *testing.T) {
		value, err := json.Marshal(MESSAGE)
		assert.NoError(t, err)

		producer.EXPECT().Produce(gomock.AssignableToTypeOf(&kafka.Message{}), nil).DoAndReturn(func(message *kafka.Message, _ interface{}) error {
			assert.Equal(t, &TOPIC, message.TopicPartition.Topic)
			assert.Equal(t, kafka.PartitionAny, message.TopicPartition.Partition)
			assert.Equal(t, []byte(TOPIC_KEY), message.Key)
			assert.Equal(t, value, message.Value)
			assert.WithinDuration(t, time.Now(), message.Timestamp, 10*time.Millisecond)
			return nil
		}).Times(1)
		producer.EXPECT().Flush(15000).Times(1)
		err = kafkaClient.SendEvents([]string{TOPIC}, TOPIC_KEY, "Just for testing")
		assert.NoError(t, err)
	})

	t.Run("Failed produce message", func(t *testing.T) {
		producer.EXPECT().Produce(gomock.AssignableToTypeOf(&kafka.Message{}), nil).Return(error_produce_message).Times(1)
		producer.EXPECT().Flush(15000).Times(0)
		err := kafkaClient.SendEvents([]string{TOPIC}, TOPIC_KEY, MESSAGE)
		assert.Error(t, err)
		assert.Equal(t, error_produce_message, err)
	})
}

func TestConsumerListenEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	kafkaClient, _, consumer := initKafkaClient(ctrl)

	t.Run("Success consume message", func(t *testing.T) {
		consumer.EXPECT().Subscribe(TOPIC, nil).Return(nil).Times(1)
		consumer.EXPECT().Close().Times(1)
		consumer.EXPECT().ReadMessage(gomock.Any()).Return(&kafka.Message{Value: []byte(MESSAGE)}, nil).Times(1)
		kafkaClient.ListenEvent(TOPIC, func(payload any, errMsg error, close func()) {
			close()
			msg, isKafkaMsg := payload.(*kafka.Message)
			assert.Equal(t, true, isKafkaMsg)
			assert.Nil(t, errMsg)
			assert.NotNil(t, payload)
			assert.Equal(t, []byte(MESSAGE), msg.Value)
		})
	})

	t.Run("Failed consume message (when subscribe)", func(t *testing.T) {
		consumer.EXPECT().Subscribe(TOPIC, nil).Return(error_consumer_message).Times(1)
		consumer.EXPECT().Close().Times(1)

		err := kafkaClient.ListenEvent(TOPIC, func(payload any, errMsg error, close func()) {
			close()
			msg, isKafkaMsg := payload.(*kafka.Message)
			assert.Equal(t, true, isKafkaMsg)
			assert.Nil(t, errMsg)
			assert.Nil(t, msg)
		})

		assert.NotNil(t, err)
	})

	t.Run("Failed consume message (when reading)", func(t *testing.T) {
		consumer.EXPECT().Subscribe(TOPIC, nil).Times(1)
		consumer.EXPECT().Close().Times(1)
		consumer.EXPECT().ReadMessage(gomock.Any()).Return(&kafka.Message{}, error_consumer_message).Times(1)

		err := kafkaClient.ListenEvent(TOPIC, func(payload any, errMsg error, close func()) {
			close()
			assert.NotNil(t, errMsg)
			assert.Nil(t, payload)
		})

		assert.Nil(t, err)
	})
}
