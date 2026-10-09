package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerPrimaryCategory начальный обработчик событий объекта 'primary_category.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerPrimaryCategory(t *datamodels.BiZoneIRPCategory) map[string][]func(any) error {
	return map[string][]func(any) error{
		"data.primary_category.id":          {t.SetAnyID},
		"data.primary_category.title.value": {t.SetAnyTitleElement},
	}
}
