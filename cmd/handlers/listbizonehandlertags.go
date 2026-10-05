package handlers

// NewListBiZoneHandlerTags обработчик для значений типа 'tags.*' основного объекта
func NewListBiZoneHandlerTags(tg *SupportingStructureForTagsType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- name ---
		"data.tags.name": {
			func(a any) error {
				return tg.HandlerValue(
					"data.tags.name",
					a,
					tg.tagTmp.SetAnyName,
				)
			}},
		//--- color ---
		"data.tags.color": {
			func(a any) error {
				return tg.HandlerValue(
					"data.tags.color",
					a,
					tg.tagTmp.SetAnyColor,
				)
			}},
		//--- created ---
		"data.tags.created": {
			func(a any) error {
				return tg.HandlerValue(
					"data.tags.created",
					a,
					tg.tagTmp.SetAnyCreated,
				)
			}},
		//--- created_by.id ---
		"data.tags.created_by.id": {
			func(a any) error {
				return tg.HandlerValue(
					"data.tags.created_by.id",
					a,
					tg.tagTmp.SetAnyCreatedByID,
				)
			}},
		//--- created_by.username ---
		"data.tags.created_by.username": {
			func(a any) error {
				return tg.HandlerValue(
					"data.tags.created_by.username",
					a,
					tg.tagTmp.SetAnyCreatedByUsername,
				)
			}},
		//--- is_visible_for_customer ---
		"data.tags.is_visible_for_customer": {
			func(a any) error {
				return tg.HandlerValue(
					"data.tags.is_visible_for_customer",
					a,
					tg.tagTmp.SetAnyCreatedByID,
				)
			}},
	}
}
