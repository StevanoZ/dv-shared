package message

import (
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
)

func SetRetryOrSetDataToDB(config *shrd_utils.BaseConfig, msg *pubsub.Message, cb func()) {
	if msg.DeliveryAttempt != nil && *msg.DeliveryAttempt <= 4 {
		time.Sleep(config.RETRY_TIME)
		log.Printf("retry message with messageID: %s, orderingKey: %s \n", msg.ID, msg.OrderingKey)
		msg.Nack()
	}

	if msg.DeliveryAttempt != nil && *msg.DeliveryAttempt > 4 {
		cb()
		log.Printf("acknowledged message")
		msg.Ack()
	}
}

func BuildDescErrorMsg(desc string, err error) string {
	return fmt.Sprintf("%s, Error: %s", desc, err.Error())
}
