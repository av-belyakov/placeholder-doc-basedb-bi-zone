package datamodels

// ReplacingOldIRPSecondaryCategoryRefs заменяет или добавляет значения свойства []CategoryRefType
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldIRPSecondaryCategoryRefs(incomingTypes []CategoryRefType) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareCategoryRef(incomingType, vi.SecondaryCategoryRef)
		if isExist {
			countReplacingFields += vi.SecondaryCategoryRef[index].ReplacingOldIRPIncidentCategoryRef(incomingType)

			continue
		}

		vi.SecondaryCategoryRef = append(vi.SecondaryCategoryRef, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}
