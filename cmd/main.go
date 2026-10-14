package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-base_db module is stop")
	}()

	NewApp().Start(ctx)
}
