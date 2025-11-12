package notify

import (
	"errors"

	"github.com/containrrr/shoutrrr"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Notifier sends notifications using a [Sender].
type Notifier struct {
	// Sender is the [Sender] used to send notifications.
	Sender Sender
}

// NotifyInput contains options for the [Notifier.Notify] function.
type NotifyInput struct {
	// Title is the title of the notification.
	Title string

	// Body is the content of the notification.
	Body string
}

// Notify sends a push notification with the Notifier's [Sender].
func (n Notifier) Notify(input NotifyInput) error {
	params := types.Params{}
	params.SetTitle(input.Title)

	errs := n.Sender.Send(input.Body, &params)

	return errors.Join(errs...)
}

// NewNotifier creates a [Notifier] that sends notifications to the given URLs.
func NewNotifier(urls []string) (Notifier, error) {
	sender, err := shoutrrr.CreateSender(urls...)
	if err != nil {
		return Notifier{}, err
	}

	return Notifier{
		Sender: sender,
	}, nil
}
