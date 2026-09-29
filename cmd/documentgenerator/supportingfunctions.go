package documentgenerator

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// injectSContentToDataSecurity дополняет объект 'data_security' списком 's_content'
func injectSContentToDataSecurity(lds map[string][]datamodels.BiZoneIRPDataSecurity, lsc []datamodels.BiZoneIRPSContent) error {
	for classType, dataSecurity := range lds {
		for k, v := range dataSecurity {
			contents, err := supportingfunctions.GetContentFromSnortRule(v.SRuleBody)
			if err != nil {
				return err
			}

			for _, content := range contents {
				index, isExist := searchSContentFromSContents(content, lsc)
				if !isExist {
					continue
				}

				v.SContent = append(v.SContent, lsc[index])
				dataSecurity[k] = v
				lds[classType] = dataSecurity
			}
		}
	}

	return nil
}

// searchSContentFromSContents ищет значение в срезе
func searchSContentFromSContents(content string, contents []datamodels.BiZoneIRPSContent) (int, bool) {
	for k, v := range contents {
		if strings.Contains(v.Content, content) {
			return k, true
		}
	}

	return -1, false
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

	//из свойства 'data.source_ip'
	information.AddIpAddressInformation(datamodels.IpAddressInformation{
		Ip: verifiedData.GetSourceIP(),
	})

	//из свойства 'data.target_ip'
	information.AddIpAddressInformation(datamodels.IpAddressInformation{
		Ip: verifiedData.GetTargetIP(),
	})

	//из свойства 'data.unmapped_dst_endpoint_array'
	for _, ip := range verifiedData.UnmappedDstEndpointArray {
		information.AddIpAddressInformation(datamodels.IpAddressInformation{
			Ip: ip,
		})
	}

	//из свойства 'data.unmapped_home_endpoint_array'
	for _, ipHome := range verifiedData.UnmappedHomeEndpointArray {
		tmp := strings.Split(ipHome, ":")
		if len(tmp) == 0 {
			continue
		}

		information.AddIpAddressInformation(datamodels.IpAddressInformation{
			Ip: tmp[0],
		})
	}

	return information
}

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
