package smtp

import (
	"io"
	"log/slog"
	"net/mail"
	"strings"

	"github.com/emersion/go-smtp"
)

// Session implements [github.com/emersion/go-smtp.Session].
type Session struct {
	// MessageReceiver defines a function to be run whenever a message is received.
	MessageReceiver MessageReceiver
}

// Data parses a message and forwards it to the [MessageReceiver].
func (s *Session) Data(r io.Reader) error {
	message, err := mail.ReadMessage(r)
	if err != nil {
		slog.Error("error reading message", slog.Any("error", err))
		return err
	}

	subject := message.Header.Get("Subject")
	bodyBytes, err := io.ReadAll(message.Body)
	if err != nil {
		slog.Error("error reading body", slog.Any("error", err))
		return err
	}

	body := string(bodyBytes)
	slog.Debug("Message received", slog.String("subject", subject), slog.String("body", body))

	return s.MessageReceiver.OnMessage(Message{
		Subject: subject,
		Body:    strings.TrimSpace(string(body)),
	})
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error { return nil }
func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error   { return nil }
func (s *Session) Reset()                                         {}
func (s *Session) Logout() error                                  { return nil }
