package main

import (
	"log/slog"

	"github.com/mattbun/smtn/internal/notify"
	"github.com/mattbun/smtn/internal/smtp"
)

// NotifierMessageReceiver is an implementation of [smtp.MessageReceiver] that sends messages to a [notify.Notifier].
type NotifierMessageReceiver struct {
	// Notifier is a [notify.Notifier] to send messages to.
	Notifier notify.Notifier
}

// OnMessage forwards an SMTP message to a [notify.Notifier].
func (r NotifierMessageReceiver) OnMessage(message smtp.Message) error {
	slog.Debug("Sending notification")

	err := r.Notifier.Notify(notify.NotifyInput{
		Title: message.Subject,
		Body:  message.Body,
	})
	if err != nil {
		slog.Error("Error sending notification", slog.Any("error", err))
	}

	slog.Debug("Successfully sent notification")
	return err
}

// NewNotifierMessageReceiver creates a new [NotifierMessageReceiver].
func NewNotifierMessageReceiver(urls []string) (NotifierMessageReceiver, error) {
	notifier, err := notify.NewNotifier(urls)
	if err != nil {
		return NotifierMessageReceiver{}, err
	}

	return NotifierMessageReceiver{
		Notifier: notifier,
	}, nil
}
