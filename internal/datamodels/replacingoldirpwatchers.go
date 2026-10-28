package datamodels

import "reflect"

// ReplacingOldIRPWatchers заменяет или добавляет значения свойства []WatchersType
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldIRPWatchers(incomingTypes []WatchersType) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareWatcher(incomingType, vi.Watchers)
		if isExist {
			countReplacingFields += vi.Watchers[index].ReplacingOldIRPWatcher(incomingType)

			continue
		}

		vi.Watchers = append(vi.Watchers, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}

// ReplacingOldIRPWatcher заменяет или добавляет значения свойства WatchersType
func (w *WatchersType) ReplacingOldIRPWatcher(incomingType WatchersType) int {
	var countReplacingFields int

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
				countReplacingFields++
			}
		}
	}

	return countReplacingFields
}
