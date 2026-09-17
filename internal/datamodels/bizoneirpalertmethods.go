package datamodels

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ******* Основная структура Alert ********
// *****************************************

// NewVerifiedBiZoneIRPAlert новый объект alert
func NewVerifiedBiZoneIRPAlert() *VerifiedBiZoneIRPAlert {
	oldTime := time.Date(1970, time.January, 1, 0, 0, 0, 1, time.UTC).Format(time.RFC3339)

	return &VerifiedBiZoneIRPAlert{
		CreatedTime:        oldTime,
		UpdatedTime:        oldTime,
		EventEndTime:       oldTime,
		EventStartTime:     oldTime,
		LastDetectionTime:  oldTime,
		FirstDetectionTime: oldTime,
		Tags:               []BiZoneIRPTag(nil),
		Snapshots:          []BiZoneIRPSnapshot(nil),
	}
}

func (va *VerifiedBiZoneIRPAlert) Get() *VerifiedBiZoneIRPAlert {
	return va
}

// GetID уникальный идентификатор
func (va *VerifiedBiZoneIRPAlert) GetID() uint64 {
	return va.ID
}

// SetID для поля id
func (va *VerifiedBiZoneIRPAlert) SetID(id uint64) error {
	va.ID = id

	return nil
}

// SetAnyID для поля id
func (va *VerifiedBiZoneIRPAlert) SetAnyID(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return va.SetID(v)
}

// GetUUID для поля uuid
func (va *VerifiedBiZoneIRPAlert) GetUUID() string {
	return va.UUID
}

// SetUUID для поля uuid
func (va *VerifiedBiZoneIRPAlert) SetUUID(UUID string) error {
	va.UUID = UUID

	return nil
}

// SetAnyUUID для поля uuid
func (va *VerifiedBiZoneIRPAlert) SetAnyUUID(a any) error {
	return va.SetUUID(fmt.Sprint(a))
}

// GetExternalID для поля external_id
func (va *VerifiedBiZoneIRPAlert) GetExternalID() string {
	return va.ExternalID
}

// SetExternalID для поля external_id
func (va *VerifiedBiZoneIRPAlert) SetExternalID(externalID string) error {
	va.ExternalID = externalID

	return nil
}

// SetAnyExternalID для поля external_id
func (va *VerifiedBiZoneIRPAlert) SetAnyExternalID(a any) error {
	return va.SetExternalID(fmt.Sprint(a))
}

// GetCustomerSystem для поля customer_system
func (va *VerifiedBiZoneIRPAlert) GetCustomerSystem() string {
	return va.CustomerSystem
}

// SetCustomerSystem для поля customer_system
func (va *VerifiedBiZoneIRPAlert) SetCustomerSystem(customerSystem string) error {
	va.CustomerSystem = customerSystem

	return nil
}

// SetAnyCustomerSystem для поля customer_system
func (va *VerifiedBiZoneIRPAlert) SetAnyCustomerSystem(a any) error {
	return va.SetCustomerSystem(fmt.Sprint(a))
}

// GetPlatformType для поля platform_type
func (va *VerifiedBiZoneIRPAlert) GetPlatformType() string {
	return va.PlatformType
}

// SetPlatformType для поля platform_type
func (va *VerifiedBiZoneIRPAlert) SetPlatformType(platformType string) error {
	va.PlatformType = platformType

	return nil
}

// SetAnyPlatformType для поля platform_type
func (va *VerifiedBiZoneIRPAlert) SetAnyPlatformType(a any) error {
	return va.SetPlatformType(fmt.Sprint(a))
}

// GetConfidence для поля confidence
func (va *VerifiedBiZoneIRPAlert) GetConfidence() string {
	return va.Confidence
}

// SetConfidence для поля confidence
func (va *VerifiedBiZoneIRPAlert) SetConfidence(confidence string) error {
	va.Confidence = confidence

	return nil
}

// SetAnyConfidence для поля confidence
func (va *VerifiedBiZoneIRPAlert) SetAnyConfidence(a any) error {
	return va.SetConfidence(fmt.Sprint(a))
}

