// Модуль для взаимодействия с API IRP
package irpapi

import (
	"errors"
	"net/http"
	"time"
)

// New настраивает модуль взаимодействия с API IRP
func New(opts ...options) (*IrpApi, error) {
	api := &IrpApi{
		settings: ApiSettings{
			connectionTimeout: 10,
		},
	}

	for _, opt := range opts {
		if err := opt(api); err != nil {
			return api, err
		}
	}

	api.client = &http.Client{
		Timeout: time.Duration(api.settings.connectionTimeout) * time.Second,
	}

	return api, nil
}

// WithAPIKey идентификатор-ключ для API
func WithAPIKey(v string) options {
	return func(th *IrpApi) error {
		if v == "" {
			return errors.New("the value of 'apiKey' cannot be empty")
		}

		th.settings.apiKey = v

		return nil
	}
}

// WithHost имя или ip адрес хоста API
func WithHost(v string) options {
	return func(th *IrpApi) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		th.settings.host = v

		return nil
	}
}

// WithPort сетевой порт API
func WithPort(v int) options {
	return func(th *IrpApi) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		th.settings.port = v

		return nil
	}
}

// WithConnectionTimeout время ожидания до разрыва сетевого соединения с API IRP
// в секундах в диапазоне от 1 секунды до 900 (15 минут)
func WithConnectionTimeout(v int) options {
	return func(th *IrpApi) error {
		if v <= 0 || v > 900 {
			return errors.New("an incorrect network port value was received")
		}

		th.settings.connectionTimeout = v

		return nil
	}
}
