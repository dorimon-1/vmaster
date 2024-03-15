/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dorimon-1/vmaster/cmd/app"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGPIPE,
	)
	defer stop()

	app.Run(ctx)
}
