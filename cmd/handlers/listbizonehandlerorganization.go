package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerOrganization начальный обработчик событий объекта 'organization.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerOrganization(o *datamodels.BiZoneIRPOrganization) map[string][]func(any) error {
	return map[string][]func(any) error{
		"data.organization.id":    {o.SetAnyID},
		"data.organization.title": {o.SetAnyTitle},
	}
}
