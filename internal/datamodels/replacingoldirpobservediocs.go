package datamodels

import (
	"fmt"
	"reflect"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ReplacingOldIRPObservedIocs заменяет или добавляет значения свойства []ObservedIocsType
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldIRPObservedIocs(incomingTypes []ObservedIocsType) []string {
	var replacingFields []string

	for _, incomingType := range incomingTypes {
		index, isExist := compareObservedIocs(incomingType, vi.ObservedIocs)
		if isExist {
			replacingFields = append(replacingFields, vi.ObservedIocs[index].ReplacingOldIRPObservedIocsType(incomingType)...)

			continue
		}

		vi.ObservedIocs = append(vi.ObservedIocs, incomingType)
		replacingFields = append(replacingFields, fmt.Sprintf("added object with ioc '%s'", incomingType.Ioc))
	}

	return replacingFields
}

// ReplacingOldIRPObservedIocsType заменяет старые значения типа ObservedIocs новыми
func (iocs *ObservedIocsType) ReplacingOldIRPObservedIocsType(incomingType ObservedIocsType) []string {
	var replacingFields []string

	currentStruct := reflect.ValueOf(iocs).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Category" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); !ok {
					currentStruct.Field(i).Set(list)
					replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
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
				replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
			}
		}
	}

	return replacingFields
}
