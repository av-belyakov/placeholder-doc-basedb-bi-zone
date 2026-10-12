package datamodels

// ReplacingOldBiZoneSnapshots заменяет или добавляет значения свойства VerifiedBiZoneAlert.Snapshots
func (va *VerifiedBiZoneIRPAlert) ReplacingOldBiZoneSnapshots(incomingTypes []BiZoneIRPSnapshot) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareSnapshot(incomingType, va.Snapshots)
		if isExist {
			countReplacingFields += va.Snapshots[index].ReplacingOldBiZOneSnapshot(incomingType)

			continue
		}

		va.Snapshots = append(va.Snapshots, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}
