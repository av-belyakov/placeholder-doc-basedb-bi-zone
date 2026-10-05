package documentgenerator

import (
	"fmt"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// BiZoneCasesGenerator генерирует верифицированный объект типа 'cases'.
// Вернет первым элементом основной уникальный идентификатор события (UUID).
// Вторым, проверенный объект типа 'cases'.
// Третьим список полей которые не были обработаны.
// Четвёртым ошибку.
func BiZoneCasesGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *datamodels.VerifiedBiZoneIRPCase, map[string]string, error) {
	// список не обработанных полей
	var (
		listRawFields map[string]string = make(map[string]string)

		err error
	)

	verifiedMainObject := datamodels.NewVerifiedBiZoneIRPCase()
	verifiedData := datamodels.NewBiZoneIRPCaseData()
	verifiedTypeObject := datamodels.NewBiZoneIRPType()
	verifiedStatusObject := datamodels.NewBiZoneIRPStatus()
	verifiedTenantObject := datamodels.NewBiZoneIRPTenant()
	verifiedPriorityObject := datamodels.NewBiZoneIRPPriority()
	verifiedCreatedByObject := datamodels.NewBiZoneIRPCreatedBy()
	verifiedPrimaryCategoryObject := datamodels.NewBiZoneIRPCategory()

	//********* основные обработчики **********
	listHandlerCases := handlers.NewListBiZoneHandlerCases(verifiedMainObject)
	listHandlerData := handlers.NewListBiZoneHandlerCaseData(verifiedData)

	//******** вспомогательные объекты ********
	supportObjectTags := handlers.NewSupportingStructureForTagsType()
	supportObjectCategores := handlers.NewSupportingStructureForCategoresType()

	// ********* обработчики для вспомогательных объектов ***********
	listHandlerTags := handlers.NewListBiZoneHandlerTags(supportObjectTags)
	listHandlerSecondaryCategores := handlers.NewListBiZoneHandlerSecondaryCategory(supportObjectCategores)

	listHandlerType := handlers.NewListBiZoneHandlerType(verifiedTypeObject)
	listHandlerStatus := handlers.NewListBiZoneHandlerStatus(verifiedStatusObject)
	listHandlerTenant := handlers.NewListBiZoneHandlerTenant(verifiedTenantObject)
	listHandlerPriority := handlers.NewListBiZoneHandlerPriority(verifiedPriorityObject)
	listHandlerCreatedBy := handlers.NewListBiZoneHandlerCreatedBy(verifiedCreatedByObject)
	listHandlerPrimaryCategory := handlers.NewListBiZoneHandlerPrimaryCategory(verifiedPrimaryCategoryObject)

	for msg := range chInput {
		var handlerIsExist bool

		//*** обработчик для объекта 'cases' ***
		if funcs, ok := listHandlerCases[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта 'cases.data' ***
		if funcs, ok := listHandlerData[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//***** обработчики для вспомогательных объектов *****
		//****************************************************
		// объект 'cases.tags'
		if funcs, ok := listHandlerTags[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'cases.secondary_category'
		if funcs, ok := listHandlerSecondaryCategores[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'cases.type'
		if funcs, ok := listHandlerType[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'cases.status'
		if funcs, ok := listHandlerStatus[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'cases.priority'
		if funcs, ok := listHandlerPriority[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'cases.primary_category'
		if funcs, ok := listHandlerPrimaryCategory[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'cases.tenant'
		if funcs, ok := listHandlerTenant[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'cases.created_by'
		if funcs, ok := listHandlerCreatedBy[msg.GetFieldBranch()]; ok {
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
	if errTmp := verifiedData.SetTags(supportObjectTags.GetTags()); errTmp != nil {
		err = errTmp
	}
	if errTmp := verifiedData.SetSecondaryCategory(supportObjectCategores.GetCategores()); errTmp != nil {
		err = errTmp
	}

	// собираем объект 'type'
	if errTmp := verifiedData.SetType(*verifiedTypeObject); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'status'
	if errTmp := verifiedData.SetStatus(*verifiedStatusObject); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'priority'
	if errTmp := verifiedData.SetPriority(*verifiedPriorityObject); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'primary_category'
	if errTmp := verifiedData.SetPrimaryCategory(*verifiedPrimaryCategoryObject); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'tenant'
	if errTmp := verifiedData.SetTenant(*verifiedTenantObject); errTmp != nil {
		err = errTmp
	}
	// собираем объект 'created_by'
	if errTmp := verifiedData.SetCreatedBy(*verifiedCreatedByObject); errTmp != nil {
		err = errTmp
	}

	if errTmp := verifiedMainObject.SetData(*verifiedData); errTmp != nil {
		err = errTmp
	}

	/*
		формирование дополнительной информации пока что не выполняется так как в кейсе нет ip адресов и id сенсоров

			// объект с дополнительной информацией по сенсорам и ip адресам
			additionalInformation := datamodels.AdditionalInformation{
				Sensors:     []datamodels.SensorInformation(nil),
				IpAddresses: []datamodels.IpAddressInformation(nil),
			}

			// формируем дополнительную информацию с идентификаторами сенсоров
			//additionalInformation.SetSensorInformation(CreateListSensorsForAlerts(verifiedData).GetSensorsInformation())
			// формируем дополнительную информацию с ip адресами
			//additionalInformation.SetIpAddressesInformation(CreateListIpAddresesForAlerts(verifiedData).GetIpAddressesInformation())
			// добавляем специальный, общий идентификатор
			// так как в alert и case uuid хранятся в разных местах и под разными названиями необходимо унифицировать
			// место хранения идентификатора для разных объектов, это позволит унифицировать обработчик отвечающий
			// за наполнение полей с доп. информацией по ip адресам и сенсорам
			additionalInformation.SetSpecialUUID(verifiedMainObject.GetUUID())

			if errTmp := verifiedMainObject.SetAdditionalInformation(additionalInformation); errTmp != nil {
				err = errTmp
			}
	*/

	return verifiedMainObject.GetUUID(), verifiedMainObject, listRawFields, err
}
