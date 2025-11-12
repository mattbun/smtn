package smtp

import (
	"fmt"
	"log/slog"

	"github.com/emersion/go-smtp"
)

// StartServerInput contains options for [StartServer].
type StartServerInput struct {
	// Address is the address that the server should listen on.
	Address string

	// Port is the port that the server should listen on.
	Port int

	// AllowInsecureAuth determines whether to allow insecure authentication.
	AllowInsecureAuth bool

	// MessageReceiver is a MessageReceiver to send new messages to.
	MessageReceiver MessageReceiver
}

// StartServer starts the SMTP server.
func StartServer(input StartServerInput) error {
	server := smtp.NewServer(&Backend{MessageReceiver: input.MessageReceiver})

	server.Addr = fmt.Sprintf("%s:%d", input.Address, input.Port)
	server.AllowInsecureAuth = input.AllowInsecureAuth

	slog.Info("Starting server", slog.String("listen-addr", input.Address), slog.Int("port", input.Port))
	slog.Debug("Full server configuration", slog.Any("config", input))

	return server.ListenAndServe()
}
