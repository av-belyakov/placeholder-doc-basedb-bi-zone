package datamodels

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewVerifiedBiZoneIRPIncident новый объект типа 'Incident'
func NewVerifiedBiZoneIRPIncident() *VerifiedBiZoneIRPIncident {
	oldTime := time.Date(1970, time.January, 1, 0, 0, 0, 1, time.UTC).Format(time.RFC3339)

	return &VerifiedBiZoneIRPIncident{
		Tags:                   []IRPTagType(nil),
		ObservedIocs:           []ObservedIocsType(nil),
		SecondaryCategoryRef:   []CategoryRefType(nil),
		ActivityDetected:       []string(nil),
		SecondaryCategory:      []string(nil),
		PlatformHostname:       []string(nil),
		KeyField:               []string(nil),
		DetectionRules:         []uint64(nil),
		ParsedActivityDetected: []any(nil),
		Attachments:            []any(nil),
		Emls:                   []any(nil),
		Slas:                   []any(nil),
		Badges:                 []any(nil),
		Watchers:               []any(nil),
		WatchersId:             []any(nil),
		Created:                oldTime,
		Updated:                oldTime,
		Timestamp:              oldTime,
		UpdatedAll:             oldTime,
		DetectionDate:          oldTime,
	}
}

func (i *VerifiedBiZoneIRPIncident) Get() *VerifiedBiZoneIRPIncident {
	return i
}

// GetTags для поля 'tags'
func (i *VerifiedBiZoneIRPIncident) GetTags() []IRPTagType {
	return i.Tags
}

// SetTags для поля 'tags'
func (i *VerifiedBiZoneIRPIncident) SetTags(v []IRPTagType) error {
	i.Tags = v

	return nil
}

// GetObservedIocs для поля 'observed_iocs'
func (i *VerifiedBiZoneIRPIncident) GetObservedIocs() []ObservedIocsType {
	return i.ObservedIocs
}

// SetObservedIocs для поля 'observed_iocs'
func (i *VerifiedBiZoneIRPIncident) SetObservedIocs(v []ObservedIocsType) error {
	i.ObservedIocs = v

	return nil
}

// GetSecondaryCategoryRef для поля 'secondary_category_ref'
func (i *VerifiedBiZoneIRPIncident) GetSecondaryCategoryRef() []CategoryRefType {
	return i.SecondaryCategoryRef
}

// SetSecondaryCategoryRef для поля 'secondary_category_ref'
func (i *VerifiedBiZoneIRPIncident) SetSecondaryCategoryRef(v []CategoryRefType) error {
	i.SecondaryCategoryRef = v

	return nil
}

// GetKeyFieli для поля 'keyfield'
func (i *VerifiedBiZoneIRPIncident) GetKeyField() []string {
	return i.KeyField
}

// SetKeyFieli для поля 'keyfield'
func (i *VerifiedBiZoneIRPIncident) SetKeyField(v []string) error {
	i.KeyField = v

	return nil
}

// SetKeyFieldElement добавляет значение 'keyfield' в список
func (i *VerifiedBiZoneIRPIncident) SetKeyFieldElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, i.PlatformHostname); !isExist {
		i.KeyField = append(i.KeyField, v)
	}

	return nil
}

// SetAnyKeyFieldElement добавляет некоторое значение в список 'keyfield'
func (i *VerifiedBiZoneIRPIncident) SetAnyKeyFieldElement(a any) error {
	return i.SetKeyFieldElement(fmt.Sprint(a))
}

// GetActivityDetected для поля 'activity_detected'
func (i *VerifiedBiZoneIRPIncident) GetActivityDetected() []string {
	return i.ActivityDetected
}

// SetActivityDetected для поля 'activity_detected'
func (i *VerifiedBiZoneIRPIncident) SetActivityDetected(v []string) error {
	i.ActivityDetected = v

	return nil
}

// SetActivityDetectedElement добавляет значение 'activity_detected' в список
func (i *VerifiedBiZoneIRPIncident) SetActivityDetectedElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, i.PlatformHostname); !isExist {
		i.ActivityDetected = append(i.ActivityDetected, v)
	}

	return nil
}

// SetAnyActivityDetectedElement добавляет некоторое значение в список 'activity_detected'
func (i *VerifiedBiZoneIRPIncident) SetAnyActivityDetectedElement(a any) error {
	return i.SetActivityDetectedElement(fmt.Sprint(a))
}

// GetSecondaryCategory для поля 'secondary_category'
func (i *VerifiedBiZoneIRPIncident) GetSecondaryCategory() []string {
	return i.SecondaryCategory
}

// SetSecondaryCategory для поля 'secondary_category'
func (i *VerifiedBiZoneIRPIncident) SetSecondaryCategory(v []string) error {
	i.SecondaryCategory = v

	return nil
}

// SetSecondaryCategoryElement добавляет значение 'secondary_category' в список
func (i *VerifiedBiZoneIRPIncident) SetSecondaryCategoryElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, i.PlatformHostname); !isExist {
		i.SecondaryCategory = append(i.SecondaryCategory, v)
	}

	return nil
}

// SetAnySecondaryCategoryElement добавляет некоторое значение в список 'secondary_category'
func (i *VerifiedBiZoneIRPIncident) SetAnySecondaryCategoryElement(a any) error {
	return i.SetSecondaryCategoryElement(fmt.Sprint(a))
}

// GetPlatformHostname для поля 'platform_hostname'
func (i *VerifiedBiZoneIRPIncident) GetPlatformHostname() []string {
	return i.PlatformHostname
}

// SetPlatformHostname для поля 'platform_hostname'
func (i *VerifiedBiZoneIRPIncident) SetPlatformHostname(v []string) error {
	i.PlatformHostname = v

	return nil
}

// SetPlatformHostnameElement добавляет значение 'platform_hostname' в список
func (i *VerifiedBiZoneIRPIncident) SetPlatformHostnameElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, i.PlatformHostname); !isExist {
		i.PlatformHostname = append(i.PlatformHostname, v)
	}

	return nil
}

// SetAnyPlatformHostnameElement добавляет некоторое значение в список 'detection_rules'
func (i *VerifiedBiZoneIRPIncident) SetAnyPlatformHostnameElement(a any) error {
	return i.SetPlatformHostnameElement(fmt.Sprint(a))
}

// GetDetectionRules для поля 'detection_rules'
func (i *VerifiedBiZoneIRPIncident) GetDetectionRules() []uint64 {
	return i.DetectionRules
}

// SetDetectionRules для поля 'detection_rules'
func (i *VerifiedBiZoneIRPIncident) SetDetectionRules(v []uint64) error {
	i.DetectionRules = v

	return nil
}

// SetDetectionRulesElement добавляет значение 'detection_rules' в список
func (i *VerifiedBiZoneIRPIncident) SetDetectionRulesElement(v uint64) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, i.DetectionRules); !isExist {
		i.DetectionRules = append(i.DetectionRules, v)
	}

	return nil
}

