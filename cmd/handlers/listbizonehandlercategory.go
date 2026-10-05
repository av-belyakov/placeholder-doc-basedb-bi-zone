package handlers

// NewListBiZoneHandlerSecondaryCategory начальный обработчик событий объекта 'secondary_category.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerSecondaryCategory(cg *SupportingStructureForCategoresType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- id ---
		"secondary_category.id": {
			func(a any) error {
				return cg.HandlerValue(
					"secondary_category.id",
					a,
					cg.categoryTmp.SetAnyID,
				)
			}},
		//--- title ---
		"secondary_category.title": {
			func(a any) error {
				return cg.HandlerValue(
					"secondary_category.title",
					a,
					cg.categoryTmp.SetAnyTitleElement,
				)
			}},
	}
}
