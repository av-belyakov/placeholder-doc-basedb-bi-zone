package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerCaseData начальный обработчик событий объекта 'data.*' топика 'socd-soar-prod-issue-v1'
func NewListBiZoneHandlerCaseData(data *datamodels.BiZoneIRPCaseData) map[string][]func(any) error {
	return map[string][]func(any) error{
		"data.created":                      {data.SetAnyCreated},
		"data.updated":                      {data.SetAnyUpdated},
		"data.detection_date":               {data.SetAnyDetectionDate},
		"data.resolutiondate":               {data.SetAnyResolutionDate},
		"data.id":                           {data.SetAnyID},
		"data.tlp":                          {data.SetAnyTLP},
		"data.fp_type":                      {data.SetAnyFpType},
		"data.summary":                      {data.SetAnySummary},
		"data.external_id":                  {data.SetAnyExternalID},
		"data.description":                  {data.SetAnyDescription},
		"data.recommendations":              {data.SetAnyRecommendations},
		"data.status_description":           {data.SetAnyStatusDescription},
		"data.resolution_detailed":          {data.SetAnyResolutionDetailed},
		"data.customer_star_rating_comment": {data.SetAnyCustomerStarRatingComment},
		"data.assignee":                     {data.SetAnyAssignee},
		"data.mitre_cov":                    {data.SetAnyMitreCov},
		"data.resolution":                   {data.SetAnyResolution},
		"data.response_team":                {data.SetAnyResponseTeam},
		"data.customer_assignee":            {data.SetAnyCustomerAssignee},
		"data.customer_star_rating":         {data.SetAnyCustomerStarRating},
		"data.is_public":                    {data.SetAnyIsPublic},
		//ниже работа со срезам содержащими простые типы
		"data.activity_detected": {data.SetAnyActivityDetected},
		"data.detection_rules":   {data.SetAnyDetectionRules},
		"data.platform_hostname": {data.SetAnyPlatformHostname},
	}
}
