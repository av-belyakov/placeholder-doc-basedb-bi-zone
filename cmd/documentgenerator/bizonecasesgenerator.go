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
	supportObjectSecondaryCategores := handlers.NewSupportingStructureForCategoresType()

	//supportObjectType := handlers.NewSupportingStructureForTitlesType()
	//supportObjectStatus := handlers.NewSupportingStructureForTitlesType()
	//supportObjectPriority := handlers.NewSupportingStructureForTitlesType()
	//supportObjectPrimaryCategory := handlers.NewSupportingStructureForTitlesType()

	// ********* обработчики для вспомогательных объектов ***********
	listHandlerTags := handlers.NewListBiZoneHandlerTags(supportObjectTags)
	listHandlerSecondaryCategores := handlers.NewListBiZoneHandlerSecondaryCategory(supportObjectSecondaryCategores)

	listHandlerType := handlers.NewListBiZoneHandlerType(verifiedTypeObject)
	listHandlerStatus := handlers.NewListBiZoneHandlerStatus(verifiedStatusObject)
	listHandlerTenant := handlers.NewListBiZoneHandlerTenant(verifiedTenantObject)
	listHandlerPriority := handlers.NewListBiZoneHandlerPriority(verifiedPriorityObject)
	listHandlerCreatedBy := handlers.NewListBiZoneHandlerCreatedBy(verifiedCreatedByObject)
	listHandlerPrimaryCategory := handlers.NewListBiZoneHandlerPrimaryCategory(verifiedPrimaryCategoryObject)

	for msg := range chInput {
		var handlerIsExist bool

		//*** обработчик для основного объекта 'cases' ***
		if funcs, ok := listHandlerCases[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта 'data' ***
		if funcs, ok := listHandlerData[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			continue
		}

		//***** обработчики для вспомогательных объектов *****
		//****************************************************
		// объект 'data.type'
		if funcs, ok := listHandlerType[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'data.tags'
		if funcs, ok := listHandlerTags[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'data.secondary_category'
		if funcs, ok := listHandlerSecondaryCategores[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range funcs {
				f(msg.GetValue())
			}

			//			continue
		}

		// объект 'data.status'
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

	//fmt.Println("--------------------------")
	//fmt.Println("supportObjectSecondaryCategores:")

	/*categores := supportObjectSecondaryCategores.GetCategores()
	titles := supportObjectSecondaryCategoryTitle.GetTitles()
	for k, v := range categores {
		//fmt.Printf("%d.\n\tID:'%s', Title:'%v'\n", k, v.GetID(), v.GetTitle())

		index := slices.IndexFunc(titles, func(ttemp handlers.TitleTemporary) bool {
			return ttemp.GetID() == v.GetID()
		})
		if index != -1 {
			categores[k].Title = append(categores[k].Title, datamodels.BiZoneIRPTitle{
				Value:    titles[index].Value,
				Language: titles[index].Language,
			})
		}
	}*/
	//fmt.Printf("func 'BiZoneCasesGenerator', supportObjectSecondaryCategores:'%+v'\n", supportObjectSecondaryCategores.GetCategores())

	//fmt.Println("\nsupportObjectSecondaryCategoryTitle verifiedSecondaryCategoryType:")
	//for k, v := range supportObjectSecondaryCategoryTitle.GetIds() {
	//for k, v := range verifiedSecondaryCategoryType {
	//	fmt.Printf("%d.\n  %s\n", k, v.ID)
	//	for i, item := range v.Title {
	//		fmt.Printf("    %d. '%s'\n", i, item)
	//	}
	//}
	//for k, v := range supportObjectSecondaryCategoryTitle.GetTitles() {
	//	fmt.Printf("%d.\n\t'%s'\n", k, v)
	//}
	//fmt.Printf("func 'BiZoneCasesGenerator', supportObjectSecondaryCategoryTitle:'%+v'\n", supportObjectSecondaryCategoryTitle.GetTitles())
	//fmt.Println("--------------------------")

	// собираем все объекты в один
	// собираем объект 'tags'
	if errTmp := verifiedData.SetTags(supportObjectTags.GetTags()); errTmp != nil {
		err = errTmp
	}

	// верифицируем тип 'data.secondary_catigory' на основе правил установленных в функции verificationSecondaryCategoryType
	verifiedSecondaryCategoryType := verificationSecondaryCategoryType(supportObjectSecondaryCategores.GetCategores())
	// собираем объект 'secondary_category'
	if errTmp := verifiedData.SetSecondaryCategory(verifiedSecondaryCategoryType); errTmp != nil {
		err = errTmp
	}

	// собираем объект 'type'
	if errTmp := verifiedData.SetType(*verifiedTypeObject); errTmp != nil {
		err = errTmp
	}
	// верифицируем тип 'data.status' на основе правил установленных в функции verificationStatusType
	verifiedStatusType := verificationStatusType(*verifiedStatusObject.Get())
	// собираем объект 'status'
	/*
			почему что здесь что в priority получается пустой срез если использовать функции верификации
			и ещё следующие поля JSON не обрабатываются, видимо нет обработчиков
				data.created_by.id:4
		        data.created_by.username:803.a.egorov@cert.gov.ru
		        data.tenant.id:8030064
		        data.tenant.name:Рязань Радиозавод
	*/
	if errTmp := verifiedData.SetStatus(verifiedStatusType); errTmp != nil {
		err = errTmp
	}
	// верифицируем тип 'data.priority' на основе правил установленных в функции verificationPriorityType
	verifiedPriorityType := verificationPriorityType(*verifiedPriorityObject.Get())
	// собираем объект 'priority'
	if errTmp := verifiedData.SetPriority(verifiedPriorityType); errTmp != nil {
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
