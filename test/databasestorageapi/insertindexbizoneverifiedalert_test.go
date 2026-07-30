package databasestorageapi_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

func TestIInsertIndexBiZoneAlert(t *testing.T) {
	var (
		verifedBiZoneAlert *datamodels.VerifiedBiZoneIRPAlert
		apiDBS             *databasestorageapi.DatabaseStorage
		fb                 []byte
		id                 string

		err error
	)

	//загружаем ключи и пароли
	if err := godotenv.Load(".env"); err != nil {
		t.Log(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
	t.Cleanup(func() {
		stop()
	})

	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-basedb-bi-zone module is Stop")

		stop()
	}()

	logging := helpers.NewLogging()
	counting := helpers.NewCounting()

	go func(ctx context.Context, l interfaces.Logger, c *helpers.Counting) {
		for {
			select {
			case <-ctx.Done():
				return

			case count := <-c.GetChan():
				fmt.Println("count.message:", count.Message, " count.Count", count.Count)

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

				if msg.GetType() == "error" {
					log.Fatal(msg.GetMessage())
				}

			}
		}
	}(ctx, logging, counting)

	t.Run("Тест 1. Подключение к БД", func(t *testing.T) {
		apiDBS, err = helpers.DataBaseConnectionInitialization(
			ctx,
			logging,
			counting,
			databasestorageapi.WithHost("192.168.9.208"),
			databasestorageapi.WithPort(9200),
			databasestorageapi.WithUser("placeholder-docbase-db"),
			databasestorageapi.WithPasswd(os.Getenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD")),
			databasestorageapi.WithStorage(map[string]string{
				"alerts":      "module_placeholderdb_bizone_alerts",
				"soar-alerts": "module_placeholderdb_bizone_soar-alerts",
			}),
		)
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Чтение тестового файла с объектом типа 'alerts'", func(t *testing.T) {
		fb, err = os.ReadFile("../kafka/incoming_alert_1.json")
		//fb, err = os.ReadFile("../kafka/replacing_alert_1.json")
		assert.NoError(t, err)
		assert.NotEqual(t, len(fb), 0)
	})

	t.Run("Тест 3. Генерирование нового объекта типа datamodels.VerifiedBiZoneAlert", func(t *testing.T) {
		decoder := decoderjsondocuments.New(counting, logging)
		id, verifedBiZoneAlert, _ = documentgenerator.BiZoneAlertsGenerator(decoder.Start(fb))

		t.Logf("\nID:'%s'\nVerifiedBiZoneAlert:'%#v'\n", id, verifedBiZoneAlert)

		assert.NotEqual(t, id, "")
		assert.NotEmpty(t, verifedBiZoneAlert)
	})

	t.Run("Тест 4. Вставка нового индекса в БД или замена старого", func(t *testing.T) {
		apiDBS.TestAddBiZoneAlerts(ctx, verifedBiZoneAlert)

		time.Sleep(time.Second * 3)

	})
}
