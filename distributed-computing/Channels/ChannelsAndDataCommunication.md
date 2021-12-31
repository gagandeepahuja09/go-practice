* We can use any primitive or user created type for sending events in channel(even functions).

Messages And Events
Events ==> Signal channels

Buffered Channel
* Using unbuffered channels can significantly slow down the program because we will have to wait for each message to be processed before we can put in another message.
* A buffered channel has a limited capacity and maintains a queue of messages which the goroutine will consume at its own pace.
* Case
* Buffered channel empty: Receiving is blocked until a message is sent.
* Buffered full: Sending is blocked until a message is received.
* Partial: Both possible.

Closing Channels
* When we no longer want to send message on any of the said channel. The behaviour on closing channels will depend on the type of channel. On sending messages, it will cause panic in all cases. For receive:
* Unbuffered closed channel: Will yield an immediate zero value of the element type.
* Buffered closed channel: Will first yield all the values in the channel's queue. Once all the values are exhausted, then will start returning zero values.
* In order to determine if a channel has been closed, we can try using close. val, ok := <-ch. If closed, ok will be false.
* We can also range over a channel. When a channel closes, the range loop ends.
* Closing a closed channel, nil channel or receive only channel will cause panic. Only a bidirectional channel or send-only channel can be closed.
* Closing a channel is not mandatory and not relevant for the garbage collector(GC). If a GC determines that a channel is unreachable irrespective of whether it is open/closed it will be garbage collected.