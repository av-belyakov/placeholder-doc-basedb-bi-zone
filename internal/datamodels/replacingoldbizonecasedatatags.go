package datamodels

// ReplacingOldBiZoneCaseDataTags заменяет или добавляет значения свойства BiZoneIRPCaseData.Tags
func (cd *BiZoneIRPCaseData) ReplacingOldBiZoneCaseDataTags(incomingTypes []BiZoneIRPTag) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, cd.Tags)
		if isExist {
			countReplacingFields += cd.Tags[index].ReplacingOldBiZoneCaseDataTag(incomingType)

			continue
		}

		cd.Tags = append(cd.Tags, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}
