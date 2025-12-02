A solution would be to add a new channel specifically for shutting down the consumer. The loop in `Start()` now first checks if the consumer has been stopped. Only if not, it is allowed to fetch a message from Kafka.

On top of that the work channel is never closed, which means that even if work is started after `Stop()` was called, the program will not panic.