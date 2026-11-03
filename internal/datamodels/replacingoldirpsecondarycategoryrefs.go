package datamodels

import "fmt"

// ReplacingOldIRPSecondaryCategoryRefs заменяет или добавляет значения свойства []CategoryRefType
func (vi *VerifiedBiZoneIRPIncident) ReplacingOldIRPSecondaryCategoryRefs(incomingTypes []CategoryRefType) []string {
	var replacingFields []string

	for _, incomingType := range incomingTypes {
		index, isExist := compareCategoryRef(incomingType, vi.SecondaryCategoryRef)
		if isExist {
			replacingFields = append(replacingFields, vi.SecondaryCategoryRef[index].ReplacingOldIRPIncidentCategoryRef(incomingType)...)

			continue
		}

		vi.SecondaryCategoryRef = append(vi.SecondaryCategoryRef, incomingType)
		replacingFields = append(replacingFields, fmt.Sprintf("added object with id '%s'", incomingType.Id))
	}

	return replacingFields
}
