package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerData начальный обработчик событий объекта 'data.*' топика 'alertmgr-alerts'
func NewListBiZoneHandlerData(data *datamodels.BiZoneIRPData) map[string][]func(any) error {
	return map[string][]func(any) error{
		"data.s_rule_body":            {data.SetAnyAgent},
		"data.severity_id":            {data.SetAnySeverityID},
		"data.desc":                   {data.SetAnyDesc},
		"data.event_uid":              {data.SetAnyEventUid},
		"data.job_title":              {data.SetAnyJobTitle},
		"data.source_ip":              {data.SetAnySourceIP},
		"data.target_ip":              {data.SetAnyTargetIP},
		"data.first_seen_time":        {data.SetAnyFirstSeenTime},
		"data.last_seen_time":         {data.SetAnyLastSeenTime},
		"data.metadata_product_name":  {data.SetAnyMetadataProductName},
		"data.unmapped_sensor_ip":     {data.SetAnyUnmappedSensorIP},
		"data.unmapped_sensor_name":   {data.SetAnyUnmappedSensorName},
		"data.unmapped_hive_alert_id": {data.SetAnyUnmappedHiveAlertID},
		//ниже работа со срезам содержащими простые типы
		"data.tags":                         {data.SetAnyTag},
		"data.detection_pattern":            {data.SetAnyDetectionPatternElement},
		"data.unmapped_agent_array":         {data.SetAnyUnmappedAgentArrayElement},
		"data.unmapped_dst_endpoint_array":  {data.SetAnyUnmappedDstEndpointArray},
		"data.unmapped_home_endpoint_array": {data.SetAnyUnmappedHomeEndpointArray},
	}
}
