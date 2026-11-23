package datamodels

func compareCategory(currentValue BiZoneIRPCategory, incomingValue []BiZoneIRPCategory) (int, bool) {
	for k, v := range incomingValue {
		if v.ID == currentValue.ID {
			return k, true
		}
	}

	return -1, false
}

func compareCategoryRef(currentValue CategoryRefType, incomingValue []CategoryRefType) (int, bool) {
	for k, v := range incomingValue {
		if v.Id == currentValue.Id {
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

func compareWatcher(currentValue WatchersType, incomingValue []WatchersType) (int, bool) {
	for k, v := range incomingValue {
		if v.Id == currentValue.Id {
			return k, true
		}
	}

	return -1, false
}

func compareObservedIocs(currentValue ObservedIocsType, incomingValue []ObservedIocsType) (int, bool) {
	for k, v := range incomingValue {
		if v.Ioc == currentValue.Ioc {
			return k, true
		}
	}

	return -1, false
}
