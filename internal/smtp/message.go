package smtp

// Message is a message received by the SMTP server.
type Message struct {
	// Subject is a the subject of the email message.
	Subject string

	// Body is the contents of the email message.
	Body string
}

// MessageReceiver is an interface for receiving messages from the SMTP server.
type MessageReceiver interface {
	// OnMessage is called whenever the SMTP server receives a message.
	OnMessage(message Message) error
}
