package documentgenerator

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// GetListIPAddr список ip адресов из элементов объекта
func GetListIPAddr(objects []datamodels.IpAddressInformation) []string {
	newList := make([]string, 0, len(objects))

	for _, v := range objects {
		if slices.ContainsFunc(newList, func(elem string) bool {
			return elem == v.GetIpAddrString()
		}) {
			continue
		}

		newList = append(newList, v.GetIpAddrString())
	}

	return newList
}

// GetListSensorId список идентификаторов сенсоров
func GetListSensorId(objects []datamodels.SensorInformation) []string {
	newList := make([]string, 0, len(objects))

	for _, v := range objects {
		if slices.ContainsFunc(newList, func(elem string) bool {
			return elem == v.GetSensorId()
		}) {
			continue
		}

		newList = append(newList, v.GetSensorId())
	}

	return newList
}

// CreateListSensorsForAlerts список с идентификаторами сенсоров из объекта 'alerts'
func CreateListSensorsForAlerts(verifiedData *datamodels.BiZoneIRPData) *datamodels.AdditionalInformation {
	information := &datamodels.AdditionalInformation{
		Sensors: []datamodels.SensorInformation(nil),
	}

	//из свойства 'data.unmapped_agent_array'
	for _, sensor := range verifiedData.UnmappedAgentArray {
		information.AddSensorInformation(datamodels.SensorInformation{
			SensorId: fmt.Sprint(sensor),
		})
	}

	//из свойства 'data.unmapped_home_endpoint_array'
	for _, ipHome := range verifiedData.UnmappedHomeEndpointArray {
		tmp := strings.Split(ipHome, ":")
		if len(tmp) == 0 {
			continue
		}

		information.AddSensorInformation(datamodels.SensorInformation{SensorId: tmp[0]})

	}

	//из свойства 'data.agent'
	if verifiedData.Agent != 0 {
		information.AddSensorInformation(datamodels.SensorInformation{
			SensorId: fmt.Sprint(verifiedData.Agent),
		})
	}

	return information
}

// CreateListIpAddreses список с ip адресами из объекта 'alerts'
func CreateListIpAddresesForAlerts(verifiedData *datamodels.BiZoneIRPData) *datamodels.AdditionalInformation {
	information := &datamodels.AdditionalInformation{
		IpAddresses: []datamodels.IpAddressInformation(nil),
	}

	//из свойства 'data.unmapped_dst_endpoint_array'
	for _, ip := range verifiedData.UnmappedDstEndpointArray {
		information.AddIpAddressInformation(datamodels.IpAddressInformation{
			Ip: ip,
		})
	}

	//из свойства 'data.source_ip'
	information.AddIpAddressInformation(datamodels.IpAddressInformation{
		Ip: verifiedData.GetSourceIP(),
	})

	return information
}

/*
// getSensorIdFromDescription выполняет поиск идентификатора сенсора в поле description
func getSensorIdFromDescription(v string) (string, error) {
	rexSensorId := regexp.MustCompile(`СОА:\s-\s\*\*\x60(\d+)\x60\*\*`)
	tmp := rexSensorId.FindStringSubmatch(v)

	if len(tmp) <= 1 {
		return "", errors.New("there is no sensor ID in the accepted line")
	}

	return tmp[1], nil
}
*/
