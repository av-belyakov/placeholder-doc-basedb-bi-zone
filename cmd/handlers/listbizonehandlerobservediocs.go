package handlers

// NewListBiZoneHandlerObservedIocs начальный обработчик событий объекта 'observed_iocs.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerObservedIocs(iocs *SupportingStructureForObservedIocsType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- ioc ---
		"observed_iocs.ioc": {
			func(a any) error {
				return iocs.HandlerValue(
					"observed_iocs.ioc",
					a,
					iocs.observedIocsTmp.SetAnyIoc,
				)
			}},
		//--- ioc_type ---
		"observed_iocs.ioc_type": {
			func(a any) error {
				return iocs.HandlerValue(
					"observed_iocs.ioc_type",
					a,
					iocs.observedIocsTmp.SetAnyIocType,
				)
			}},
		//ниже работа со срезам содержащими простые типы
		"observed_iocs.category": {iocs.observedIocsTmp.SetAnyCategoryElement},
	}
}
