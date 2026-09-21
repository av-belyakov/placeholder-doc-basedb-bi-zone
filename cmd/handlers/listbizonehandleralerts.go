package handlers

import (
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// NewListBiZoneHandlerAlerts начальный обработчик событий топика 'alertmgr-alerts'
func NewListBiZoneHandlerAlerts(alert *datamodels.VerifiedBiZoneIRPAlert) map[string][]func(any) error {
	return map[string][]func(any) error{
		"id":                   {alert.SetAnyID},
		"uuid":                 {alert.SetAnyUUID},
		"title":                {alert.SetAnyTitle},
		"severity":             {alert.SetAnySeverity},
		"confidence":           {alert.SetAnyConfidence},
		"description":          {alert.SetAnyDescription},
		"priority_id":          {alert.SetAnyPriorityID},
		"external_id":          {alert.SetAnyExternalID},
		"created_time":         {alert.SetAnyCreatedTime},
		"updated_time":         {alert.SetAnyUpdatedTime},
		"response_team":        {alert.SetAnyResponseTeam},
		"platform_type":        {alert.SetAnyPlatformType},
		"event_end_time":       {alert.SetAnyEventEndTime},
		"detection_rule":       {alert.SetAnyDetectionRule},
		"recommendations":      {alert.SetAnyRecommendations},
		"customer_system":      {alert.SetAnyCustomerSystem},
		"event_start_time":     {alert.SetAnyEventStartTime},
		"platform_hostname":    {alert.SetAnyPlatformHostname},
		"last_detection_time":  {alert.SetAnyLastDetectionTime},
		"affected_log_sources": {alert.SetAnyAffectedLogSource},
		"first_detection_time": {alert.SetAnyFirstDetectionTime},
	}
}
