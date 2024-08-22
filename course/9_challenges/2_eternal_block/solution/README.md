There are multiple ways we can fix this. We could use a locking mechanism to check if the consumer is already stopping, and if so, just return instead of trying to send an error into the critical errors channel. We could also create a buffered channel, which would allow multiple errors to be sent into the channel. It's not a problem that only the first error is handled. At least the code won't be blocking. The channel will be garbage collected after the consumer goes out of memory. Probably we are shutting down the application, anyway.

We could, of course, also call `waitGroup.Done()` _before_ sending the value onto the channel. But then we would have to call it in multiple places, which means also maintaining it in multiple places, which is not as nice.