// SetAnyDetectionRulesElement добавляет некоторое значение в список 'detection_rules'
func (i *VerifiedBiZoneIRPIncident) SetAnyDetectionRulesElement(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return i.SetDetectionRulesElement(v)
}

// GetParsedActivityDetected для поля 'parsed_activity_detected'
func (i *VerifiedBiZoneIRPIncident) GetParsedActivityDetected() []any {
	return i.ParsedActivityDetected
}

// SetParsedActivityDetected для поля 'parsed_activity_detected'
func (i *VerifiedBiZoneIRPIncident) SetParsedActivityDetected(v []any) error {
	i.ParsedActivityDetected = v

	return nil
}

// SetParsedActivityDetectedElement добавляет значение 'parsed_activity_detected' в список
func (i *VerifiedBiZoneIRPIncident) SetParsedActivityDetectedElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, i.ActivityDetected); !isExist {
	i.ParsedActivityDetected = append(i.ParsedActivityDetected, v)
	//}

	return nil
}

// SetAnyParsedActivityDetected добавляет некоторое значение в список 'parsed_activity_detected'
func (i *VerifiedBiZoneIRPIncident) SetAnyParsedActivityDetectedElement(a any) error {
	return i.SetParsedActivityDetectedElement(fmt.Sprint(a))
}

// GetAttachments для поля 'attachments'
func (i *VerifiedBiZoneIRPIncident) GetAttachments() []any {
	return i.Attachments
}

// SetAttachments для поля 'attachments'
func (i *VerifiedBiZoneIRPIncident) SetAttachments(v []any) error {
	i.Attachments = v

	return nil
}

// SetAttachmentsElement добавляет значение 'attachments' в список
func (i *VerifiedBiZoneIRPIncident) SetAttachmentsElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, i.ActivityDetected); !isExist {
	i.Attachments = append(i.Attachments, v)
	//}

	return nil
}

// SetAnyAttachmentsElement добавляет некоторое значение в список 'attachments'
func (i *VerifiedBiZoneIRPIncident) SetAnyAttachmentsElement(a any) error {
	return i.SetAttachmentsElement(fmt.Sprint(a))
}

// GetEmls для поля 'emls'
func (i *VerifiedBiZoneIRPIncident) GetEmls() []any {
	return i.Emls
}

// SetEmls для поля 'emls'
func (i *VerifiedBiZoneIRPIncident) SetEmls(v []any) error {
	i.Emls = v

	return nil
}

// SetEmlsElement добавляет значение 'emls' в список
func (i *VerifiedBiZoneIRPIncident) SetEmlsElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, i.ActivityDetected); !isExist {
	i.Emls = append(i.Emls, v)
	//}

	return nil
}

// SetAnyEmlsElement добавляет некоторое значение в список 'emls'
func (i *VerifiedBiZoneIRPIncident) SetAnyEmlsElement(a any) error {
	return i.SetEmlsElement(fmt.Sprint(a))
}

// GetSlas для поля 'slas'
func (i *VerifiedBiZoneIRPIncident) GetSlas() []any {
	return i.Slas
}

// SetSlas для поля 'slas'
func (i *VerifiedBiZoneIRPIncident) SetSlas(v []any) error {
	i.Slas = v

	return nil
}

// SetSlasElement добавляет значение 'slas' в список
func (i *VerifiedBiZoneIRPIncident) SetSlasElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, i.ActivityDetected); !isExist {
	i.Slas = append(i.Slas, v)
	//}

	return nil
}

// SetAnySlasElement добавляет некоторое значение в список 'slas'
func (i *VerifiedBiZoneIRPIncident) SetAnySlasElement(a any) error {
	return i.SetSlasElement(fmt.Sprint(a))
}

// GetBadges для поля 'badges'
func (i *VerifiedBiZoneIRPIncident) GetBadges() []any {
	return i.Badges
}

// SetBadges для поля 'badges'
func (i *VerifiedBiZoneIRPIncident) SetBadges(v []any) error {
	i.Badges = v

	return nil
}

// SetBadgesElement добавляет значение 'badges' в список
func (i *VerifiedBiZoneIRPIncident) SetBadgesElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, i.ActivityDetected); !isExist {
	i.Badges = append(i.Badges, v)
	//}

	return nil
}

// SetAnyBadgesElement добавляет некоторое значение в список 'badges'
func (i *VerifiedBiZoneIRPIncident) SetAnyBadgesElement(a any) error {
	return i.SetBadgesElement(fmt.Sprint(a))
}

// GetWatchers для поля 'watchers'
func (i *VerifiedBiZoneIRPIncident) GetWatchers() []any {
	return i.Watchers
}

// SetWatchers для поля 'watchers'
func (i *VerifiedBiZoneIRPIncident) SetWatchers(v []any) error {
	i.Watchers = v

	return nil
}

// SetWatchersElement добавляет значение 'watchers' в список
func (i *VerifiedBiZoneIRPIncident) SetWatchersElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, i.ActivityDetected); !isExist {
	i.Watchers = append(i.Watchers, v)
	//}

	return nil
}

// SetAnyWatchersElement добавляет некоторое значение в список 'watchers'
func (i *VerifiedBiZoneIRPIncident) SetAnyWatchersElement(a any) error {
	return i.SetWatchersElement(fmt.Sprint(a))
}

// GetWatchersId для поля 'watchers_id'
func (i *VerifiedBiZoneIRPIncident) GetWatchersId() []any {
	return i.WatchersId
}

// SetWatchersId для поля 'watchers_id'
func (i *VerifiedBiZoneIRPIncident) SetWatchersId(v []any) error {
	i.WatchersId = v

	return nil
}

// SetWatchersIdElement добавляет значение 'watchers_id' в список
func (i *VerifiedBiZoneIRPIncident) SetWatchersIdElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, i.ActivityDetected); !isExist {
	i.WatchersId = append(i.WatchersId, v)
	//}

	return nil
}

// SetAnyWatchersIdElement добавляет некоторое значение в список 'watchers_id'
func (i *VerifiedBiZoneIRPIncident) SetAnyWatchersIdElement(a any) error {
	return i.SetWatchersIdElement(fmt.Sprint(a))
}

// GetIssueTypeRef для поля 'issue_type_ref'
func (i *VerifiedBiZoneIRPIncident) GetIssueTypeRef() *IssueTypeRefType {
	return &i.IssueTypeRef
}

// SetIssueTypeRef для поля 'issue_type_ref'
func (i *VerifiedBiZoneIRPIncident) SetIssueTypeRef(v IssueTypeRefType) error {
	i.IssueTypeRef = v

	return nil
}

// GetPrimaryCategoryRef для поля 'primary_category_ref'
func (i *VerifiedBiZoneIRPIncident) GetPrimaryCategoryRef() *CategoryRefType {
	return &i.PrimaryCategoryRef
}

// SetPrimaryCategoryRef для поля 'primary_category_ref'
func (i *VerifiedBiZoneIRPIncident) SetPrimaryCategoryRef(v CategoryRefType) error {
	i.PrimaryCategoryRef = v

	return nil
}