// GetDescription для поля description
func (va *VerifiedBiZoneIRPAlert) GetDescription() string {
	return va.Description
}

// SetDescription для поля description
func (va *VerifiedBiZoneIRPAlert) SetDescription(description string) error {
	description = strings.ReplaceAll(description, "\t", "")
	description = strings.ReplaceAll(description, "\n", "")

	va.Description = description

	return nil
}

// SetAnyDescription для поля description
func (va *VerifiedBiZoneIRPAlert) SetAnyDescription(a any) error {
	return va.SetDescription(fmt.Sprint(a))
}

// GetDetectionRule для поля detection_rule
func (va *VerifiedBiZoneIRPAlert) GetDetectionRule() string {
	return va.DetectionRule
}

// SetDetectionRule для поля detection_rule
func (va *VerifiedBiZoneIRPAlert) SetDetectionRule(detectionRule string) error {
	va.DetectionRule = detectionRule

	return nil
}

// SetAnyDetectionRule для поля detection_rule
func (va *VerifiedBiZoneIRPAlert) SetAnyDetectionRule(a any) error {
	return va.SetDetectionRule(fmt.Sprint(a))
}

// GetCreatedTime для поля created_time (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetCreatedTime() string {
	return va.CreatedTime
}

// SetCreatedTime для поля created_time (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetCreatedTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.CreatedTime = timeStr

	return nil
}

// SetAnyCreatedTime для поля created_time
func (va *VerifiedBiZoneIRPAlert) SetAnyCreatedTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetCreatedTime(v)
	}

	return errors.New("type conversion error")
}

// GetUpdatedTime для поля updated_time (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetUpdatedTime() string {
	return va.UpdatedTime
}

// SetUpdatedTime для поля updated_time (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetUpdatedTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.UpdatedTime = timeStr

	return nil
}

// SetAnyUpdatedTime для поля updated_time
func (va *VerifiedBiZoneIRPAlert) SetAnyUpdatedTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetUpdatedTime(v)
	}

	return errors.New("type conversion error")
}

// GetEventStartTime для поля event_start_time (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetEventStartTime() string {
	return va.EventStartTime
}

// SetEventStartTime для поля event_start_time (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetEventStartTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.EventStartTime = timeStr

	return nil
}

// SetAnyEventStartTime для поля event_start_time
func (va *VerifiedBiZoneIRPAlert) SetAnyEventStartTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetEventStartTime(v)
	}

	return errors.New("type conversion error")
}

// GetVerifiedBiZoneIRPAlertEndTime для поля event_end_time (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetEventEndTime() string {
	return va.EventEndTime
}

// SetEventEndTime для поля event_end_time (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetEventEndTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.EventEndTime = timeStr

	return nil
}

// SetAnyEventEndTime для поля event_end_time
func (va *VerifiedBiZoneIRPAlert) SetAnyEventEndTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetEventEndTime(v)
	}

	return errors.New("type conversion error")
}

// GetFirstDetectionTime для поля first_detection_time (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetFirstDetectionTime() string {
	return va.FirstDetectionTime
}

// SetFirstDetectionTime для поля first_detection_time (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetFirstDetectionTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.FirstDetectionTime = timeStr

	return nil
}

// SetAnyFirstDetectionTime для поля first_detection_time
func (va *VerifiedBiZoneIRPAlert) SetAnyFirstDetectionTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetFirstDetectionTime(v)
	}

	return errors.New("type conversion error")
}

// GetLastDetectionTime для поля last_detection_time (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetLastDetectionTime() string {
	return va.LastDetectionTime
}

// SetLastDetectionTime для поля last_detection_time (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetLastDetectionTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.LastDetectionTime = timeStr

	return nil

}

// SetAnyLastDetectionTime для поля last_detection_time
func (va *VerifiedBiZoneIRPAlert) SetAnyLastDetectionTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetLastDetectionTime(v)
	}

	return errors.New("type conversion error")
}

