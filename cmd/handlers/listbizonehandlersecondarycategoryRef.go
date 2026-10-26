package handlers

// NewListBiZoneHandlerSecondaryCategoryRef начальный обработчик событий объекта 'secondary_category.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerSecondaryCategoryRef(cr *SupportingStructureForCategoreRefsType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- id ---
		"secondary_category_ref.id": {
			func(a any) error {
				return cr.HandlerValue(
					"secondary_category_ref.id",
					a,
					cr.categoryRefTmp.SetAnyId,
				)
			}},
		//--- title ---
		"secondary_category_ref.title": {
			func(a any) error {
				return cr.HandlerValue(
					"secondary_category_ref.title",
					a,
					cr.categoryRefTmp.SetAnyTitle,
				)
			}},
		//--- description ---
		"secondary_category_ref.description": {
			func(a any) error {
				return cr.HandlerValue(
					"secondary_category_ref.description",
					a,
					cr.categoryRefTmp.SetAnyDescription,
				)
			}},
	}
}
