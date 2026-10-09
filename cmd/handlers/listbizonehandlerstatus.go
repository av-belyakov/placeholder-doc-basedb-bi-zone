package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerStatus начальный обработчик событий объекта 'status.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerStatus(t *datamodels.BiZoneIRPStatus) map[string][]func(any) error {
	return map[string][]func(any) error{
		"data.status.id":          {t.SetAnyID},
		"data.status.title.value": {t.SetAnyTitleElement},
	}
}
