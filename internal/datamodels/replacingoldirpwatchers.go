package datamodels

import (
	"fmt"
	"reflect"
)

// ReplacingOldIRPWatchers заменяет или добавляет значения свойства []WatchersType
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldIRPWatchers(incomingTypes []WatchersType) []string {
	var replacingFields []string

	for _, incomingType := range incomingTypes {
		index, isExist := compareWatcher(incomingType, vi.Watchers)
		if isExist {
			replacingFields = append(replacingFields, vi.Watchers[index].ReplacingOldIRPWatcher(incomingType)...)

			continue
		}

		vi.Watchers = append(vi.Watchers, incomingType)
		replacingFields = append(replacingFields, fmt.Sprintf("added object with id '%d'", incomingType.Id))
	}

	return replacingFields
}

// ReplacingOldIRPWatcher заменяет или добавляет значения свойства WatchersType
func (w *WatchersType) ReplacingOldIRPWatcher(incomingType WatchersType) []string {
	var replacingFields []string

	currentStruct := reflect.ValueOf(w).Elem()
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
