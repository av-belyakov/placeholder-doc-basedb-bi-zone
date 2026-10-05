package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerTenant начальный обработчик событий объекта 'tenant.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerTenant(t *datamodels.BiZoneIRPTenant) map[string][]func(any) error {
	return map[string][]func(any) error{
		"t.id":   {t.SetAnyID},
		"t.name": {t.SetAnyName},
	}
}
