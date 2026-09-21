package handlers

// NewListBiZoneHandlerTags обработчик для значений типа 'tags.*' основного объекта
func NewListBiZoneHandlerTags(tg *SupportingStructureForTagsType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- name ---
		"tags.name": {
			func(a any) error {
				return tg.HandlerValue(
					"tags.name",
					a,
					tg.GetTagTmp().SetAnyName,
				)
			}},
		//--- color ---
		"tags.color": {
			func(a any) error {
				return tg.HandlerValue(
					"tags.color",
					a,
					tg.GetTagTmp().SetAnyColor,
				)
			}},
		//--- created ---
		"tags.created": {
			func(a any) error {
				return tg.HandlerValue(
					"tags.created",
					a,
					tg.GetTagTmp().SetAnyCreated,
				)
			}},
		//--- created_by.id ---
		"tags.created_by.id": {
			func(a any) error {
				return tg.HandlerValue(
					"tags.created_by.id",
					a,
					tg.GetTagTmp().SetAnyCreatedByID,
				)
			}},
		//--- created_by.username ---
		"tags.created_by.username": {
			func(a any) error {
				return tg.HandlerValue(
					"tags.created_by.username",
					a,
					tg.GetTagTmp().SetAnyCreatedByUsername,
				)
			}},
		//--- is_visible_for_customer ---
		"tags.is_visible_for_customer": {
			func(a any) error {
				return tg.HandlerValue(
					"tags.is_visible_for_customer",
					a,
					tg.GetTagTmp().SetAnyCreatedByID,
				)
			}},
	}
}
