package main

import (
	"log/slog"

	"github.com/mattbun/smtprrr/internal/notify"
	"github.com/mattbun/smtprrr/internal/smtp"
)

type NotifierMessageReceiver struct {
	Notifier notify.Notifier
}

func (r NotifierMessageReceiver) OnMessage(message smtp.Message) error {
	slog.Debug("Sending notification")

	err := r.Notifier.Notify(notify.NotifyInput{
		Title: message.Subject,
		Body:  message.Body,
	})
	if err != nil {
		slog.Error("Error sending notification", slog.Any("error", err))
	}

	return err
}

func NewNotifierMessageReceiver(urls []string) (NotifierMessageReceiver, error) {
	notifier, err := notify.NewNotifier(urls)
	if err != nil {
		return NotifierMessageReceiver{}, err
	}

	return NotifierMessageReceiver{
		Notifier: notifier,
	}, nil
}