// GetResolutionNameRef для поля 'resolution_name_ref'
func (i *VerifiedBiZoneIRPIncident) GetResolutionNameRef() *CategoryRefType {
	return &i.ResolutionNameRef
}

// SetResolutionNameRef для поля 'resolution_name_ref'
func (i *VerifiedBiZoneIRPIncident) SetResolutionNameRef(v CategoryRefType) error {
	i.ResolutionNameRef = v

	return nil
}

// GetCreated для поля 'created' (формат RFC3339)
func (i *VerifiedBiZoneIRPIncident) GetCreated() string {
	return i.Created
}

// SetCreated для поля 'created' (преобразует в формат времени RFC3339)
func (i *VerifiedBiZoneIRPIncident) SetCreated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	i.Created = timeStr

	return nil
}

// SetAnyCreated для поля 'created'
func (i *VerifiedBiZoneIRPIncident) SetAnyCreated(a any) error {
	if v, ok := a.(string); ok {
		return i.SetCreated(v)
	}

	return errors.New("type conversion error for field 'created'")
}

// GetUpdated для поля 'updated' (формат RFC3339)
func (i *VerifiedBiZoneIRPIncident) GetUpdated() string {
	return i.Updated
}

// SetUpdated для поля 'updated' (преобразует в формат времени RFC3339)
func (i *VerifiedBiZoneIRPIncident) SetUpdated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	i.Updated = timeStr

	return nil
}

// SetAnyUpdated для поля 'updated'
func (i *VerifiedBiZoneIRPIncident) SetAnyUpdated(a any) error {
	if v, ok := a.(string); ok {
		return i.SetUpdated(v)
	}

	return errors.New("type conversion error for field 'updated'")
}

// GetTimestamp для поля 'timestamp' (формат RFC3339)
func (i *VerifiedBiZoneIRPIncident) GetTimestamp() string {
	return i.Timestamp
}

// SetTimestamp для поля 'timestamp' (преобразует в формат времени RFC3339)
func (i *VerifiedBiZoneIRPIncident) SetTimestamp(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	i.Timestamp = timeStr

	return nil
}

// SetAnyTimestamp для поля 'timestamp'
func (i *VerifiedBiZoneIRPIncident) SetAnyTimestamp(a any) error {
	if v, ok := a.(string); ok {
		return i.SetTimestamp(v)
	}

	return errors.New("type conversion error for field 'timestamp'")
}

// GetUpdatedAll для поля 'updated_all' (формат RFC3339)
func (i *VerifiedBiZoneIRPIncident) GetUpdatedAll() string {
	return i.UpdatedAll
}

// Set для поля 'updated_all' (преобразует в формат времени RFC3339)
func (i *VerifiedBiZoneIRPIncident) SetUpdatedAll(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	i.UpdatedAll = timeStr

	return nil
}

// SetAnyUpdatedAll для поля 'updated_all'
func (i *VerifiedBiZoneIRPIncident) SetAnyUpdatedAll(a any) error {
	if v, ok := a.(string); ok {
		return i.SetUpdatedAll(v)
	}

	return errors.New("type conversion error for field 'updated_all'")
}

// GetDetectionDate для поля 'detection_date' (формат RFC3339)
func (i *VerifiedBiZoneIRPIncident) GetDetectionDate() string {
	return i.DetectionDate
}

// SetDetectionDate для поля 'detection_date' (преобразует в формат времени RFC3339)
func (i *VerifiedBiZoneIRPIncident) SetDetectionDate(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	i.DetectionDate = timeStr

	return nil
}

// SetAnyDetectionDate для поля 'detection_date'
func (i *VerifiedBiZoneIRPIncident) SetAnyDetectionDate(a any) error {
	if v, ok := a.(string); ok {
		return i.SetDetectionDate(v)
	}

	return errors.New("type conversion error for field 'detection_date'")
}

// GetFirstSentNotificationTime для поля 'first_sent_notification_time' (формат RFC3339)
func (i *VerifiedBiZoneIRPIncident) GetFirstSentNotificationTime() string {
	return i.FirstSentNotificationTime
}

// SetFirstSentNotificationTime для поля 'first_sent_notification_time' (преобразует в формат времени RFC3339)
func (i *VerifiedBiZoneIRPIncident) SetFirstSentNotificationTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	i.FirstSentNotificationTime = timeStr

	return nil
}

// SetAnyFirstSentNotificationTime для поля 'first_sent_notification_time'
func (i *VerifiedBiZoneIRPIncident) SetAnyFirstSentNotificationTime(a any) error {
	if v, ok := a.(string); ok {
		return i.SetFirstSentNotificationTime(v)
	}

	return errors.New("type conversion error for field 'first_sent_notification_time'")
}

// GetID для поля 'id'
func (i *VerifiedBiZoneIRPIncident) GetID() uint64 {
	return i.ID
}

// SetID для поля 'id'
func (i *VerifiedBiZoneIRPIncident) SetID(v uint64) error {
	i.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (i *VerifiedBiZoneIRPIncident) SetAnyID(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return i.SetID(v)
}

// GetSystem для поля 'system'
func (i *VerifiedBiZoneIRPIncident) GetSystem() uint64 {
	return i.System
}

// SetSystem для поля 'system'
func (i *VerifiedBiZoneIRPIncident) SetSystem(v uint64) error {
	i.System = v

	return nil
}

// SetAnySystem для поля 'system'
func (i *VerifiedBiZoneIRPIncident) SetAnySystem(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return i.SetSystem(v)
}

// GetIsVisibleForCustomer для поля 'is_visible_for_customer'
func (i *VerifiedBiZoneIRPIncident) GetIsVisibleForCustomer() bool {
	return i.IsVisibleForCustomer
}

// SetIsVisibleForCustomer для поля 'is_visible_for_customer'
func (i *VerifiedBiZoneIRPIncident) SetIsVisibleForCustomer(v bool) error {
	i.IsVisibleForCustomer = v

	return nil
}

// SetAnyIsVisibleForCustomer для поля 'is_visible_for_customer'
func (i *VerifiedBiZoneIRPIncident) SetAnyIsVisibleForCustomer(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'is_visible_for_customer'")
	}

	return i.SetIsVisibleForCustomer(v)
}

// GetShowGossopkaButton для поля 'show_gossopka_button'
func (i *VerifiedBiZoneIRPIncident) GetShowGossopkaButton() bool {
	return i.ShowGossopkaButton
}

// SetShowGossopkaButton для поля 'show_gossopka_button'
func (i *VerifiedBiZoneIRPIncident) SetShowGossopkaButton(v bool) error {
	i.ShowGossopkaButton = v

	return nil
}

// SetAnyShowGossopkaButton для поля 'show_gossopka_button'
func (i *VerifiedBiZoneIRPIncident) SetAnyShowGossopkaButton(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'show_gossopka_button'")
	}

	return i.SetShowGossopkaButton(v)
}

// GetCancelGossopkaSendButton для поля 'cancel_gossopka_send_button'
func (i *VerifiedBiZoneIRPIncident) GetCancelGossopkaSendButton() bool {
	return i.CancelGossopkaSendButton
}

