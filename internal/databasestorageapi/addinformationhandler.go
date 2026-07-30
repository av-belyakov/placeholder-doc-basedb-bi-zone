package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/response"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// addGeoIPInformation дополнительная информация о географическом местоположении ip адресов
func (dbs *DatabaseStorage) addGeoIPInformation(ctx context.Context, a any) {
	newData, ok := a.([]byte)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	var newDocument response.ResponseGeoIpInformation
	if err := json.Unmarshal(newData, &newDocument); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("section:'information handling', command:'add geoip information', accepted object:'%#v'", newDocument))

	//если в принятом ответе от модуля обогащения информацией о географическом местоположении
	//ip адресов есть глобальная ошибка
	if newDocument.Error != "" {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(newDocument.Error)).Error())

		return
	}

	//получаем наименование хранилища
	// так как id для выполяемой задачи по поиску информации о месторасположении ip адресов
	// был сформирован по следующему шаблону id := fmt.Sprintf("alerts:%s", newDocument.GetUUID())
	tmp := strings.Split(newDocument.TaskId, ":")
	if len(tmp) <= 1 {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("no value was found to determine the index name")).Error())

		return
	}

	//получаем имя индекса из настроек конфигурации
	indexName, isExist := dbs.settings.storages[tmp[0]]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	specialUuid := tmp[1]
	t := time.Now()
	month := int(t.Month())
	//текущий индекс
	indexCurrent := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), month)

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, time.Second*15)
	defer ctxCancel()

	//добавляем небольшую задержку что бы СУБД успела добавить индекс
	//***************************************************************
	time.Sleep(3 * time.Second)
	//***************************************************************

	//формируется список с информацией по ip адресам
	var ipInfoList []datamodels.IpAddressInformation
	for _, ipAddress := range newDocument.Informations {
		if ipAddress.Error != "" {
			dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(ipAddress.Error)).Error())

			ipInfoList = append(ipInfoList, datamodels.IpAddressInformation{
				Ip: ipAddress.IpAddr,
			})

			continue
		}

		ipInfoList = append(ipInfoList, datamodels.IpAddressInformation{
			Ip:          ipAddress.IpAddr,
			City:        ipAddress.City,
			Country:     ipAddress.Country,
			CountryCode: ipAddress.Code,
		})
	}

	//обновляемый список не должен быть пустым, а то получается что полу пустой
	//список в котором есть хотя бы ip адреса затирается вообще пустым списком
	if len(ipInfoList) == 0 {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the list with information on ip addresses should not be empty")).Error())

		return
	}

	underlineId, geoIpInfo, err := dbs.SearchGeoIPInformation(ctxTimeout, indexCurrent, specialUuid)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//выполняем сравнение принятого списка со списком из БД и добавляем недостающие
	//ip адреса или заменяем информацию по уже имеющимся ip адресам на более свежую
	for _, v := range ipInfoList {
		num, ok := supportingfunctions.SliceContainsElementFunc(geoIpInfo, func(num int) bool {
			if v.Ip == geoIpInfo[num].Ip {
				return true
			}

			return false
		})
		if ok {
			geoIpInfo[num] = v
		} else {
			geoIpInfo = append(geoIpInfo, v)
		}
	}

	request, err := json.MarshalIndent(datamodels.AdditionalInformationIpAddress{IpAddresses: geoIpInfo}, "", " ")
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("@special_uuid:'%s', '%w'", specialUuid, err)).Error())

		return
	}

	//выполняется обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err := dbs.client.Update(indexCurrent, underlineId, bodyUpdate)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("@specail_uuid:'%s', '%w'", specialUuid, err)).Error())

		return
	}
	defer res.Body.Close()

	if res != nil && res.StatusCode != http.StatusOK {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("@special_uuid:'%s', '%w'", specialUuid, err)).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("the geoip information for @special_uuid:'%s' it was added successfully", newDocument.TaskId))
}

// addSensorInformation дополнительная информация о местоположении и принадлежности сенсоров
func (dbs *DatabaseStorage) addSensorInformation(ctx context.Context, a any) {
	newData, ok := a.([]byte)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	var newDocument response.ResponseSensorsInformation
	if err := json.Unmarshal(newData, &newDocument); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("section:'information handling', command:'add sensor information', accepted object:'%#v'", newDocument))

	//если в принятом ответе от модуля обогащения информацией о сенсорах есть глобальная ошибка
	if newDocument.Error != "" {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(newDocument.Error)).Error())

		return
	}

	//получаем наименование хранилища
	// так как id для выполяемой задачи по поиску информации о месторасположении ip адресов
	// был сформирован по следующему шаблону id := fmt.Sprintf("alerts:%s", newDocument.GetUUID())
	tmp := strings.Split(newDocument.TaskId, ":")
	if len(tmp) <= 1 {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("no value was found to determine the index name")).Error())

		return
	}

	//получаем имя индекса из настроек конфигурации
	indexName, isExist := dbs.settings.storages[tmp[0]]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	specialUuid := tmp[1]
	t := time.Now()
	month := int(t.Month())
	//текущий индекс
	indexCurrent := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), month)

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, time.Second*15)
	defer ctxCancel()

	//добавляем небольшую задержку что бы СУБД успела добавить индекс
	//***************************************************************
	time.Sleep(3 * time.Second)
	//***************************************************************

	//поиск _id объекта по его '@special_uuid'
	underlineId, err := dbs.GetUnderlineId(ctxTimeout, indexCurrent, specialUuid)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//формируется список с информацией по сенсорам
	var sensorInfoList []datamodels.SensorInformation
	for _, sensor := range newDocument.Informations {
		if sensor.Error != "" {
			dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(sensor.Error)).Error())
		}

		sensorInfoList = append(sensorInfoList, datamodels.SensorInformation{
			SensorId:    sensor.SensorID,
			GeoCode:     sensor.GeoCode,
			ObjectArea:  sensor.ObjectArea,
			SubjectRF:   sensor.SubjectRussianFederation,
			INN:         sensor.INN,
			HomeNet:     sensor.HomeNet,
			OrgName:     sensor.OrganizationName,
			FullOrgName: sensor.FullOrganizationName,
		})
	}

	//обновляемый список не должен быть пустым, а то получается что пустой
	//список с информацией по сенсорам затирается вообще пустым списком
	if len(sensorInfoList) == 0 {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the list with information on sensors should not be empty")).Error())

		return
	}

	request, err := json.MarshalIndent(datamodels.AdditionalInformationSensors{Sensors: sensorInfoList}, "", " ")
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}

	//обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err := dbs.client.Update(indexCurrent, underlineId, bodyUpdate)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}
	defer res.Body.Close()

	if res != nil && res.StatusCode != http.StatusOK {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("the sensors information for root id:'%s' it was added successfully", newDocument.TaskId))
}
