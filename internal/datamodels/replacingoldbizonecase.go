package datamodels

import (
	"reflect"
)

// RepalcingOldBiZoneCase заменяет старые значения типа VerifiedBiZoneCase новыми
func (vc *VerifiedBiZoneIRPCase) RepalcingOldBiZoneCase(incomingType VerifiedBiZoneIRPCase) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(vc).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Data" {
				if data, ok := newStruct.Field(j).Interface().(BiZoneIRPCaseData); ok {
					countReplacingFields += vc.Data.ReplacingOldBiZoneCaseData(data)
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
