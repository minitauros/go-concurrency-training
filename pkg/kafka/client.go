package kafka

import (
	"sync/atomic"
)

var defaultMessageBody = []byte(`{"foo":"bar","bar":"baz"}`)

type Message struct {
	Topic     string
	Offset    int64
	Partition int
	Body      []byte
}

type Client struct {
	offset atomic.Int64
}

// ReadMessage fetches one message from Kafka and automatically commits the offset.
func (c *Client) ReadMessage() Message {
	msg := c.FetchMessage()
	c.offset.Add(1)
	return msg
}

// FetchMessage fetches one message from Kafka but does not commit the offset.
func (c *Client) FetchMessage() Message {
	currentOffset := c.offset.Load()
	return Message{
		Topic:     "foo.bar.baz",
		Offset:    currentOffset,
		Partition: 1,
		Body:      defaultMessageBody,
	}
}

// CommitOffset commits the offset.
func (c *Client) CommitOffset(topic string, partition int, offset int64) {
	c.offset.Store(offset)
}
