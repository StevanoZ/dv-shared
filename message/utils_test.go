package message

import (
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/stretchr/testify/assert"
)

const data = "TESTING"

func TestSetRetryOrSetDataToDB(t *testing.T) {
	config := shrd_utils.LoadBaseConfig("../app", "test")
	config.RETRY_TIME = 100 * time.Millisecond
	t.Run("Shoud retry", func(t *testing.T) {
		attempt := 3
		msg := pubsub.Message{
			Data:            []byte(data),
			DeliveryAttempt: &attempt,
		}
		SetRetryOrSetDataToDB(config, &msg, func() {})
	})
	t.Run("Shoud not retry and ack message", func(t *testing.T) {
		attempt := 5
		msg := pubsub.Message{
			Data:            []byte(data),
			DeliveryAttempt: &attempt,
		}
		SetRetryOrSetDataToDB(config, &msg, func() {})
	})
}

func TestBuildDescErrorMsg(t *testing.T) {
	msg := pubsub.Message{
		ID: "123",
	}

	descMsg := BuildDescErrorMsg("error description", &msg)
	assert.Equal(t, "error description, messageID: 123", descMsg)
}
