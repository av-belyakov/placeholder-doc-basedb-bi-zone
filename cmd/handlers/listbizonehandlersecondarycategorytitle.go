package handlers

// NewListBiZoneHandlerSecondaryCategory начальный обработчик событий объекта 'data.secondary_category.title.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerSecondaryCategoryTitle(cgt *SupportingStructureForTitlesType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- id ---
		"data.secondary_category.id": {
			func(a any) error {
				return cgt.HandlerValue(
					"data.secondary_category.id",
					a,
					cgt.titleTmp.SetAnyID,
				)
			}},
		//--- value ---
		"data.secondary_category.title.value": {
			func(a any) error {
				return cgt.HandlerValue(
					"data.secondary_category.title.value",
					a,
					cgt.titleTmp.SetAnyValue,
				)
			}},
		//--- language ---
		"data.secondary_category.title.language": {
			func(a any) error {
				return cgt.HandlerValue(
					"data.secondary_category.title.language",
					a,
					cgt.titleTmp.SetAnyLanguage,
				)
			}},
	}
}
