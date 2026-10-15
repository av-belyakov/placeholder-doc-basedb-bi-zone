package kafkaapi

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"maps"
	"os"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

// StartConsumer инициализирует новый модуль-потребитель Kafka,
// при инициализации возращается канал для взаимодействия с модулем, все запросы к модулю выполняются через него
func (api *kafkaApiModule) StartConsumer(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	var topics []string

	mapTopics := maps.Values(api.topics)
	for topic := range mapTopics {
		topics = append(topics, topic)
	}

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true, // использовать или IPv4 или IPv6
	}

	if strings.ToUpper(api.settings.saslMechanism) != "ANONYMOUS" {
		tlsCfg, err := api.createTLSConfig()
		if err != nil {
			return err
		}

		dialer.TLS = tlsCfg
	}

	switch strings.ToUpper(api.settings.saslMechanism) {
	case "PLAIN":
		dialer.SASLMechanism = plain.Mechanism{
			Username: api.settings.sslUsername,
			Password: api.settings.sslPassword,
		}

	case "SCRAM-SHA-256", "SCRAM-SHA-512":
		mechanism, err := scram.Mechanism(scram.SHA512, api.settings.sslUsername, api.settings.sslPassword)
		if err != nil {
			return err
		}

		dialer.SASLMechanism = mechanism
	}

	kafkaReaderCfg := kafka.ReaderConfig{
		Brokers:     []string{fmt.Sprintf("%s:%d", api.settings.host, api.settings.port)},
		GroupID:     api.settings.groupId,
		GroupTopics: topics,
		Dialer:      dialer,
	}

	api.consumer = kafka.NewReader(kafkaReaderCfg)

	api.wg.Add(1)
	//обработчик подписок
	go api.topicsHandler(ctx)

	return nil
}

func (api *kafkaApiModule) StopConsumer(ctx context.Context) error {
	var err error

	api.stopOnce.Do(func() {
		log.Printf("%vkafka consumer stopped%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)

		done := make(chan struct{})
		go func() {
			api.wg.Wait()
			close(done)
		}()

		if api.consumer != nil {
			if err := api.consumer.Close(); err != nil {
				err = fmt.Errorf("consumer Kafka closure error: %w", err)
			}
		}

		select {
		case <-done:
			log.Println("kafka consumer was successfully stopped")

		case <-ctx.Done():
			err = ctx.Err()

		}

		close(api.chFromModule)
	})

	return err
}

// GetChannelFromModule канал для приёма данных из модуля
func (api *kafkaApiModule) GetChannelFromModule() <-chan SettingsChanOutput {
	return api.chFromModule
}

// GetChannelToModule канал для передачи данных в модуль
func (api *kafkaApiModule) GetChannelToModule() chan<- SettingsChanInput {
	return api.chToModule
}

// GetConsumer потребитель для kafka
func (api *kafkaApiModule) GetConsumer() *kafka.Reader {
	return api.consumer
}

// GetTopics топики Kafka
func (api *kafkaApiModule) GetTopics() map[string]string {
	return api.topics
}

// createTLSConfig создает TLS конфигурацию
func (api *kafkaApiModule) createTLSConfig() (*tls.Config, error) {
	cfg := &tls.Config{
		MinVersion:         tls.VersionTLS13,
		ServerName:         api.settings.sslServerName,
		InsecureSkipVerify: true,
	}

	if api.settings.sslCeFile != "" {
		caCert, err := os.ReadFile(api.settings.sslCertFile)
		if err != nil {
			return cfg, fmt.Errorf("reading CA: %w", err)
		}

		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			return cfg, fmt.Errorf("parsing CA")
		}

		cfg.RootCAs = caCertPool
	}

	if api.settings.sslCertFile != "" && api.settings.sslKeyFile != "" {
		cert, err := tls.LoadX509KeyPair(api.settings.sslCertFile, api.settings.sslKeyFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load client certificate: %w", err)
		}

		cfg.Certificates = []tls.Certificate{cert}
	}

	return cfg, nil
}
