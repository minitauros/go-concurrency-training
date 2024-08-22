We could handle the critical errors in a separate goroutine so that it can never block, OR we can actually implement a tokenized worker pool to ensure that we will only try to fetch the next message from Kafka if one of the worker functions actually has completely finished its work and is ready for new work.

We'll go with the latter.