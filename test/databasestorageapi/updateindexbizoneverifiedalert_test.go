package databasestorageapi

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
	"time"

	"github.com/goforj/godump"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

func TestUpdateIndexBiZoneVerifiedAlert(t *testing.T) {
	const Index_ID = "eb62b40e-067a-4261-824f-56c3ff280db3"

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
				fmt.Println("\tcount.message:", count.Message, " count.Count", count.Count)

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
			t.Context(),
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
		/*
		*
		* пока что используем этот файл для замены уже имеющихся данных
		* потом заменим его на такой же, но где данные отличаются от данных в БД
		*
		 */
		fb, err = os.ReadFile("../kafka/replacing_alert_1.json")
		assert.NoError(t, err)
		assert.NotEqual(t, len(fb), 0)
	})

	t.Run("Тест 3. Генерирование нового объекта типа datamodels.VerifiedBiZoneAlert", func(t *testing.T) {
		decoder := decoderjsondocuments.New(counting, logging)
		id, verifedBiZoneAlert, _ = documentgenerator.BiZoneAlertsGenerator(decoder.Start(fb))

		fmt.Println("\tID:", id)
		fmt.Println("\verifedBiZoneAlert:")
		godump.Dump(verifedBiZoneAlert)

		assert.NotEqual(t, id, "")
		assert.NotEmpty(t, verifedBiZoneAlert)
	})

	t.Run("Тест 4. Поиск и замена индекса в БД", func(t *testing.T) {
		//получаем существующие индексы
		existingIndexes, err := apiDBS.GetExistingIndexes(ctx, "module_placeholderdb_bizone_alerts")
		assert.NoError(t, err)

		//ищем похожие индексы
		indexesOnlyCurrentYear := []string(nil)
		for _, v := range existingIndexes {
			if strings.Contains(v, fmt.Sprint(time.Now().Year())) {
				indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
			}
		}
		assert.Greater(t, len(indexesOnlyCurrentYear), 0)

		//ищем объект с таким же идентификатором как и принятый в обработку объект
		currentQuery := strings.NewReader(
			fmt.Sprintf(
				"{\"query\": {\"bool\": {\"must\": [{\"match\": {\"uuid\": \"%s\"}}]}}}",
				Index_ID))
		res, err := apiDBS.Search(ctx, indexesOnlyCurrentYear, currentQuery)
		assert.NoError(t, err)

		//обрабатываем принятую от базы данных информацию
		response := databasestorageapi.ResponseVerifiedBiZoneAlerts{}
		err = json.NewDecoder(res.Body).Decode(&response)
		assert.NoError(t, err)

		//проверяем наличие документа
		assert.Greater(t, len(response.Options.Hits), 0)

		//выполняем замену искомого документа
		var countReplacingFields int
		listDeleting := []databasestorageapi.ServiseOption(nil)
		updateVerified := datamodels.NewVerifiedBiZoneIRPAlert()
		//заполняем новый объект информацией из базы данных
		for _, v := range response.Options.Hits {
			countReplacingFields += updateVerified.RepalcingOldBiZoneAlert(*v.Source.Get())
			listDeleting = append(listDeleting, databasestorageapi.ServiseOption{
				ID:    v.ID,
				Index: v.Index,
			})

			//заполняем новый объект дополнительной информацией о сенсорах и
			// месторасположении ip адресов
			updateVerified.SetAdditionalInformation(*v.Source.GetAdditionalInformation())
		}

		//выполняем обновление нового объекта данными полученными от брокера сообщений
		num := updateVerified.RepalcingOldBiZoneAlert(*verifedBiZoneAlert.Get())
		assert.NotEqual(t, num, 0)

		t.Log("\nCOUNT replacing:", num)
		fmt.Println("\tupdateVerified:")
		godump.Dump(updateVerified)
	})
}
