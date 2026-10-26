package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerPrimaryCategoryRef начальный обработчик событий объекта 'primary_category_ref.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerPrimaryCategoryRef(c *datamodels.CategoryRefType) map[string][]func(any) error {
	return map[string][]func(any) error{
		"primary_category_ref.id":          {c.SetAnyId},
		"primary_category_ref.title":       {c.SetAnyTitle},
		"primary_category_ref.description": {c.SetAnyDescription},
	}
}
