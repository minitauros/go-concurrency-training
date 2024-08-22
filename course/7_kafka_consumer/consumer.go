package consumer

// Consumer consumes messages and handles them while running. Once stopped, it stops consumption and has a graceful
// shutdown.
type Consumer interface {
	Start() error
	Stop() error
}
