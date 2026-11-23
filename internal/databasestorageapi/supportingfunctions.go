package databasestorageapi

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9/esapi"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/request"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// sendAdditionalInformationRequest
func sendAdditionalInformationRequest(id string, doc interfaces.SpecialAdditionalInformation, getChan func() chan SettingsChanOutput, logger interfaces.Logger) {
	// обогащение кейса дополнительной информацией о локальном место положении ip адресов
	listIp := documentgenerator.GetListIPAddr(doc.GetAdditionalInformation().GetIpAddressesInformation())
	ok, err := sendGeoIpRequest(id, listIp, getChan)
	if err != nil {
		logger.Send("error", supportingfunctions.CustomError(err).Error())
	}
	if ok {
		logger.Send("info", fmt.Sprintf("database storage - note for type 'incident', we are sending a request to search for information on the following list of ip addresses: %#v", listIp))
	}

	// обогащение кейса дополнительной информацией по расположению сенсоров
	listSensor := documentgenerator.GetListSensorId(doc.GetAdditionalInformation().GetSensorsInformation())
	ok, err = sendSensorInformationRequest(id, listSensor, getChan)
	if err != nil {
		logger.Send("error", supportingfunctions.CustomError(err).Error())
	}
	if ok {
		logger.Send("info", fmt.Sprintf("database storage - note for type 'incident', we are sending a request to search for information on the following list of sensors: %#v", listSensor))
	}
}

// sendGeoIpRequest запрос для обогащения кейса дополнительной информацией о локальном местоположении ip адресов
func sendGeoIpRequest(id string, list []string, getChan func() chan SettingsChanOutput) (bool, error) {
	if len(list) == 0 {
		return false, nil
	}

	reqGeoIP, err := json.Marshal(request.RequestGeoIP{
		Source:          "placeholder-doc-basedb-bizone",
		TaskId:          id,
		ListIpAddresses: list,
	})
	if err != nil {
		return false, err
	}

	getChan() <- SettingsChanOutput{
		Command: "get_geoip_info",
		Id:      id,
		Data:    reqGeoIP,
	}

	return true, nil
}

// sendSensorInformationRequest запрос для обогащения кейса дополнительной информацией по расположению сенсоров
func sendSensorInformationRequest(id string, list []string, getChan func() chan SettingsChanOutput) (bool, error) {
	if len(list) == 0 {
		return false, nil
	}

	reqSensorId, err := json.Marshal(request.RequestSensorInformation{
		Source:      "placeholder-doc-basedb-bizone",
		TaskId:      id,
		ListSensors: list,
	})
	if err != nil {
		return false, err
	}

	getChan() <- SettingsChanOutput{
		Command: "get_sensor_info",
		Id:      id,
		Data:    reqSensorId,
	}

	return true, nil
}

// bodyClose закрывает ответ с предварительной проверкой
func bodyClose(res *esapi.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}
