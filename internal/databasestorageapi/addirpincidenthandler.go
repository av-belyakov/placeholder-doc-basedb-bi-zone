package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// TestAddIRPIncident использовать только в тестах
func (dbs *DatabaseStorage) TestAddIRPIncident(ctx context.Context, a any) error {
	return dbs.addIRPIncident(ctx, a)
}

// addIRPIncident добавление объекта типа 'incident'
func (dbs *DatabaseStorage) addIRPIncident(ctx context.Context, a any) error {
	t := time.Now()

	newDocument, ok := a.(*datamodels.VerifiedBiZoneIRPIncident)
	if !ok {
		err := supportingfunctions.CustomError(errors.New("note for type 'incident', type conversion error"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	dbs.logger.Send("info", fmt.Sprintf("database storage - note for type 'incident', accepted document with key (id):'%s'", newDocument.Key))

	binaryDocument, err := json.Marshal(newDocument.Get())
	if err != nil {
		err := supportingfunctions.CustomError(err)
		dbs.logger.Send("error", err.Error())

		return err
	}

	// получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["incidents"]
	if !isExist {
		err := supportingfunctions.CustomError(errors.New("database storage - note for type 'incident', the identifier of the index name was not found"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	// получаем существующие индексы
	existingIndexes, err := dbs.GetExistingIndexes(ctx, indexName)
	if err != nil {
		err := supportingfunctions.CustomError(err)
		dbs.logger.Send("error", err.Error())

		return err
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

	// создаём уникальный идентификатор дополнительной информации для поиска инцидентов по нему в БД
	// однако, вновь созданный идентификатор можно использовать только в том случае если инцидента
	// ещё нет в БД, если инциденто уже загружен в БД нужно брать уже сущетсвующий id
	idAddInfo := fmt.Sprintf("%s:%s", currentIndex, newDocument.GetSpecialUUID())

	// если похожих индексов нет
	if len(indexesOnlyCurrentYear) == 0 {
		//
		// вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, binaryDocument)
		if err != nil {
			err := supportingfunctions.CustomError(err)
			dbs.logger.Send("error", err.Error())

			return err
		}

		existingIndexes = append(existingIndexes, currentIndex)
		// устанавливаем максимальный лимит количества полей для всех индексов которые
		// содержат значение по умолчанию в 1000 полей
		err = dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}

		dbs.logger.Send("info", fmt.Sprintf("database storage - note for type 'incident', insert new document to incidents key (id) '%s', special_uuid '%s', status code '%d'", newDocument.GetKey(), newDocument.GetSpecialUUID(), statusCode))

		// запрос на обогащение документа дополнительной информацией
		go sendAdditionalInformationRequest(idAddInfo, newDocument, dbs.GetChannelFromModule, dbs.logger)

		return supportingfunctions.CustomError(err)
	}

	// устанавливаем максимальный лимит количества полей для всех индексов которые
	// содержат значение по умолчанию в 1000 полей
	if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	}

	// ищем объект с таким же идентификатором как и принятый в обработку объект
	query := strings.NewReader(fmt.Sprintf(
		`{
		    "query": {
			  "term": {
			    "key.keyword": "%s"
			  }
		    }
		}`,
		newDocument.GetKey()))
	res, err := dbs.GetDocument(ctx, indexesOnlyCurrentYear, query, 0)
	if err != nil {
		err = supportingfunctions.CustomError(err)
		dbs.logger.Send("error", err.Error())

		return err
	}

	fmt.Printf("---- ищем инцидент с идентификатором '%s'\n", newDocument.Key)

	response := ResponseVerifiedIRPIncidents{}
	if err = json.Unmarshal(res, &response); err != nil {
		err := supportingfunctions.CustomError(err)
		dbs.logger.Send("error", err.Error())

		return err
	}

	fmt.Printf("---- в результате поиска инцидента '%s' совпадений найдено %d\n", newDocument.Key, response.Options.Total.Value)

	// вставка выполняется только когда не найден искомый инцидент
	if response.Options.Total.Value == 0 {
		//
		// вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, binaryDocument)
		if err != nil {
			err := supportingfunctions.CustomError(err)
			dbs.logger.Send("error", err.Error())

			return err
		}

		dbs.logger.Send("info", fmt.Sprintf("database storage - note for type 'incident', insert new document to incident key (id) '%s', status code '%d'", newDocument.GetKey(), statusCode))

		// запрос на обогащение документа дополнительной информацией
		go sendAdditionalInformationRequest(idAddInfo, newDocument, dbs.GetChannelFromModule, dbs.logger)

		//*******************************************
		//*** здесь установка тегов, под вопросом ***
		//*******************************************

		return err
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
	replasedField := updateVerified.ReplacingOldIRPIncident(*newDocument.Get())

	// готовим обновлённый объект к загрузке в базу данных
	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		err := supportingfunctions.CustomError(err)
		dbs.logger.Send("error", err.Error())

		return err
	}

	// обновление в базе данных уже существующего документа
	statusCode, countDel, err := dbs.UpdateDocument(ctx, currentIndex, listDeleting, nvbyte)
	if err != nil {
		err := supportingfunctions.CustomError(fmt.Errorf("key (id) '%s' '%s'", newDocument.GetKey(), err.Error()))
		dbs.logger.Send("error", err.Error())

		return err
	}

	dbs.logger.Send(
		"info",
		fmt.Sprintf(
			"database storage - note for type 'incident', update document, count delete documents '%d', count replacing fields '%d' for object with key (id) '%s', status code '%d'",
			countDel,
			len(replasedField),
			newDocument.GetKey(),
			statusCode,
		))

	for _, v := range replasedField {
		dbs.logger.Send("info", fmt.Sprintf("database storage - update field '%s' to document 'incident' with key (id) '%s'", v, newDocument.GetKey()))
	}

	idAddInfo = fmt.Sprintf("%s:%s", currentIndex, updateVerified.GetAdditionalInformation().SpecialUUID)

	// запрос на обогащение документа дополнительной информацией
	go sendAdditionalInformationRequest(idAddInfo, newDocument, dbs.GetChannelFromModule, dbs.logger)

	return nil
}