// SetCancelGossopkaSendButton для поля 'cancel_gossopka_send_button'
func (i *VerifiedBiZoneIRPIncident) SetCancelGossopkaSendButton(v bool) error {
	i.CancelGossopkaSendButton = v

	return nil
}

// SetAnyCancelGossopkaSendButton для поля 'cancel_gossopka_send_button'
func (i *VerifiedBiZoneIRPIncident) SetAnyCancelGossopkaSendButton(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'cancel_gossopka_send_button'")
	}

	return i.SetCancelGossopkaSendButton(v)
}

// GetShowGtiButton для поля 'show_gti_button'
func (i *VerifiedBiZoneIRPIncident) GetShowGtiButton() bool {
	return i.ShowGtiButton
}

// SetShowGtiButton для поля 'show_gti_button'
func (i *VerifiedBiZoneIRPIncident) SetShowGtiButton(v bool) error {
	i.ShowGtiButton = v

	return nil
}

// SetAnyShowGtiButton для поля 'show_gti_button'
func (i *VerifiedBiZoneIRPIncident) SetAnyShowGtiButton(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'show_gti_button'")
	}

	return i.SetShowGtiButton(v)
}

// GetKey для поля 'key'
func (i *VerifiedBiZoneIRPIncident) GetKey() string {
	return i.Key
}

// SetKey для поля 'key'
func (i *VerifiedBiZoneIRPIncident) SetKey(v string) error {
	i.Key = v

	return nil
}

// SetAnyKey для поля 'key'
func (i *VerifiedBiZoneIRPIncident) SetAnyKey(a any) error {
	return i.SetKey(fmt.Sprint(a))
}

// GetStatus для поля 'status'
func (i *VerifiedBiZoneIRPIncident) GetStatus() string {
	return i.Status
}

// SetStatus для поля 'status'
func (i *VerifiedBiZoneIRPIncident) SetStatus(v string) error {
	i.Status = v

	return nil
}

// SetAnyStatus для поля 'status'
func (i *VerifiedBiZoneIRPIncident) SetAnyStatus(a any) error {
	return i.SetStatus(fmt.Sprint(a))
}

// GetSummary для поля 'summary'
func (i *VerifiedBiZoneIRPIncident) GetSummary() string {
	return i.Summary
}

// SetSummary для поля 'summary'
func (i *VerifiedBiZoneIRPIncident) SetSummary(v string) error {
	i.Summary = v

	return nil
}

// SetAnySummary для поля 'summary'
func (i *VerifiedBiZoneIRPIncident) SetAnySummary(a any) error {
	return i.SetSummary(fmt.Sprint(a))
}

// GetPriority для поля 'priority'
func (i *VerifiedBiZoneIRPIncident) GetPriority() string {
	return i.Priority
}

// SetPriority для поля 'priority'
func (i *VerifiedBiZoneIRPIncident) SetPriority(v string) error {
	i.Priority = v

	return nil
}

// SetAnyPriority для поля 'priority'
func (i *VerifiedBiZoneIRPIncident) SetAnyPriority(a any) error {
	return i.SetPriority(fmt.Sprint(a))
}

// GetTenantId для поля 'tenant_id'
func (i *VerifiedBiZoneIRPIncident) GetTenantId() string {
	return i.TenantId
}

// SetTenantId для поля 'tenant_id'
func (i *VerifiedBiZoneIRPIncident) SetTenantId(v string) error {
	i.TenantId = v

	return nil
}

// SetAnyTenantId для поля 'tenant_id'
func (i *VerifiedBiZoneIRPIncident) SetAnyTenantId(a any) error {
	return i.SetTenantId(fmt.Sprint(a))
}

// GetIssueType для поля 'issue_type'
func (i *VerifiedBiZoneIRPIncident) GetIssueType() string {
	return i.IssueType
}

// SetIssueType для поля 'issue_type'
func (i *VerifiedBiZoneIRPIncident) SetIssueType(v string) error {
	i.IssueType = v

	return nil
}

// SetAnyIssueType для поля 'issue_type'
func (i *VerifiedBiZoneIRPIncident) SetAnyIssueType(a any) error {
	return i.SetIssueType(fmt.Sprint(a))
}

// GetDescription для поля 'description'
func (i *VerifiedBiZoneIRPIncident) GetDescription() string {
	return i.Description
}

// SetDescription для поля 'description'
func (i *VerifiedBiZoneIRPIncident) SetDescription(v string) error {
	i.Description = v

	return nil
}

// SetAnyDescription для поля 'description'
func (i *VerifiedBiZoneIRPIncident) SetAnyDescription(a any) error {
	return i.SetDescription(fmt.Sprint(a))
}

// GetRenderedDescription для поля 'rendered_description'
func (i *VerifiedBiZoneIRPIncident) GetRenderedDescription() string {
	return i.RenderedDescription
}

// SetRenderedDescription для поля 'rendered_description'
func (i *VerifiedBiZoneIRPIncident) SetRenderedDescription(v string) error {
	i.RenderedDescription = v

	return nil
}

// SetAnyRenderedDescription для поля 'rendered_description'
func (i *VerifiedBiZoneIRPIncident) SetAnyRenderedDescription(a any) error {
	return i.SetRenderedDescription(fmt.Sprint(a))
}

// GetStatusDescription для поля 'status_description'
func (i *VerifiedBiZoneIRPIncident) GetStatusDescription() string {
	return i.StatusDescription
}

// SetStatusDescription для поля 'status_description'
func (i *VerifiedBiZoneIRPIncident) SetStatusDescription(v string) error {
	i.StatusDescription = v

	return nil
}

// SetAnyStatusDescription для поля 'status_description'
func (i *VerifiedBiZoneIRPIncident) SetAnyStatusDescription(a any) error {
	return i.SetStatusDescription(fmt.Sprint(a))
}

// GetAssigneeDisplayName для поля 'assignee_displayName'
func (i *VerifiedBiZoneIRPIncident) GetAssigneeDisplayName() string {
	return i.AssigneeDisplayName
}

// SetAssigneeDisplayName для поля 'assignee_displayName'
func (i *VerifiedBiZoneIRPIncident) SetAssigneeDisplayName(v string) error {
	i.AssigneeDisplayName = v

	return nil
}

// SetAnyAssigneeDisplayName для поля 'assignee_displayName'
func (i *VerifiedBiZoneIRPIncident) SetAnyAssigneeDisplayName(a any) error {
	return i.SetAssigneeDisplayName(fmt.Sprint(a))
}

// GetReporterDisplayName для поля 'reporter_displayName'
func (i *VerifiedBiZoneIRPIncident) GetReporterDisplayName() string {
	return i.ReporterDisplayName
}

// SetReporterDisplayName для поля 'reporter_displayName'
func (i *VerifiedBiZoneIRPIncident) SetReporterDisplayName(v string) error {
	i.ReporterDisplayName = v

	return nil
}

// SetAnyReporterDisplayName для поля 'reporter_displayName'
func (i *VerifiedBiZoneIRPIncident) SetAnyReporterDisplayName(a any) error {
	return i.SetReporterDisplayName(fmt.Sprint(a))
}

