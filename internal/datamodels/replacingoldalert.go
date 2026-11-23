package datamodels

import (
	"reflect"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ReplacingOldBiZoneAlert заменяет старые значения типа VerifiedBiZoneAlert новыми
func (va *VerifiedBiZoneIRPAlert) ReplacingOldBiZoneAlert(incomingType VerifiedBiZoneIRPAlert) []string {
	var replacingFields []string

	currentStruct := reflect.ValueOf(va).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Snapshots" {
				if snapshots, ok := newStruct.Field(j).Interface().([]BiZoneIRPSnapshot); ok {
					replacingFields = append(replacingFields, va.ReplacingOldBiZoneSnapshots(snapshots)...)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if tags, ok := newStruct.Field(j).Interface().([]IRPTagType); ok {
					replacingFields = append(replacingFields, va.ReplacingOldTags(tags)...)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "AffectedLogSources" {
				//if data, ok := newStruct.Field(j).Interface().([]string); ok {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
				}
				//}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Data" {
				if data, ok := newStruct.Field(j).Interface().(BiZoneIRPAlertData); ok {
					replacingFields = append(replacingFields, va.Data.ReplacingOldBiZoneData(data)...)
				}

				continue
			}

			//поле с дополнительной информацией пропускаем, так как заменять эту
			//информацию не будем, а будет выполнятся добавление новой информации
			if typeOfCurrentStruct.Field(i).Name == "AdditionalInformation" {
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
				replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
			}
		}
	}

	return replacingFields
}

/*
		replacingFields = append(replacingFields, fmt.Sprintf("added new mac_address to snapshots with cmdb_id '%s'", incomingType.CmdbId))
	replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
*/
