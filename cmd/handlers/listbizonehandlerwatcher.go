package handlers

// NewListBiZoneHandlerWatchers начальный обработчик событий объекта 'watchers.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerWatchers(w *SupportingStructureForWatchersType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- id ---
		"watchers.id": {
			func(a any) error {
				return w.HandlerValue(
					"watchers.id",
					a,
					w.watchersTmp.SetAnyId,
				)
			}},
		//--- is_active ---
		"watchers.is_active": {
			func(a any) error {
				return w.HandlerValue(
					"watchers.is_active",
					a,
					w.watchersTmp.SetAnyIsActive,
				)
			}},
		//--- email ---
		"watchers.email": {
			func(a any) error {
				return w.HandlerValue(
					"watchers.email",
					a,
					w.watchersTmp.SetAnyEmail,
				)
			}},
		//--- last_name ---
		"watchers.last_name": {
			func(a any) error {
				return w.HandlerValue(
					"watchers.last_name",
					a,
					w.watchersTmp.SetAnyLastName,
				)
			}},
		//--- username ---
		"watchers.username": {
			func(a any) error {
				return w.HandlerValue(
					"watchers.username",
					a,
					w.watchersTmp.SetAnyUsername,
				)
			}},
		//--- first_name ---
		"watchers.first_name": {
			func(a any) error {
				return w.HandlerValue(
					"watchers.first_name",
					a,
					w.watchersTmp.SetAnyFirstName,
				)
			}},
		//--- patronimic---
		"watchers.patronimic": {
			func(a any) error {
				return w.HandlerValue(
					"watchers.patronimic",
					a,
					w.watchersTmp.SetAnyPatronimic,
				)
			}},
	}
}
