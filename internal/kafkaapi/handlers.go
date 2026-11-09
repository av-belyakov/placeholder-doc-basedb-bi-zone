package kafkaapi

import (
	"context"
	"fmt"
	"log"
	"time"

	//"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// topicsHandler обработчик топиков (подписок)
func (api *kafkaApiModule) topicsHandler(ctx context.Context) {
	defer api.wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Printf("%vthe kafka module, topics handler is stopped%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)

			return

		default:
			// таймаут для чтения, чтобы периодически проверять контекст
			readCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
			msg, err := api.consumer.FetchMessage(readCtx)
			cancel()

			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded {
					// отмена контекста, проверяем основной контекст
					if ctx.Err() != nil {
						return
					}

					continue
				}

				api.logger.Send("error", fmt.Errorf("kafka reading message:%w", err).Error())

				continue
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

			// подтверждение получения сообщения (commit offset)
			if err := api.consumer.CommitMessages(ctx, msg); err != nil {
				api.logger.Send("error", fmt.Errorf("kafka commiting message:%w", err).Error())
			}
		}
	}
}
