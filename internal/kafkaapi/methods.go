package kafkaapi

import (
	"context"
	"fmt"
	"maps"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// StartConsumer инициализирует новый модуль-потребитель Kafka,
// при инициализации возращается канал для взаимодействия с модулем, все запросы к модулю выполняются через него
func (api *kafkaApiModule) StartConsumer(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	cfg := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", api.settings.host, api.settings.port),
		"group.id":          fmt.Sprintf("%s-group", api.settings.nameRegionalObject), // Идентификатор группы
		"auto.offset.reset": "earliest",                                               // Читать с начала
	}

switch api.settings.

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
func (api *kafkaApiModule) GetConsumer() *kafka.Consumer {
	return api.consumer
}

// GetTopics топики Kafka
func (api *kafkaApiModule) GetTopics() map[string]string {
	return api.topics
}
