package documentgenerator

import (
	"fmt"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// IRPIncidentGenerator генерирует верифицированный объект типа 'incident' получаемый при запросе к IRP, не от Kafka
func IRPIncidentGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *datamodels.VerifiedBiZoneIRPCase, map[string]string, error) {
	// список не обработанных полей
	var (
		listRawFields map[string]string = make(map[string]string)

		err error
	)

	/*

	   Всё что ниже, нужно заменить на новые обработчики

	*/

	verifiedMainObject := datamodels.NewVerifiedBiZoneIRPIncident()
	verifiedIssueTypeRef := datamodels.NewIssueTypeRefType()
	verifiedResolutionNameRef := datamodels.NewCategoryRefType()
	verifiedPrimaryCategoryRef := datamodels.NewCategoryRefType()

	//********* основные обработчики **********
	listHandlerIncident := handlers.NewListBiZoneHandlerIncident(verifiedMainObject)
	listHandlerIssueTypeRef := handlers.NewListBiZoneHandlerIssueTypeRef(verifiedIssueTypeRef)
	listHandlerResolutionNameRef := handlers.NewListBiZoneHandlerResolutionNameRef(verifiedResolutionNameRef)
	listHandlerPrimaryCategoryRef := handlers.NewListBiZoneHandlerPrimaryCategoryRef(verifiedPrimaryCategoryRef)

	//******** вспомогательные объекты ********
	supportObjectTags := handlers.NewSupportingStructureForTagsType()
	supportObjectObservedIocs := handlers.NewSupportingStructureForObservedIocsType()
	supportObjectSecondaryCategoreRefs := handlers.NewSupportingStructureForCategoreRefsType()

	// ********* обработчики для вспомогательных объектов ***********
	listHandlerTags := handlers.NewListBiZoneHandlerTags(supportObjectTags)
	listHandlerObservedIocs := handlers.NewListBiZoneHandlerObservedIocs(supportObjectObservedIocs)
	listHandlerSecondaryCategoresRef := handlers.NewListBiZoneHandlerSecondaryCategoryRef(supportObjectSecondaryCategoreRefs)

	// объект с дополнительной информацией по сенсорам и ip адресам
	additionalInformation := datamodels.AdditionalInformation{
		Sensors:     []datamodels.SensorInformation(nil),
		IpAddresses: []datamodels.IpAddressInformation(nil),
	}

	for msg := range chInput {
		var handlerIsExist bool

		//*** обработчик для основного объекта ***
		if funcs, ok := listHandlerIncident[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		// объект 'tags.*'
		if funcs, ok := listHandlerTags[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'observed_iocs.*'
		if funcs, ok := listHandlerObservedIocs[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'secondary_category.*'
		if funcs, ok := listHandlerSecondaryCategoresRef[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		//*** обработчик для под объекта 'issue_type_ref.*' ***
		if funcs, ok := listHandlerIssueTypeRef[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта 'primary_category_ref.*' ***
		if funcs, ok := listHandlerPrimaryCategoryRef[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта 'resolution_name_ref.*' ***
		if funcs, ok := listHandlerResolutionNameRef[msg.GetFieldBranch()]; ok {
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
	// объект 'tags'
	if errTmp := verifiedMainObject.SetTags(supportObjectTags.GetTags()); errTmp != nil {
		err = errTmp
	}

	// объект 'observed_iocs'
	if errTmp := verifiedMainObject.SetObservedIocs(supportObjectObservedIocs.GetObservedIocs()); errTmp != nil {
		err = errTmp
	}

	// объект 'secondary_category_ref'
	if errTmp := verifiedMainObject.SetSecondaryCategoryRef(supportObjectSecondaryCategoreRefs.GetCategoreRefs()); errTmp != nil {
		err = errTmp
	}

	// объект 'issue_type_ref'
	if errTmp := verifiedMainObject.SetIssueTypeRef(*verifiedIssueTypeRef); errTmp != nil {
		err = errTmp
	}

	// объект 'primary_category_ref'
	if errTmp := verifiedMainObject.SetPrimaryCategoryRef(*verifiedPrimaryCategoryRef); errTmp != nil {
		err = errTmp
	}

	// объект 'resolution_name_ref'
	if errTmp := verifiedMainObject.SetResolutionNameRef(*verifiedResolutionNameRef); errTmp != nil {
		err = errTmp
	}

	/*
		надо посмотреть возможно какие то поля надо верефицировать или преобразовывать
		ещё надо надо получить список сенсоров и ip адресов
	*/

	// формируем дополнительную информацию с идентификаторами сенсоров
	additionalInformation.SetSensorInformation(CreateListSensorsForCases(verifiedData).GetSensorsInformation())

	// формируем дополнительную информацию с ip адресами
	additionalInformation.SetIpAddressesInformation(CreateListIpAddresesForCases(verifiedData).GetIpAddressesInformation())

	// добавляем специальный, общий идентификатор
	additionalInformation.SetSpecialUUID(verifiedMainObject.GetID())

	if errTmp := verifiedMainObject.SetAdditionalInformation(additionalInformation); errTmp != nil {
		err = errTmp
	}

	return verifiedMainObject.GetID(), verifiedMainObject, listRawFields, err
}
