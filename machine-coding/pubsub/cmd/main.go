package main

import (
	"fmt"

	"go-practice.com/machine-coding/pubsub/pubsub"
)

func main() {
	broker := pubsub.NewBroker()
	err := broker.CreateTopic("topicA")
	fmt.Printf("create topic err: %v\n", err)
	err = broker.CreateTopic("topicA")
	fmt.Printf("create topic  err: %v\n", err)

	for i := 0; i < 10; i++ {
		err = broker.Publish(fmt.Sprintf("message_%d", i), "topicA")
		fmt.Printf("publish err: %v\n", err)
	}
	err = broker.Publish("message_10", "topicB")
	fmt.Printf("publish err 2: %v\n", err)

	err = broker.Subscribe("consumerA", "topicB")
	fmt.Printf("subscribe err: %v\n", err)

	err = broker.Subscribe("consumerA", "topicA")
	fmt.Printf("subscribe err 2: %v\n", err)

	err = broker.Subscribe("consumerA", "topicA")
	fmt.Printf("subscribe err 3: %v\n", err)

	err = broker.Unsubscribe("consumerA", "topicA")
	fmt.Printf("Unsubscribe err 1: %v\n", err)

	err = broker.Unsubscribe("consumerA", "topicA")
	fmt.Printf("Unsubscribe err 2: %v\n", err)

	err = broker.Subscribe("consumerA", "topicA")
	fmt.Printf("Subscribe err 1: %v\n", err)

	messages, updatedLastConsumedOffset, err := broker.Consume("consumerA", "topicA", 0)
	fmt.Printf("Consume result print: %+v %+v %+v\n\n", messages, updatedLastConsumedOffset, err)

	messages, updatedLastConsumedOffset, err = broker.Consume("consumerA", "topicA", 0)
	fmt.Printf("Consume result print: %+v %+v %+v\n\n", messages, updatedLastConsumedOffset, err)

	// todo: processing logic here to process all messages

	err = broker.Acknowledge("consumerA", "topicA", updatedLastConsumedOffset)
	fmt.Printf("Acknowledge err: %v\n", err)

	messages, updatedLastConsumedOffset, err = broker.Consume("consumerA", "topicA", 0)
	fmt.Printf("Consume result print: %+v %+v %+v\n\n", messages, updatedLastConsumedOffset, err)

	err = broker.Acknowledge("consumerA", "topicB", updatedLastConsumedOffset)
	fmt.Printf("Acknowledge err: %v\n", err)

	err = broker.Acknowledge("consumerB", "topicA", updatedLastConsumedOffset)
	fmt.Printf("Acknowledge err: %v\n\n", err)

	for i := 0; i < 11; i++ {
		err = broker.Publish(fmt.Sprintf("message_%d", i), "topicA")
		fmt.Printf("publish err: %v\n", err)
	}

	messages, updatedLastConsumedOffset, err = broker.Consume("consumerA", "topicA", 4)
	fmt.Printf("Consume result print: %+v %+v %+v\n\n", messages, updatedLastConsumedOffset, err)

	err = broker.Acknowledge("consumerA", "topicA", updatedLastConsumedOffset)
	fmt.Printf("Acknowledge err: %v\n", err)

	messages, updatedLastConsumedOffset, err = broker.Consume("consumerA", "topicA", 4)
	fmt.Printf("Consume result print: %+v %+v %+v\n\n", messages, updatedLastConsumedOffset, err)

	err = broker.Acknowledge("consumerA", "topicA", updatedLastConsumedOffset)
	fmt.Printf("Acknowledge err: %v\n", err)

	messages, updatedLastConsumedOffset, err = broker.Consume("consumerA", "topicA", 4)
	fmt.Printf("Consume result print: %+v %+v %+v\n\n", messages, updatedLastConsumedOffset, err)

	err = broker.Acknowledge("consumerA", "topicA", updatedLastConsumedOffset)
	fmt.Printf("Acknowledge err: %v\n", err)

	messages, updatedLastConsumedOffset, err = broker.Consume("consumerA", "topicA", 4)
	fmt.Printf("Consume result print: %+v %+v %+v\n\n", messages, updatedLastConsumedOffset, err)

	err = broker.Acknowledge("consumerA", "topicA", updatedLastConsumedOffset)
	fmt.Printf("Acknowledge err: %v\n", err)
}
