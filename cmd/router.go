package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/dicontainer"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/natsapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/storageobjects"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewRouter маршрутизатор сообщений внутри приложения
func NewRouter(dic *dicontainer.DiContainer) *ApplicationRouter {
	return &ApplicationRouter{
		diContainer: dic,
	}
}

func (r *ApplicationRouter) Start(ctx context.Context) error {
	counter := r.diContainer.Counter(ctx)
	logger := r.diContainer.Logger(ctx)

	chToNats := r.diContainer.NatsConnecter(ctx).GetChannelToModule()
	chFromNats := r.diContainer.NatsConnecter(ctx).GetChannelFromModule()
	chFromKafka := r.diContainer.KafkaConnecter(ctx).GetChannelFromModule()
	chToDB := r.diContainer.DbConnecter(ctx).GetChannelToModule()
	chFromDB := r.diContainer.DbConnecter(ctx).GetChannelFromModule()

	//очередь для хранения объектов с отложенной отправкой в канал
	storage, err := storageobjects.New(
		storageobjects.WithChannelSize[any](4),
		storageobjects.WithTimeTick[any](1),
		storageobjects.WithTimeToLive[any](r.diContainer.Configer().GetTemporaryStorage().StorageObjectTTL),
	)
	if err != nil {
		return supportingfunctions.CustomError(err)
	}

	storage.Start(ctx)

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
					/*
							!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
							Обработка алертов пока выключена

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
					*/

				case "incidents":
					data, err := handlers.UnmarshalIncidentDocument(msg.Data)
					if err != nil {
						logger.Send("error", fmt.Sprintf("unpacking a json incident document - %s", err.Error()))

						return
					}

					// сохраняем объект для отложеного выполнения
					// ожидание истечения времени из параметра storageObject[n].timeSending
					storage.AddObject(
						r.diContainer.Configer().GetTemporaryStorage().StorageDelayToSendObject,
						storageobjects.StorageObjectDataSettings[any]{
							Id:         data.Data.Id,
							ObjectType: data.Type,
							Data:       data,
						})
				}

			case data := <-storage.GetObjects():
				// пока не буду разделять инциденты по их типам data.Type:
				// soar.core.incident.created.v1
				// case "soar.core.incident.changed.v1":
				go func() {
					startTime := time.Now()
					// делаем запроос к IRP что бы получить дополнительную информации по инциденту
					response, statusCode, err := r.diContainer.IRPConnecter().GetIncident(ctx, data.Id)
					if err != nil {
						logger.Send("error", fmt.Sprintf("incident request id '%s' from the IRP - %s", data.Id, err.Error()))

						return
					}

					logger.Send("info", fmt.Sprintf("the request to irp to receive the full incident '%s' was completed in %dms", data.Id, (time.Since(startTime)/time.Millisecond)))

					if statusCode != http.StatusOK {
						logger.Send("warning", fmt.Sprintf("requesting incident '%s' from the IRP returned the status code %d", data.Id, statusCode))

						return
					}

					// генерируем проверенную структуру
					_, verifiedBiZoneIncident, listRawFields, err := documentgenerator.GeneratingVerifiedDocumentFromIncident(decoder.Start(response))
					if err != nil {
						logger.Send("error", supportingfunctions.CustomError(err).Error())
					}

					logger.Send("info", fmt.Sprintf("the incident '%s' document was successfully generated and transferred to the database", data.Id))

					if len(listRawFields) > 0 {
						logger.Send("incident_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "incident key (id)", data.Id))
					}

					// передача объекта в модуль взаимодействия с базой данных для дальнейшей загрузки данных в базу
					chToDB <- databasestorageapi.SettingsChanInput{
						Section: "handling irp incident",
						Command: "add incident",
						Data:    verifiedBiZoneIncident,
					}
				}()

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

	return nil
}
