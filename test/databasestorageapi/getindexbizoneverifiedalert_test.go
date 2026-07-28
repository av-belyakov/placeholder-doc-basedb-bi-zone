package databasestorageapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

func TestGetIndexBiZoneVerifiedAlerts(t *testing.T) {
	const Search_UUID = "eb62b40e-067a-4261-824f-56c3ff280db3"

	var (
		indexes []string = []string{"module_placeholderdb_bizone_alerts_2025_8"}
		apiDBS  *databasestorageapi.DatabaseStorage

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

				return

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

				if msg.GetType() == "error" {
					log.Fatal(msg.GetMessage())

					return
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
		apiDBS.Start(ctx)
	})

	t.Run("Тест 2. Получить документ", func(t *testing.T) {
		//ищем объект с таким же идентификатором как и принятый в обработку объект
		res, err := apiDBS.GetDocument(
			ctx,
			indexes,
			strings.NewReader(
				fmt.Sprintf(
					`{
						"query": {
							"bool": {
								"must": [{"match": {"uuid": "%s"}}]}}}`,
					Search_UUID,
				)),
			0)
		assert.NoError(t, err)

		//обрабатываем принятую от базы данных информацию
		response := databasestorageapi.ResponseVerifiedBiZoneAlerts{}
		err = json.Unmarshal(res, &response)
		assert.NoError(t, err)
		assert.Equal(t, response.Options.Hits[0].Source.UUID, Search_UUID)
	})

	t.Run("Тест 3. Получить некоторое количество документов", func(t *testing.T) {
		fmt.Println("\nFound documents:")

		for num := range 3 {
			res, err := apiDBS.GetDocument(
				ctx,
				indexes,
				strings.NewReader(`{ "query": { "match_all": {} } }`),
				num)
			assert.NoError(t, err)

			//fmt.Println("---------------------------------------")
			//fmt.Println(string(res))

			//обрабатываем принятую от базы данных информацию
			response := databasestorageapi.ResponseVerifiedBiZoneAlerts{}
			err = json.Unmarshal(res, &response)
			assert.NoError(t, err)

			for k, v := range response.Options.Hits {
				fmt.Printf("%d.%d. UUID:%s\n", num, k, v.Source.UUID)
			}
		}
	})
}
