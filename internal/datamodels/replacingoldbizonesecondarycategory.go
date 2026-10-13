package datamodels

// ReplacingOldBiZoneCaseDataSecondaryCategory заменяет или добавляет значения свойства BiZoneIRPCaseData.SecondaryCategory
func (cd *BiZoneIRPCaseData) ReplacingOldBiZoneCaseDataSecondaryCategory(incomingTypes []BiZoneIRPCategory) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareCategory(incomingType, cd.SecondaryCategory)
		if isExist {
			countReplacingFields += cd.SecondaryCategory[index].ReplacingOldBiZoneCaseDataCategory(incomingType)

			continue
		}

		cd.SecondaryCategory = append(cd.SecondaryCategory, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}
