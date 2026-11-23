package datamodels

import (
	"fmt"
	"reflect"
)

// ReplacingOldTags заменяет или добавляет значения свойства []IRPTagType
func (va *VerifiedBiZoneIRPAlert) ReplacingOldTags(incomingTypes []IRPTagType) []string {
	var replacingFields []string

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, va.Tags)
		if isExist {
			replacingFields = append(replacingFields, va.Tags[index].ReplacingOldTag(incomingType)...)

			continue
		}

		va.Tags = append(va.Tags, incomingType)
		replacingFields = append(replacingFields, fmt.Sprintf("added object with name '%s'", incomingType.Name))
	}

	return replacingFields
}

// ReplacingOldTags заменяет или добавляет значения свойства []IRPTagType
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldTags(incomingTypes []IRPTagType) []string {
	var replacingFields []string

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, vi.Tags)
		if isExist {
			replacingFields = append(replacingFields, vi.Tags[index].ReplacingOldTag(incomingType)...)

			continue
		}

		vi.Tags = append(vi.Tags, incomingType)
		replacingFields = append(replacingFields, fmt.Sprintf("added object with name '%s'", incomingType.Name))
	}

	return replacingFields
}

// ReplacingOldTag заменяет старые значения типа BiZoneTag новыми
func (t *IRPTagType) ReplacingOldTag(incomingType IRPTagType) []string {
	var replacingFields []string

	currentStruct := reflect.ValueOf(t).Elem()
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
				replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
			}
		}
	}

	return replacingFields
}
