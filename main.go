package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	pubsub_client "github.com/StevanoZ/dv-shared/pubsub"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
)

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	//	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "service-account.json")
	fmt.Println(os.Getenv("PUBSUB_EMULATOR_HOST"))
	ctx := context.Background()
	config := shrd_utils.LoadBaseConfig("app", "app")
	gPubSub, _ := pubsub_client.NewGooglePubSub(config)
	client := pubsub_client.NewPubSubClient(config, gPubSub)

	topic, err := client.CreateTopicIfNotExists(ctx, "DLQ")
	fmt.Println("ERROR TOPIC", err)
	// _, err = client.CreateTopicIfNotExists(ctx, "DLQ_USER")
	// fmt.Println("ERROR TOPIC", err)
	// _, err = client.CreateTopicIfNotExists(ctx, "DLQ_USER-IMAGE")
	// fmt.Println("ERROR TOPIC", err)

	client.PullMessages(ctx, "SUBZ", topic, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("DATA", msg.Data)
	})
}
