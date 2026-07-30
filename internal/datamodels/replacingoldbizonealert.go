package datamodels

import (
	"reflect"
)

// RepalcingOldBiZoneAlert заменяет старые значения типа VerifiedBiZoneAlert новыми
func (va *VerifiedBiZoneIRPAlert) RepalcingOldBiZoneAlert(incomingType VerifiedBiZoneIRPAlert) int {
	var countReplacingFields int

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
					countReplacingFields += va.ReplacingOldBiZoneSnapshots(snapshots)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if tags, ok := newStruct.Field(j).Interface().([]BiZoneIRPTag); ok {
					countReplacingFields += va.ReplacingOldBiZoneTags(tags)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Data" {
				if data, ok := newStruct.Field(j).Interface().(BiZoneIRPData); ok {
					countReplacingFields += va.Data.ReplacingOldBiZoneData(data)
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
				countReplacingFields++
			}
		}
	}

	return countReplacingFields
}
