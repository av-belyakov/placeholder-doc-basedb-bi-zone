package datamodels

import (
	"errors"
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewBiZoneIRPData новый объект Data
func NewBiZoneIRPCaseData() *BiZoneIRPCaseData {
	return &BiZoneIRPCaseData{
		Tags:              []BiZoneIRPTag(nil),
		DetectionRules:    []uint64(nil),
		PlatformHostname:  []string(nil),
		SecondaryCategory: []BiZoneIRPCategory(nil),
		ActivityDetected:  []any(nil),
	}
}

func (d *BiZoneIRPCaseData) Get() *BiZoneIRPCaseData {
	return d
}

// GetCreated для поля 'created' (формат RFC3339)
func (d *BiZoneIRPCaseData) GetCreated() string {
	return d.Created
}

// SetCreated для поля 'created' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPCaseData) SetCreated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.Created = timeStr

	return nil
}

// SetAnyCreated для поля 'created'
func (d *BiZoneIRPCaseData) SetAnyCreated(a any) error {
	if v, ok := a.(string); ok {
		return d.SetCreated(v)
	}

	return errors.New("type conversion error for field 'created'")
}

// GetUpdated для поля 'updated' (формат RFC3339)
func (d *BiZoneIRPCaseData) GetUpdated() string {
	return d.Updated
}

// SetUpdated для поля 'updated' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPCaseData) SetUpdated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.Updated = timeStr

	return nil
}

// SetAnyUpdated для поля 'updated'
func (d *BiZoneIRPCaseData) SetAnyUpdated(a any) error {
	if v, ok := a.(string); ok {
		return d.SetUpdated(v)
	}

	return errors.New("type conversion error for field 'updated'")
}

// GetDetectionDate для поля 'detection_date' (формат RFC3339)
func (d *BiZoneIRPCaseData) GetDetectionDate() string {
	return d.DetectionDate
}

// SetDetectionDate для поля 'detection_date' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPCaseData) SetDetectionDate(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.DetectionDate = timeStr

	return nil
}

// SetAnyDetectionDate для поля 'detection_date'
func (d *BiZoneIRPCaseData) SetAnyDetectionDate(a any) error {
	if v, ok := a.(string); ok {
		return d.SetDetectionDate(v)
	}

	return errors.New("type conversion error for field 'detection_date'")
}

// GetResolutionDate для поля 'resolutiondate' (формат RFC3339)
func (d *BiZoneIRPCaseData) GetResolutionDate() string {
	return d.DetectionDate
}

// SetResolutionDate для поля 'resolutiondate' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPCaseData) SetResolutionDate(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.DetectionDate = timeStr

	return nil
}

// SetAnyResolutionDate для поля 'resolutiondate'
func (d *BiZoneIRPCaseData) SetAnyResolutionDate(a any) error {
	if v, ok := a.(string); ok {
		return d.SetResolutionDate(v)
	}

	return errors.New("type conversion error for field 'resolutiondate'")
}

// GetID для поля 'id'
func (d *BiZoneIRPCaseData) GetID() string {
	return d.ID
}

