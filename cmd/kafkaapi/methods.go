package kafkaapi

import (
	"context"
	"fmt"
	"maps"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Start инициализирует новый модуль взаимодействия с API Kafka,
// при инициализации возращается канал для взаимодействия с модулем,
// все запросы к модулю выполняются через данный канал
func (api *kafkaApiModule) Start(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", api.settings.host, api.settings.port),
		"group.id":          fmt.Sprintf("%s-group", api.settings.nameRegionalObject), // Идентификатор группы
		"auto.offset.reset": "earliest",                                               // Читать с начала
	})
	if err != nil {
		return err
	}
	api.consumer = consumer
	context.AfterFunc(ctx, func() {
		consumer.Close()
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

// GetChanDataToModule канал для передачи данных в модуль
func (api *kafkaApiModule) GetChanDataToModule() chan SettingsChanInput {
	return api.chToModule
}

// GetChanDataFromModule канал для приёма данных из модуля
func (api *kafkaApiModule) GetChanDataFromModule() chan SettingsChanOutput {
	return api.chFromModule
}