// GetPlatformHostname для поля platform_hostname
func (va *VerifiedBiZoneIRPAlert) GetPlatformHostname() string {
	return va.PlatformHostname
}

// SetPlatformHostname для поля platform_hostname
func (va *VerifiedBiZoneIRPAlert) SetPlatformHostname(platformHostname string) error {
	va.PlatformHostname = platformHostname

	return nil
}

// SetAnyPlatformHostname для поля platform_hostname
func (va *VerifiedBiZoneIRPAlert) SetAnyPlatformHostname(a any) error {
	return va.SetPlatformHostname(fmt.Sprint(a))
}

// GetTitle для поля title
func (va *VerifiedBiZoneIRPAlert) GetTitle() string {
	return va.Title
}

// SetTitle для поля title
func (va *VerifiedBiZoneIRPAlert) SetTitle(title string) error {
	va.Title = title

	return nil
}

// SetAnyTitle для поля title
func (va *VerifiedBiZoneIRPAlert) SetAnyTitle(a any) error {
	return va.SetTitle(fmt.Sprint(a))
}

// GetSeverity для поля severity
func (va *VerifiedBiZoneIRPAlert) GetSeverity() string {
	return va.Severity
}

// SetSeverity для поля severity
func (va *VerifiedBiZoneIRPAlert) SetSeverity(severity string) error {
	va.Severity = severity

	return nil
}

// SetAnySeverity для поля severity
func (va *VerifiedBiZoneIRPAlert) SetAnySeverity(a any) error {
	return va.SetSeverity(fmt.Sprint(a))
}

// GetRecommendations для поля recommendations
func (va *VerifiedBiZoneIRPAlert) GetRecommendations() string {
	return va.Recommendations
}

// SetRecommendations для поля recommendations
func (va *VerifiedBiZoneIRPAlert) SetRecommendations(recommendations string) error {
	va.Recommendations = recommendations

	return nil
}

// SetAnyRecommendations для поля recommendations
func (va *VerifiedBiZoneIRPAlert) SetAnyRecommendations(a any) error {
	return va.SetRecommendations(fmt.Sprint(a))
}

// GetPriorityID для поля priority_id
func (va *VerifiedBiZoneIRPAlert) GetPriorityID() string {
	return va.PriorityID
}

// SetPriorityID для поля priority_id
func (va *VerifiedBiZoneIRPAlert) SetPriorityID(id string) error {
	va.PriorityID = id

	return nil
}

// SetAnyPriorityID для поля priority_id
func (va *VerifiedBiZoneIRPAlert) SetAnyPriorityID(a any) error {
	return va.SetPriorityID(fmt.Sprint(a))
}

// GetResponseTeam для поля response_team
func (va *VerifiedBiZoneIRPAlert) GetResponseTeam() uint64 {
	return va.ResponseTeam
}

// SetResponseTeam для поля response_team
func (va *VerifiedBiZoneIRPAlert) SetResponseTeam(resTeam uint64) error {
	va.ResponseTeam = resTeam

	return nil
}

// SetAnyResponseTeam для поля response_team
func (va *VerifiedBiZoneIRPAlert) SetAnyResponseTeam(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return va.SetID(v)
}

// GetData для поля data
func (va *VerifiedBiZoneIRPAlert) GetData() *BiZoneIRPData {
	return &va.Data
}

// SetData для поля data
func (va *VerifiedBiZoneIRPAlert) SetData(data BiZoneIRPData) error {
	va.Data = data

	return nil
}

// GetSnapshots для поля snapshots
func (va *VerifiedBiZoneIRPAlert) GetSnapshots() []BiZoneIRPSnapshot {
	return va.Snapshots
}

// SetSnapshots для поля snapshots
func (va *VerifiedBiZoneIRPAlert) SetSnapshots(snapshots []BiZoneIRPSnapshot) error {
	va.Snapshots = snapshots

	return nil
}

