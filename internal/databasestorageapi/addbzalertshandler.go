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

// addBiZoneAlerts добавление объекта типа 'alerts'
func (dbs *DatabaseStorage) addBiZoneAlerts(ctx context.Context, a any) {
	t := time.Now()

	newDocument, ok := a.(*datamodels.VerifiedBiZoneIRPAlert)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	binaryDocument, err := json.Marshal(newDocument.Get())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["alerts"]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//получаем существующие индексы
	existingIndexes, err := dbs.GetExistingIndexes(ctx, indexName)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	defer func(document *datamodels.VerifiedBiZoneIRPAlert, getChan func() chan SettingsChanOutput, logger interfaces.Logger) {
		id := fmt.Sprintf("alerts:%s", newDocument.GetSpecialUUID())

		//обогащение кейса дополнительной информацией о локальном место положении ip адресов
		listIp := documentgenerator.GetListIPAddr(document.GetAdditionalInformation().GetIpAddressesInformation())
		ok, err := sendGeoIpRequest(id, listIp, getChan)
		if err != nil {
			logger.Send("error", supportingfunctions.CustomError(err).Error())
		}
		if ok {
			logger.Send("info", fmt.Sprintf("we are sending a request to search for information on the following list of ip addresses: %#v", listIp))
		}

		//обогащение кейса дополнительной информацией по расположению сенсоров
		listSensor := documentgenerator.GetListSensorId(newDocument.GetAdditionalInformation().GetSensorsInformation())
		ok, err = sendSensorInformationRequest(id, listSensor, getChan)
		if err != nil {
			logger.Send("error", supportingfunctions.CustomError(err).Error())
		}
		if ok {
			logger.Send("info", fmt.Sprintf("we are sending a request to search for information on the following list of sensors: %#v", listSensor))
		}
	}(newDocument, dbs.GetChannelFromModule, dbs.logger)

	//формируем наименование индекса
	currentIndex := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), int(t.Month()))

	//*****************************
	//*** здесь установка тегов ***
	//пока не знаю нужно ли это в BiZone
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//caseId := fmt.Sprint(newDocument.GetEvent().GetObject().CaseId)
	//reqSetTag := fmt.Appendf(nil, `{
	//					  "service": "placeholder_doc-basedb-bi-zone",
	//					  "command": "add_case_tag",
	//					  "root_id": "%s",
	//					  "case_id": "%s",
	//					  "value": "Webhook: send=\"ElasticsearchDB"
	//					}`,
	//	newDocument.GetEvent().GetRootId(),
	//	caseId)
	//*****************************

	//будет выполнятся поиск по индексам только в текущем году так как при
	//накоплении большого количества индексов, поиск по всем серьезно замедлит работу
	indexesOnlyCurrentYear := []string(nil)
	for _, v := range existingIndexes {
		if strings.Contains(v, fmt.Sprint(t.Year())) {
			indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
		}
	}

	// если похожих индексов нет
	if len(indexesOnlyCurrentYear) == 0 {
		//
		//вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, binaryDocument)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		existingIndexes = append(existingIndexes, currentIndex)
		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject alerts to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document to alerts id:'%d', uuid:'%s', status code:'%d'", newDocument.GetIDNum(), newDocument.GetUUID(), statusCode))

		//*******************************************
		//*** здесь установка тегов, под вопросом ***
		//*******************************************

		return
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	//if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
	//	dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	//}

	//ищем объект с таким же идентификатором как и принятый в обработку объект
	res, err := dbs.GetDocument(
		ctx,
		indexesOnlyCurrentYear,
		strings.NewReader(
			fmt.Sprintf(
				`{
					"query": {
						"bool": {
							"must": [{"match": {"uuid": "%s"}}]}}}`,
				newDocument.GetUUID())),
		0)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}
	response := ResponseVerifiedBiZoneAlerts{}
	if err = json.Unmarshal(res, &response); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//вставка выполняется только когда не найден искомый документ
	if response.Options.Total.Value == 0 {
		//
		//вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, binaryDocument)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject alerts to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document to alerts id:'%d', uuid:'%s', status code:'%d'", newDocument.GetIDNum(), newDocument.GetUUID(), statusCode))

		//*******************************************
		//*** здесь установка тегов, под вопросом ***
		//*******************************************

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	var countReplacingFields int
	listDeleting := []ServiseOption(nil)
	updateVerified := datamodels.NewVerifiedBiZoneIRPAlert()
	//заполняем новый объект информацией из базы данных
	for _, v := range response.Options.Hits {
		countReplacingFields += updateVerified.RepalcingOldBiZoneAlert(*v.Source.Get())
		listDeleting = append(listDeleting, ServiseOption{
			ID:    v.ID,
			Index: v.Index,
		})

		//заполняем новый объект дополнительной информацией о сенсорах и
		// месторасположении ip адресов
		updateVerified.SetAdditionalInformation(*v.Source.GetAdditionalInformation())
	}

	//выполняем обновление нового объекта данными полученными от брокера сообщений
	updateVerified.RepalcingOldBiZoneAlert(*newDocument.Get())

	//готовим обновлённый объект к загрузке в базу данных
	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//обновление в базе данных уже существующего документа
	statusCode, countDel, err := dbs.UpdateDocument(ctx, currentIndex, listDeleting, nvbyte)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("uuid '%s' '%s'", newDocument.GetUUID(), err.Error())).Error())

		return
	}

	dbs.counter.SendMessage("update count insert subject alerts to db", 1)
	dbs.logger.Send("info", fmt.Sprintf("update document 'alerts' type, count delete:'%d', count replacing fields:'%d' for alerts with uuid:'%s', status code:'%d'", countDel, countReplacingFields, newDocument.GetUUID(), statusCode))
}

// TestAddBiZoneAlerts использовать только в тестах
func (dbs *DatabaseStorage) TestAddBiZoneAlerts(ctx context.Context, a any) {
	dbs.addBiZoneAlerts(ctx, a)
}
