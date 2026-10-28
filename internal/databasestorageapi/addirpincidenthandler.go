package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// addIRPIncident добавление объекта типа 'incident'
func (dbs *DatabaseStorage) addIRPIncident(ctx context.Context, a any) {
	t := time.Now()

	newDocument, ok := a.(*datamodels.VerifiedBiZoneIRPIncident)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("note for type 'incident', type conversion error")).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("database storage - note for type 'incident', accepted document with key (id):'%s'", newDocument.Key))

	binaryDocument, err := json.Marshal(newDocument.Get())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	// получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["incidents"]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("database storage - note for type 'incidents', the identifier of the index name was not found")).Error())

		return
	}

	// получаем существующие индексы
	existingIndexes, err := dbs.GetExistingIndexes(ctx, indexName)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	// формируем наименование индекса
	currentIndex := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), int(t.Month()))

	// будет выполнятся поиск по индексам только в текущем году так как при
	// накоплении большого количества индексов, поиск по всем серьезно замедлит работу
	indexesOnlyCurrentYear := []string(nil)
	for _, v := range existingIndexes {
		if strings.Contains(v, fmt.Sprint(t.Year())) {
			indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
		}
	}

	defer func(document *datamodels.VerifiedBiZoneIRPIncident, getChan func() chan SettingsChanOutput, logger interfaces.Logger) {
		// уникальный идентификатор для поиска данных в БД
		id := fmt.Sprintf("%s:%s", currentIndex, newDocument.GetSpecialUUID())

		// обогащение кейса дополнительной информацией о локальном место положении ip адресов
		listIp := documentgenerator.GetListIPAddr(document.GetAdditionalInformation().GetIpAddressesInformation())
		ok, err := sendGeoIpRequest(id, listIp, getChan)
		if err != nil {
			logger.Send("error", supportingfunctions.CustomError(err).Error())
		}
		if ok {
			logger.Send("info", fmt.Sprintf("database storage - note for type 'incidents', we are sending a request to search for information on the following list of ip addresses: %#v", listIp))
		}

		// обогащение кейса дополнительной информацией по расположению сенсоров
		listSensor := documentgenerator.GetListSensorId(newDocument.GetAdditionalInformation().GetSensorsInformation())
		ok, err = sendSensorInformationRequest(id, listSensor, getChan)
		if err != nil {
			logger.Send("error", supportingfunctions.CustomError(err).Error())
		}
		if ok {
			logger.Send("info", fmt.Sprintf("database storage - note for type 'incidents', we are sending a request to search for information on the following list of sensors: %#v", listSensor))
		}
	}(newDocument, dbs.GetChannelFromModule, dbs.logger)

	// если похожих индексов нет
	if len(indexesOnlyCurrentYear) == 0 {
		//
		// вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, binaryDocument)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		existingIndexes = append(existingIndexes, currentIndex)
		// устанавливаем максимальный лимит количества полей для всех индексов которые
		// содержат значение по умолчанию в 1000 полей
		if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}

		// счетчик
		dbs.counter.SendMessage("update count insert subject incidents to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("database storage - note for type 'incidents', insert new document to incidents key (id) '%s', status code '%d'", newDocument.GetKey(), statusCode))

		return
	}

	// устанавливаем максимальный лимит количества полей для всех индексов которые
	// содержат значение по умолчанию в 1000 полей
	if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	}

	// ищем объект с таким же идентификатором как и принятый в обработку объект
	res, err := dbs.GetDocument(
		ctx,
		indexesOnlyCurrentYear,
		strings.NewReader(
			fmt.Sprintf(
				`{
					"query": {
						"bool": {
							"must": [{"match": {"data.key": "%s"}}]}}}`,
				newDocument.GetKey())),
		0)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	response := ResponseVerifiedIRPIncidents{}
	if err = json.Unmarshal(res, &response); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	fmt.Printf("--------========= DatabaseStorage.addIRPIncident, ResponseVerifiedIRPIncidents:'%+v'\n", response)

	// вставка выполняется только когда не найден искомый документ
	if response.Options.Total.Value == 0 {
		//
		// вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, binaryDocument)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		// счетчик
		dbs.counter.SendMessage("update count insert subject incident to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("database storage - note for type 'incident', insert new document to incident key (id) '%s', status code '%d'", newDocument.GetKey(), statusCode))

		//*******************************************
		//*** здесь установка тегов, под вопросом ***
		//*******************************************

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	listDeleting := []ServiseOption(nil)
	updateVerified := datamodels.NewVerifiedBiZoneIRPIncident()
	// заполняем новый объект информацией из БАЗЫ ДАННЫХ
	for _, v := range response.Options.Hits {
		updateVerified.ReplacingOldIRPIncident(*v.Source.Get())
		listDeleting = append(listDeleting, ServiseOption{
			ID:    v.ID,
			Index: v.Index,
		})

		// заполняем новый объект дополнительной информацией о сенсорах и месторасположении ip адресов
		updateVerified.SetAdditionalInformation(*v.Source.GetAdditionalInformation())
	}

	// выполняем обновление нового объекта данными полученными от IRP
	countReplasedField := updateVerified.ReplacingOldIRPIncident(*newDocument.Get())

	// готовим обновлённый объект к загрузке в базу данных
	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	// обновление в базе данных уже существующего документа
	statusCode, countDel, err := dbs.UpdateDocument(ctx, currentIndex, listDeleting, nvbyte)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("key (id) '%s' '%s'", newDocument.GetKey(), err.Error())).Error())

		return
	}

	dbs.counter.SendMessage("update count insert subject incident to db", 1)
	dbs.logger.Send(
		"info",
		fmt.Sprintf(
			"database storage - note for type 'incident', update document 'incident' type, count delete '%d', count replacing fields '%d' for object with key (id) '%s', status code '%d'",
			countDel,
			countReplasedField,
			newDocument.GetKey(),
			statusCode,
		))
}

// TestAddIRPIncident использовать только в тестах
func (dbs *DatabaseStorage) TestAddIRPIncident(ctx context.Context, a any) {
	dbs.addIRPIncident(ctx, a)
}
