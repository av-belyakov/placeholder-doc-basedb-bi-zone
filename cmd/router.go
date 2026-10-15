package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/natsapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewRouter маршрутизатор сообщений внутри приложения
func NewRouter(logger interfaces.Logger, counter interfaces.Counter, settings ApplicationRouterSettings) *ApplicationRouter {
	return &ApplicationRouter{
		logger:         logger,
		counter:        counter,
		chToNatsApi:    settings.ChanToNats,
		chFromNatsApi:  settings.ChanFromNats,
		chToKafkaApi:   settings.ChanToKafka,
		chFromKafkaApi: settings.ChanFromKafka,
		chToDBSApi:     settings.ChanToDBS,
		chFromDBSApi:   settings.ChanFromDBS,
	}
}

func (r *ApplicationRouter) Start(ctx context.Context) {
	go func() {
		decoder := decoderjsondocuments.New(r.counter, r.logger)

		for {
			select {
			case <-ctx.Done():
				log.Printf("%vthe major router is stopped%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)

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
						id, verifedBiZoneAlert, listRawFields, err := documentgenerator.BiZoneAlertsGenerator(decoder.Start(msg.Data))
						if err != nil {
							r.logger.Send("error", supportingfunctions.CustomError(err).Error())
						}

						r.logger.Send("info", fmt.Sprintf("an 'ALERT' document has been generated, and the document has been transferred to the database (uuid document '%s')", id))

						if len(listRawFields) > 0 {
							r.logger.Send("alert_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "id", id))
						}

						//передача объекта в модуль взаимодействия с базой данных для
						//дальнейшей загрузки данных в базу
						r.chToDBSApi <- databasestorageapi.SettingsChanInput{
							Section: "handling bizone alerts",
							Command: "add alerts",
							Data:    verifedBiZoneAlert,
						}
					}()

				case "cases":
					go func() {
						id, verifedBiZoneCase, listRawFields, err := documentgenerator.BiZoneCasesGenerator(decoder.Start(msg.Data))
						if err != nil {
							r.logger.Send("error", supportingfunctions.CustomError(err).Error())
						}

						r.logger.Send("info", fmt.Sprintf("an 'CASE' document has been generated, and the document has been transferred to the database (uuid document '%s')", id))

						if len(listRawFields) > 0 {
							r.logger.Send("case_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "id", id))
						}

						//передача объекта в модуль взаимодействия с базой данных для
						//дальнейшей загрузки данных в базу
						r.chToDBSApi <- databasestorageapi.SettingsChanInput{
							Section: "handling bizone cases",
							Command: "add cases",
							Data:    verifedBiZoneCase,
						}
					}()
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
