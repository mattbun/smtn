package notify

import (
	"errors"

	"github.com/containrrr/shoutrrr"
	"github.com/containrrr/shoutrrr/pkg/types"
)

type Notifier struct {
	Sender Sender
}

type NotifyInput struct {
	Title string
	Body  string
}

func (n Notifier) Notify(input NotifyInput) error {
	params := types.Params{}
	params.SetTitle(input.Title)

	errs := n.Sender.Send(input.Body, &params)

	return errors.Join(errs...)
}

func NewNotifier(urls []string) (Notifier, error) {
	sender, err := shoutrrr.CreateSender(urls...)
	if err != nil {
		return Notifier{}, err
	}

	return Notifier{
		Sender: sender,
	}, nil
}