// GetTags для поля tags
func (va *VerifiedBiZoneIRPAlert) GetTags() []BiZoneIRPTag {
	return va.Tags
}

// SetTags для поля tags
func (va *VerifiedBiZoneIRPAlert) SetTags(tags []BiZoneIRPTag) error {
	va.Tags = tags

	return nil
}

// GetAffectedLogSources для поля affected_log_sources
func (va *VerifiedBiZoneIRPAlert) GetAffectedLogSources() []string {
	return va.AffectedLogSources
}

// SetAffectedLogSources для поля affected_log_sources
func (va *VerifiedBiZoneIRPAlert) SetAffectedLogSources(affectedLogSources []string) error {
	va.AffectedLogSources = affectedLogSources

	return nil
}

// GetAdditionalInformation поле дополнительной информации
func (va *VerifiedBiZoneIRPAlert) GetAdditionalInformation() *AdditionalInformation {
	return &va.AdditionalInformation
}

// SetAdditionalInformation для поля дополнительной информации
func (va *VerifiedBiZoneIRPAlert) SetAdditionalInformation(ai AdditionalInformation) error {
	va.AdditionalInformation = ai

	return nil
}

// ToStringBeautiful форматированный вывод
func (va *VerifiedBiZoneIRPAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	str.WriteString(fmt.Sprintf("%s'id': '%d'\n", ws, va.ID))
	str.WriteString(fmt.Sprintf("%s'uuid': '%s'\n", ws, va.UUID))
	str.WriteString(fmt.Sprintf("%s'Title': '%s'\n", ws, va.Title))
	str.WriteString(fmt.Sprintf("%s'severity': '%s'\n", ws, va.Severity))
	str.WriteString(fmt.Sprintf("%s'external_id': '%s'\n", ws, va.ExternalID))
	str.WriteString(fmt.Sprintf("%s'confidence': '%s'\n", ws, va.Confidence))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, va.Description))
	str.WriteString(fmt.Sprintf("%s'created_time': '%s'\n", ws, va.CreatedTime))
	str.WriteString(fmt.Sprintf("%s'updated_time': '%s'\n", ws, va.UpdatedTime))
	str.WriteString(fmt.Sprintf("%s'event_start_time': '%s'\n", ws, va.EventStartTime))
	str.WriteString(fmt.Sprintf("%s'event_end_time': '%s'\n", ws, va.EventEndTime))
	str.WriteString(fmt.Sprintf("%s'first_detection_time': '%s'\n", ws, va.FirstDetectionTime))
	str.WriteString(fmt.Sprintf("%s'last_detection_time': '%s'\n", ws, va.LastDetectionTime))
	str.WriteString(fmt.Sprintf("%s'detection_rule': '%s'\n", ws, va.DetectionRule))
	str.WriteString(fmt.Sprintf("%s'customer_system': '%s'\n", ws, va.CustomerSystem))
	str.WriteString(fmt.Sprintf("%s'recommendations': '%s'\n", ws, va.Recommendations))
	str.WriteString(fmt.Sprintf("%s'platform_type': '%s'\n", ws, va.PlatformType))
	str.WriteString(fmt.Sprintf("%s'platform_hostname': '%s'\n", ws, va.PlatformHostname))
	str.WriteString(fmt.Sprintf("%s'affected_log_sources': '%s'\n", ws, va.AffectedLogSources))
	str.WriteString(fmt.Sprintf("%s'data':\n%s", ws, va.Data.ToStringBeautiful(num+1)))
	str.WriteString(fmt.Sprintf("%s'tags':\n", ws))
	for k, v := range va.Tags {
		str.WriteString(fmt.Sprintf("%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2)))
	}
	str.WriteString(fmt.Sprintf("%s'snapshots':\n", ws))
	for k, v := range va.Snapshots {
		str.WriteString(fmt.Sprintf("%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2)))
	}
	str.WriteString(fmt.Sprintf("%s'@additional_information':\n%s", ws, va.AdditionalInformation.ToStringBeautiful(num+1)))

	return str.String()
}