// SetID для поля 'id'
func (d *BiZoneIRPCaseData) SetID(v string) error {
	d.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (d *BiZoneIRPCaseData) SetAnyID(a any) error {
	return d.SetID(fmt.Sprint(a))
}

// GetSummary для поля 'summary'
func (d *BiZoneIRPCaseData) GetSummary() string {
	return d.Summary
}

// SetSummary для поля 'summary'
func (d *BiZoneIRPCaseData) SetSummary(v string) error {
	d.Summary = v

	return nil
}

// SetAnySummary для поля 'summary'
func (d *BiZoneIRPCaseData) SetAnySummary(a any) error {
	return d.SetSummary(fmt.Sprint(a))
}

// GetDescription для поля 'description'
func (d *BiZoneIRPCaseData) GetDescription() string {
	return d.Description
}

// SetDescription для поля 'description'
func (d *BiZoneIRPCaseData) SetDescription(v string) error {
	d.Description = v

	return nil
}

// SetAnyDescription для поля 'description'
func (d *BiZoneIRPCaseData) SetAnyDescription(a any) error {
	return d.SetDescription(fmt.Sprint(a))
}

// GetRecommendations для поля 'recommendations'
func (d *BiZoneIRPCaseData) GetRecommendations() string {
	return d.Recommendations
}

// SetRecommendations для поля 'recommendations'
func (d *BiZoneIRPCaseData) SetRecommendations(v string) error {
	d.Recommendations = v

	return nil
}

// SetAnyRecommendations для поля 'recommendations'
func (d *BiZoneIRPCaseData) SetAnyRecommendations(a any) error {
	return d.SetRecommendations(fmt.Sprint(a))
}

// GetExternalID для поля 'external_id'
func (d *BiZoneIRPCaseData) GetExternalID() string {
	return d.ExternalID
}

// SetExternalID для поля 'external_id'
func (d *BiZoneIRPCaseData) SetExternalID(v string) error {
	d.ExternalID = v

	return nil
}

// SetAnyExternalID для поля 'external_id'
func (d *BiZoneIRPCaseData) SetAnyExternalID(a any) error {
	return d.SetExternalID(fmt.Sprint(a))
}

// GetResolutionDetailed для поля 'resolution_detailed'
func (d *BiZoneIRPCaseData) GetResolutionDetailed() string {
	return d.ResolutionDetailed
}

// SetResolutionDetailed для поля 'resolution_detailed'
func (d *BiZoneIRPCaseData) SetResolutionDetailed(v string) error {
	d.ResolutionDetailed = v

	return nil
}

// SetAnyResolutionDetailed для поля 'resolution_detailed'
func (d *BiZoneIRPCaseData) SetAnyResolutionDetailed(a any) error {
	return d.SetResolutionDetailed(fmt.Sprint(a))
}

// GetCustomerStarRatingComment для поля 'customer_star_rating_comment'
func (d *BiZoneIRPCaseData) GetCustomerStarRatingComment() string {
	return d.CustomerStarRatingComment
}

// SetCustomerStarRatingComment для поля 'customer_star_rating_comment'
func (d *BiZoneIRPCaseData) SetCustomerStarRatingComment(v string) error {
	d.CustomerStarRatingComment = v

	return nil
}

// SetAnyCustomerStarRatingComment для поля 'customer_star_rating_comment'
func (d *BiZoneIRPCaseData) SetAnyCustomerStarRatingComment(a any) error {
	return d.SetCustomerStarRatingComment(fmt.Sprint(a))
}

// GetStatusDescription для поля 'status_description'
func (d *BiZoneIRPCaseData) GetStatusDescription() string {
	return d.StatusDescription
}

// SetStatusDescription для поля 'status_description'
func (d *BiZoneIRPCaseData) SetStatusDescription(v string) error {
	d.StatusDescription = v

	return nil
}

// SetAnyStatusDescription для поля 'status_description'
func (d *BiZoneIRPCaseData) SetAnyStatusDescription(a any) error {
	return d.SetStatusDescription(fmt.Sprint(a))
}

// GetFpType для поля 'fp_type'
func (d *BiZoneIRPCaseData) GetFpType() string {
	return d.FpType
}

// SetFpType для поля 'fp_type'
func (d *BiZoneIRPCaseData) SetFpType(v string) error {
	d.FpType = v

	return nil
}

// SetAnyFpType для поля 'fp_type'
func (d *BiZoneIRPCaseData) SetAnyFpType(a any) error {
	return d.SetFpType(fmt.Sprint(a))
}

// GetTLP для поля 'tlp'
func (d *BiZoneIRPCaseData) GetTLP() string {
	return d.TLP
}

// SetTLP для поля 'tlp'
func (d *BiZoneIRPCaseData) SetTLP(v string) error {
	d.TLP = v

	return nil
}

// SetAnyTLP для поля 'tlp'
func (d *BiZoneIRPCaseData) SetAnyTLP(a any) error {
	return d.SetTLP(fmt.Sprint(a))
}

// GetAssignee для поля 'assignee'
func (d *BiZoneIRPCaseData) GetAssignee() *string {
	return d.Assignee
}

// SetAssignee для поля 'assignee'
func (d *BiZoneIRPCaseData) SetAssignee(v *string) error {
	d.Assignee = v

	return nil
}

// SetAnyAssignee для поля 'assignee'
func (d *BiZoneIRPCaseData) SetAnyAssignee(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'assignee'")
	}

	return d.SetAssignee(&v)
}

