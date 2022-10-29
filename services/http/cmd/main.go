package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"imperial-fleet-inventory/services/http/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := app.NewApplication().Run(ctx); err != nil {
		log.Fatal(err)
	}
}
