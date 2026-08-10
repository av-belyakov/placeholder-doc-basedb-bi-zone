package kafkaapi

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"maps"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

// StartConsumer инициализирует новый модуль-потребитель Kafka,
// при инициализации возращается канал для взаимодействия с модулем, все запросы к модулю выполняются через него
func (api *kafkaApiModule) StartConsumer(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	var (
		dialer    *kafka.Dialer
		tlsCfg    *tls.Config
		mechanism sasl.Mechanism
		topics    []string
		err       error
	)

	mapTopics := maps.Values(api.topics)
	for topic := range mapTopics {
		topics = append(topics, topic)
	}

	tlsCfg, err = api.createTLSConfig()
	if err != nil {
		return err
	}

	switch strings.ToUpper(api.settings.saslMechanism) {
	case "PLAIN":
		mechanism = plain.Mechanism{
			Username: api.settings.sslUsername,
			Password: api.settings.sslPassword,
		}

	case "SCRAM-SHA-256", "SCRAM-SHA-512":
		mechanism, err = scram.Mechanism(scram.SHA512, api.settings.sslUsername, api.settings.sslPassword)
		if err != nil {
			return err
		}
	}

	dialer = &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true, // использовать или IPv4 или IPv6
		SASLMechanism: mechanism,
		TLS:           tlsCfg,
	}

	api.consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{fmt.Sprintf("%s:%d", api.settings.host, api.settings.port)},
		GroupID:     api.settings.groupId,
		GroupTopics: topics,
		Dialer:      dialer,
	})

	//обработчик подписок
	go api.topicsHandler(ctx)

	return nil
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
