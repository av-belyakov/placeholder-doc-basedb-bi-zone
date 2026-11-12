package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/appversion"
)

var version string

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	done := make(chan struct{})
	go func() {
		<-ctx.Done()

		var count int
		t := time.NewTicker(time.Second * 1)
		defer t.Stop()
		for range t.C {
			count++
			fmt.Println(count)

			if count == 3 {
				break
			}
		}

		close(done)
	}()

	version, _ = appversion.GetAppVersion()
	NewApp().Start(ctx)

	<-done
	log.Printf("%vstopping module 'placeholder_doc-base_db'...%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)
}
