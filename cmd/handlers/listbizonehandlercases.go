package handlers

import (
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// NewListBiZoneHandlerCases начальный обработчик событий топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerCases(c *datamodels.VerifiedBiZoneIRPCase) map[string][]func(any) error {
	return map[string][]func(any) error{
		"id":          {c.SetAnyUUID},
		"type":        {c.SetAnyType},
		"source":      {c.SetAnySource},
		"specversion": {c.SetAnySpecVersion},
		"time":        {c.SetAnyTime},
		"subject":     {c.SetAnySubject},
		"dataschema":  {c.SetAnyDataSchema},
	}
}
