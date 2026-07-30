// Модуль для взаимодействия с API NATS
package natsapi

import (
	"errors"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
)

// New настраивает новый модуль взаимодействия с API NATS
func New(logger interfaces.Logger, counter interfaces.Counter, opts ...NatsApiOptions) (*apiNatsModule, error) {
	api := &apiNatsModule{
		settings: apiNatsSettings{
			cachettl: 10,
		},
		//для подсчёта
		counter: counter,
		//для логирования
		logger: logger,
		//запросы в модуль
		chFromModule: make(chan SettingsChanOutput),
		//события из модуля
		chToModule: make(chan SettingsChanInput),
	}

	for _, opt := range opts {
		if err := opt(api); err != nil {
			return api, err
		}
	}

	return api, nil
}

// WithHost имя или ip адрес хоста API
func WithHost(v string) NatsApiOptions {
	return func(n *apiNatsModule) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		n.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) NatsApiOptions {
	return func(n *apiNatsModule) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		n.settings.port = v

		return nil
	}
}

// WithCacheTTL время жизни для кэша хранящего функции-обработчики запросов к модулю
func WithCacheTTL(v int) NatsApiOptions {
	return func(th *apiNatsModule) error {
		if v <= 10 || v > 86400 {
			return errors.New("the lifetime of a cache entry should be between 10 and 86400 seconds")
		}

		th.settings.cachettl = v

		return nil
	}
}

// WithNameRegionalObject наименование которое будет отображатся в статистике подключений NATS
func WithNameRegionalObject(v string) NatsApiOptions {
	return func(n *apiNatsModule) error {
		n.settings.nameRegionalObject = v

		return nil
	}
}

// WithRequests запросы к сторонним сервисам
func WithRequests(v map[string]string) NatsApiOptions {
	return func(n *apiNatsModule) error {
		if len(v) == 0 {
			return errors.New("the value of 'requests' cannot be empty")
		}

		n.requests = v

		return nil
	}
}

// WithSubscriptions 'слушатель' разных типов сообщений
func WithSubscriptions(v map[string]string) NatsApiOptions {
	return func(n *apiNatsModule) error {
		if len(v) == 0 {
			return errors.New("the value of 'subscriptions' cannot be empty")
		}

		n.subscriptions = v

		return nil
	}
}

// WithSendCommand команду отправляемая в NATS
func WithSendCommand(v string) NatsApiOptions {
	return func(n *apiNatsModule) error {
		if v == "" {
			return errors.New("the value of 'command' cannot be empty")
		}

		n.settings.command = v

		return nil
	}
}
