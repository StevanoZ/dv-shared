package kafka_client

import (
	"encoding/json"
	"fmt"
	"time"

	shrd_service "github.com/StevanoZ/dv-shared/service"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang.org/x/sync/errgroup"
)

type KafkaClientImpl struct {
	producer KafkaProducer
	consumer KafkaConsumer
}

func NewKafkaProducer(config *shrd_utils.BaseConfig) (*kafka.Producer, error) {
	var producer *kafka.Producer
	var err error

	if config.IsRemoteBroker {
		producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers":  config.KafkaBroker,
			"security.protocol":  "SASL_PLAINTEXT",
			"sasl.username":      config.KafkaUsername,
			"sasl.password":      config.KafkaPassword,
			"sasl.mechanism":     "PLAIN",
			"enable.idempotence": true,
		})
	} else {
		producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers":  config.KafkaBroker,
			"enable.idempotence": true,
		})
	}

	return producer, err
}

func NewKafkaConsumer(config *shrd_utils.BaseConfig) (*kafka.Consumer, error) {
	var consumer *kafka.Consumer
	var err error

	if config.IsRemoteBroker {
		consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": config.KafkaBroker,
			"security.protocol": "SASL_PLAINTEXT",
			"sasl.username":     config.KafkaUsername,
			"sasl.password":     config.KafkaPassword,
			"sasl.mechanism":    "PLAIN",
			"group.id":          config.ServiceName,
			"auto.offset.reset": "earliest",
		})
		return consumer, err
	} else {
		consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":     config.KafkaBroker,
			"group.id":              config.ServiceName,
			"auto.offset.reset":     "earliest",
			"broker.address.family": "v4",
		})
		return consumer, err
	}
}

func NewKafkaClient(producer KafkaProducer, consumer KafkaConsumer) shrd_service.MessageBrokerClient {
	return &KafkaClientImpl{
		producer: producer,
		consumer: consumer,
	}
}

func (c *KafkaClientImpl) SendEvents(topics []string, key string, message interface{}) error {
	ewg := errgroup.Group{}
	value, _ := json.Marshal(message)
	for i := range topics {
		topic := topics[i]
		ewg.Go(func() error {

			err := c.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &topic,
					Partition: kafka.PartitionAny,
				},
				Key:       []byte(key),
				Value:     value,
				Timestamp: time.Now(),
			}, nil)

			if err != nil {
				return err
			}

			fmt.Println("send message to topic: ", topic)
			c.producer.Flush(15000)
			return nil
		})
	}

	if err := ewg.Wait(); err != nil {
		return err
	}
	return nil
}

func (c *KafkaClientImpl) ListenEvent(topic string, cb func(payload any, errMsg error, close func())) error {
	err := c.consumer.Subscribe(topic, nil)
	defer c.consumer.Close()

	if err != nil {
		fmt.Println("error when subscribe topic", err)
		return err
	}

	fmt.Println("started consuming topic: ", topic)

	isListen := true
	close := func() {
		isListen = false
	}
	for isListen {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			cb(nil, err, close)

		} else {
			fmt.Printf("message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			cb(msg, nil, close)
		}

	}

	return nil
}
