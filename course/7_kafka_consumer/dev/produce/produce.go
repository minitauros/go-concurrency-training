package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	consumer "github.com/minitauros/go-concurrency-training/course/7_kafka_consumer"
	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Balancer:  &kafka.Hash{},
		Transport: &kafka.Transport{},
	}

	messages := make([]kafka.Message, 100)
	for i := 0; i < 100; i++ {
		messages[i] = kafka.Message{
			Topic: "foo",
			Value: msgBytes(i),
		}
	}

	err := w.WriteMessages(context.Background(), messages...)
	if err != nil {
		panic(err)
	}
	fmt.Println("produced 100 messages to topic 'foo'")
}

func msgBytes(num int) []byte {
	jsonBytes, err := json.Marshal(consumer.Message{
		Foo: "foo" + strconv.Itoa(num),
		Bar: "bar" + strconv.Itoa(num),
	})
	if err != nil {
		panic(err)
	}
	return jsonBytes
}
