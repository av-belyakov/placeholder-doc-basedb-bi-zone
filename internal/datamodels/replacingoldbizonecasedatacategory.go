package datamodels

import (
	"reflect"
)

// ReplacingOldBiZoneCaseDataCategory заменяет старые значения типа BiZoneIRPCategory новыми
func (t *BiZoneIRPCategory) ReplacingOldBiZoneCaseDataCategory(incomingType BiZoneIRPCategory) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(t).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// свойство 'title' пока сравнивать не будем
			if typeOfCurrentStruct.Field(i).Name == "Title" {
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
