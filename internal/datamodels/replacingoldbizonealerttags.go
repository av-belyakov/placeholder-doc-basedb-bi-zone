package datamodels

// ReplacingOldBiZoneAlertTags заменяет или добавляет значения свойства VerifiedBiZoneAlert.Tags
func (va *VerifiedBiZoneIRPAlert) ReplacingOldBiZoneAlertTags(incomingTypes []BiZoneIRPTag) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, va.Tags)
		if isExist {
			countReplacingFields += va.Tags[index].ReplacingOldBiZoneAlertTag(incomingType)

			continue
		}

		va.Tags = append(va.Tags, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}
