package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerIssueTypeRef начальный обработчик событий объекта 'issue_type_ref.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerIssueTypeRef(itr *datamodels.IssueTypeRefType) map[string][]func(any) error {
	return map[string][]func(any) error{
		"issue_type_ref.id":          {itr.SetAnyId},
		"issue_type_ref.title":       {itr.SetAnyTitle},
		"issue_type_ref.description": {itr.SetAnyDescription},
	}
}
