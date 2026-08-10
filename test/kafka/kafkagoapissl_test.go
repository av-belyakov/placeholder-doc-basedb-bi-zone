package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"maps"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
)

func TestKafkaGoApiSsl(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "test")

	cfg, err := confighandler.New(constants.Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

	mechanism, err := scram.Mechanism(scram.SHA512, cfg.Kafka.SSLUsername, cfg.Kafka.SSLPassword)
	if err != nil {
		log.Fatalln(err)
	}

	var topics []string
	mapTopics := maps.Values(cfg.Kafka.Topics)
	for topic := range mapTopics {
		topics = append(topics, topic)
	}

	publicCert, err := os.ReadFile("../../secrets/ca.crt")
	if err != nil {
		log.Fatalln(err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(publicCert) {
		log.Fatalln("certificate CA parsing error")
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port)},
		GroupID:     "", //fmt.Sprintf("%s-group", cfg.Common.RegionalObject),
		GroupTopics: topics,
		Dialer: &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true, // уточнить для чего
			SASLMechanism: mechanism,
			TLS: &tls.Config{
				RootCAs:            caCertPool,
				MinVersion:         tls.VersionTLS13,
				ServerName:         "kafka-cluster-kafka-bootstrap",
				InsecureSkipVerify: false, // если true то проверка сертификата не выполняется
			},
		},
	})

	msg, err := r.ReadMessage(t.Context())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Topic name:%s\n, Data:%s\n", msg.Topic, string(msg.Value))

	t.Cleanup(func() {
		r.Close()
	})
}