// GetCreatorDisplayName для поля 'creator_displayName'
func (i *VerifiedBiZoneIRPIncident) GetCreatorDisplayName() string {
	return i.CreatorDisplayName
}

// SetCreatorDisplayName для поля 'creator_displayName'
func (i *VerifiedBiZoneIRPIncident) SetCreatorDisplayName(v string) error {
	i.CreatorDisplayName = v

	return nil
}

// SetAnyCreatorDisplayName для поля 'creator_displayName'
func (i *VerifiedBiZoneIRPIncident) SetAnyCreatorDisplayName(a any) error {
	return i.SetCreatorDisplayName(fmt.Sprint(a))
}

// GetReporterEmailAddress для поля 'reporter_emailAddress'
func (i *VerifiedBiZoneIRPIncident) GetReporterEmailAddress() string {
	return i.ReporterEmailAddress
}

// SetReporterEmailAddress для поля 'reporter_emailAddress'
func (i *VerifiedBiZoneIRPIncident) SetReporterEmailAddress(v string) error {
	i.ReporterEmailAddress = v

	return nil
}

// SetAnyReporterEmailAddress для поля 'reporter_emailAddress'
func (i *VerifiedBiZoneIRPIncident) SetAnyReporterEmailAddress(a any) error {
	return i.SetReporterEmailAddress(fmt.Sprint(a))
}

// GetPrimaryCategory для поля 'primary_category'
func (i *VerifiedBiZoneIRPIncident) GetPrimaryCategory() string {
	return i.PrimaryCategory
}

// SetPrimaryCategory для поля 'primary_category'
func (i *VerifiedBiZoneIRPIncident) SetPrimaryCategory(v string) error {
	i.PrimaryCategory = v

	return nil
}

// SetAnyPrimaryCategory для поля 'primary_category'
func (i *VerifiedBiZoneIRPIncident) SetAnyPrimaryCategory(a any) error {
	return i.SetPrimaryCategory(fmt.Sprint(a))
}

// GetResolutionDate для поля 'resolutiondate'
func (i *VerifiedBiZoneIRPIncident) GetResolutionDate() string {
	return i.ResolutionDate
}

// SetResolutionDate для поля 'resolutiondate'
func (i *VerifiedBiZoneIRPIncident) SetResolutionDate(v string) error {
	i.ResolutionDate = v

	return nil
}

// SetAnyResolutionDate для поля 'resolutiondate'
func (i *VerifiedBiZoneIRPIncident) SetAnyResolutionDate(a any) error {
	return i.SetResolutionDate(fmt.Sprint(a))
}

// GetResolutionName для поля 'resolution_name'
func (i *VerifiedBiZoneIRPIncident) GetResolutionName() string {
	return i.ResolutionName
}

// SetResolutionName для поля 'resolution_name'
func (i *VerifiedBiZoneIRPIncident) SetResolutionName(v string) error {
	i.ResolutionName = v

	return nil
}

// SetAnyResolutionName для поля 'resolution_name'
func (i *VerifiedBiZoneIRPIncident) SetAnyResolutionName(a any) error {
	return i.SetResolutionName(fmt.Sprint(a))
}

// GetResolutionDetailed для поля 'resolution_detailed'
func (i *VerifiedBiZoneIRPIncident) GetResolutionDetailed() string {
	return i.ResolutionDetailed
}

// SetResolutionDetailed для поля 'resolution_detailed'
func (i *VerifiedBiZoneIRPIncident) SetResolutionDetailed(v string) error {
	i.ResolutionDetailed = v

	return nil
}

// SetAnyResolutionDetailed для поля 'resolution_detailed'
func (i *VerifiedBiZoneIRPIncident) SetAnyResolutionDetailed(a any) error {
	return i.SetResolutionDetailed(fmt.Sprint(a))
}

// GetPlatformUrl для поля 'platform_url'
func (i *VerifiedBiZoneIRPIncident) GetPlatformUrl() string {
	return i.PlatformUrl
}

// SetPlatformUrl для поля 'platform_url'
func (i *VerifiedBiZoneIRPIncident) SetPlatformUrl(v string) error {
	i.PlatformUrl = v

	return nil
}

// SetAnyPlatformUrl для поля 'platform_url'
func (i *VerifiedBiZoneIRPIncident) SetAnyPlatformUrl(a any) error {
	return i.SetPlatformUrl(fmt.Sprint(a))
}

// GetAttackType для поля 'attack_type'
func (i *VerifiedBiZoneIRPIncident) GetAttackType() string {
	return i.AttackType
}

// SetAttackType для поля 'attack_type'
func (i *VerifiedBiZoneIRPIncident) SetAttackType(v string) error {
	i.AttackType = v

	return nil
}

// SetAnyAttackType для поля 'attack_type'
func (i *VerifiedBiZoneIRPIncident) SetAnyAttackType(a any) error {
	return i.SetAttackType(fmt.Sprint(a))
}

// GetGossopkaSendTime для поля 'gossopka_send_time'
func (i *VerifiedBiZoneIRPIncident) GetGossopkaSendTime() string {
	return i.GossopkaSendTime
}

// SetGossopkaSendTime для поля 'gossopka_send_time'
func (i *VerifiedBiZoneIRPIncident) SetGossopkaSendTime(v string) error {
	i.GossopkaSendTime = v

	return nil
}

// SetAnyGossopkaSendTime для поля 'gossopka_send_time'
func (i *VerifiedBiZoneIRPIncident) SetAnyGossopkaSendTime(a any) error {
	return i.SetGossopkaSendTime(fmt.Sprint(a))
}

// GetGtiSendTime для поля 'gti_send_time'
func (i *VerifiedBiZoneIRPIncident) GetGtiSendTime() string {
	return i.GtiSendTime
}

// SetGtiSendTime для поля 'gti_send_time'
func (i *VerifiedBiZoneIRPIncident) SetGtiSendTime(v string) error {
	i.GtiSendTime = v

	return nil
}

// SetAnyGtiSendTime для поля 'gti_send_time'
func (i *VerifiedBiZoneIRPIncident) SetAnyGtiSendTime(a any) error {
	return i.SetGtiSendTime(fmt.Sprint(a))
}

// GetCustomerStarRatingComment для поля 'customer_star_rating_comment'
func (i *VerifiedBiZoneIRPIncident) GetCustomerStarRatingComment() string {
	return i.CustomerStarRatingComment
}

// SetCustomerStarRatingComment для поля 'customer_star_rating_comment'
func (i *VerifiedBiZoneIRPIncident) SetCustomerStarRatingComment(v string) error {
	i.CustomerStarRatingComment = v

	return nil
}

// SetAnyCustomerStarRatingComment для поля 'customer_star_rating_comment'
func (i *VerifiedBiZoneIRPIncident) SetAnyCustomerStarRatingComment(a any) error {
	return i.SetCustomerStarRatingComment(fmt.Sprint(a))
}

// GetRecommendations для поля 'recommendations'
func (i *VerifiedBiZoneIRPIncident) GetRecommendations() string {
	return i.Recommendations
}

