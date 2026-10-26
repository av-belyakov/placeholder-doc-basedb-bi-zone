package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerResolutionNameRef начальный обработчик событий объекта 'resolution_name_ref.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerResolutionNameRef(c *datamodels.CategoryRefType) map[string][]func(any) error {
	return map[string][]func(any) error{
		"resolution_name_ref.id":          {c.SetAnyId},
		"resolution_name_ref.title":       {c.SetAnyTitle},
		"resolution_name_ref.description": {c.SetAnyDescription},
	}
}
