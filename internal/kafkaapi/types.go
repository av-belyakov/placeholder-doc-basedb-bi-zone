package kafkaapi

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
)

// kafkaApiModule настройки для API Kafka
type kafkaApiModule struct {
	counter      interfaces.Counter
	logger       interfaces.Logger
	consumer     *kafka.Consumer
	topics       map[string]string
	settings     kafkaApiSettings
	chFromModule chan SettingsChanOutput
	chToModule   chan SettingsChanInput
}

type kafkaApiSettings struct {
	nameRegionalObject string
	command            string
	host               string
	cachettl           int
	port               int
}

// KafkaApiOptions функциональные опции
type KafkaApiOptions func(*kafkaApiModule) error

// SettingsChanOutput канал вывода данных из модуля
type SettingsChanOutput struct {
	Data        []byte
	TaskId      string
	SubjectType string
}

// SettingsChanInput канал для приема данных в модуль
type SettingsChanInput struct {
	Data    []byte
	Command string
	TaskId  string
	CaseId  string
	RootId  string
}
