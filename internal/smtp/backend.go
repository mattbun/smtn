package smtp

import "github.com/emersion/go-smtp"

// Backend implements [github.com/emersion/go-smtp.Backend].
type Backend struct {
	// MessageReceiver is an interface that defines an OnMessage method, which is called whenever a message is received.
	MessageReceiver MessageReceiver
}

// NewSession creates a new [Session] with the [Backend]'s [MessageReceiver].
func (b *Backend) NewSession(conn *smtp.Conn) (smtp.Session, error) {
	return &Session{MessageReceiver: b.MessageReceiver}, nil
}
