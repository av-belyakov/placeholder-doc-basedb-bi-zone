package datamodels

import (
	"reflect"
)

// ReplacingOldBiZoneCaseDataTenant заменяет старые значения типа BiZoneIRPTenant новыми
func (tnt *BiZoneIRPTenant) ReplacingOldBiZoneCaseDataTenant(incomingType BiZoneIRPTenant) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(tnt).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
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