// GetMitreCov для поля 'mitre_cov'
func (d *BiZoneIRPCaseData) GetMitreCov() *string {
	return d.MitreCov
}

// SetMitreCov для поля 'mitre_cov'
func (d *BiZoneIRPCaseData) SetMitreCov(v *string) error {
	d.MitreCov = v

	return nil
}

// SetAnyMitreCov для поля 'mitre_cov'
func (d *BiZoneIRPCaseData) SetAnyMitreCov(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'mitre_cov'")
	}

	return d.SetMitreCov(&v)
}

// GetResolution для поля 'resolution'
func (d *BiZoneIRPCaseData) GetResolution() *string {
	return d.Resolution
}

// SetResolution для поля 'resolution'
func (d *BiZoneIRPCaseData) SetResolution(v *string) error {
	d.Resolution = v

	return nil
}

// SetAnyResolution для поля 'resolution'
func (d *BiZoneIRPCaseData) SetAnyResolution(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'resolution'")
	}

	return d.SetResolution(&v)
}

// GetResponseTeam для поля 'response_team'
func (d *BiZoneIRPCaseData) GetResponseTeam() *string {
	return d.ResponseTeam
}

// SetResponseTeam для поля 'response_team'
func (d *BiZoneIRPCaseData) SetResponseTeam(v *string) error {
	d.ResponseTeam = v

	return nil
}

// SetAnyResponseTeam для поля 'response_team'
func (d *BiZoneIRPCaseData) SetAnyResponseTeam(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'response_team'")
	}

	return d.SetResponseTeam(&v)
}

// GetCustomerAssignee для поля 'assignee'
func (d *BiZoneIRPCaseData) GetCustomerAssignee() *string {
	return d.CustomerAssignee
}

// SetCustomerAssignee для поля 'customer_assignee'
func (d *BiZoneIRPCaseData) SetCustomerAssignee(v *string) error {
	d.CustomerAssignee = v

	return nil
}

// SetAnyCustomerAssignee для поля 'customer_assignee'
func (d *BiZoneIRPCaseData) SetAnyCustomerAssignee(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'customer_assignee'")
	}

	return d.SetCustomerAssignee(&v)
}

// GetCustomerStarRating для поля 'customer_star_rating'
func (d *BiZoneIRPCaseData) GetCustomerStarRating() *uint64 {
	return d.CustomerStarRating
}

// SetCustomerStarRating для поля 'customer_star_rating'
func (d *BiZoneIRPCaseData) SetCustomerStarRating(v *uint64) error {
	d.CustomerStarRating = v

	return nil
}

// SetAnyCustomerStarRating для поля 'customer_star_rating'
func (d *BiZoneIRPCaseData) SetAnyCustomerStarRating(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetCustomerStarRating(&v)
}

// GetIsPublic для поля 'is_public'
func (d *BiZoneIRPCaseData) GetIsPublic() bool {
	return d.IsPublic
}

// SetIsPublic для поля 'is_public'
func (d *BiZoneIRPCaseData) SetIsPublic(v bool) error {
	d.IsPublic = v

	return nil
}

// SetAnyIsPublic для поля 'is_public'
func (d *BiZoneIRPCaseData) SetAnyIsPublic(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'is_public'")
	}

	return d.SetIsPublic(v)
}

// GetDesc для поля 'desc'
func (d *BiZoneIRPCaseData) GetDesc() string {
	return d.Desc
}

