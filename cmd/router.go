package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/natsapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewRouter маршрутизатор сообщений внутри приложения
func NewRouter(counter interfaces.Counter, logger interfaces.Logger, settings ApplicationRouterSettings) *ApplicationRouter {
	return &ApplicationRouter{
		counter:        counter,
		logger:         logger,
		chToNatsApi:    settings.ChanToNats,
		chFromNatsApi:  settings.ChanFromNats,
		chToKafkaApi:   settings.ChanToKafka,
		chFromKafkaApi: settings.ChanFromKafka,
		chToDBSApi:     settings.ChanToDBS,
		chFromDBSApi:   settings.ChanFromDBS,
	}
}

func (r *ApplicationRouter) Router(ctx context.Context) {
	go func() {
		decoder := decoderjsondocuments.New(r.counter, r.logger)

		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-r.chFromNatsApi:
				switch msg.SubjectType {
				case "geoip information":
					//передача информации о географическом местоположении ip адресов
					r.chToDBSApi <- databasestorageapi.SettingsChanInput{
						Section: "information handling",
						Command: "add geoip information",
						Data:    msg.Data,
					}

				case "sensor information":
					// местоположении и принадлежности сенсоров
					r.chToDBSApi <- databasestorageapi.SettingsChanInput{
						Section: "information handling",
						Command: "add sensor information",
						Data:    msg.Data,
					}

				default:
					r.logger.Send("error", supportingfunctions.CustomError(errors.New("undefined subscription type")).Error())

				}

			case msg := <-r.chFromKafkaApi:
				switch msg.SubjectType {
				case "alerts":
					go func() {
						id, verifedBiZoneAlert, listRawFields := documentgenerator.BiZoneAlertsGenerator(decoder.Start(msg.Data))

						r.logger.Send("info", fmt.Sprintf("an 'alerts' document has been generated, and the document has been transferred to the database (id document '%s')", id))

						if len(listRawFields) > 0 {
							r.logger.Send("alert_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "id", id))
						}

						//передача объекта в модуль взаимодействия с базой данных для
						//дальнейшей загрузки данных в базу
						r.chToDBSApi <- databasestorageapi.SettingsChanInput{
							Section: "handling bz_alerts",
							Command: "add alerts",
							Data:    verifedBiZoneAlert,
						}
					}()

				case "cases":
					// *************
					// Здесь нужны разные обработчики
					// *************

					//soaralert_raw_fields

					fmt.Println("\treceived message from topic: 'socd-soar-prod-issue-v1' (case)")

					if str, err := supportingfunctions.NewReadReflectJSONSprint(msg.Data); err == nil {
						if str != "" {
							r.logger.Send("events", fmt.Sprintf("\n%s\n", str))
						}
					}

					//для этого типа событий пока нет обработчиков, потому что я
					// пока что не видел события этого типа

				}

			case msg := <-r.chFromDBSApi:
				//пересылаются запросы на установку тега в TheHive, geoip информации
				// информации о место располажения и принадлежности сенсоров
				r.chToNatsApi <- natsapi.SettingsChanInput{
					Command: msg.Command,
					RootId:  msg.Id,
					Data:    msg.Data,
				}
			}
		}
	}()
}
