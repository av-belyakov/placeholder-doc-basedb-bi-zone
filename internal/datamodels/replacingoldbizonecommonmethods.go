package datamodels

func compareTag(currentValue BiZoneIRPTag, incomingValue []BiZoneIRPTag) (int, bool) {
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
