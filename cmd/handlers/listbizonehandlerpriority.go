package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerPriority начальный обработчик событий объекта 'priority.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerPriority(t *datamodels.BiZoneIRPPriority) map[string][]func(any) error {
	return map[string][]func(any) error{
		"data.priority.id":          {t.SetAnyID},
		"data.priority.title.value": {t.SetAnyTitleElement},
	}
}
