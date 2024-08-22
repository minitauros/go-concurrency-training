package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
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
			Value: msgBytes(),
		}
	}

	err := w.WriteMessages(context.Background(), messages...)
	if err != nil {
		panic(err)
	}
	fmt.Println("done")
}

func msgBytes() []byte {
	jsonBytes, err := json.Marshal(consumer.Message{
		Foo: "foo" + strconv.Itoa(rand.IntN(100)),
		Bar: "bar" + strconv.Itoa(rand.IntN(100)),
	})
	if err != nil {
		panic(err)
	}
	return jsonBytes
}
