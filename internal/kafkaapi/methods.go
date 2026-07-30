package kafkaapi

import (
	"context"
)

// Start инициализирует новый модуль взаимодействия с API Kafka,
// при инициализации возращается канал для взаимодействия с модулем,
// все запросы к модулю выполняются через данный канал
func (api *kafkaApiModule) Start(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	//
	// сделать настроку для kafka с сертификатом и авторизацией
	//
}

// GetChannelFromModule канал для приёма данных из модуля
func (api *kafkaApiModule) GetChannelFromModule() <-chan SettingsChanOutput {
	return api.chFromModule
}

// GetChannelToModule канал для передачи данных в модуль
func (api *kafkaApiModule) GetChannelToModule() chan<- SettingsChanInput {
	return api.chToModule
}
