package notify

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

type Sender interface {
	Send(message string, params *types.Params) []error
}
