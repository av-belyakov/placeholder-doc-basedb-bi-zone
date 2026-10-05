package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerType начальный обработчик событий объекта 'type.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerType(t *datamodels.BiZoneIRPType) map[string][]func(any) error {
	return map[string][]func(any) error{
		"t.id":    {t.SetAnyID},
		"t.title": {t.SetAnyTitleElement},
	}
}
