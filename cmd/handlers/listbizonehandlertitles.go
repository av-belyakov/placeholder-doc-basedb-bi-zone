package handlers

// NewListBiZoneHandlerTitles обработчик для значений типа 'title.*' основного объекта
func NewListBiZoneHandlerTitles(tg *SupportingStructureForTitlesType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- value ---
		"title.value": {
			func(a any) error {
				return tg.HandlerValue(
					"title.value",
					a,
					tg.titleTmp.SetAnyValue,
				)
			}},
		//--- language ---
		"title.language": {
			func(a any) error {
				return tg.HandlerValue(
					"title.language",
					a,
					tg.titleTmp.SetAnyLanguage,
				)
			}},
	}
}
