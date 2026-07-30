package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

func TestKafkaProducerFile(t *testing.T) {
	var (
		id                 string
		fb                 []byte
		producer           *kafka.Producer
		verifedBiZoneAlert *datamodels.VerifiedBiZoneIRPAlert

		topic string = "object.topicalerttype.test"

		err error
	)

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

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

				if msg.GetType() == "error" {
					log.Fatal(msg.GetMessage())
				}

			}
		}
	}(ctx, logging, counting)

	t.Run("Тест 2. Чтение тестового файла с объектом типа 'alerts'", func(t *testing.T) {
		fb, err = os.ReadFile("../kafka/incoming_alert_1.json")
		//fb, err = os.ReadFile("../kafka/replacing_alert_1.json")
		assert.NoError(t, err)
		assert.NotEqual(t, len(fb), 0)
	})

	t.Run("Тест 3. Генерирование нового объекта типа datamodels.VerifiedBiZoneAlert", func(t *testing.T) {
		decoder := decoderjsondocuments.New(counting, logging)
		id, verifedBiZoneAlert, _ = documentgenerator.BiZoneAlertsGenerator(decoder.Start(fb))
		assert.NotEqual(t, id, "")
		assert.NotEmpty(t, verifedBiZoneAlert)

		//t.Logf("\nID:'%s'\nVerifiedBiZoneAlert:'%#v'\n", id, verifedBiZoneAlert)
	})

	t.Run("Тест 4. Передача полученного объекта VerifiedBiZoneAlert в kafka", func(t *testing.T) {
		producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
			"group.id":          fmt.Sprintf("%s-group", "GCM-test"),
		})
		assert.NoError(t, err)

		// Delivery report handler for produced messages
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

		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: fb,
		}, nil)

		producer.Flush(15 * 1000)
	})
}
