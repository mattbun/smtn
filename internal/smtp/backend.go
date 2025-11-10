package smtp

import "github.com/emersion/go-smtp"

type Backend struct {
	MessageReceiver MessageReceiver
}

func (b *Backend) NewSession(conn *smtp.Conn) (smtp.Session, error) {
	return &Session{MessageReceiver: b.MessageReceiver}, nil
}
