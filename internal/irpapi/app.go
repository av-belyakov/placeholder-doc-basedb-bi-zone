// Модуль для взаимодействия с API IRP
package irpapi

import (
	"errors"
)

// New настраивает модуль взаимодействия с API IRP
func New(opts ...options) (*IrpApi, error) {
	api := &IrpApi{
		settings: ApiSettings{},
	}

	for _, opt := range opts {
		if err := opt(api); err != nil {
			return api, err
		}
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
