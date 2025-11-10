package smtp

type Message struct {
	Subject string
	Body    string
}

type MessageReceiver interface {
	OnMessage(message Message) error
}
