package kafkaapi

import (
	"errors"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
)

// New настраивает модуль взаимодействия с API Kafka
func New(logger interfaces.Logger, counter interfaces.Counter, opts ...KafkaApiOptions) (*kafkaApiModule, error) {
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
	return func(api *kafkaApiModule) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		api.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		api.settings.port = v

		return nil
	}
}

// WithCacheTTL время жизни для кэша хранящего функции-обработчики запросов к модулю
func WithCacheTTL(v int) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		if v <= 10 || v > 86400 {
			return errors.New("the lifetime of a cache entry should be between 10 and 86400 seconds")
		}

		api.settings.cachettl = v

		return nil
	}
}

// WithNameRegionalObject наименование которое будет отображатся в статистике подключений NATS
func WithNameRegionalObject(v string) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		api.settings.nameRegionalObject = v

		return nil
	}
}

// WithTopicsSubscription 'слушатель' разных топиков
func WithTopicsSubscription(v map[string]string) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		if len(v) == 0 {
			return errors.New("the value of 'topics' cannot be empty")
		}

		api.topics = v

		return nil
	}
}

// WithClientName имя клиента
func WithClientName(v string) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		api.settings.userName = v

		return nil
	}
}

// WithPassword пароль для доступа к Kafka
func WithPassword(v string) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		api.settings.passwd = v

		return nil
	}
}

// WithCertPath путь к сертификату
func WithCertPath(v string) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		api.settings.certPath = v

		return nil
	}
}

// WithTruststoragePath путь к ключу доверенного хранилища
func WithTruststoragePath(v string) KafkaApiOptions {
	return func(api *kafkaApiModule) error {
		api.settings.truststorePath = v

		return nil
	}
}