// SetRecommendations для поля 'recommendations'
func (i *VerifiedBiZoneIRPIncident) SetRecommendations(v string) error {
	i.Recommendations = v

	return nil
}

// SetAnyRecommendations для поля 'recommendations'
func (i *VerifiedBiZoneIRPIncident) SetAnyRecommendations(a any) error {
	return i.SetRecommendations(fmt.Sprint(a))
}

// GetAffectedService для поля 'affected_service'
func (i *VerifiedBiZoneIRPIncident) GetAffectedService() string {
	return i.AffectedService
}

// SetAffectedService для поля 'affected_service'
func (i *VerifiedBiZoneIRPIncident) SetAffectedService(v string) error {
	i.AffectedService = v

	return nil
}

// SetAnyAffectedService для поля 'affected_service'
func (i *VerifiedBiZoneIRPIncident) SetAnyAffectedService(a any) error {
	return i.SetAffectedService(fmt.Sprint(a))
}

// GetFakeAsPath для поля 'fake_as_path'
func (i *VerifiedBiZoneIRPIncident) GetFakeAsPath() string {
	return i.FakeAsPath
}

// SetFakeAsPath для поля 'fake_as_path'
func (i *VerifiedBiZoneIRPIncident) SetFakeAsPath(v string) error {
	i.FakeAsPath = v

	return nil
}

// SetAnyFakeAsPath для поля 'fake_as_path'
func (i *VerifiedBiZoneIRPIncident) SetAnyFakeAsPath(a any) error {
	return i.SetFakeAsPath(fmt.Sprint(a))
}

// GetFakePrefix для поля 'fake_prefix'
func (i *VerifiedBiZoneIRPIncident) GetFakePrefix() string {
	return i.FakePrefix
}

// SetFakePrefix для поля 'fake_prefix'
func (i *VerifiedBiZoneIRPIncident) SetFakePrefix(v string) error {
	i.FakePrefix = v

	return nil
}

// SetAnyFakePrefix для поля 'fake_prefix'
func (i *VerifiedBiZoneIRPIncident) SetAnyFakePrefix(a any) error {
	return i.SetFakePrefix(fmt.Sprint(a))
}

// GetFpType для поля 'fp_type'
func (i *VerifiedBiZoneIRPIncident) GetFpType() string {
	return i.FpType
}

// SetFpType для поля 'fp_type'
func (i *VerifiedBiZoneIRPIncident) SetFpType(v string) error {
	i.FpType = v

	return nil
}

// SetAnyFpType для поля 'fp_type'
func (i *VerifiedBiZoneIRPIncident) SetAnyFpType(a any) error {
	return i.SetFpType(fmt.Sprint(a))
}

// GetLookingGlass для поля 'looking_glass'
func (i *VerifiedBiZoneIRPIncident) GetLookingGlass() string {
	return i.LookingGlass
}

// SetLookingGlass для поля 'looking_glass'
func (i *VerifiedBiZoneIRPIncident) SetLookingGlass(v string) error {
	i.LookingGlass = v

	return nil
}

// SetAnyLookingGlass для поля 'looking_glass'
func (i *VerifiedBiZoneIRPIncident) SetAnyLookingGlass(a any) error {
	return i.SetLookingGlass(fmt.Sprint(a))
}

// GetSpamRecipients для поля 'spam_recipients'
func (i *VerifiedBiZoneIRPIncident) GetSpamRecipients() string {
	return i.SpamRecipients
}

// SetSpamRecipients для поля 'spam_recipients'
func (i *VerifiedBiZoneIRPIncident) SetSpamRecipients(v string) error {
	i.SpamRecipients = v

	return nil
}

// SetAnySpamRecipients для поля 'spam_recipients'
func (i *VerifiedBiZoneIRPIncident) SetAnySpamRecipients(a any) error {
	return i.SetSpamRecipients(fmt.Sprint(a))
}

// GetTLP для поля 'tlp'
func (i *VerifiedBiZoneIRPIncident) GetTLP() string {
	return i.TLP
}

// SetTLP для поля 'tlp'
func (i *VerifiedBiZoneIRPIncident) SetTLP(v string) error {
	i.TLP = v

	return nil
}

// SetAnyTLP для поля 'tlp'
func (i *VerifiedBiZoneIRPIncident) SetAnyTLP(a any) error {
	return i.SetTLP(fmt.Sprint(a))
}

// GetUsualPrefix для поля 'usual_prefix'
func (i *VerifiedBiZoneIRPIncident) GetUsualPrefix() string {
	return i.UsualPrefix
}

// SetUsualPrefix для поля 'usual_prefix'
func (i *VerifiedBiZoneIRPIncident) SetUsualPrefix(v string) error {
	i.UsualPrefix = v

	return nil
}

// SetAnyUsualPrefix для поля 'usual_prefix'
func (i *VerifiedBiZoneIRPIncident) SetAnyUsualPrefix(a any) error {
	return i.SetUsualPrefix(fmt.Sprint(a))
}

// GetUsualAsPath для поля 'usual_as_path'
func (i *VerifiedBiZoneIRPIncident) GetUsualAsPath() string {
	return i.UsualAsPath
}

// SetUsualAsPath для поля 'usual_as_path'
func (i *VerifiedBiZoneIRPIncident) SetUsualAsPath(v string) error {
	i.UsualAsPath = v

	return nil
}

// SetAnyUsualAsPath для поля 'usual_as_path'
func (i *VerifiedBiZoneIRPIncident) SetAnyUsualAsPath(a any) error {
	return i.SetUsualAsPath(fmt.Sprint(a))
}

// GetSource для поля 'source'
func (i *VerifiedBiZoneIRPIncident) GetSource() string {
	return i.Source
}

// SetSource для поля 'source'
func (i *VerifiedBiZoneIRPIncident) SetSource(v string) error {
	i.Source = v

	return nil
}

// SetAnySource для поля 'source'
func (i *VerifiedBiZoneIRPIncident) SetAnySource(a any) error {
	return i.SetSource(fmt.Sprint(a))
}

// GetExternalId для поля 'external_id'
func (i *VerifiedBiZoneIRPIncident) GetExternalId() string {
	return i.ExternalId
}

// SetExternalId для поля 'external_id'
func (i *VerifiedBiZoneIRPIncident) SetExternalId(v string) error {
	i.ExternalId = v

	return nil
}

// SetAnyExternalId для поля 'external_id'
func (i *VerifiedBiZoneIRPIncident) SetAnyExternalId(a any) error {
	return i.SetExternalId(fmt.Sprint(a))
}

// GetUnderlineSource для поля '_source'
func (i *VerifiedBiZoneIRPIncident) GetUnderlineSource() string {
	return i.UnderlineSource
}

// SetUnderlineSource для поля '_source'
func (i *VerifiedBiZoneIRPIncident) SetUnderlineSource(v string) error {
	i.UnderlineSource = v

	return nil
}

// SetAnyUnderlineSource для поля '_source'
func (i *VerifiedBiZoneIRPIncident) SetAnyUnderlineSource(a any) error {
	return i.SetUnderlineSource(fmt.Sprint(a))
}

