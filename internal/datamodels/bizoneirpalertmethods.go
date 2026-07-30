package datamodels

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ******* Основная структура Alert ********
// *****************************************

// NewVerifiedBiZoneIRPAlert новый объект
func NewVerifiedBiZoneIRPAlert() *VerifiedBiZoneIRPAlert {
	oldTime := time.Date(1970, time.January, 1, 0, 0, 0, 1, time.UTC).Format(time.RFC3339)

	return &VerifiedBiZoneIRPAlert{
		SpecialUUID:        uuid.NewString(),
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

// GetSpecialUUID уникальный идентификатор типа UUID
func (va *VerifiedBiZoneIRPAlert) GetSpecialUUID() string {
	return va.SpecialUUID
}

// GetIDNum
func (va *VerifiedBiZoneIRPAlert) GetIDNum() uint64 {
	return va.IDNum
}

// SetIDNum для поля IDNum
func (va *VerifiedBiZoneIRPAlert) SetIDNum(idNum uint64) error {
	va.IDNum = idNum

	return nil
}

// SetAnyIDNum для поля IDNum
func (va *VerifiedBiZoneIRPAlert) SetAnyIDNum(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return va.SetIDNum(v)
}

// GetUUID для поля UUID
func (va *VerifiedBiZoneIRPAlert) GetUUID() string {
	return va.UUID
}

// SetUUID для поля UUID
func (va *VerifiedBiZoneIRPAlert) SetUUID(UUID string) error {
	va.UUID = UUID

	return nil
}

// SetAnyUUID для поля UUID
func (va *VerifiedBiZoneIRPAlert) SetAnyUUID(a any) error {
	return va.SetUUID(fmt.Sprint(a))
}

// GetExternalID для поля ExternalID
func (va *VerifiedBiZoneIRPAlert) GetExternalID() string {
	return va.ExternalID
}

// SetExternalID для поля ExternalID
func (va *VerifiedBiZoneIRPAlert) SetExternalID(externalID string) error {
	va.ExternalID = externalID

	return nil
}

// SetAnyExternalID для поля ExternalID
func (va *VerifiedBiZoneIRPAlert) SetAnyExternalID(a any) error {
	return va.SetExternalID(fmt.Sprint(a))
}

// GetCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneIRPAlert) GetCustomerSystem() string {
	return va.CustomerSystem
}

// SetCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneIRPAlert) SetCustomerSystem(customerSystem string) error {
	va.CustomerSystem = customerSystem

	return nil
}

// SetAnyCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneIRPAlert) SetAnyCustomerSystem(a any) error {
	return va.SetCustomerSystem(fmt.Sprint(a))
}

// GetPlatformType для поля PlatformType
func (va *VerifiedBiZoneIRPAlert) GetPlatformType() string {
	return va.PlatformType
}

// SetPlatformType для поля PlatformType
func (va *VerifiedBiZoneIRPAlert) SetPlatformType(platformType string) error {
	va.PlatformType = platformType

	return nil
}

// SetAnyPlatformType для поля PlatformType
func (va *VerifiedBiZoneIRPAlert) SetAnyPlatformType(a any) error {
	return va.SetPlatformType(fmt.Sprint(a))
}

// GetAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneIRPAlert) GetAffectedLogSources() string {
	return va.AffectedLogSources
}

// SetAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneIRPAlert) SetAffectedLogSources(affectedLogSources string) error {
	va.AffectedLogSources = affectedLogSources

	return nil
}

// SetAnyAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneIRPAlert) SetAnyAffectedLogSources(a any) error {
	return va.SetAffectedLogSources(fmt.Sprint(a))
}

// GetConfidence для поля Confidence
func (va *VerifiedBiZoneIRPAlert) GetConfidence() string {
	return va.Confidence
}

// SetConfidence для поля Confidence
func (va *VerifiedBiZoneIRPAlert) SetConfidence(confidence string) error {
	va.Confidence = confidence

	return nil
}

// SetAnyConfidence для поля Confidence
func (va *VerifiedBiZoneIRPAlert) SetAnyConfidence(a any) error {
	return va.SetConfidence(fmt.Sprint(a))
}

// GetDescription для поля Description
func (va *VerifiedBiZoneIRPAlert) GetDescription() string {
	return va.Description
}

// SetDescription для поля Description
func (va *VerifiedBiZoneIRPAlert) SetDescription(description string) error {
	description = strings.ReplaceAll(description, "\t", "")
	description = strings.ReplaceAll(description, "\n", "")

	va.Description = description

	return nil
}

// SetAnyDescription для поля Description
func (va *VerifiedBiZoneIRPAlert) SetAnyDescription(a any) error {
	return va.SetDescription(fmt.Sprint(a))
}

// GetDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneIRPAlert) GetDetectionRule() string {
	return va.DetectionRule
}

// SetDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneIRPAlert) SetDetectionRule(detectionRule string) error {
	va.DetectionRule = detectionRule

	return nil
}

// SetAnyDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneIRPAlert) SetAnyDetectionRule(a any) error {
	return va.SetDetectionRule(fmt.Sprint(a))
}

// GetCreatedTime для поля CreatedTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetCreatedTime() string {
	return va.CreatedTime
}

// SetCreatedTime для поля CreatedTime (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetCreatedTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.CreatedTime = timeStr

	return nil
}

// SetAnyCreatedTime для поля CreatedTime
func (va *VerifiedBiZoneIRPAlert) SetAnyCreatedTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetCreatedTime(v)
	}

	return errors.New("type conversion error")
}

// GetUpdatedTime для поля UpdatedTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetUpdatedTime() string {
	return va.UpdatedTime
}

// SetUpdatedTime для поля UpdatedTime (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetUpdatedTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.UpdatedTime = timeStr

	return nil
}

