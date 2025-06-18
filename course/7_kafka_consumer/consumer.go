package consumer

// Consumer consumes messages and handles them while running. Once stopped, it stops consumption and has a graceful
// shutdown.
type Consumer struct {
}

func (m *Consumer) Start() error {
}

func (m *Consumer) Stop() error {
}
