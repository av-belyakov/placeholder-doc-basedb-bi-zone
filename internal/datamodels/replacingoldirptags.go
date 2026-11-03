package datamodels

import (
	"fmt"
	"reflect"
)

// ReplacingOldIRPTags заменяет или добавляет значения свойства []IRPTagType
func (cd *BiZoneIRPCaseData) ReplacingOldIRPTags(incomingTypes []IRPTagType) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, cd.Tags)
		if isExist {
			countReplacingFields += cd.Tags[index].ReplacingOldIrpTag(incomingType)

			continue
		}

		cd.Tags = append(cd.Tags, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}

// ReplacingOldIrpTag заменяет старые значения типа BiZoneTag новыми
func (t *IRPTagType) ReplacingOldIrpTag(incomingType IRPTagType) int {
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

// ReplacingOldIRPTags заменяет или добавляет значения свойства []IRPTagType
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldIRPTags(incomingTypes []IRPTagType) []string {
	var replacingFields []string

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, vi.Tags)
		if isExist {
			replacingFields = append(replacingFields, vi.Tags[index].ReplacingOldIRPTag(incomingType)...)

			continue
		}

		vi.Tags = append(vi.Tags, incomingType)
		replacingFields = append(replacingFields, fmt.Sprintf("added object with name '%s'", incomingType.Name))
	}

	return replacingFields
}

// ReplacingOldIRPTag заменяет старые значения типа BiZoneTag новыми
func (t *IRPTagType) ReplacingOldIRPTag(incomingType IRPTagType) []string {
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
