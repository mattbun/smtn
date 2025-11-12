package notify

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Sender is a subset of [github.com/containrrr/shoutrrr/pkg/router.ServiceRouter].
type Sender interface {
	// Send sends a notification.
	// It matches the signature of [github.com/containrrr/shoutrrr/pkg/router.ServiceRouter.Send].
	Send(message string, params *types.Params) []error
}
