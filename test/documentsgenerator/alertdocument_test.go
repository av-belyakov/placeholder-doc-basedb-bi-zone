package documentsgenerator

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	kafkaapi "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/kafkaapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supporting"
)

var (
	f *os.File

	err error
)

func TestAlertDocument(t *testing.T) {
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

	os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "development")

	t.Run("Тест 1. Инициализация файла для записи сгенерированных, на основании входящих данных, документов", func(t *testing.T) {
		f, err = os.OpenFile(
			fmt.Sprintf("generate_documents_%s.log", time.Now().Format(time.RFC3339)),
			os.O_RDWR|os.O_CREATE,
			0644,
		)
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Инициализация модуля взаимодействия с kafka", func(t *testing.T) {
		//инициализация модуля
		kafkaApiModule, err := kafkaapi.New(
			logging,
			counting,
			kafkaapi.WithNameRegionalObject("GCM-TEST"),
			kafkaapi.WithHost("192.168.9.71"),
			kafkaapi.WithPort(31792),
			kafkaapi.WithTopicsSubscription(map[string]string{
				"alerts":      "alertmgr-alerts",
				"soar-alerts": "soar-alertmgr-alerts-default-topic",
			}),
			kafkaapi.WithCacheTTL(3600),
		)
		assert.NoError(t, err)

		//запуск модуля
		err = kafkaApiModule.Start(ctx)
		assert.NoError(t, err)

		//обработка входящих сообщений
		for msg := range kafkaApiModule.GetChannelFromModule() {
			t.Log("Received new message...")
			t.Logf("Received message:%s\n\n", string(msg.Data))
			decoder := decoderjsondocuments.New(counting, logging)
			id, verifedBiZoneAlert, listRawFields := documentgenerator.BiZoneAlertsGenerator(decoder.Start(msg.Data))

			t.Log("with id:", id)
			//t.Logf("Received message:\n%s\n", str)
			if _, err = fmt.Fprintf(
				f,
				"--- TOPIC - '%s' (id:%s) %s\n%s\n___ List raw fields ___:\n%+v\n\n",
				msg.SubjectType,
				id,
				time.Now().Format(time.RFC3339),
				verifedBiZoneAlert.ToStringBeautiful(0),
				listRawFields); err != nil {
				t.Logf("Error: %+v\n", err)
			}
		}
	})

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDBBZ_MAIN")
		f.Close()
	})
}
