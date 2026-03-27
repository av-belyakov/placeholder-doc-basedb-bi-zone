package kafka_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	kafkaapi "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/kafkaapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supporting"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func TestKafkaApi(t *testing.T) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
 
	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-basedb-bi-zone module is Stop")

		stop()
	}()

	logging := supporting.NewLogging()
	counting := &supporting.Counting{}

	go func(ctx context.Context, l *supporting.Logging) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging)

	//инициализация файла для записи входящих сообщений
	f, err := os.OpenFile(
		fmt.Sprintf("incoming_messages_%s.log", time.Now().Format(time.RFC3339)),
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//инициализация модуля
	kafkaApiModule, err := kafkaapi.New(
		counting,
		logging,
		kafkaapi.WithNameRegionalObject("GCM-TEST"),
		kafkaapi.WithHost("192.168.9.71"),
		kafkaapi.WithPort(31792),
		kafkaapi.WithTopicsSubscription(map[string]string{
			"alerts": "alertmgr-alerts",
			"case":   "socd-soar-prod-issue-v1",
		}),
		kafkaapi.WithCacheTTL(3600),
	)
	if err != nil {
		log.Fatal(err)
	}

	//запуск модуля
	if err := kafkaApiModule.Start(ctx); err != nil {
		log.Fatal(err)
	}

	for msg := range kafkaApiModule.GetChanDataFromModule() {
		//t.Logf("Received message:%s\n\n", string(msg.Data))
		str, err := supportingfunctions.NewReadReflectJSONSprint(msg.Data)
		if err != nil {
			t.Logf("Error: %+v\n", err)

			continue
		}

		//t.Logf("Received message:\n%s\n", str)
		if _, err = fmt.Fprintf(f, "--- TOPIC - '%s' %s\n%s\n\n", msg.SubjectType, time.Now().Format(time.RFC3339), str); err != nil {
			t.Logf("Error: %+v\n", err)
		}
	}
}
