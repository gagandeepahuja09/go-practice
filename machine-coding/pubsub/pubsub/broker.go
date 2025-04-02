package pubsub

import (
	"errors"
)

type consumerNameTopicName struct {
	consumerName string
	topicName    string
}

type Broker struct {
	topicNameVsMessages                       map[string][]string
	consumerNameTopicNameVsLastConsumedOffset map[consumerNameTopicName]int
}

func NewBroker() Broker {
	return Broker{
		topicNameVsMessages:                       make(map[string][]string),
		consumerNameTopicNameVsLastConsumedOffset: make(map[consumerNameTopicName]int),
	}
}

func (b *Broker) CreateTopic(topicName string) error {
	if _, ok := b.topicNameVsMessages[topicName]; ok {
		return errors.New("topic name already exists")
	}
	b.topicNameVsMessages[topicName] = []string{}
	return nil
}

func (b *Broker) Publish(message, topicName string) error {
	if _, ok := b.topicNameVsMessages[topicName]; !ok {
		return errors.New("topic name not found")
	}
	b.topicNameVsMessages[topicName] = append(b.topicNameVsMessages[topicName], message)
	return nil
}

func (b *Broker) Subscribe(consumerName, topicName string) error {
	if _, ok := b.topicNameVsMessages[topicName]; !ok {
		return errors.New("topic name not found")
	}
	consumerNameTopicNameStruct := consumerNameTopicName{
		consumerName: consumerName,
		topicName:    topicName,
	}
	if _, ok := b.consumerNameTopicNameVsLastConsumedOffset[consumerNameTopicNameStruct]; ok {
		return errors.New("consumer already subscribed to the topic")
	}

	b.consumerNameTopicNameVsLastConsumedOffset[consumerNameTopicNameStruct] = -1
	return nil
}

func (b *Broker) Unsubscribe(consumerName, topicName string) error {
	if _, ok := b.topicNameVsMessages[topicName]; !ok {
		return errors.New("topic name not found")
	}
	consumerNameTopicNameStruct := consumerNameTopicName{
		consumerName: consumerName,
		topicName:    topicName,
	}
	if _, ok := b.consumerNameTopicNameVsLastConsumedOffset[consumerNameTopicNameStruct]; !ok {
		return errors.New("consumer already unsubscribed to the topic")
	}

	delete(b.consumerNameTopicNameVsLastConsumedOffset, consumerNameTopicNameStruct)
	return nil
}

// returns the list of all unprocessed messages along with the updatedLastConsumedOffset: which is the
// offset value to be updated after all the messages are processed.
// offset value is updated via Acknowledge function
// if batchSize is 0, then all messages are returned. Else only the no. of messages as per batchSize
// are processed
func (b *Broker) Consume(consumerName, topicName string, batchSize int) ([]string, int, error) {
	topicMessages, ok := b.topicNameVsMessages[topicName]
	if !ok {
		return []string{}, 0, errors.New("topic name not found")
	}
	consumerNameTopicNameStruct := consumerNameTopicName{
		consumerName: consumerName,
		topicName:    topicName,
	}
	lastConsumedOffset, ok := b.consumerNameTopicNameVsLastConsumedOffset[consumerNameTopicNameStruct]
	if !ok {
		return []string{}, 0, errors.New("consumer is not subscribed to the topic")
	}
	// a[low:high]

	if batchSize == 0 {
		return topicMessages[lastConsumedOffset+1:], len(topicMessages) - 1, nil
	}

	updatedLastConsumedOffset := len(topicMessages) - 1
	if lastConsumedOffset+batchSize < updatedLastConsumedOffset {
		updatedLastConsumedOffset = lastConsumedOffset + batchSize
	}

	return topicMessages[lastConsumedOffset+1 : lastConsumedOffset+1+batchSize], updatedLastConsumedOffset, nil
}

func (b *Broker) Acknowledge(consumerName, topicName string, updatedLastConsumedOffset int) error {
	topicMessages, ok := b.topicNameVsMessages[topicName]
	if !ok {
		return errors.New("topic name not found")
	}
	consumerNameTopicNameStruct := consumerNameTopicName{
		consumerName: consumerName,
		topicName:    topicName,
	}
	currentLastConsumedOffset, ok := b.consumerNameTopicNameVsLastConsumedOffset[consumerNameTopicNameStruct]
	if !ok {
		return errors.New("consumer is not subscribed to the topic")
	}

	if updatedLastConsumedOffset > len(topicMessages) || updatedLastConsumedOffset < currentLastConsumedOffset {
		return errors.New("updatedLastConsumedOffset should be in the range of currentLastConsumedOffset and no. of messages within the topic")
	}

	b.consumerNameTopicNameVsLastConsumedOffset[consumerNameTopicNameStruct] = updatedLastConsumedOffset
	return nil
}
