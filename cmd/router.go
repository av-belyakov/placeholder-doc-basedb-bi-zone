package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/dicontainer"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/natsapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewRouter маршрутизатор сообщений внутри приложения
func NewRouter(dic *dicontainer.DiContainer) *ApplicationRouter {
	return &ApplicationRouter{
		diContainer: dic,
	}
}

func (r *ApplicationRouter) Start(ctx context.Context) {
	counter := r.diContainer.Counter(ctx)
	logger := r.diContainer.Logger(ctx)

	chToNats := r.diContainer.NatsConnecter(ctx).GetChannelToModule()
	chFromNats := r.diContainer.NatsConnecter(ctx).GetChannelFromModule()
	chFromKafka := r.diContainer.KafkaConnecter(ctx).GetChannelFromModule()
	chToDB := r.diContainer.DbConnecter(ctx).GetChannelToModule()
	chFromDB := r.diContainer.DbConnecter(ctx).GetChannelFromModule()

	go func() {
		decoder := decoderjsondocuments.New(counter, logger)

		for {
			select {
			case <-ctx.Done():
				log.Printf("%vthe major router is stopped%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)

				return

			case msg := <-chFromNats:
				switch msg.SubjectType {
				case "geoip information":
					// передача информации о географическом местоположении ip адресов
					chToDB <- databasestorageapi.SettingsChanInput{
						Section: "information handling",
						Command: "add geoip information",
						Data:    msg.Data,
					}

				case "sensor information":
					// местоположении и принадлежности сенсоров
					chToDB <- databasestorageapi.SettingsChanInput{
						Section: "information handling",
						Command: "add sensor information",
						Data:    msg.Data,
					}

				default:
					logger.Send("error", supportingfunctions.CustomError(errors.New("undefined subscription type")).Error())

				}

			case msg := <-chFromKafka:
				switch msg.SubjectType {
				case "alerts":
					go func() {
						id, verifedBiZoneAlert, listRawFields, err := documentgenerator.BiZoneAlertsGenerator(decoder.Start(msg.Data))
						if err != nil {
							logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("generating a document of type irp bizone alert - %w", err)).Error())
						}

						logger.Send("info", fmt.Sprintf("an 'ALERT' document has been generated, and the document has been transferred to the database (uuid document '%s')", id))

						if len(listRawFields) > 0 {
							logger.Send("alert_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "id", id))
						}

						// передача объекта в модуль взаимодействия с базой данных для дальнейшей загрузки данных в базу
						chToDB <- databasestorageapi.SettingsChanInput{
							Section: "handling bizone alerts",
							Command: "add alerts",
							Data:    verifedBiZoneAlert,
						}
					}()

				case "incidents":
					data, err := handlers.GetTypeIncidentDocument(msg.Data)
					if err != nil {
						logger.Send("error", fmt.Sprintf("unpacking a json incident document - %s", err.Error()))

						return
					}

					switch data.Type {
					case "soar.core.incident.created.v1":
						go func() {
							// делаем запроос к IRP что бы получить дополнительную информации по инциденту
							response, statusCode, err := r.diContainer.IRPConnecter().GetIncident(ctx, data.Id)
							if err != nil {
								logger.Send("error", fmt.Sprintf("incident request id '%s' from the IRP - %s", data.Id, err.Error()))

								return
							}

							if statusCode != http.StatusOK {
								logger.Send("warning", fmt.Sprintf("requesting incident id '%s' from the IRP returned the status code %s", data.Id, statusCode))

								return
							}

							// генерируем проверенную структуру
							generatedUuid, verifiedBiZoneIncident, listRawFields, err := documentgenerator.GeneratingVerifiedDocumentFromIncident(decoder.Start(response))
							if err != nil {
								logger.Send("error", supportingfunctions.CustomError(err).Error())
							}

							logger.Send("info", fmt.Sprintf("an 'INCIDENT' document has been generated, and the document has been transferred to the database (generated uuid document '%s')", generatedUuid))

							if len(listRawFields) > 0 {
								logger.Send("incident_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "id", generatedUuid))
							}

							// передача объекта в модуль взаимодействия с базой данных для дальнейшей загрузки данных в базу
							chToDB <- databasestorageapi.SettingsChanInput{
								Section: "handling irp incident",
								Command: "add incident",
								Data:    verifiedBiZoneIncident,
							}
						}()

					case "soar.core.incident.changed.v1":
						log.Printf("Получена информация об изменении инцидента '%s'\n", data.Id)
					}
					/*
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
					*/
				}

			case msg := <-chFromDB:
				// пересылаются запросы на установку тега в TheHive, geoip информации
				// информации о место располажения и принадлежности сенсоров
				chToNats <- natsapi.SettingsChanInput{
					Command: msg.Command,
					RootId:  msg.Id,
					Data:    msg.Data,
				}
			}
		}
	}()
}
