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
	var (
		listRawFields map[string]string = make(map[string]string)

		err error
	)

	verifiedMainObject := datamodels.NewVerifiedBiZoneIRPAlert()
	verifiedData := datamodels.NewBiZoneIRPAlertData()

	//********* основные обработчики **********
	listHandlerAlerts := handlers.NewListBiZoneHandlerAlerts(verifiedMainObject)
	listHandlerData := handlers.NewListBiZoneHandlerAlertData(verifiedData)

	//******** вспомогательные объекты ********
	supportObjectTags := handlers.NewSupportingStructureForTagsType()
	supportObjectSnapshot := handlers.NewSupportingStructureForSnapshotsType()
	supportObjectSContent := handlers.NewSupportingStructureForSContentsType()
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

		//*** обработчик для объекта 'alerts' ***
		if funcs, ok := listHandlerAlerts[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта 'alerts.data' ***
		if funcs, ok := listHandlerData[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//***** обработчики для вспомогательных объектов *****
		//****************************************************
		// объект 'tags'
		if funcs, ok := listHandlerTags[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}
		// объект 'snapshots'
		if funcs, ok := listHandlerSnapshots[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'data.datasecurity.scontent'
		if funcs, ok := listHandlerScontents[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'data.datasecurity'
		if funcs, ok := listHandlerDataSecurity[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
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
	err = injectSContentToDataSecurity(listDataSecurity, listSContent)

	// собираем объект 'data_security'
	if errTmp := verifiedData.SetDataSecurity(listDataSecurity); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'data'
	if errTmp := verifiedMainObject.SetData(*verifiedData.Get()); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'tags'
	if errTmp := verifiedMainObject.SetTags(supportObjectTags.GetTags()); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'snapshots'
	if errTmp := verifiedMainObject.SetSnapshots(supportObjectSnapshot.GetSnapshots()); errTmp != nil {
		err = errTmp
	}

	// формируем дополнительную информацию с идентификаторами сенсоров
	additionalInformation.SetSensorInformation(CreateListSensorsForAlerts(verifiedData).GetSensorsInformation())
	// формируем дополнительную информацию с ip адресами
	additionalInformation.SetIpAddressesInformation(CreateListIpAddresesForAlerts(verifiedData).GetIpAddressesInformation())
	// добавляем специальный, общий идентификатор
	// так как в alert и case uuid хранятся в разных местах и под разными названиями необходимо унифицировать
	// место хранения идентификатора для разных объектов, это позволит унифицировать обработчик отвечающий
	// за наполнение полей с доп. информацией по ip адресам и сенсорам
	additionalInformation.SetSpecialUUID(verifiedMainObject.GetUUID())

	if errTmp := verifiedMainObject.SetAdditionalInformation(additionalInformation); errTmp != nil {
		err = errTmp
	}

	return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, err
}