// SetDesc для поля 'desc'
func (d *BiZoneIRPCaseData) SetDesc(v string) error {
	d.Desc = v

	return nil
}

// SetAnyDesc для поля 'desc'
func (d *BiZoneIRPCaseData) SetAnyDesc(a any) error {
	return d.SetDesc(fmt.Sprint(a))
}

// GetQuery для поля 'query'
func (d *BiZoneIRPCaseData) GetQuery() string {
	return d.Query
}

// SetQuery для поля 'query'
func (d *BiZoneIRPCaseData) SetQuery(v string) error {
	d.Query = v

	return nil
}

// SetAnyQuery для поля 'query'
func (d *BiZoneIRPCaseData) SetAnyQuery(a any) error {
	return d.SetQuery(fmt.Sprint(a))
}

// GetEventUid для поля 'event_uid'
func (d *BiZoneIRPCaseData) GetEventUid() string {
	return d.EventUID
}

// SetEventUid для поля 'event_uid'
func (d *BiZoneIRPCaseData) SetEventUid(v string) error {
	d.EventUID = v

	return nil
}

// SetAnyEventUid для поля 'event_uid'
func (d *BiZoneIRPCaseData) SetAnyEventUid(a any) error {
	return d.SetEventUid(fmt.Sprint(a))
}

// GetJobTitle для поля 'job_title'
func (d *BiZoneIRPCaseData) GetJobTitle() string {
	return d.JobTitle
}

// SetJobTitle для поля 'job_title'
func (d *BiZoneIRPCaseData) SetJobTitle(v string) error {
	d.JobTitle = v

	return nil
}

// SetAnyJobTitle для поля 'job_title'
func (d *BiZoneIRPCaseData) SetAnyJobTitle(a any) error {
	return d.SetJobTitle(fmt.Sprint(a))
}

// GetMetadataProductName для поля 'metadata_product_name'
func (d *BiZoneIRPCaseData) GetMetadataProductName() string {
	return d.MetadataProductName
}

// SetMetadataProductName для поля 'metadata_product_name'
func (d *BiZoneIRPCaseData) SetMetadataProductName(v string) error {
	d.MetadataProductName = v

	return nil
}

// SetAnyMetadataProductName для поля 'metadata_product_name'
func (d *BiZoneIRPCaseData) SetAnyMetadataProductName(a any) error {
	return d.SetMetadataProductName(fmt.Sprint(a))
}

// GetSourceIP для поля 'source_ip'
func (d *BiZoneIRPCaseData) GetSourceIP() string {
	return d.SourceIP
}

// SetSourceIP для поля 'source_ip'
func (d *BiZoneIRPCaseData) SetSourceIP(v string) error {
	d.SourceIP = v

	return nil
}

// SetAnySourceIP для поля 'source_ip'
func (d *BiZoneIRPCaseData) SetAnySourceIP(a any) error {
	return d.SetSourceIP(fmt.Sprint(a))
}

// GetTargetIP для поля 'target_ip'
func (d *BiZoneIRPCaseData) GetTargetIP() string {
	return d.TargetIP
}

// SetTargetIP для поля 'target_ip'
func (d *BiZoneIRPCaseData) SetTargetIP(v string) error {
	d.TargetIP = v

	return nil
}

// SetAnyTargetIP для поля 'target_ip'
func (d *BiZoneIRPCaseData) SetAnyTargetIP(a any) error {
	return d.SetTargetIP(fmt.Sprint(a))
}

// GetUnmappedEventCard для поля 'unmapped_event_card'
func (d *BiZoneIRPCaseData) GetUnmappedEventCard() string {
	return d.UnmappedEventCard
}

// SetUnmappedEventCard для поля 'unmapped_event_card'
func (d *BiZoneIRPCaseData) SetUnmappedEventCard(v string) error {
	d.UnmappedEventCard = v

	return nil
}

// SetAnyUnmappedEventCard для поля 'unmapped_event_card'
func (d *BiZoneIRPCaseData) SetAnyUnmappedEventCard(a any) error {
	return d.SetUnmappedEventCard(fmt.Sprint(a))
}

