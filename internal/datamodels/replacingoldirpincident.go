package datamodels

import (
	"reflect"
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ReplacingOldIRPIncident заменяет старые значения типа VerifiedBiZoneIRPIncident новыми
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldIRPIncident(incomingType VerifiedBiZoneIRPIncident) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(vi).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "IssueTypeRef" {
				if data, ok := newStruct.Field(j).Interface().(IssueTypeRefType); ok {
					countReplacingFields += vi.IssueTypeRef.ReplacingOldIRPIncidentIssueTypeRef(data)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "ResolutionNameRef" {
				if data, ok := newStruct.Field(j).Interface().(CategoryRefType); ok {
					countReplacingFields += vi.ResolutionNameRef.ReplacingOldIRPIncidentCategoryRef(data)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "PrimaryCategoryRef" {
				if data, ok := newStruct.Field(j).Interface().(CategoryRefType); ok {
					countReplacingFields += vi.PrimaryCategoryRef.ReplacingOldIRPIncidentCategoryRef(data)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if tags, ok := newStruct.Field(j).Interface().([]IRPTagType); ok {
					countReplacingFields += vi.ReplacingOldIRPTags(tags)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Watchers" {
				if watchers, ok := newStruct.Field(j).Interface().([]WatchersType); ok {
					countReplacingFields += vi.ReplacingOldIRPWatchers(watchers)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "ObservedIocs" {
				if observedIocs, ok := newStruct.Field(j).Interface().([]ObservedIocsType); ok {
					countReplacingFields += vi.ReplacingOldIRPObservedIocs(observedIocs)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "SecondaryCategoryRef" {
				if secondaryCategoryRefs, ok := newStruct.Field(j).Interface().([]CategoryRefType); ok {
					countReplacingFields += vi.ReplacingOldIRPSecondaryCategoryRefs(secondaryCategoryRefs)
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

			stringTypeFields := []string{"KeyField", "ActivityDetected", "PlatformHostname", "SecondaryCategory"}
			if slices.Contains(stringTypeFields, typeOfCurrentStruct.Field(i).Name) {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для всех полей которые ниже мы пока пропускаем сравнение так как не ясен тип хранящийся в них, а бех этого
			// корректного сравнения всё равно не получится
			anyTypeFields := []string{"Emls", "Slas", "Badges", "WatchersId", "Attachments", "ParsedActivityDetected"}
			if slices.Contains(anyTypeFields, typeOfCurrentStruct.Field(i).Name) {
				continue
			}

			//поле с дополнительной информацией пропускаем, так как заменять эту
			//информацию не будем, а будет выполнятся добавление новой информации
			if typeOfCurrentStruct.Field(i).Name == "AdditionalInformation" {
				continue
			}

			// далее анализ только строковых полей, числовые или булевы поля пока не рассматриваем так как либо они
			// не должны менятся, например, поле ID, либо никак не влияют на подсчет статистики по инциденту
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