// GetGossopkaKey для поля 'gossopka_key'
func (i *VerifiedBiZoneIRPIncident) GetGossopkaKey() string {
	return i.GossopkaKey
}

// SetGossopkaKey для поля 'gossopka_key'
func (i *VerifiedBiZoneIRPIncident) SetGossopkaKey(v string) error {
	i.GossopkaKey = v

	return nil
}

// SetAnyGossopkaKey для поля 'gossopka_key'
func (i *VerifiedBiZoneIRPIncident) SetAnyGossopkaKey(a any) error {
	return i.SetGossopkaKey(fmt.Sprint(a))
}

// GetService для поля 'service'
func (i *VerifiedBiZoneIRPIncident) GetService() string {
	return i.Service
}

// SetService для поля 'service'
func (i *VerifiedBiZoneIRPIncident) SetService(v string) error {
	i.Service = v

	return nil
}

// SetAnyService для поля 'service'
func (i *VerifiedBiZoneIRPIncident) SetAnyService(a any) error {
	return i.SetService(fmt.Sprint(a))
}

// GetAssignee для поля 'assignee'
func (i *VerifiedBiZoneIRPIncident) GetAssignee() string {
	return i.Assignee
}

// SetAssignee для поля 'assignee'
func (i *VerifiedBiZoneIRPIncident) SetAssignee(v string) error {
	i.Assignee = v

	return nil
}

// SetAnyAssignee для поля 'assignee'
func (i *VerifiedBiZoneIRPIncident) SetAnyAssignee(a any) error {
	return i.SetAssignee(fmt.Sprint(a))
}

// GetGtiId для поля 'gti_id'
func (i *VerifiedBiZoneIRPIncident) GetGtiId() string {
	return i.GtiId
}

// SetGtiId для поля 'gti_id'
func (i *VerifiedBiZoneIRPIncident) SetGtiId(v string) error {
	i.GtiId = v

	return nil
}

// SetAnyGtiId для поля 'gti_id'
func (i *VerifiedBiZoneIRPIncident) SetAnyGtiId(a any) error {
	return i.SetGtiId(fmt.Sprint(a))
}

// GetCustomerStarRating для поля 'customer_star_rating'
func (i *VerifiedBiZoneIRPIncident) GetCustomerStarRating() string {
	return i.CustomerStarRating
}

// SetCustomerStarRating для поля 'customer_star_rating'
func (i *VerifiedBiZoneIRPIncident) SetCustomerStarRating(v string) error {
	i.CustomerStarRating = v

	return nil
}

// SetAnyCustomerStarRating для поля 'customer_star_rating'
func (i *VerifiedBiZoneIRPIncident) SetAnyCustomerStarRating(a any) error {
	return i.SetCustomerStarRating(fmt.Sprint(a))
}

// GetGossopkaId для поля 'gossopka_id'
func (i *VerifiedBiZoneIRPIncident) GetGossopkaId() string {
	return i.GossopkaId
}

// SetGossopkaId для поля 'gossopka_id'
func (i *VerifiedBiZoneIRPIncident) SetGossopkaId(v string) error {
	i.GossopkaId = v

	return nil
}

// SetAnyGossopkaId для поля 'gossopka_id'
func (i *VerifiedBiZoneIRPIncident) SetAnyGossopkaId(a any) error {
	return i.SetGossopkaId(fmt.Sprint(a))
}

// GetMitreCov для поля 'mitre_cov'
func (i *VerifiedBiZoneIRPIncident) GetMitreCov() string {
	return i.MitreCov
}

// SetMitreCov для поля 'mitre_cov'
func (i *VerifiedBiZoneIRPIncident) SetMitreCov(v string) error {
	i.MitreCov = v

	return nil
}

// SetAnyMitreCov для поля 'mitre_cov'
func (i *VerifiedBiZoneIRPIncident) SetAnyMitreCov(a any) error {
	return i.SetMitreCov(fmt.Sprint(a))
}

// GetResponseTeam для поля 'response_team'
func (i *VerifiedBiZoneIRPIncident) GetResponseTeam() string {
	return i.ResponseTeam
}

// SetResponseTeam для поля 'response_team'
func (i *VerifiedBiZoneIRPIncident) SetResponseTeam(v string) error {
	i.ResponseTeam = v

	return nil
}

// SetAnyResponseTeam для поля 'response_team'
func (i *VerifiedBiZoneIRPIncident) SetAnyResponseTeam(a any) error {
	return i.SetResponseTeam(fmt.Sprint(a))
}

// GetGossopkaErrors для поля 'gossopka_errors'
func (i *VerifiedBiZoneIRPIncident) GetGossopkaErrors() string {
	return i.GossopkaErrors
}

// SetGossopkaErrors для поля 'gossopka_errors'
func (i *VerifiedBiZoneIRPIncident) SetGossopkaErrors(v string) error {
	i.GossopkaErrors = v

	return nil
}

// SetAnyGossopkaErrors для поля 'gossopka_errors'
func (i *VerifiedBiZoneIRPIncident) SetAnyGossopkaErrors(a any) error {
	return i.SetGossopkaErrors(fmt.Sprint(a))
}

// GetGossopkaKeyLink для поля 'gossopka_key_link'
func (i *VerifiedBiZoneIRPIncident) GetGossopkaKeyLink() string {
	return i.GossopkaKeyLink
}

// SetGossopkaKeyLink для поля 'gossopka_key_link'
func (i *VerifiedBiZoneIRPIncident) SetGossopkaKeyLink(v string) error {
	i.GossopkaKeyLink = v

	return nil
}

