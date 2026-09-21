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
func (va *VerifiedBiZoneIRPAlert) SetID(v uint64) error {
	va.ID = v

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
func (va *VerifiedBiZoneIRPAlert) SetUUID(v string) error {
	va.UUID = v

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
func (va *VerifiedBiZoneIRPAlert) SetExternalID(v string) error {
	va.ExternalID = v

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
func (va *VerifiedBiZoneIRPAlert) SetCustomerSystem(v string) error {
	va.CustomerSystem = v

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
func (va *VerifiedBiZoneIRPAlert) SetPlatformType(v string) error {
	va.PlatformType = v

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
func (va *VerifiedBiZoneIRPAlert) SetConfidence(v string) error {
	va.Confidence = v

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
func (va *VerifiedBiZoneIRPAlert) SetDescription(v string) error {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	va.Description = v

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
func (va *VerifiedBiZoneIRPAlert) SetDetectionRule(v string) error {
	va.DetectionRule = v

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

	return errors.New("type conversion error for field 'created_time'")
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

	return errors.New("type conversion error for field 'updated_time'")
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

	return errors.New("type conversion error for field 'event_start_time'")
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

	return errors.New("type conversion error for field 'event_end_time'")
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

	return errors.New("type conversion error for field 'first_detection_time'")
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

	return errors.New("type conversion error for field 'last_detection_time'")
}

// GetPlatformHostname для поля platform_hostname
func (va *VerifiedBiZoneIRPAlert) GetPlatformHostname() string {
	return va.PlatformHostname
}

// SetPlatformHostname для поля platform_hostname
func (va *VerifiedBiZoneIRPAlert) SetPlatformHostname(v string) error {
	va.PlatformHostname = v

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
func (va *VerifiedBiZoneIRPAlert) SetTitle(v string) error {
	va.Title = v

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
func (va *VerifiedBiZoneIRPAlert) SetSeverity(v string) error {
	va.Severity = v

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
func (va *VerifiedBiZoneIRPAlert) SetRecommendations(v string) error {
	va.Recommendations = v

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
func (va *VerifiedBiZoneIRPAlert) SetPriorityID(v string) error {
	va.PriorityID = v

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
func (va *VerifiedBiZoneIRPAlert) SetResponseTeam(v uint64) error {
	va.ResponseTeam = v

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
func (va *VerifiedBiZoneIRPAlert) SetData(v BiZoneIRPData) error {
	va.Data = v

	return nil
}

// GetSnapshots для поля snapshots
func (va *VerifiedBiZoneIRPAlert) GetSnapshots() []BiZoneIRPSnapshot {
	return va.Snapshots
}

// SetSnapshots для поля snapshots
func (va *VerifiedBiZoneIRPAlert) SetSnapshots(v []BiZoneIRPSnapshot) error {
	va.Snapshots = v

	return nil
}

// GetTags для поля tags
func (va *VerifiedBiZoneIRPAlert) GetTags() []BiZoneIRPTag {
	return va.Tags
}

// SetTags для поля tags
func (va *VerifiedBiZoneIRPAlert) SetTags(v []BiZoneIRPTag) error {
	va.Tags = v

	return nil
}

// GetAffectedLogSources для поля affected_log_sources
func (va *VerifiedBiZoneIRPAlert) GetAffectedLogSources() []string {
	return va.AffectedLogSources
}

// SetAffectedLogSources для поля affected_log_sources
func (va *VerifiedBiZoneIRPAlert) SetAffectedLogSources(v []string) error {
	va.AffectedLogSources = v

	return nil
}

// SetAffectedLogSource одно значение для поля affected_log_sources
func (va *VerifiedBiZoneIRPAlert) SetAffectedLogSource(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, va.AffectedLogSources); !isExist {
		va.AffectedLogSources = append(va.AffectedLogSources, v)
	}

	return nil
}

// SetAnyAffectedLogSource одно значение для поля affected_log_sources
func (va *VerifiedBiZoneIRPAlert) SetAnyAffectedLogSource(a any) error {
	return va.SetAffectedLogSource(fmt.Sprint(a))
}

// GetAdditionalInformation поле дополнительной информации
func (va *VerifiedBiZoneIRPAlert) GetAdditionalInformation() *AdditionalInformation {
	return &va.AdditionalInformation
}

// SetAdditionalInformation для поля дополнительной информации
func (va *VerifiedBiZoneIRPAlert) SetAdditionalInformation(v AdditionalInformation) error {
	va.AdditionalInformation = v

	return nil
}

// ToStringBeautiful форматированный вывод
func (va *VerifiedBiZoneIRPAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	fmt.Fprintf(&str, "%s'id': '%d'\n", ws, va.ID)
	fmt.Fprintf(&str, "%s'uuid': '%s'\n", ws, va.UUID)
	fmt.Fprintf(&str, "%s'Title': '%s'\n", ws, va.Title)
	fmt.Fprintf(&str, "%s'severity': '%s'\n", ws, va.Severity)
	fmt.Fprintf(&str, "%s'external_id': '%s'\n", ws, va.ExternalID)
	fmt.Fprintf(&str, "%s'confidence': '%s'\n", ws, va.Confidence)
	fmt.Fprintf(&str, "%s'description': '%s'\n", ws, va.Description)
	fmt.Fprintf(&str, "%s'created_time': '%s'\n", ws, va.CreatedTime)
	fmt.Fprintf(&str, "%s'updated_time': '%s'\n", ws, va.UpdatedTime)
	fmt.Fprintf(&str, "%s'event_start_time': '%s'\n", ws, va.EventStartTime)
	fmt.Fprintf(&str, "%s'event_end_time': '%s'\n", ws, va.EventEndTime)
	fmt.Fprintf(&str, "%s'first_detection_time': '%s'\n", ws, va.FirstDetectionTime)
	fmt.Fprintf(&str, "%s'last_detection_time': '%s'\n", ws, va.LastDetectionTime)
	fmt.Fprintf(&str, "%s'detection_rule': '%s'\n", ws, va.DetectionRule)
	fmt.Fprintf(&str, "%s'customer_system': '%s'\n", ws, va.CustomerSystem)
	fmt.Fprintf(&str, "%s'recommendations': '%s'\n", ws, va.Recommendations)
	fmt.Fprintf(&str, "%s'platform_type': '%s'\n", ws, va.PlatformType)
	fmt.Fprintf(&str, "%s'platform_hostname': '%s'\n", ws, va.PlatformHostname)
	fmt.Fprintf(&str, "%s'affected_log_sources': '%s'\n", ws, va.AffectedLogSources)
	fmt.Fprintf(&str, "%s'data':\n%s", ws, va.Data.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'tags':\n", ws)
	for k, v := range va.Tags {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	}
	fmt.Fprintf(&str, "%s'snapshots':\n", ws)
	for k, v := range va.Snapshots {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	}
	fmt.Fprintf(&str, "%s'@additional_information':\n%s", ws, va.AdditionalInformation.ToStringBeautiful(num+1))

	return str.String()
}
