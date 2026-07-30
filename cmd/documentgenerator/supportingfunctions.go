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

// CreateListSensors список с идентификаторами сенсоров
func CreateListSensors(verifiedData *datamodels.BiZoneIRPData) *datamodels.AdditionalInformation {
	information := &datamodels.AdditionalInformation{
		Sensors: []datamodels.SensorInformation(nil),
	}

	//из свойства 'data.all_sensors'
	for _, sensor := range verifiedData.AllSensors {
		information.AddSensorInformation(datamodels.SensorInformation{
			SensorId: fmt.Sprint(sensor),
		})
	}

	//из свойства 'data.all__ip_home'
	for _, ipHome := range verifiedData.AllIPHome {
		tmp := strings.Split(ipHome, ":")
		if len(tmp) == 0 {
			continue
		}

		information.AddSensorInformation(datamodels.SensorInformation{SensorId: tmp[0]})

	}

	//из свойства 'data.sensor'
	if verifiedData.Sensor != 0 {
		information.AddSensorInformation(datamodels.SensorInformation{
			SensorId: fmt.Sprint(verifiedData.Sensor),
		})
	}

	return information
}

// CreateListIpAddreses список с ip адресами
func CreateListIpAddreses(verifiedMainObject *datamodels.VerifiedBiZoneIRPAlert) *datamodels.AdditionalInformation {
	information := &datamodels.AdditionalInformation{
		IpAddresses: []datamodels.IpAddressInformation(nil),
	}

	//из свойства 'snapshots'
	for _, v := range verifiedMainObject.GetSnapshots() {
		if len(v.IPAddresses) == 0 {
			continue
		}

		for _, ip := range v.IPAddresses {
			if strings.Contains(ip, "[.]") {
				ip = strings.ReplaceAll(ip, "[.]", ".")
			}

			information.AddIpAddressInformation(datamodels.IpAddressInformation{
				Ip: ip,
			})
		}
	}

	//из свойства 'data.ip_exter'
	information.AddIpAddressInformation(datamodels.IpAddressInformation{
		Ip: verifiedMainObject.GetData().GetIPExter(),
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
