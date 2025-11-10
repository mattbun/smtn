package smtp

import (
	"fmt"
	"log/slog"

	"github.com/emersion/go-smtp"
)

type StartServerInput struct {
	Address           string
	Port              int
	AllowInsecureAuth bool

	MessageReceiver MessageReceiver
}

func StartServer(input StartServerInput) error {
	server := smtp.NewServer(&Backend{MessageReceiver: input.MessageReceiver})

	server.Addr = fmt.Sprintf("%s:%d", input.Address, input.Port)
	server.AllowInsecureAuth = input.AllowInsecureAuth

	slog.Info("Starting server", slog.String("listen-addr", input.Address), slog.Int("port", input.Port))
	slog.Debug("Full server configuration", slog.Any("config", input))

	return server.ListenAndServe()
}
