package databasestorageapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/elastic/go-elasticsearch/v9/esapi"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/request"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

var idLike = regexp.MustCompile(`^\d+$|^[0-9a-f]{8}-[0-9a-f]{4}|^[A-Za-z0-9+/]{20,}`)

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

func normalizePath(p string) string {
	if p == "" {
		return "/"
	}

	segs := splitPath(p)
	for i, s := range segs {
		if s == "" || (len(s) > 0 && s[0] == '_') {
			continue
		}
		if idLike.MatchString(s) {
			segs[i] = "<id>"
		}
	}

	return "/" + joinPath(segs)
}

func splitPath(path string) []string {
	var out []string
	for _, p := range bytes.Split([]byte(path), []byte("/")) {
		if len(p) > 0 {
			out = append(out, string(p))
		}
	}

	return out
}

func joinPath(segs []string) string {
	if len(segs) == 0 {
		return "/"
	}

	var result strings.Builder

	// Если первый сегмент пустой — значит, путь был абсолютным (начинался с '/')
	isAbsolute := len(segs) > 0 && segs[0] == ""
	if isAbsolute {
		result.WriteByte('/')
		// Начинаем со второго элемента, потому что первый — это маркер абсолютности
		segs = segs[1:]
	}

	for i, seg := range segs {
		if seg == "" {
			// Пропускаем пустые сегменты (они возникают, если было //)
			continue
		}
		if i > 0 || isAbsolute {
			// Добавляем слэш перед сегментом, если:
			// - это не самый первый сегмент в относительном пути,
			// - или путь абсолютный (тогда мы уже написали начальный '/')
			result.WriteByte('/')
		}
		result.WriteString(seg)
	}

	// Если после пропуска пустых сегментов ничего не осталось — возвращаем "/"
	if result.Len() == 0 {
		return "/"
	}

	return result.String()
}

/*
func indexFromPath(path string) string {
	// путь вида /{index}/_search или /{index}/_doc/{id}
	parts := splitPath(path)
	if len(parts) > 0 && parts[0] != "" && parts[0][0] != '_' {
		return parts[0]
	}

	return "cluster"
}

func indexFromOpName(op string) string {
	// op в bulk-ответе: "index", "create", "update", "delete"
	return op
}
*/
