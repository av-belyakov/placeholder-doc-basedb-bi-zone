package kafkaapi

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// topicsHandler обработчик топиков (подписок)
func (api *kafkaApiModule) topicsHandler(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			msg, err := api.consumer.ReadMessage(ctx)
			if err != nil {
				if !err.(kafka.Error).IsTimeout() {
					api.logger.Send("error", supportingfunctions.CustomError(err).Error())

					continue
				}
			}

			subjectType := "undefined_type"
			topicKey, ok := supportingfunctions.SearchValue(api.topics, msg.Topic)
			if ok {
				subjectType = topicKey
			}

			api.chFromModule <- SettingsChanOutput{
				SubjectType: subjectType,
				Data:        msg.Value,
			}
		}
	}
}
