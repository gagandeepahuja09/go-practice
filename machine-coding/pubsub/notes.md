## Functions
* createTopic(topicName string) returns (error)
* publish(message string, topicName string) returns (error)
* subscribe(consumerName, topicName string) returns (error)
* unsubscribe(consumerName, topicName string) returns (error)
* consume(consumerName, topicName string) returns ([]string, error)
    * this will return an array of all messages which are not yet consumed by the consumer.

## Entities
* One class: Broker

## Data Structures
* `lastConsumedOffset` OR `consumerTopicNameVsLastConsumedOffset`: map[string]int (key: consumerName_topicName) ==> (value: lastConsumedOffset which is the array index of the topicName array)
* `topicNameVsMessages` => map[string][]string
    * key topicName, value: list of messages


## What Will Functions Do?
* publish function will update the `topicNameVsMessages` map.
* subscribe function
    * checks if already subscribed first
    * set `consumerTopicNameVsLastConsumedOffset` map for consumerName_topicName to 0
* unsubsribe function
    * unset key for consumerName_topicName from `consumerTopicNameVsLastConsumedOffset` map
* consume function 
    * this will update the `consumerTopicNameVsLastConsumedOffset` to the length of the `topicNameVsMessages` for that specific topic indicating that all messages till now are processed.
    * we can either do batching for updating or directly update

## How to handle availability issue with the consumer?
* It will restart from the lastConsumedOffset. Hence few message might be restarted.

## How to handle availability issue with the broker?
* The current design is inmemory, so all messages will be lost. In order to avoid that, we require some persistence.
* During bootup, we can read everything from the file.
* We can also do async updates to the file in batches.

## How to handle filtered message consumption?
* As of now, the message is a string. It will be hard to apply filtering on a string. We will have to change to some data structure.