package datamodels

import "fmt"

// ReplacingOldBiZoneSnapshots заменяет или добавляет значения свойства VerifiedBiZoneAlert.Snapshots
func (va *VerifiedBiZoneIRPAlert) ReplacingOldBiZoneSnapshots(incomingTypes []BiZoneIRPSnapshot) []string {
	var replacingFields []string

	for _, incomingType := range incomingTypes {
		index, isExist := compareSnapshot(incomingType, va.Snapshots)
		if isExist {
			replacingFields = append(replacingFields, va.Snapshots[index].ReplacingOldBiZOneSnapshot(incomingType)...)

			continue
		}

		va.Snapshots = append(va.Snapshots, incomingType)
		replacingFields = append(replacingFields, fmt.Sprintf("added new snapshots with cmdb_id '%s'", incomingType.CmdbId))
	}

	return replacingFields
}
