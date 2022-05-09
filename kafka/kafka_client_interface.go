package kafka_client

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
	Flush(timeoutMs int) int
	Close()
}

type KafkaConsumer interface {
	Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
	Close() error
}
