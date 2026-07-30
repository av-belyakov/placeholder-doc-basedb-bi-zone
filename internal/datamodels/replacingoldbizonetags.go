package datamodels

// ReplacingOldBiZoneTags заменяет или добавляет значения свойства VerifiedBiZoneAlert.Tags
func (va *VerifiedBiZoneIRPAlert) ReplacingOldBiZoneTags(incomingTypes []BiZoneIRPTag) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, va.Tags)
		if isExist {
			countReplacingFields += va.Tags[index].ReplacingOldBiZoneTag(incomingType)

			continue
		}

		va.Tags = append(va.Tags, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}

func compareTag(currentValue BiZoneIRPTag, incomingValue []BiZoneIRPTag) (int, bool) {
	for k, v := range incomingValue {
		if v.Name == currentValue.Name {
			return k, true
		}
	}

	return -1, false
}
