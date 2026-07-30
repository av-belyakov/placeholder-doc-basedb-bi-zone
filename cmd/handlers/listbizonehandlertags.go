package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerTags обработчик для значений типа 'data.tags.*' основного объекта
func NewListBiZoneHandlerTags(sst *datamodels.SupportingStructureForTags) map[string][]func(any) {
	return map[string][]func(any){
		//--- name ---
		"tags.name": {func(a any) {
			sst.HandlerValue("data.tags.name", a, sst.GetTagTmp().SetAnyName)
		}},
		//--- color ---
		"tags.color": {func(a any) {
			sst.HandlerValue("data.tags.color", a, sst.GetTagTmp().SetAnyColor)
		}},
		//--- created ---
		"tags.created": {func(a any) {
			sst.HandlerValue("data.tags.created", a, func(a any) {
				//sst.GetTagTmp().SetAnyCreated может возвращать ошибку, которая пока не как бы обрабатывается
				_ = sst.GetTagTmp().SetAnyCreated(a)
			})
		}},
		//--- created_by.id ---
		"tags.created_by.id": {func(a any) {
			sst.HandlerValue("data.tags.created_by.id", a, sst.GetTagTmp().SetAnyCreatedByID)
		}},
		//--- created_by.username ---
		"tags.created_by.username": {func(a any) {
			sst.HandlerValue("data.tags.created_by.username", a, sst.GetTagTmp().SetAnyCreatedByUsername)
		}},
	}
}
