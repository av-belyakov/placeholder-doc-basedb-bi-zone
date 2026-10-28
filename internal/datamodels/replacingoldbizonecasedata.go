package datamodels

import (
	"reflect"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ReplacingOldBiZoneCaseData заменяет старые значения типа BiZoneIRPCaseData новыми
func (cd *BiZoneIRPCaseData) ReplacingOldBiZoneCaseData(incomingType BiZoneIRPCaseData) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(cd).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Type" {
				if cdType, ok := newStruct.Field(j).Interface().(BiZoneIRPType); ok {
					countReplacingFields += cd.Type.ReplacingOldBiZoneCaseDataType(cdType)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Tenant" {
				if tenant, ok := newStruct.Field(j).Interface().(BiZoneIRPTenant); ok {
					countReplacingFields += cd.Tenant.ReplacingOldBiZoneCaseDataTenant(tenant)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Priority" {
				if priority, ok := newStruct.Field(j).Interface().(BiZoneIRPPriority); ok {
					countReplacingFields += cd.Priority.ReplacingOldBiZoneCaseDataPriority(priority)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Status" {
				if status, ok := newStruct.Field(j).Interface().(BiZoneIRPStatus); ok {
					countReplacingFields += cd.Status.ReplacingOldBiZoneCaseDataStatus(status)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Organization" {
				if organization, ok := newStruct.Field(j).Interface().(BiZoneIRPOrganization); ok {
					countReplacingFields += cd.Organization.ReplacingOldBiZoneCaseDataOrganization(organization)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "PrimaryCategory" {
				if pCategory, ok := newStruct.Field(j).Interface().(BiZoneIRPCategory); ok {
					countReplacingFields += cd.PrimaryCategory.ReplacingOldBiZoneCaseDataCategory(pCategory)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "CreatedBy" {
				if createdBy, ok := newStruct.Field(j).Interface().(BiZoneIRPCreatedBy); ok {
					countReplacingFields += cd.CreatedBy.ReplacingOldBiZoneCaseDataCreatedBy(createdBy)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "SecondaryCategory" {
				if secondaryCategory, ok := newStruct.Field(j).Interface().([]BiZoneIRPCategory); ok {
					countReplacingFields += cd.ReplacingOldBiZoneCaseDataSecondaryCategory(secondaryCategory)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "PlatformHostname" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "DetectionRules" {
				if list, ok := supportingfunctions.ReplacingSlice[uint64](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if tags, ok := newStruct.Field(j).Interface().([]IRPTagType); ok {
					countReplacingFields += cd.ReplacingOldIRPTags(tags)
				}

				continue
			}

			if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
				if !currentStruct.Field(i).CanSet() {
					continue
				}

				if str, ok := newStruct.Field(j).Interface().(string); ok {
					//не обновлять текущие значения новыми пустыми значениями
					if str == "" {
						continue
					}
				}

				currentStruct.Field(i).Set(newStruct.Field(j))
				countReplacingFields++
			}
		}
	}

	return countReplacingFields
}
