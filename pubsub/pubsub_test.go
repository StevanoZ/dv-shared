package pubsub_client

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/pubsub"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type PubSubTopic struct {
	*pubsub.Topic
	exist bool
	err   error
}

func (p *PubSubTopic) Exists(ctx context.Context) (bool, error) {
	return p.exist, p.err
}

func loadBaseConfig() *shrd_utils.BaseConfig {
	return shrd_utils.LoadBaseConfig("../app", "test")
}

func initPubSubClient(t *testing.T, ctrl *gomock.Controller) (PubSubClient, *MockGooglePubSub) {
	gPubSub := NewMockGooglePubSub(ctrl)
	pubSubClient := NewPubSubClient(gPubSub)
	assert.NotNil(t, pubSubClient)
	return pubSubClient, gPubSub
}

func TestNewGooglePubSub(t *testing.T) {
	// JUST LOAD FAKE CREDENTIALS
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "service-account.json")

	gPubSub, err := NewGooglePubSub(loadBaseConfig())
	assert.NoError(t, err)
	assert.NotNil(t, gPubSub)
}

func TestNewPubSubClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	initPubSubClient(t, ctrl)
}

// func TestCreateTopicIfNotExists(t *testing.T) {
// 	ctx := context.Background()
// 	topic := "TEST-TOPIC"
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	client, gPubSub := initPubSubClient(t, ctrl)

// 	// pubSubTopic := PubSubTopic{
// 	// 	exist: true,
// 	// 	err:   nil,
// 	// }

// 	gPubSub.EXPECT().Topic(topic).Return(&pubsub.Topic{}).Times(1)

// 	_, err := client.CreateTopicIfNotExists(ctx, topic)
// 	assert.NoError(t, err)
// }

