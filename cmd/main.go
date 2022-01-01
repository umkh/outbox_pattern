package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/umkh/outbox_pattern/internal/bootstrap"
)

func main() {
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	app := bootstrap.New()

	go func() {
		<-sigc
		cancel()
	}()

	app.Run(ctx)
}