// SetAnyUpdatedTime для поля UpdatedTime
func (va *VerifiedBiZoneIRPAlert) SetAnyUpdatedTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetUpdatedTime(v)
	}

	return errors.New("type conversion error")
}

// GetEventStartTime для поля EventStartTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetEventStartTime() string {
	return va.EventStartTime
}

// SetEventStartTime для поля EventStartTime (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetEventStartTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.EventStartTime = timeStr

	return nil
}

// SetAnyEventStartTime для поля EventStartTime
func (va *VerifiedBiZoneIRPAlert) SetAnyEventStartTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetEventStartTime(v)
	}

	return errors.New("type conversion error")
}

// GetVerifiedBiZoneIRPAlertEndTime для поля EventEndTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetEventEndTime() string {
	return va.EventEndTime
}

// SetEventEndTime для поля EventEndTime (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetEventEndTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.EventEndTime = timeStr

	return nil
}

// SetAnyEventEndTime для поля EventEndTime
func (va *VerifiedBiZoneIRPAlert) SetAnyEventEndTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetEventEndTime(v)
	}

	return errors.New("type conversion error")
}

// GetFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetFirstDetectionTime() string {
	return va.FirstDetectionTime
}

// SetFirstDetectionTime для поля FirstDetectionTime (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetFirstDetectionTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.FirstDetectionTime = timeStr

	return nil
}

// SetAnyFirstDetectionTime для поля FirstDetectionTime
func (va *VerifiedBiZoneIRPAlert) SetAnyFirstDetectionTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetFirstDetectionTime(v)
	}

	return errors.New("type conversion error")
}

// GetLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetLastDetectionTime() string {
	return va.LastDetectionTime
}

// SetLastDetectionTime для поля LastDetectionTime (преобразует в формат времени RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetLastDetectionTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	va.LastDetectionTime = timeStr

	return nil

}

// SetAnyLastDetectionTime для поля LastDetectionTime
func (va *VerifiedBiZoneIRPAlert) SetAnyLastDetectionTime(a any) error {
	if v, ok := a.(string); ok {
		return va.SetLastDetectionTime(v)
	}

	return errors.New("type conversion error")
}

// GetPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneIRPAlert) GetPlatformHostname() string {
	return va.PlatformHostname
}

// SetPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneIRPAlert) SetPlatformHostname(platformHostname string) error {
	va.PlatformHostname = platformHostname

	return nil
}

// SetAnyPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneIRPAlert) SetAnyPlatformHostname(a any) error {
	return va.SetPlatformHostname(fmt.Sprint(a))
}

// GetTitle для поля Title
func (va *VerifiedBiZoneIRPAlert) GetTitle() string {
	return va.Title
}

// SetTitle для поля Title
func (va *VerifiedBiZoneIRPAlert) SetTitle(title string) error {
	va.Title = title

	return nil
}

// SetAnyTitle для поля Title
func (va *VerifiedBiZoneIRPAlert) SetAnyTitle(a any) error {
	return va.SetTitle(fmt.Sprint(a))
}

// GetSeverity для поля Severity
func (va *VerifiedBiZoneIRPAlert) GetSeverity() string {
	return va.Severity
}

// SetSeverity для поля Severity
func (va *VerifiedBiZoneIRPAlert) SetSeverity(severity string) error {
	va.Severity = severity

	return nil
}

// SetAnySeverity для поля Severity
func (va *VerifiedBiZoneIRPAlert) SetAnySeverity(a any) error {
	return va.SetSeverity(fmt.Sprint(a))
}

// GetRecommendations для поля Recommendations
func (va *VerifiedBiZoneIRPAlert) GetRecommendations() string {
	return va.Recommendations
}

// SetRecommendations для поля Recommendations
func (va *VerifiedBiZoneIRPAlert) SetRecommendations(recommendations string) error {
	va.Recommendations = recommendations

	return nil
}

// SetAnyRecommendations для поля Recommendations
func (va *VerifiedBiZoneIRPAlert) SetAnyRecommendations(a any) error {
	return va.SetRecommendations(fmt.Sprint(a))
}

// GetData для поля Data
func (va *VerifiedBiZoneIRPAlert) GetData() *BiZoneIRPData {
	return &va.Data
}

// SetData для поля Data
func (va *VerifiedBiZoneIRPAlert) SetData(data BiZoneIRPData) error {
	va.Data = data

	return nil
}

// GetSnapshots для поля Snapshots
func (va *VerifiedBiZoneIRPAlert) GetSnapshots() []BiZoneIRPSnapshot {
	return va.Snapshots
}

// SetSnapshots для поля Snapshots
func (va *VerifiedBiZoneIRPAlert) SetSnapshots(snapshots []BiZoneIRPSnapshot) error {
	va.Snapshots = snapshots

	return nil
}

// GetTags для поля Tags
func (va *VerifiedBiZoneIRPAlert) GetTags() []BiZoneIRPTag {
	return va.Tags
}

// SetTags для поля Tags
func (va *VerifiedBiZoneIRPAlert) SetTags(tags []BiZoneIRPTag) error {
	va.Tags = tags

	return nil
}

// GetAdditionalInformation для поля AdditionalInformation
func (va *VerifiedBiZoneIRPAlert) GetAdditionalInformation() *AdditionalInformation {
	return &va.AdditionalInformation
}

// SetAdditionalInformation для поля AdditionalInformation
func (va *VerifiedBiZoneIRPAlert) SetAdditionalInformation(ai AdditionalInformation) error {
	va.AdditionalInformation = ai

	return nil
}

// ToStringBeautiful форматированный вывод
func (va *VerifiedBiZoneIRPAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	str.WriteString(fmt.Sprintf("%s'id': '%d'\n", ws, va.IDNum))
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
