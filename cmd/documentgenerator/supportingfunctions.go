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

// verificationSecondaryCategoryType верифицирует BiZone тип 'data.secondary_category'
// Функция ищет срезы 'data.secondary_category.title' равные nil и заполняет их ОДНИМ значением
// взятым из предыдущего среза 'data.secondary_category.title'. Это делается из-за того, что поступающие
// в потоке данные и пути получаемые через метод msg.GetFieldBranch() содержат в себе значение
// 'data.secondary_category.id' и значения 'data.secondary_category.title.value', одно из них содержит
// наименование компьютерной атаки на английском языке, а второе перечень наименований на русском или
// других языках. Из-за потоковой обработки JSON и при количестве значений в 'data.secondary_category.title'
// иногда возникает путаница и последнее значение в 'data.secondary_category.title' для 'data.secondary_category.id'
// должно принадлежать следующей категории с 'data.secondary_category.id', но этого не происходит.
// Например, как в этом отрывке JSON:
//
//	{
//	  "id": "Social_Engineering",
//	  "title": [
//	    "Социальная инженерия"
//	  ]
//	},
//	{
//	  "id": "Network_Scanning",
//	  "title": [
//	    "Сетевое сканирование",
//	    "Публикация мошеннической информации"
//	  ]
//	},
//	{
//	  "id": "Fraudulent_Content_Published",
//	    "title": null
//	},
//
// где id = "Fraudulent_Content_Published" должно содержать значение "Публикация мошеннической информации"
func verificationSecondaryCategoryType(categores []datamodels.BiZoneIRPCategory) []datamodels.BiZoneIRPCategory {
	for i := range categores {
		if categores[i].Title != nil || i == 0 {
			continue
		}

		sizePreviousTitle := len(categores[i-1].Title)
		if sizePreviousTitle <= 1 {
			continue
		}

		// забираем значение title из предыдущего id
		categores[i].Title = append(categores[i].Title, categores[i-1].Title[sizePreviousTitle-1])
		categores[i-1].Title = categores[i-1].Title[: sizePreviousTitle-1 : sizePreviousTitle-1]
	}

	return categores
}

// verificationStatusType оставляет список 'data.status.title' содержащий только уникальные значения,
// так как в исходном JSON, полученном от IRP BiZone, часто встречаются title содержащий дублирующеся
// значения. Например,
// "title": [
//
//		{
//	       "value": "В работе",
//		   "language": "ru"
//		},
//		{
//		   "value": "В работе",
//		   "language": "en"
//		}
//
// ]
func verificationStatusType(status datamodels.BiZoneIRPStatus) datamodels.BiZoneIRPStatus {
	status.Title = supportingfunctions.SliceUniqueValues(status.Title)

	fmt.Println("func 'verificationStatusType', status:", status)

	return status
}

// verificationPriorityType оставляет список 'data.priority.title' содержащий только уникальные значения,
// так как в исходном JSON, полученном от IRP BiZone, часто встречаются title содержащий дублирующеся
// значения. Например,
// "title": [
//
//		{
//	       "value": "High",
//		   "language": "ru"
//		},
//		{
//		   "value": "High",
//		   "language": "en"
//		}
//
// ]
func verificationPriorityType(priority datamodels.BiZoneIRPPriority) datamodels.BiZoneIRPPriority {
	priority.Title = supportingfunctions.SliceUniqueValues(priority.Title)

	fmt.Println("func 'verificationPriorityType', priority:", priority)

	return priority
}

// CreateListSensorsForAlerts список с идентификаторами сенсоров из объекта 'alerts'
func CreateListSensorsForAlerts(verifiedData *datamodels.BiZoneIRPAlertData) *datamodels.AdditionalInformation {
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
func CreateListIpAddresesForAlerts(verifiedData *datamodels.BiZoneIRPAlertData) *datamodels.AdditionalInformation {
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
