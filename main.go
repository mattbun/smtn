package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/mattbun/smtprrr/internal/smtp"
	"github.com/urfave/cli/v3"
)

const (
	FlagAllowInsecure   = "allow-insecure"
	FlagListenAddr      = "listen-addr"
	FlagNotificationUrl = "notification-url"
	FlagPort            = "port"
	FlagVerbose         = "verbose"
)

var cmd *cli.Command = &cli.Command{
	Name:  "smtprrr",
	Usage: "Run a SMTP server that turns messages into notifications",

	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    FlagAllowInsecure,
			Aliases: []string{"i"},
			Usage:   "Allow insecure auth",
			Value:   false,
			Sources: cli.EnvVars("ALLOW_INSECURE"),
		},

		&cli.StringFlag{
			Name:    FlagListenAddr,
			Aliases: []string{"l"},
			Usage:   "Address that the SMTP server should listen on",
			Value:   "127.0.0.1",
			Sources: cli.EnvVars("LISTEN_ADDR"),
		},

		&cli.StringSliceFlag{
			Name:     FlagNotificationUrl,
			Aliases:  []string{"n"},
			Usage:    "Shoutrrr notification url(s)",
			Sources:  cli.EnvVars("NOTIFICATION_URL"),
			Required: true,
		},

		&cli.IntFlag{
			Name:    FlagPort,
			Aliases: []string{"p"},
			Usage:   "Port that the SMTP server should listen on",
			Value:   25,
			Sources: cli.EnvVars("PORT"),
		},

		&cli.BoolFlag{
			Name:    FlagVerbose,
			Aliases: []string{"v"},
			Usage:   "Enable verbose logging",
			Value:   false,
			Sources: cli.EnvVars("VERBOSE"),
		},
	},

	Action: func(_ context.Context, cmd *cli.Command) error {
		configureLogging(cmd.Bool(FlagVerbose))

		notifier, err := NewNotifierMessageReceiver(cmd.StringSlice(FlagNotificationUrl))
		if err != nil {
			return err
		}

		return smtp.StartServer(smtp.StartServerInput{
			Address:           cmd.String(FlagListenAddr),
			Port:              cmd.Int(FlagPort),
			AllowInsecureAuth: cmd.Bool(FlagAllowInsecure),

			MessageReceiver: &notifier,
		})
	},
}

func main() {
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
