package documentgenerator

import (
	"fmt"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// BiZoneAlertsGenerator генерирует верифицированный объект типа 'alerts'.
// Вернет первым элементом основной уникальный идентификатор события (UUID).
// Вторым, проверенный объект типа 'alerts'. Третьим список полей которые не были
// обработаны
func BiZoneAlertsGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *datamodels.VerifiedBiZoneIRPAlert, map[string]string) {
	// список не обработанных полей
	var listRawFields map[string]string = make(map[string]string)

	verifiedMainObject := datamodels.NewVerifiedBiZoneIRPAlert()
	verifiedData := datamodels.NewBiZoneIRPData()

	//********* основные обработчики **********
	//--- alerts ---
	listHandlerAlerts := handlers.NewListBiZoneHandlerAlerts(verifiedMainObject)

	//--- data ---
	listHandlerData := handlers.NewListBiZoneHandlerData(verifiedData)

	//******** вспомогательные объекты ********
	supportObjectTags := datamodels.NewSupportingStructureForTags()
	supportObjectSnapshot := datamodels.NewSupportingStructureForSnapshots()

	// ********* обработчики для вспомогательных объектов ***********
	listHandlerTags := handlers.NewListBiZoneHandlerTags(supportObjectTags)
	listHandlerSnapshots := handlers.NewListBiZoneHandlerSnapshots(supportObjectSnapshot)

	//объект с дополнительной информацией по сенсорам и ip адресам
	additionalInformation := datamodels.AdditionalInformation{
		Sensors:     []datamodels.SensorInformation(nil),
		IpAddresses: []datamodels.IpAddressInformation(nil),
	}

	for msg := range chInput {
		var handlerIsExist bool

		//*** обработчик для объекта alerts ***
		if lf, ok := listHandlerAlerts[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта alerts.data ***
		if lf, ok := listHandlerData[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}

		//***** обработчики для вспомогательных объектов *****
		//****************************************************
		//объект tags
		if lf, ok := listHandlerTags[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}
		//объект snapshots
		if lf, ok := listHandlerSnapshots[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}

		//записываем в лог-файл поля, которые не были обработаны
		if !handlerIsExist {
			listRawFields[msg.GetFieldBranch()] = fmt.Sprint(msg.GetValue())
		}
	}

	//собираем все объекты в один
	verifiedMainObject.SetData(*verifiedData.Get())
	verifiedMainObject.SetTags(supportObjectTags.GetTags())
	verifiedMainObject.SetSnapshots(supportObjectSnapshot.GetSnapshots())

	// формируем дополнительную информацию с идентификаторами сенсоров
	additionalInformation.SetSensorInformation(CreateListSensors(verifiedData).GetSensorsInformation())

	// формируем дополнительную информацию с ip адресами
	additionalInformation.SetIpAddressesInformation(CreateListIpAddreses(verifiedMainObject).GetIpAddressesInformation())

	verifiedMainObject.SetAdditionalInformation(additionalInformation)

	return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields
}
