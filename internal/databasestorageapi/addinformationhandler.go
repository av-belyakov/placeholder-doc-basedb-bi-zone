package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/response"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// AddGeoIPInformation_Test дополнительная информация о географическом местоположении ip адресов
func (dbs *DatabaseStorage) AddGeoIPInformation_Test(ctx context.Context, a any) error {
	return dbs.addGeoIPInformation(ctx, a)
}

// addGeoIPInformation дополнительная информация о географическом местоположении ip адресов
func (dbs *DatabaseStorage) addGeoIPInformation(ctx context.Context, a any) error {
	newData, ok := a.([]byte)
	if !ok {
		err := supportingfunctions.CustomError(errors.New("type conversion error"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	var newDocument response.ResponseGeoIpInformation
	if err := json.Unmarshal(newData, &newDocument); err != nil {
		err := supportingfunctions.CustomError(err)
		dbs.logger.Send("error", err.Error())

		return err
	}

	dbs.logger.Send("info", fmt.Sprintf("section:'information handling', command:'add geoip information', accepted object:'%#v'", newDocument))

	//если в принятом ответе от модуля обогащения информацией о географическом местоположении
	//ip адресов есть глобальная ошибка
	if newDocument.Error != "" {
		err := supportingfunctions.CustomError(errors.New(newDocument.Error))
		dbs.logger.Send("error", err.Error())

		return err
	}

	//получаем наименование хранилища
	// так как id для выполяемой задачи по поиску информации о месторасположении ip адресов
	// был сформирован по следующему шаблону id := fmt.Sprintf("<index_elasticsearch>:%s", newDocument.GetUUID())
	tmp := strings.Split(newDocument.TaskId, ":")
	if len(tmp) <= 1 {
		err := supportingfunctions.CustomError(errors.New("no value was found to determine the index name"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	currentIndex, specialUUID := tmp[0], tmp[1]

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
		err := supportingfunctions.CustomError(errors.New("the list with information on ip addresses should not be empty"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	underlineId, geoIpInfo, err := dbs.SearchGeoIPInformation(ctx, currentIndex, specialUUID)
	if err != nil {
		err := supportingfunctions.CustomError(fmt.Errorf("the identifier of the index name '%s' was not found", currentIndex))
		dbs.logger.Send("error", err.Error())

		return err
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
		err := supportingfunctions.CustomError(fmt.Errorf("@special_uuid:'%s', '%w'", specialUUID, err))
		dbs.logger.Send("error", err.Error())

		return err
	}

	//выполняется обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err := dbs.client.Update(currentIndex, underlineId, bodyUpdate)
	if err != nil {
		err := supportingfunctions.CustomError(fmt.Errorf("@special_uuid:'%s', '%w'", specialUUID, err))
		dbs.logger.Send("error", err.Error())

		return err
	}
	defer bodyClose(res)

	if res != nil && res.StatusCode != http.StatusOK {
		err := supportingfunctions.CustomError(fmt.Errorf("@special_uuid:'%s', '%w'", specialUUID, err))
		dbs.logger.Send("error", err.Error())

		return err
	}

	dbs.logger.Send("info", fmt.Sprintf("the geoip information for @special_uuid:'%s' it was added successfully", newDocument.TaskId))

	return nil
}

// AddSensorInformation_Test дополнительная информация о местоположении и принадлежности сенсоров, только для тестов
func (dbs *DatabaseStorage) AddSensorInformation_Test(ctx context.Context, a any) error {
	return dbs.addSensorInformation(ctx, a)
}

// addSensorInformation дополнительная информация о местоположении и принадлежности сенсоров
func (dbs *DatabaseStorage) addSensorInformation(ctx context.Context, a any) error {
	newData, ok := a.([]byte)
	if !ok {
		err := supportingfunctions.CustomError(errors.New("type conversion error"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	var newDocument response.ResponseSensorsInformation
	if err := json.Unmarshal(newData, &newDocument); err != nil {
		err = supportingfunctions.CustomError(err)
		dbs.logger.Send("error", err.Error())

		return err
	}

	dbs.logger.Send("info", fmt.Sprintf("section:'information handling', command:'add sensor information', accepted object:'%#v'", newDocument))

	//если в принятом ответе от модуля обогащения информацией о сенсорах есть глобальная ошибка
	if newDocument.Error != "" {
		err := supportingfunctions.CustomError(errors.New(newDocument.Error))
		dbs.logger.Send("error", err.Error())

		return err
	}

	//получаем наименование хранилища
	// так как id для выполяемой задачи по поиску информации о месторасположении ip адресов
	// был сформирован по следующему шаблону id := fmt.Sprintf("<index_elasticsearch>:%s", newDocument.GetUUID())
	tmp := strings.Split(newDocument.TaskId, ":")
	if len(tmp) <= 1 {
		err := supportingfunctions.CustomError(errors.New("no value was found to determine the index name"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	currentIndex, specialUUID := tmp[0], tmp[1]

	fmt.Printf("\t---DatabaseStorage.addSensorInformation elasticsearch index:'%s', UUID:'%s'\n", currentIndex, specialUUID)

	//добавляем небольшую задержку что бы СУБД успела добавить индекс
	//***************************************************************
	time.Sleep(3 * time.Second)
	//***************************************************************

	//поиск _id объектов по их '@special_uuid', вообще должно быть найдено только одно значение
	underlineIds, err := dbs.GetUnderlineIds(ctx, currentIndex, specialUUID)
	if err != nil {
		err := supportingfunctions.CustomError(errors.New("the identifier of the index name was not found"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	fmt.Printf("\t---DatabaseStorage.addSensorInformation в базе данных был найден объект с underlineId:'%+v'\n", underlineIds)

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
		err := supportingfunctions.CustomError(errors.New("the list with information on sensors should not be empty"))
		dbs.logger.Send("error", err.Error())

		return err
	}

	request, err := json.MarshalIndent(datamodels.AdditionalInformationSensors{Sensors: sensorInfoList}, "", " ")
	if err != nil {
		err := supportingfunctions.CustomError(fmt.Errorf("'rootId '%s', json unmarshal - '%w'", newDocument.TaskId, err))
		dbs.logger.Send("error", err.Error())

		return err
	}

	//обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))

	fmt.Printf("\t---DatabaseStorage.addSensorInformation обновление информации о сенсорах, готовый объект:'%s'\n", fmt.Sprintf("{\"doc\": %s}", string(request)))

	for _, underlineId := range underlineIds {
		res, err := dbs.client.Update(currentIndex, underlineId, bodyUpdate)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId '%s', update database - '%w'", newDocument.TaskId, err)).Error())
		}

		fmt.Printf("\t---DatabaseStorage.addSensorInformation статус код ответа:'%d', response:'%+v'\n", res.StatusCode, res.Body)

		if res != nil && res.StatusCode != http.StatusOK {
			if res.Body == nil {
				dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId '%s', update database status code '%d'", newDocument.TaskId, res.StatusCode)).Error())

				continue
			}

			resInfo, err := io.ReadAll(res.Body)
			if err != nil {
				dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId '%s', status code '%d', body reading - '%w'", newDocument.TaskId, res.StatusCode, err)).Error())
			} else {
				dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId '%s', status code '%d', data base responsed '%s'", newDocument.TaskId, res.StatusCode, string(resInfo))).Error())
			}

			continue
		}

		bodyClose(res)
	}

	dbs.logger.Send("info", fmt.Sprintf("the sensors information for root id:'%s' it was added successfully", newDocument.TaskId))

	return nil
}
