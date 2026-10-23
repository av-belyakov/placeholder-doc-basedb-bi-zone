package datamodels

func compareCategory(currentValue BiZoneIRPCategory, incomingValue []BiZoneIRPCategory) (int, bool) {
	for k, v := range incomingValue {
		if v.ID == currentValue.ID {
			return k, true
		}
	}

	return -1, false
}

func compareTag(currentValue IRPTagType, incomingValue []IRPTagType) (int, bool) {
	for k, v := range incomingValue {
		if v.Name == currentValue.Name {
			return k, true
		}
	}

	return -1, false
}

func compareSnapshot(currentValue BiZoneIRPSnapshot, incomingValue []BiZoneIRPSnapshot) (int, bool) {
	for k, v := range incomingValue {
		if v.CmdbId == currentValue.CmdbId {
			return k, true
		}
	}

	return -1, false
}