// GetUnmappedHiveAlertID для поля 'unmapped_hive_alert_id'
func (d *BiZoneIRPCaseData) GetUnmappedHiveAlertID() string {
	return d.UnmappedHiveAlertID
}

// SetUnmappedHiveAlertID для поля 'unmapped_hive_alert_id'
func (d *BiZoneIRPCaseData) SetUnmappedHiveAlertID(v string) error {
	d.UnmappedHiveAlertID = v

	return nil
}

// SetAnyUnmappedHiveAlertID для поля 'unmapped_hive_alert_id'
func (d *BiZoneIRPCaseData) SetAnyUnmappedHiveAlertID(a any) error {
	return d.SetUnmappedHiveAlertID(fmt.Sprint(a))
}

// GetUnmappedSensorIP для поля 'unmapped_sensor_ip'
func (d *BiZoneIRPCaseData) GetUnmappedSensorIP() string {
	return d.UnmappedSensorIP
}

// SetUnmappedSensorIP для поля 'unmapped_sensor_ip'
func (d *BiZoneIRPCaseData) SetUnmappedSensorIP(v string) error {
	d.UnmappedSensorIP = v

	return nil
}

// SetAnyUnmappedSensorIP для поля 'unmapped_sensor_ip'
func (d *BiZoneIRPCaseData) SetAnyUnmappedSensorIP(a any) error {
	return d.SetUnmappedSensorIP(fmt.Sprint(a))
}

// GetUnmappedSensorName для поля 'unmapped_sensor_name'
func (d *BiZoneIRPCaseData) GetUnmappedSensorName() string {
	return d.UnmappedSensorName
}

// SetUnmappedSensorName для поля 'unmapped_sensor_name'
func (d *BiZoneIRPCaseData) SetUnmappedSensorName(v string) error {
	d.UnmappedSensorName = v

	return nil
}

// SetAnyUnmappedSensorName для поля 'unmapped_sensor_name'
func (d *BiZoneIRPCaseData) SetAnyUnmappedSensorName(a any) error {
	return d.SetUnmappedSensorName(fmt.Sprint(a))
}

// GetFirstSeenTime для поля 'first_seen_time' (формат RFC3339)
func (d *BiZoneIRPCaseData) GetFirstSeenTime() string {
	return d.FirstSeenTime
}

// SetFirstSeenTime для поля 'first_seen_time' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPCaseData) SetFirstSeenTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.FirstSeenTime = timeStr

	return nil
}

// SetAnyFirstSeenTime для поля 'first_seen_time'
func (d *BiZoneIRPCaseData) SetAnyFirstSeenTime(a any) error {
	if v, ok := a.(string); ok {
		return d.SetFirstSeenTime(v)
	}

	return errors.New("type conversion error for field 'first_seen_time'")
}

// GetLastSeenTime для поля 'last_seen_time' (формат RFC3339)
func (d *BiZoneIRPCaseData) GetLastSeenTime() string {
	return d.LastSeenTime
}

// SetLastSeenTime для поля 'last_seen_time' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPCaseData) SetLastSeenTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.LastSeenTime = timeStr

	return nil
}

// SetAnyLastSeenTime для поля 'last_seen_time'
func (d *BiZoneIRPCaseData) SetAnyLastSeenTime(a any) error {
	if v, ok := a.(string); ok {
		return d.SetLastSeenTime(v)
	}

	return errors.New("type conversion error for field 'last_seen_time'")
}

// GetTags для поля 'tags'
func (d *BiZoneIRPCaseData) GetTags() []string {
	return d.Tags
}

// SetTags для поля 'tags'
func (d *BiZoneIRPCaseData) SetTags(v []string) error {
	d.Tags = v

	return nil
}

// SetTag добавляет значение 'tag' в список
func (d *BiZoneIRPCaseData) SetTag(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.Tags); !isExist {
		d.Tags = append(d.Tags, v)
	}

	return nil
}

