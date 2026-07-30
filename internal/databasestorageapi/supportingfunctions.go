package databasestorageapi

import (
	"encoding/json"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/request"
)

// sendGeoIpRequest запрос для обогащения кейса дополнительной информацией о
// локальном местоположении ip адресов
func sendGeoIpRequest(
	id string,
	list []string,
	getChan func() chan SettingsChanOutput) (bool, error) {
	if len(list) == 0 {
		return false, nil
	}

	reqGeoIP, err := json.Marshal(request.RequestGeoIP{
		Source:          "placeholder_doc-basedb-bi-zone",
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

// sendSensorInformationRequest запрос для обогащения кейса дополнительной информацией по
// расположению сенсоров
func sendSensorInformationRequest(
	id string,
	list []string,
	getChan func() chan SettingsChanOutput) (bool, error) {
	if len(list) == 0 {
		return false, nil
	}

	reqSensorId, err := json.Marshal(request.RequestSensorInformation{
		Source:      "placeholder_doc-basedb-bi-zone",
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
