package documentgenerator

import (
	"fmt"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// BiZoneAlertsGenerator генерирует верифицированный объект типа 'alerts'.
// Вернет первым элементом основной уникальный идентификатор события (UUID).
// Вторым, проверенный объект типа 'alerts'.
// Третьим список полей которые не были обработаны.
// Четвёртым ошибку.
func BiZoneAlertsGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *datamodels.VerifiedBiZoneIRPAlert, map[string]string, error) {
	// список не обработанных полей
	var listRawFields map[string]string = make(map[string]string)

	verifiedMainObject := datamodels.NewVerifiedBiZoneIRPAlert()
	verifiedData := datamodels.NewBiZoneIRPData()

	//********* основные обработчики **********
	listHandlerAlerts := handlers.NewListBiZoneHandlerAlerts(verifiedMainObject)
	listHandlerData := handlers.NewListBiZoneHandlerData(verifiedData)

	//******** вспомогательные объекты ********
	supportObjectTags := handlers.NewSupportingStructureForTagsType()
	supportObjectSnapshot := handlers.NewSupportingStructureForSnapshotsType()
	supportObjectSContent := handlers.NewSupportingStructureForSContentType()
	supportObjectDataSecurity := handlers.NewSupportingStructureForDataSecurityType()

	// ********* обработчики для вспомогательных объектов ***********
	listHandlerTags := handlers.NewListBiZoneHandlerTags(supportObjectTags)
	listHandlerSnapshots := handlers.NewListBiZoneHandlerSnapshots(supportObjectSnapshot)
	listHandlerScontents := handlers.NewListBiZoneHandlerSContents(supportObjectSContent)
	listHandlerDataSecurity := handlers.NewListBiZoneHandlerDataSecurity(supportObjectDataSecurity)

	// объект с дополнительной информацией по сенсорам и ip адресам
	additionalInformation := datamodels.AdditionalInformation{
		Sensors:     []datamodels.SensorInformation(nil),
		IpAddresses: []datamodels.IpAddressInformation(nil),
	}

	for msg := range chInput {
		var handlerIsExist bool

		//*** обработчик для объекта alerts ***
		if funcs, ok := listHandlerAlerts[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта alerts.data ***
		if funcs, ok := listHandlerData[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//***** обработчики для вспомогательных объектов *****
		//****************************************************
		// объект tags
		if funcs, ok := listHandlerTags[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}
		// объект snapshots
		if funcs, ok := listHandlerSnapshots[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		// объект scontent
		if funcs, ok := listHandlerScontents[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//////continue
		}

		// объект datasecurity
		if funcs, ok := listHandlerDataSecurity[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		// записываем в лог-файл поля, которые не были обработаны
		if !handlerIsExist {
			listRawFields[msg.GetFieldBranch()] = fmt.Sprint(msg.GetValue())
		}
	}

	// собираем все объекты в один
	listDataSecurity := supportObjectDataSecurity.GetDataSecurity()
	listSContent := supportObjectSContent.GetSContent()
	// дополняем объект 'data_security' списком 's_content'
	for k, v := range listDataSecurity {
		for item, dataSecurity := range v {
			if sContents, ok := listSContent[dataSecurity.ISid]; ok {
				dataSecurity.SContent = sContents
				listDataSecurity[k][item] = dataSecurity
			}
		}
	}

	// собираем объект 'data_security'
	if err := verifiedData.SetDataSecurity(listDataSecurity); err != nil {
		return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, err
	}
	// собираем объект 'data'
	if err := verifiedMainObject.SetData(*verifiedData.Get()); err != nil {
		return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, err
	}
	// собираем объект 'tags'
	if err := verifiedMainObject.SetTags(supportObjectTags.GetTags()); err != nil {
		return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, err
	}
	// собираем объект 'snapshots'
	if err := verifiedMainObject.SetSnapshots(supportObjectSnapshot.GetSnapshots()); err != nil {
		return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, err
	}

	// формируем дополнительную информацию с идентификаторами сенсоров
	additionalInformation.SetSensorInformation(CreateListSensorsForAlerts(verifiedData).GetSensorsInformation())

	// формируем дополнительную информацию с ip адресами
	additionalInformation.SetIpAddressesInformation(CreateListIpAddresesForAlerts(verifiedData).GetIpAddressesInformation())

	if err := verifiedMainObject.SetAdditionalInformation(additionalInformation); err != nil {
		return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, err
	}

	return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, nil
}