// SetAnyTag добавляет некоторое значение в список 'tags'
func (d *BiZoneIRPCaseData) SetAnyTag(a any) error {
	return d.SetTag(fmt.Sprint(a))
}

// GetUnmappedDstEndpointArray для поля 'unmapped_dst_endpoint_array'
func (d *BiZoneIRPCaseData) GetUnmappedDstEndpointArray() []string {
	return d.UnmappedDstEndpointArray
}

// SetUnmappedDstEndpointArray для поля 'unmapped_dst_endpoint_array'
func (d *BiZoneIRPCaseData) SetUnmappedDstEndpointArray(v []string) error {
	d.UnmappedDstEndpointArray = v

	return nil
}

// SetUnmappedDstEndpointArrayElement добавляет значение 'unmapped_dst_endpoint_array' в список
func (d *BiZoneIRPCaseData) SetUnmappedDstEndpointArrayElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.UnmappedDstEndpointArray); !isExist {
		d.UnmappedDstEndpointArray = append(d.UnmappedDstEndpointArray, v)
	}

	return nil
}

// SetAnyUnmappedDstEndpointArray добавляет некоторое значение в список 'unmapped_dst_endpoint_array'
func (d *BiZoneIRPCaseData) SetAnyUnmappedDstEndpointArray(a any) error {
	return d.SetUnmappedDstEndpointArrayElement(fmt.Sprint(a))
}

// GetUnmappedHomeEndpointArray для поля 'unmapped_home_endpoint_array'
func (d *BiZoneIRPCaseData) GetUnmappedHomeEndpointArray() []string {
	return d.UnmappedHomeEndpointArray
}

// SetUnmappedHomeEndpointArray для поля 'unmapped_home_endpoint_array'
func (d *BiZoneIRPCaseData) SetUnmappedHomeEndpointArray(v []string) error {
	d.UnmappedHomeEndpointArray = v

	return nil
}

// SetUnmappedHomeEndpointArrayElement добавляет значение 'unmapped_home_endpoint_array' в список
func (d *BiZoneIRPCaseData) SetUnmappedHomeEndpointArrayElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.UnmappedHomeEndpointArray); !isExist {
		d.UnmappedHomeEndpointArray = append(d.UnmappedHomeEndpointArray, v)
	}

	return nil
}

// SetAnyUnmappedHomeEndpointArray добавляет некоторое значение в список 'unmapped_home_endpoint_array'
func (d *BiZoneIRPCaseData) SetAnyUnmappedHomeEndpointArray(a any) error {
	return d.SetUnmappedHomeEndpointArrayElement(fmt.Sprint(a))
}

// GetDetectionPattern для поля 'detection_pattern'
func (d *BiZoneIRPCaseData) GetDetectionPattern() []uint64 {
	return d.DetectionPattern
}

// SetDetectionPattern для поля 'detection_pattern'
func (d *BiZoneIRPCaseData) SetDetectionPattern(v []uint64) error {
	d.DetectionPattern = v

	return nil
}

// SetDetectionPatternElement добавляет значение 'detection_pattern' в список
func (d *BiZoneIRPCaseData) SetDetectionPatternElement(v uint64) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.DetectionPattern); !isExist {
		d.DetectionPattern = append(d.DetectionPattern, v)
	}

	return nil
}

// SetAnyDetectionPatternElement добавляет некоторое значение в список 'detection_pattern'
func (d *BiZoneIRPCaseData) SetAnyDetectionPatternElement(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetDetectionPatternElement(v)
}

// GetUnmappedAgentArray для поля 'unmapped_agent_array'
func (d *BiZoneIRPCaseData) GetUnmappedAgentArray() []uint64 {
	return d.UnmappedAgentArray
}

// SetUnmappedAgentArrayn для поля 'unmapped_agent_array'
func (d *BiZoneIRPCaseData) SetUnmappedAgentArrayn(v []uint64) error {
	d.UnmappedAgentArray = v

	return nil
}

