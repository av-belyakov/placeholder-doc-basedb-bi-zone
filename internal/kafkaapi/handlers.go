package kafkaapi

import (
	"context"
	"fmt"
	"time"

	//"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// topicsHandler обработчик топиков (подписок)
func (api *kafkaApiModule) topicsHandler(ctx context.Context) {
	defer api.wg.Done()

	fmt.Println("kafkaApiModule.topicsHandler START")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("---- Stop Kafka handler ----")
			api.logger.Send("info", "kafka handler stopped")

			return

		default:
			fmt.Println("kafkaApiModule.topicsHandler default value")

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

			fmt.Println("kafkaApiModule.topicsHandler After send message to channel")

			// подтверждение получения сообщения (commit offset)
			if err := api.consumer.CommitMessages(ctx, msg); err != nil {
				api.logger.Send("error", fmt.Errorf("kafka commiting message:%w", err).Error())
			}

			/*msg, err := api.consumer.ReadMessage(ctx)
			if err != nil {
				if kafkaErr, ok := err.(kafka.Error); ok {
					if kafkaErr.IsTimeout() {
						errTimeout := errors.New("timeout error")
						errors.Join(errTimeout, err)
						api.logger.Send("error", supportingfunctions.CustomError(errTimeout).Error())
					}
				}

				api.logger.Send("error", supportingfunctions.CustomError(err).Error())

				continue
			}*/
		}
	}
}
