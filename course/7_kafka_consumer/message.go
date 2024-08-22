package consumer

// Message defines the structure of the messages in the topic for this example implementation.
type Message struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}
