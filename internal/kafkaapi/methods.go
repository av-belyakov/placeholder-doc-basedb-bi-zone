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

	var dialer *kafka.Dialer
	switch strings.ToLower(api.settings.authType) {
	case "ssl":
	case "sasl-ssl":
		mechanism, err := scram.Mechanism(scram.SHA512, api.settings.sslUsername, api.settings.sslPassword)
		if err != nil {
			return err
		}

		dialer = &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true, // уточнить для чего
			SASLMechanism: mechanism,
			TLS:           &tls.Config{},
		}
	}

	if api.settings.sslCertFile != "" {

	}

	if api.settings.sslKeyFile != "" {

	}

	api.consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{fmt.Sprintf("%s:%d", api.settings.host, api.settings.port)},
		GroupID:     fmt.Sprintf("%s-group", api.settings.nameRegionalObject),
		GroupTopics: topics,
		Dialer:      dialer,
	})

	/*
		cfg := &kafka.ConfigMap{
			"bootstrap.servers":     fmt.Sprintf("%s:%d", api.settings.host, api.settings.port),
			"group.id":              fmt.Sprintf("%s-group", api.settings.nameRegionalObject), // Идентификатор группы
			"auto.offset.reset":     "earliest",                                               // Читать с начала
			"enable.auto.commit":    false,                                                    // не проверял
			"heartbeat.interval.ms": 3000,                                                     // не проверял
			"max.poll.interval.ms":  300000,                                                   // не проверял
		}

		switch strings.ToLower(api.settings.authType) {
		case "ssl":
			cfg.SetKey("security.protocol", "SSL")
			cfg.SetKey("ssl.ca.location", api.settings.sslCeFile)
			cfg.SetKey("ssl.certificate.location", api.settings.sslCertFile)
			cfg.SetKey("ssl.key.location", api.settings.sslKeyFile)
			cfg.SetKey("ssl.endpoint.identification.algorithm", "https")
		case "sasl-ssl":
			cfg.SetKey("security.protocol", "SASL_SSL")
			cfg.SetKey("sasl.mechanisms", api.settings.saslMechanism)
			cfg.SetKey("sasl.username", api.settings.sslUsername)
			cfg.SetKey("sasl.password", api.settings.sslPassword)
			cfg.SetKey("ssl.ca.location", filepath.Join("../../", api.settings.sslCeFile))
			cfg.SetKey("ssl.certificate.location", api.settings.sslCertFile)
			cfg.SetKey("ssl.endpoint.identification.algorithm", "none")
		}

		if api.settings.sslCertFile != "" {
			cfg.SetKey("ssl.certificate.location", filepath.Join("../../", api.settings.sslCertFile))
		}

		if api.settings.sslKeyFile != "" {
			cfg.SetKey("ssl.key.location", filepath.Join("../../", api.settings.sslKeyFile))
		}

		consumer, err := kafka.NewConsumer(cfg)
		if err != nil {
			return err
		}
		api.consumer = consumer

		context.AfterFunc(ctx, func() {
			consumer.Close()

			close(api.chToModule)
			close(api.chFromModule)
		})

		var topics []string
		mapTopics := maps.Values(api.topics)
		for topic := range mapTopics {
			topics = append(topics, topic)
		}

		// подписка на топик
		err = api.consumer.SubscribeTopics(topics, nil)
		if err != nil {
			return err
		}
	*/

	//обработчик подписок
	go api.topicsHandler(ctx)

	return nil
}

// createTLSConfig создает TLS конфигурацию
func (api *kafkaApiModule) createTLSConfig() (*tls.Config, error) {
	if api.settings.sslCertFile == "" {
		return nil, fmt.Errorf("the path to the CA certificate is not specified")
	}

	caCert, err := os.ReadFile(api.settings.sslCertFile)
	if err != nil {
		return nil, fmt.Errorf("reading CA: %w", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("parsing CA")
	}

	return &tls.Config{
		RootCAs:            caCertPool,
		MinVersion:         tls.VersionTLS13,
		ServerName:         "kafka-cluster-kafka-bootstrap", /// Это надо добавить в конфиг!!!!!!
		InsecureSkipVerify: false,
	}, nil
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