// SetAnyGossopkaKeyLink для поля 'gossopka_key_link'
func (i *VerifiedBiZoneIRPIncident) SetAnyGossopkaKeyLink(a any) error {
	return i.SetGossopkaKeyLink(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (i *VerifiedBiZoneIRPIncident) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	//wsStr := supportingfunctions.GetWhitespace(num + 1)
	wsInt := supportingfunctions.GetWhitespace(num + 2)

	fmt.Fprintf(&str, "%s'created': '%s'\n", ws, i.Created)
	fmt.Fprintf(&str, "%s'updated': '%s'\n", ws, i.Updated)
	fmt.Fprintf(&str, "%s'timestamp': '%s'\n", ws, i.Timestamp)
	fmt.Fprintf(&str, "%s'updated_all': '%s'\n", ws, i.UpdatedAll)
	fmt.Fprintf(&str, "%s'detection_date': '%s'\n", ws, i.DetectionDate)
	fmt.Fprintf(&str, "%s'first_sent_notification_time': '%s'\n", ws, i.FirstSentNotificationTime)
	fmt.Fprintf(&str, "%s'key': '%s'\n", ws, i.Key)
	fmt.Fprintf(&str, "%s'status': '%s'\n", ws, i.Status)
	fmt.Fprintf(&str, "%s'summary': '%s'\n", ws, i.Summary)
	fmt.Fprintf(&str, "%s'priority': '%s'\n", ws, i.Priority)
	fmt.Fprintf(&str, "%s'tenant_id': '%s'\n", ws, i.TenantId)
	fmt.Fprintf(&str, "%s'issue_type': '%s'\n", ws, i.IssueType)
	fmt.Fprintf(&str, "%s'description': '%s'\n", ws, i.Description)
	fmt.Fprintf(&str, "%s'rendered_description': '%s'\n", ws, i.RenderedDescription)
	fmt.Fprintf(&str, "%s'status_description': '%s'\n", ws, i.StatusDescription)
	fmt.Fprintf(&str, "%s'assignee_displayName': '%s'\n", ws, i.AssigneeDisplayName)
	fmt.Fprintf(&str, "%s'reporter_displayName': '%s'\n", ws, i.ReporterDisplayName)
	fmt.Fprintf(&str, "%s'creator_displayName': '%s'\n", ws, i.CreatorDisplayName)
	fmt.Fprintf(&str, "%s'reporter_emailAddress': '%s'\n", ws, i.ReporterEmailAddress)
	fmt.Fprintf(&str, "%s'primary_category': '%s'\n", ws, i.PrimaryCategory)
	fmt.Fprintf(&str, "%s'resolutiondate': '%s'\n", ws, i.ResolutionDate)
	fmt.Fprintf(&str, "%s'resolution_name': '%s'\n", ws, i.ResolutionName)
	fmt.Fprintf(&str, "%s'resolution_detailed': '%s'\n", ws, i.ResolutionDetailed)
	fmt.Fprintf(&str, "%s'platform_url': '%s'\n", ws, i.PlatformUrl)
	fmt.Fprintf(&str, "%s'attack_type': '%s'\n", ws, i.AttackType)
	fmt.Fprintf(&str, "%s'gossopka_send_time': '%s'\n", ws, i.GossopkaSendTime)
	fmt.Fprintf(&str, "%s'gti_send_time': '%s'\n", ws, i.GtiSendTime)
	fmt.Fprintf(&str, "%s'customer_star_rating_comment': '%s'\n", ws, i.CustomerStarRatingComment)
	fmt.Fprintf(&str, "%s'recommendations': '%s'\n", ws, i.Recommendations)
	fmt.Fprintf(&str, "%s'affected_service': '%s'\n", ws, i.AffectedService)
	fmt.Fprintf(&str, "%s'fake_as_path': '%s'\n", ws, i.FakeAsPath)
	fmt.Fprintf(&str, "%s'fake_prefix': '%s'\n", ws, i.FakePrefix)
	fmt.Fprintf(&str, "%s'fp_type': '%s'\n", ws, i.FpType)
	fmt.Fprintf(&str, "%s'looking_glass': '%s'\n", ws, i.LookingGlass)
	fmt.Fprintf(&str, "%s'spam_recipients': '%s'\n", ws, i.SpamRecipients)
	fmt.Fprintf(&str, "%s'tlp': '%s'\n", ws, i.TLP)
	fmt.Fprintf(&str, "%s'usual_prefix': '%s'\n", ws, i.UsualPrefix)
	fmt.Fprintf(&str, "%s'usual_as_path': '%s'\n", ws, i.UsualAsPath)
	fmt.Fprintf(&str, "%s'source': '%s'\n", ws, i.Source)
	fmt.Fprintf(&str, "%s'external_id': '%s'\n", ws, i.ExternalId)
	fmt.Fprintf(&str, "%s'_source': '%s'\n", ws, i.UnderlineSource)
	fmt.Fprintf(&str, "%s'gossopka_key': '%s'\n", ws, i.GossopkaKey)
	fmt.Fprintf(&str, "%s'service': '%s'\n", ws, i.Service)
	fmt.Fprintf(&str, "%s'assignee': '%s'\n", ws, i.Assignee)
	fmt.Fprintf(&str, "%s'gti_id': '%s'\n", ws, i.GtiId)
	fmt.Fprintf(&str, "%s'customer_star_rating': '%s'\n", ws, i.CustomerStarRating)
	fmt.Fprintf(&str, "%s'gossopka_id': '%s'\n", ws, i.GossopkaId)
	fmt.Fprintf(&str, "%s'mitre_cov': '%s'\n", ws, i.MitreCov)
	fmt.Fprintf(&str, "%s'response_team': '%s'\n", ws, i.ResponseTeam)
	fmt.Fprintf(&str, "%s'gossopka_errors': '%s'\n", ws, i.GossopkaErrors)
	fmt.Fprintf(&str, "%s'gossopka_key_link': '%s'\n", ws, i.GossopkaKeyLink)
	fmt.Fprintf(&str, "%s'id': '%d'\n", ws, i.ID)
	fmt.Fprintf(&str, "%s'system': '%d'\n", ws, i.System)
	fmt.Fprintf(&str, "%s'is_visible_for_customer': '%t'\n", ws, i.IsVisibleForCustomer)
	fmt.Fprintf(&str, "%s'show_gossopka_button': '%t'\n", ws, i.ShowGossopkaButton)
	fmt.Fprintf(&str, "%s'cancel_gossopka_send_button': '%t'\n", ws, i.CancelGossopkaSendButton)
	fmt.Fprintf(&str, "%s'show_gti_button': '%t'\n", ws, i.ShowGtiButton)
	fmt.Fprintf(&str, "%s'tags':\n", ws)
	for k, v := range i.Tags {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInt, k, v.ToStringBeautiful(num+2))
	}
	fmt.Fprintf(&str, "%s'observed_iocs':\n", ws)
	for k, v := range i.ObservedIocs {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInt, k, v.ToStringBeautiful(num+2))
	}
	fmt.Fprintf(&str, "%s'secondary_category_ref':\n", ws)
	for k, v := range i.SecondaryCategoryRef {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInt, k, v.ToStringBeautiful(num+2))
	}
	fmt.Fprintf(&str, "%s'keyfield': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.KeyField))
	fmt.Fprintf(&str, "%s'activity_detected': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.ActivityDetected))
	fmt.Fprintf(&str, "%s'secondary_category': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.SecondaryCategory))
	fmt.Fprintf(&str, "%s'platform_hostname': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.PlatformHostname))
	fmt.Fprintf(&str, "%s'detection_rules': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.DetectionRules))
	fmt.Fprintf(&str, "%s'parsed_activity_detected': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.ParsedActivityDetected))
	fmt.Fprintf(&str, "%s'attachments': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.Attachments))
	fmt.Fprintf(&str, "%s'emls': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.Emls))
	fmt.Fprintf(&str, "%s'slas': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.Slas))
	fmt.Fprintf(&str, "%s'badges': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.Badges))
	fmt.Fprintf(&str, "%s'watchers': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.Watchers))
	fmt.Fprintf(&str, "%s'watchers_id': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, i.WatchersId))
	fmt.Fprintf(&str, "%s'issue_type_ref':\n%s", ws, i.IssueTypeRef.ToStringBeautiful(num+1))

	fmt.Fprintf(&str, "%s'primary_category_ref':\n%s", ws, i.PrimaryCategoryRef.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'resolution_name_ref':\n%s", ws, i.ResolutionNameRef.ToStringBeautiful(num+1))

	return str.String()
}
