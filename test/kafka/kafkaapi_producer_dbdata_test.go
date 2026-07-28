package kafka

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

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

func TestProducerDBData(t *testing.T) {
	var (
		topic         string   = "object.topicalerttype.test"
		outputIndexes []string = []string{"module_placeholderdb_bizone_alerts_2025_8"}

		producer *kafka.Producer
		apiDBS   *databasestorageapi.DatabaseStorage
		response databasestorageapi.ResponseVerifiedBiZoneAlerts
		alerts   []databasestorageapi.PatternVerifiedBiZoneAlert

		err error
	)

	//загружаем ключи и пароли
	if err = godotenv.Load("./../databasestorageapi/.env"); err != nil {
		t.Log(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
	t.Cleanup(func() {
		stop()
		producer.Close()
	})

	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-basedb-bi-zone module is Stop")

		stop()
		producer.Close()
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

	t.Run("Тест 2. Получение некоторого количества документов из БД", func(t *testing.T) {
		fmt.Println("\nFound documents:")

		for num := range 3 {
			res, err := apiDBS.GetDocument(
				ctx,
				outputIndexes,
				strings.NewReader(`{ "query": { "match_all": {} } }`),
				num)
			assert.NoError(t, err)

			//fmt.Println("---------------------------------------")
			//fmt.Println(string(res))

			//обрабатываем принятую от базы данных информацию
			err = json.Unmarshal(res, &response)
			assert.NoError(t, err)

			for k, v := range response.Options.Hits {
				fmt.Printf("%d.%d. UUID:%s\n", num, k, v.Source.Get().GetUUID())

				alerts = append(alerts, v)
			}
		}
	})

	t.Run("Тест 3. Передача некоторого количества документов в kafka", func(t *testing.T) {
		producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
			"group.id":          fmt.Sprintf("%s-group", "GCM-test"),
		})
		assert.NoError(t, err)

		// информационное сообщение о доставки
		go func() {
			for e := range producer.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()

		fmt.Println("\nThe beginning of data transmission")
		for k, v := range alerts {
			fmt.Printf("%d. UUID:%s\n", k, v.Source.Get().GetUUID())

			b, err := json.Marshal(v.Source.Get())
			assert.NoError(t, err)

			producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &topic,
					Partition: kafka.PartitionAny,
				},
				Value: b,
			}, nil)
		}
		producer.Flush(15 * 1000)
	})
}
