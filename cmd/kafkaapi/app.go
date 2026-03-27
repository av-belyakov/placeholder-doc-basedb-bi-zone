package kafkaapi

import (
	"errors"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
)

// New настраивает модуль взаимодействия с API Kafka
func New(counter interfaces.Counter, logger interfaces.Logger, opts ...KafkaApiOptions) (*kafkaApiModule, error) {
	api := &kafkaApiModule{
		counter: counter,
		logger:  logger,
		settings: kafkaApiSettings{
			cachettl: 15,
		},
		chFromModule: make(chan SettingsChanOutput),
		chToModule:   make(chan SettingsChanInput),
	}

	for _, opt := range opts {
		if err := opt(api); err != nil {
			return api, err
		}
	}

	return api, nil
}

// WithHost имя или ip адрес хоста API
func WithHost(v string) KafkaApiOptions {
	return func(n *kafkaApiModule) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		n.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) KafkaApiOptions {
	return func(n *kafkaApiModule) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		n.settings.port = v

		return nil
	}
}

// WithCacheTTL время жизни для кэша хранящего функции-обработчики запросов к модулю
func WithCacheTTL(v int) KafkaApiOptions {
	return func(th *kafkaApiModule) error {
		if v <= 10 || v > 86400 {
			return errors.New("the lifetime of a cache entry should be between 10 and 86400 seconds")
		}

		th.settings.cachettl = v

		return nil
	}
}

// WithNameRegionalObject наименование которое будет отображатся в статистике подключений NATS
func WithNameRegionalObject(v string) KafkaApiOptions {
	return func(n *kafkaApiModule) error {
		n.settings.nameRegionalObject = v

		return nil
	}
}

// WithTopicsSubscription 'слушатель' разных топиков
func WithTopicsSubscription(v map[string]string) KafkaApiOptions {
	return func(n *kafkaApiModule) error {
		if len(v) == 0 {
			return errors.New("the value of 'topics' cannot be empty")
		}

		n.topics = v

		return nil
	}
}
