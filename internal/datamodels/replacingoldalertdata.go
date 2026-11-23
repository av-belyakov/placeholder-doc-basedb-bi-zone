package datamodels

import (
	"reflect"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ReplacingOldBiZoneData заменяет старые значения типа BiZoneIRPAlertData новыми
func (d *BiZoneIRPAlertData) ReplacingOldBiZoneData(incomingType BiZoneIRPAlertData) []string {
	var replacingFields []string

	currentStruct := reflect.ValueOf(d).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(incomingType)
	typeOfNewStruct := newStruct.Type()

	for i := range currentStruct.NumField() {
		for j := range newStruct.NumField() {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "Tags" ||
				typeOfCurrentStruct.Field(i).Name == "UnmappedDstEndpointArray" ||
				typeOfCurrentStruct.Field(i).Name == "UnmappedHomeEndpointArray" {
				if list, ok := supportingfunctions.ReplacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "DetectionPattern" ||
				typeOfCurrentStruct.Field(i).Name == "UnmappedAgentArray" {
				if list, ok := supportingfunctions.ReplacingSlice[uint64](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					replacingFields = append(replacingFields, typeOfCurrentStruct.Field(i).Name)
				}

				continue
			}

			if typeOfCurrentStruct.Field(i).Name == "DataSecurity" {
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