// SetUnmappedAgentArrayElement добавляет значение 'unmapped_agent_array' в список
func (d *BiZoneIRPCaseData) SetUnmappedAgentArrayElement(v uint64) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.UnmappedAgentArray); !isExist {
		d.UnmappedAgentArray = append(d.UnmappedAgentArray, v)
	}

	return nil
}

// SetAnyUnmappedAgentArrayElement добавляет некоторое значение в список 'unmapped_agent_array'
func (d *BiZoneIRPCaseData) SetAnyUnmappedAgentArrayElement(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetUnmappedAgentArrayElement(v)
}

// GetDataSecurity для поля 'data_security'
func (d *BiZoneIRPCaseData) GetDataSecurity() map[string][]BiZoneIRPDataSecurity {
	return d.DataSecurity
}

// SetDataSecurity для поля 'data_security'
func (d *BiZoneIRPCaseData) SetDataSecurity(v map[string][]BiZoneIRPDataSecurity) error {
	d.DataSecurity = v

	return nil
}

// SetDataSecurityElement добавляет значение 'data_security' в список
func (d *BiZoneIRPCaseData) SetDataSecurityElement(k string, v []BiZoneIRPDataSecurity) error {
	d.DataSecurity[k] = v

	return nil
}

// ToStringBeautiful форматированный вывод
func (d *BiZoneIRPCaseData) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsStr := supportingfunctions.GetWhitespace(num + 1)
	wsInt := supportingfunctions.GetWhitespace(num + 2)

	fmt.Fprintf(&str, "%s'desc': '%s'\n", ws, d.Desc)
	fmt.Fprintf(&str, "%s'query': '%s'\n", ws, d.Query)
	fmt.Fprintf(&str, "%s'source_ip': '%s'\n", ws, d.SourceIP)
	fmt.Fprintf(&str, "%s'target_ip': '%s'\n", ws, d.TargetIP)
	fmt.Fprintf(&str, "%s'event_uid': '%s'\n", ws, d.EventUID)
	fmt.Fprintf(&str, "%s'job_title': '%s'\n", ws, d.JobTitle)
	fmt.Fprintf(&str, "%s'first_seen_time': '%s'\n", ws, d.FirstSeenTime)
	fmt.Fprintf(&str, "%s'last_seen_time': '%s'\n", ws, d.LastSeenTime)
	fmt.Fprintf(&str, "%s'unmapped_sensor_ip': '%s'\n", ws, d.UnmappedSensorIP)
	fmt.Fprintf(&str, "%s'unmapped_event_card': '%s'\n", ws, d.UnmappedEventCard)
	fmt.Fprintf(&str, "%s'metadata_product_name': '%s'\n", ws, d.MetadataProductName)
	fmt.Fprintf(&str, "%s'unmapped_hive_alert_id': '%s'\n", ws, d.UnmappedHiveAlertID)
	fmt.Fprintf(&str, "%s'unmapped_sensor_name': '%s'\n", ws, d.UnmappedSensorName)
	fmt.Fprintf(&str, "%s'agent': '%d'\n", ws, d.Agent)
	fmt.Fprintf(&str, "%s'severity_id': '%d'\n", ws, d.SeverityID)
	fmt.Fprintf(&str, "%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.Tags))
	fmt.Fprintf(&str, "%s'unmapped_dst_endpoint_array': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.UnmappedDstEndpointArray))
	fmt.Fprintf(&str, "%s'unmapped_home_endpoint_array': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.UnmappedHomeEndpointArray))
	fmt.Fprintf(&str, "%s'detection_pattern': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.DetectionPattern))
	fmt.Fprintf(&str, "%s'unmapped_agent_array': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.UnmappedAgentArray))
	fmt.Fprintf(&str, "%s'data_security':\n", ws)
	for k, v := range d.DataSecurity {
		fmt.Fprintf(&str, "%s%s:\n", wsStr, k)

		for item, value := range v {
			fmt.Fprintf(&str, "%s%d.\n%s\n", wsInt, item, value.ToStringBeautiful(num+3))
		}
	}

	return str.String()
}
