package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerCreatedBy начальный обработчик событий объекта 'created_by.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerCreatedBy(t *datamodels.BiZoneIRPCreatedBy) map[string][]func(any) error {
	return map[string][]func(any) error{
		"data.created_by.id":       {t.SetAnyID},
		"data.created_by.username": {t.SetAnyUsername},
	}
}
