## Task

Implement the `Consumer` interface. You can use the fake Kafka client in /pkg/kafka to fetch messages. Alternatively,
you can talk to a real (local) Kafka cluster - see the "challenges" section.

## Requirements

* The consumer must:
    * Handle messages concurrently (with configurable concurrency);
    * Have configurable concurrency;
    * Have a graceful shutdown mechanism.

## Running your solution

I imagine it's nice to be able to see if the code you write actually works. You will need a `main` package for this. My
suggestion is to create a `main.go` file in the root directory, name its package `main`, have it import your Kafka
consumer package, and then run `main.go`:

```
go run main.go
```

## Challenges for the pros

* Talk to a real (local) Kafka cluster.
    * Use [this library](https://github.com/segmentio/kafka-go) to talk to Kafka.
    * Start Kafka and Kafka UI by running `docker compose up -d` in the project's root directory.
    * Access Kafka UI at https://localhost:8092.
* Make it even faster by committing offsets asynchronously.
* Make it even faster by reading multiple messages at once.
* Write tests for your implementation.
    * This way you can experience how difficult it can be to test concurrent code.
* Benchmark your implementation.

## Producing to Kafka

Those who are talking to a real Kafka cluster may want to produce messages to Kafka so that they have messages that can
be consumed and experimented with. You can run ./dev/produce/main.go to produce 100 messages to Kafka. 