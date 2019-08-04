package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"go.stevenxie.me/bincode/internal/info"
	"go.stevenxie.me/bincode/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "server"
	app.Usage = "A binary representation conversion server."
	app.UsageText = "server [global options]"

	app.Version = info.Version
	app.Author = "Steven Xie <hello@stevenxie.me>"
	app.HideHelp = true
	app.Action = run

	// Hide help command.
	app.Commands = []cli.Command{{Name: "help", Hidden: true}}

	// Configure flags.
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Usage:  "port that the server listens on",
			EnvVar: "PORT",
			Value:  3000,
		},
		cli.DurationFlag{
			Name:  "shutdown-interval",
			Usage: "maximum time to wait for graceful shutdown",
			Value: 5 * time.Second,
		},
		cli.BoolFlag{
			Name:  "help,h",
			Usage: "show help",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	// Init logger.
	log := buildLogger()

	// Create and configure server.
	log.Info("Initializing server...")
	srv := server.New(
		func(cfg *server.Config) {
			cfg.Logger = log
		},
	)

	// Shut down server gracefully upon interrupt.
	{
		var (
			interval *time.Duration
			value    = c.Duration("shutdown-interval")
		)
		if value > 0 {
			interval = &value
		}
		go shutdownUponInterrupt(srv, log, interval)
	}

	err := srv.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")))
	if err == http.ErrServerClosed {
		err = nil
	}
	return errors.Wrap(err, "starting server")
}

func shutdownUponInterrupt(
	srv *server.Server,
	log logrus.FieldLogger,
	timeout *time.Duration,
) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	// Wait for interrupt signal.
	<-sig

	const msg = "Received interrupt signal; shutting down."
	if timeout != nil {
		log.WithField("timeout", timeout.String()).Info(msg)
	} else {
		log.Info(msg)
	}

	// Prepare shutdown context.
	ctx := context.Background()
	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), *timeout)
		defer cancel()
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.WithError(err).Error("Server didn't shut down correctly.")
	}
}
