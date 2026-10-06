package handlers

// NewListBiZoneHandlerSecondaryCategory начальный обработчик событий объекта 'data.secondary_category.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerSecondaryCategory(cg *SupportingStructureForCategoresType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- id ---
		"data.secondary_category.id": {
			func(a any) error {
				return cg.HandlerValue(
					"data.secondary_category.id",
					a,
					cg.categoryTmp.SetAnyID,
				)
			}},
	}
}
