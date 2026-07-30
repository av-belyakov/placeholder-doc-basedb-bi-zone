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

func compareSnapshot(currentValue BiZoneIRPSnapshot, incomingValue []BiZoneIRPSnapshot) (int, bool) {
	for k, v := range incomingValue {
		if v.CMDBID == currentValue.CMDBID {
			return k, true
		}
	}

	return -1, false
}
