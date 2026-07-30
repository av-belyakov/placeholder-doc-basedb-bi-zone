package datamodels

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPCase() *VerifiedIRPBiZoneCase {
	return &VerifiedIRPBiZoneCase{
		GossopkaErrors:       make(map[string]any),
		ObservedIocs:         make([]BiZoneIRPObservedIocs, 0),
		SecondaryCategoryRef: make([]BiZoneIRPTypeRef, 0),
		DetectionRules:       make([]string, 0),
		SecondaryCategory:    make([]string, 0),
		PlatformHostname:     make([]string, 0),
		WatchersId:           make([]uint64, 0),
		//		ActivityDetected:     make([]any, 0),
		//		Attachments:          make([]any, 0),
		//		Badges:               make([]any, 0),
		//		Emls:                 make([]any, 0),
		//		Slas:                 make([]any, 0),
		//		Tags:                 make([]any, 0),
		//		KeyField:             make([]any, 0),
	}
}

// GetGossopkaErrors список ошибок
func (vc *VerifiedIRPBiZoneCase) GetGossopkaErrors() map[string]any {
	if vc.GossopkaErrors == nil {
		vc.GossopkaErrors = make(map[string]any)
	}

	return vc.GossopkaErrors
}

// SetGossopkaErrors список ошибок
func (vc *VerifiedIRPBiZoneCase) SetGossopkaErrors(errors map[string]any) {
	vc.GossopkaErrors = errors
}

// SetAnyGossopkaErrors список ошибок
func (vc *VerifiedIRPBiZoneCase) SetAnyGossopkaErrors(k string, v any) bool {
	vc.GetGossopkaErrors()

	if _, ok := vc.GossopkaErrors[k]; !ok {
		vc.GossopkaErrors[k] = []string(nil)
	}

	return true
}

// GetObservedIocs список наблюдаемых IOC
func (vc *VerifiedIRPBiZoneCase) GetObservedIocs() []BiZoneIRPObservedIocs {
	if vc.ObservedIocs == nil {
		vc.ObservedIocs = make([]BiZoneIRPObservedIocs, 0)
	}

	return vc.ObservedIocs
}

// SetObservedIocs список наблюдаемых IOC
func (vc *VerifiedIRPBiZoneCase) SetObservedIocs(list []BiZoneIRPObservedIocs) {
	vc.ObservedIocs = list
}

// AddObservedIocs добавить значение в список наблюдаемых IOC
func (vc *VerifiedIRPBiZoneCase) AddObservedIocs(v BiZoneIRPObservedIocs) {
	vc.GetObservedIocs()

	vc.ObservedIocs = append(vc.ObservedIocs, v)
}

// GetSecondaryCategoryRef список ссылок на категории
func (vc *VerifiedIRPBiZoneCase) GetSecondaryCategoryRef() []BiZoneIRPTypeRef {
	if vc.SecondaryCategoryRef == nil {
		vc.SecondaryCategoryRef = make([]BiZoneIRPTypeRef, 0)
	}

	return vc.SecondaryCategoryRef
}

// SetSecondaryCategoryRef список ссылок на категории
func (vc *VerifiedIRPBiZoneCase) SetSecondaryCategoryRef(list []BiZoneIRPTypeRef) {
	vc.SecondaryCategoryRef = list
}

// GetWatchers список наблюдателей
func (vc *VerifiedIRPBiZoneCase) GetWatchers() []BiZoneIRPWatcher {
	if vc.Watchers == nil {
		vc.Watchers = make([]BiZoneIRPWatcher, 0)
	}

	return vc.Watchers
}

// SetWatchers список наблюдателей
func (vc *VerifiedIRPBiZoneCase) SetWatchers(list []BiZoneIRPWatcher) {
	vc.Watchers = list
}

// AddSecondaryCategoryRef добавляет ссылку в список категорий
func (vc *VerifiedIRPBiZoneCase) AddSecondaryCategoryRef(v BiZoneIRPTypeRef) {
	vc.GetSecondaryCategoryRef()

	vc.SecondaryCategoryRef = append(vc.SecondaryCategoryRef, v)
}

// GetActivityDetected список обнаруженной активности
func (vc *VerifiedIRPBiZoneCase) GetActivityDetected() []any {
	if vc.ActivityDetected == nil {
		vc.ActivityDetected = make([]any, 0)
	}

	return vc.ActivityDetected
}

// SetActivityDetected список обнаруженной активности
func (vc *VerifiedIRPBiZoneCase) SetActivityDetected(list []any) {
	vc.ActivityDetected = list
}

// AddActivityDetected обнаруженная активность
func (vc *VerifiedIRPBiZoneCase) AddActivityDetected(a any) {
	vc.GetActivityDetected()

	vc.ActivityDetected = append(vc.ActivityDetected, a)
}

// GetAttachments список связий
func (vc *VerifiedIRPBiZoneCase) GetAttachments() []any {
	if vc.Attachments == nil {
		vc.Assignee = make([]any, 0)
	}

	return vc.Attachments
}

// SetAttachments список связей
func (vc *VerifiedIRPBiZoneCase) SetAttachments(list []any) {
	vc.Attachments = list
}

// AddAttachments связи
func (vc *VerifiedIRPBiZoneCase) AddAttachments(v any) {
	vc.GetAttachments()

	vc.Attachments = append(vc.Attachments, v)
}

// GetBadges значки
func (vc *VerifiedIRPBiZoneCase) GetBadges() []any {
	if vc.Badges == nil {
		vc.Badges = make([]any, 0)
	}

	return vc.Badges
}

// SetBadges значки
func (vc *VerifiedIRPBiZoneCase) SetBadges(list []any) {
	vc.Badges = list
}

// AddBadges значок
func (vc *VerifiedIRPBiZoneCase) AddBadges(v any) {
	vc.GetBadges()

	vc.Badges = append(vc.Badges, v)
}

// GetEmls
func (vc *VerifiedIRPBiZoneCase) GetEmls() []any {
	if vc.Emls == nil {
		vc.Emls = make([]any, 0)
	}

	return vc.Emls
}

// SetEmls
func (vc *VerifiedIRPBiZoneCase) SetEmls(list []any) {
	vc.Emls = list
}

// AddEmls
func (vc *VerifiedIRPBiZoneCase) AddEmls(a any) {
	vc.GetEmls()

	vc.Emls = append(vc.Emls, a)
}

// GetSlas
func (vc *VerifiedIRPBiZoneCase) GetSlas() []any {
	if vc.Slas == nil {
		vc.Slas = make([]any, 0)
	}

	return vc.Slas
}

// SetSlas
func (vc *VerifiedIRPBiZoneCase) SetSlas(list []any) {
	vc.Slas = list
}

// AddSlas
func (vc *VerifiedIRPBiZoneCase) AddSlas(a any) {
	vc.GetSlas()

	vc.Slas = append(vc.Slas, a)
}

// GetTags теги
func (vc *VerifiedIRPBiZoneCase) GetTags() []any {
	if vc.Tags == nil {
		vc.Tags = make([]any, 0)
	}

	return vc.Tags
}

// SetTags теги
func (vc *VerifiedIRPBiZoneCase) SetTags(list []any) {
	vc.Tags = list
}

// AddTags теги
func (vc *VerifiedIRPBiZoneCase) AddTags(a any) {
	vc.GetTags()

	vc.Tags = append(vc.Tags, a)
}

// GetKeyField ключевые поля
func (vc *VerifiedIRPBiZoneCase) GetKeyField() []any {
	if vc.KeyField == nil {
		vc.KeyField = make([]any, 0)
	}

	return vc.KeyField
}

// SetKeyField ключевые поля
func (vc *VerifiedIRPBiZoneCase) SetKeyField(list []any) {
	vc.KeyField = list
}

// AddKeyField ключевые поля
func (vc *VerifiedIRPBiZoneCase) AddKeyField(a any) {
	vc.GetKeyField()

	vc.KeyField = append(vc.KeyField, a)
}

// GetMITRECOV
func (vc *VerifiedIRPBiZoneCase) GetMITRECOV() BiZoneIRPMITRECOV {
	if vc.MITRECOV.GetPersistence() == nil {
		vc.MITRECOV = *NewBiZoneMITRECOV()
	}

	return vc.MITRECOV
}

// SetMITRECOV
func (vc *VerifiedIRPBiZoneCase) SetMITRECOV(v BiZoneIRPMITRECOV) {
	vc.MITRECOV = v
}

// GetDetectionRules правила обнаружения
func (vc *VerifiedIRPBiZoneCase) GetDetectionRules() []string {
	if vc.DetectionRules == nil {
		return make([]string, 0)
	}

	return vc.DetectionRules
}

// SetDetectionRules правила обнаружения
func (vc *VerifiedIRPBiZoneCase) SetDetectionRules(list []string) {
	vc.DetectionRules = list
}

// AddDetectionRules добавляет правило в список
func (vc *VerifiedIRPBiZoneCase) AddDetectionRules(v string) {
	vc.GetDetectionRules()

	if slices.Contains(vc.DetectionRules, v) {
		return
	}

	vc.DetectionRules = append(vc.DetectionRules, v)
}

// AddAnyDetectionRule добавляет правило в список
func (vc *VerifiedIRPBiZoneCase) AddAnyDetectionRule(a any) {
	vc.AddDetectionRules(fmt.Sprint(a))
}

// GetSecondaryCategory список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) GetSecondaryCategory() []string {
	if vc.SecondaryCategory == nil {
		vc.SecondaryCategory = make([]string, 0)
	}

	return vc.SecondaryCategory
}

// SetSecondaryCategories список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) SetSecondaryCategories(list []string) {
	vc.SecondaryCategory = list
}

// AddSecondaryCategory добавляет значение в список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) AddSecondaryCategory(v string) {
	vc.GetSecondaryCategory()

	if slices.Contains(vc.SecondaryCategory, v) {
		return
	}

	vc.SecondaryCategory = append(vc.SecondaryCategory, v)
}

// AddAnySecondaryCategory добавляет значение в список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) AddAnySecondaryCategory(a any) {
	vc.AddSecondaryCategory(fmt.Sprint(a))
}

// GetPlatformHostname список с наименованиями платформы
func (vc *VerifiedIRPBiZoneCase) GetPlatformHostname() []string {
	if vc.PlatformHostname == nil {
		vc.PlatformHostname = make([]string, 0)
	}

	return vc.PlatformHostname
}

// SetPlatformHostname список с наименованиями платформы
func (vc *VerifiedIRPBiZoneCase) SetPlatformHostname(list []string) {
	vc.PlatformHostname = list
}

// AddPlatformHostname добавляетнаименование платформы в список
func (vc *VerifiedIRPBiZoneCase) AddPlatformHostname(v string) {
	vc.GetPlatformHostname()

	if slices.Contains(vc.PlatformHostname, v) {
		return
	}

	vc.PlatformHostname = append(vc.PlatformHostname, v)
}

// AddAnyPlatformHostname добавляетнаименование платформы в список
func (vc *VerifiedIRPBiZoneCase) AddAnyPlatformHostname(a any) {
	vc.AddPlatformHostname(fmt.Sprint(a))
}

// GetWatchersId список идентификаторов наблюдения
func (vc *VerifiedIRPBiZoneCase) GetWatchersId() []uint64 {
	if vc.WatchersId == nil {
		vc.WatchersId = make([]uint64, 0)
	}

	return vc.WatchersId
}

// SetWatchersId список идентификаторов наблюдения
func (vc *VerifiedIRPBiZoneCase) SetWatchersId(list []uint64) {
	vc.WatchersId = list
}

// AddWatchersId добавить идентификатор в список идентификаторов наблюдения
func (vc *VerifiedIRPBiZoneCase) AddWatchersId(v uint64) {
	vc.GetWatchersId()

	if slices.Contains(vc.WatchersId, v) {
		return
	}

	vc.WatchersId = append(vc.WatchersId, v)
}

// AddAnyPlatformHostname добавляетнаименование платформы в список
func (vc *VerifiedIRPBiZoneCase) AddAnyWatchersId(a any) {
	switch v := a.(type) {
	case int:
		vc.AddWatchersId(uint64(v))

	case int32:
		vc.AddWatchersId(uint64(v))

	case int64:
		vc.AddWatchersId(uint64(v))

	case float32:
		vc.AddWatchersId(uint64(v))

	case float64:
		vc.AddWatchersId(uint64(v))
	}
}

// GetAssignee правоприемник
func (vc *VerifiedIRPBiZoneCase) GetAssignee() any {
	return vc.Assignee
}

// SetAssignee правоприемник
func (vc *VerifiedIRPBiZoneCase) SetAssignee(a any) {
	vc.Assignee = a
}

// GetAssigneeDisplayName отображаемое имя правопреемника
func (vc *VerifiedIRPBiZoneCase) GetAssigneeDisplayName() any {
	return vc.AssigneeDisplayName
}

// SetAssigneeDisplayName отображаемое имя правопреемника
func (vc *VerifiedIRPBiZoneCase) SetAssigneeDisplayName(a any) {
	vc.AssigneeDisplayName = a
}

// GetService сервис
func (vc *VerifiedIRPBiZoneCase) GetService() any {
	return vc.Service
}

// SetService сервис
func (vc *VerifiedIRPBiZoneCase) SetService(a any) {
	vc.Service = a
}

// GetResolutionDate дата разрешения
func (vc *VerifiedIRPBiZoneCase) GetResolutionDate() any {
	return vc.ResolutionDate
}

// SetResolutionDate дата разрешения
func (vc *VerifiedIRPBiZoneCase) SetResolutionDate(a any) {
	vc.ResolutionDate = a
}

// GetResolutionName имя разрешения
func (vc *VerifiedIRPBiZoneCase) GetResolutionName() any {
	return vc.ResolutionName
}

// SetResolutionName имя разрешения
func (vc *VerifiedIRPBiZoneCase) SetResolutionName(a any) {
	vc.ResolutionName = a
}

// GetResolutionNameRef ссылка на имя разрешения
func (vc *VerifiedIRPBiZoneCase) GetResolutionNameRef() any {
	return vc.ResolutionNameRef
}

// SetResolutionNameRef ссылка на имя разрешения
func (vc *VerifiedIRPBiZoneCase) SetResolutionNameRef(a any) {
	vc.ResolutionNameRef = a
}

// GetGossopkaId идентификатор
func (vc *VerifiedIRPBiZoneCase) GetGossopkaId() any {
	return vc.GossopkaId
}

// SetGossopkaId идентификатор
func (vc *VerifiedIRPBiZoneCase) SetGossopkaId(a any) {
	vc.GossopkaId = a
}

// GetGossopkaSendTime установленное время (возможно типа string или time)
func (vc *VerifiedIRPBiZoneCase) GetGossopkaSendTime() any {
	return vc.GossopkaSendTime
}

// SetGossopkaSendTime установленное время (возможно типа string или time)
func (vc *VerifiedIRPBiZoneCase) SetGossopkaSendTime(a any) {
	vc.GossopkaSendTime = a
}

// GetGtiId идентификатор
func (vc *VerifiedIRPBiZoneCase) GetGtiId() any {
	return vc.GtiId
}

// SetGtiId идентификатор
func (vc *VerifiedIRPBiZoneCase) SetGtiId(a any) {
	vc.GtiId = a
}

// GetGtiSendTime некое время
func (vc *VerifiedIRPBiZoneCase) GetGtiSendTime() any {
	return vc.GtiSendTime
}

// SetGtiSendTime некое время
func (vc *VerifiedIRPBiZoneCase) SetGtiSendTime(a any) {
	vc.GtiSendTime = a
}

// GetCustomerStarRating звёздный рейтинг потребителей
func (vc *VerifiedIRPBiZoneCase) GetCustomerStarRating() any {
	return vc.CustomerStarRating
}

// SetCustomerStarRating звёздный рейтинг потребителей
func (vc *VerifiedIRPBiZoneCase) SetCustomerStarRating(a any) {
	vc.CustomerStarRating = a
}

// GetFirstSentNotificationTime время первого сообщения
func (vc *VerifiedIRPBiZoneCase) GetFirstSentNotificationTime() any {
	return vc.FirstSentNotificationTime
}

// SetFirstSentNotificationTime время первого сообщения
func (vc *VerifiedIRPBiZoneCase) SetFirstSentNotificationTime(a any) {
	vc.FirstSentNotificationTime = a
}

// GetResponseTeam группа ответа
func (vc *VerifiedIRPBiZoneCase) GetResponseTeam() any {
	return vc.ResponseTeam
}

// SetResponseTeam группа ответа
func (vc *VerifiedIRPBiZoneCase) SetResponseTeam(a any) {
	vc.ResponseTeam = a
}

// GetGossopkaKeyLink ссылка на ключ
func (vc *VerifiedIRPBiZoneCase) GetGossopkaKeyLink() any {
	return vc.GossopkaKeyLink
}

// SetGossopkaKeyLink ссылка на ключ
func (vc *VerifiedIRPBiZoneCase) SetGossopkaKeyLink(a any) {
	vc.GossopkaKeyLink = a
}

// GetIssueTypeRef ссылка на тип проблеммы
func (vc *VerifiedIRPBiZoneCase) GetIssueTypeRef() BiZoneIRPTypeRef {
	return vc.IssueTypeRef
}

// SetIssueTypeRef ссылка на тип проблеммы
func (vc *VerifiedIRPBiZoneCase) SetIssueTypeRef(v BiZoneIRPTypeRef) {
	vc.IssueTypeRef = v
}

// GetPrimaryCategoryRef ссылка на первичную категорию
func (vc *VerifiedIRPBiZoneCase) GetPrimaryCategoryRef() BiZoneIRPTypeRef {
	return vc.PrimaryCategoryRef
}

// SetPrimaryCategoryRef ссылка на первичную категорию
func (vc *VerifiedIRPBiZoneCase) SetPrimaryCategoryRef(v BiZoneIRPTypeRef) {
	vc.PrimaryCategoryRef = v
}

// GetAttackType тип атаки
func (vc *VerifiedIRPBiZoneCase) GetAttackType() string {
	return vc.AttackType
}

// SetAttackType тип атаки
func (vc *VerifiedIRPBiZoneCase) SetAttackType(v string) {
	vc.AttackType = v
}

// SetAnyAttackType тип атаки
func (vc *VerifiedIRPBiZoneCase) SetAnyAttackType(a any) {
	vc.SetAttackType(fmt.Sprint(a))
}

// GetCreated время создания (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetCreated() string {
	return vc.Created
}

// SetCreated время создания (преобразует в формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetCreated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	vc.Created = timeStr

	return nil
}

// SetAnyCreated время создания
func (vc *VerifiedIRPBiZoneCase) SetAnyCreated(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetCreated(v)
	}

	return errors.New("type conversion error")
}

// GetCreatorDisplayName имя создающего дисплей
func (vc *VerifiedIRPBiZoneCase) GetCreatorDisplayName() string {
	return vc.CreatorDisplayName
}

// SetCreatorDisplayName имя создающего дисплей
func (vc *VerifiedIRPBiZoneCase) SetCreatorDisplayName(v string) {
	vc.CreatorDisplayName = v
}

// SetAnyCreatorDisplayName имя создающего дисплей
func (vc *VerifiedIRPBiZoneCase) SetAnyCreatorDisplayName(a any) {
	vc.SetCreatorDisplayName(fmt.Sprint(a))
}

// GetDetectionDate дата обнаружения (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetDetectionDate() string {
	return vc.DetectionDate
}

// SetDetectionDate дата обнаружения (преобразует в формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetDetectionDate(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	vc.DetectionDate = timeStr

	return nil
}

// SetAnyDetectionDate дата обнаружения
func (vc *VerifiedIRPBiZoneCase) SetAnyDetectionDate(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetDetectionDate(v)
	}

	return errors.New("type conversion error")
}

// GetKey ключ
func (vc *VerifiedIRPBiZoneCase) GetKey() string {
	return vc.Key
}

// SetKey ключ
func (vc *VerifiedIRPBiZoneCase) SetKey(v string) {
	vc.SetKey(v)
}

// SetAnyKey ключ
func (vc *VerifiedIRPBiZoneCase) SetAnyKey(a any) {
	vc.Key = fmt.Sprint(a)
}

// GetIssueType тип проблемы
func (vc *VerifiedIRPBiZoneCase) GetIssueType() string {
	return vc.IssueType
}

// SetIssueType тип проблемы
func (vc *VerifiedIRPBiZoneCase) SetIssueType(v string) {
	vc.SetIssueType(v)
}

// SetAnyIssueType тип проблемы
func (vc *VerifiedIRPBiZoneCase) SetAnyIssueType(a any) {
	vc.IssueType = fmt.Sprint(a)
}

// GetPriority приоритет
func (vc *VerifiedIRPBiZoneCase) GetPriority() string {
	return vc.Priority
}

// SetPriority приоритет
func (vc *VerifiedIRPBiZoneCase) SetPriority(v string) {
	vc.Priority = v
}

// SetAnyPriority приоритет
func (vc *VerifiedIRPBiZoneCase) SetAnyPriority(a any) {
	vc.SetPriority(fmt.Sprint(a))
}

// GetSummary резюме
func (vc *VerifiedIRPBiZoneCase) GetSummary() string {
	return vc.Summary
}

// SetSummary резюме
func (vc *VerifiedIRPBiZoneCase) SetSummary(v string) {
	vc.Summary = v
}

// SetAnySummary резюме
func (vc *VerifiedIRPBiZoneCase) SetAnySummary(a any) {
	vc.SetSummary(fmt.Sprint(a))
}

// GetDescription описание
func (vc *VerifiedIRPBiZoneCase) GetDescription() string {
	return vc.Description
}

// SetDescription описание
func (vc *VerifiedIRPBiZoneCase) SetDescription(v string) {
	vc.Description = v
}

// SetAnyDescription описание
func (vc *VerifiedIRPBiZoneCase) SetAnyDescription(a any) {
	vc.SetDescription(fmt.Sprint(a))
}

// GetRenderedDescription отрисованое описание
func (vc *VerifiedIRPBiZoneCase) GetRenderedDescription() string {
	return vc.RenderedDescription
}

// SetRenderedDescription отрисованое описание
func (vc *VerifiedIRPBiZoneCase) SetRenderedDescription(v string) {
	vc.RenderedDescription = v
}

// SetAnyRenderedDescription отрисованое описание
func (vc *VerifiedIRPBiZoneCase) SetAnyRenderedDescription(a any) {
	vc.SetRenderedDescription(fmt.Sprint(a))
}

// GetStatus статус
func (vc *VerifiedIRPBiZoneCase) GetStatus() string {
	return vc.Status
}

// SetStatus статус
func (vc *VerifiedIRPBiZoneCase) SetStatus(v string) {
	vc.Status = v
}

// SetAnyStatus статус
func (vc *VerifiedIRPBiZoneCase) SetAnyStatus(a any) {
	vc.SetStatus(fmt.Sprint(a))
}

// GetStatusDescription описание статуса
func (vc *VerifiedIRPBiZoneCase) GetStatusDescription() string {
	return vc.StatusDescription
}

// SetStatusDescription описание статуса
func (vc *VerifiedIRPBiZoneCase) SetStatusDescription(v string) {
	vc.StatusDescription = v
}

// SetAnyStatusDescription описание статуса
func (vc *VerifiedIRPBiZoneCase) SetAnyStatusDescription(a any) {
	vc.SetStatusDescription(fmt.Sprint(a))
}

// GetReporterDisplayName нименование докладчика
func (vc *VerifiedIRPBiZoneCase) GetReporterDisplayName() string {
	return vc.ReporterDisplayName
}

// SetReporterDisplayName нименование докладчика
func (vc *VerifiedIRPBiZoneCase) SetReporterDisplayName(v string) {
	vc.ReporterDisplayName = v
}

// SetAnyReporterDisplayName нименование докладчика
func (vc *VerifiedIRPBiZoneCase) SetAnyReporterDisplayName(a any) {
	vc.SetReporterDisplayName(fmt.Sprint(a))
}

// GetReporterEmailAddress email адрес докладчика
func (vc *VerifiedIRPBiZoneCase) GetReporterEmailAddress() string {
	return vc.ReporterEmailAddress
}

// SetReporterEmailAddress email адрес докладчика
func (vc *VerifiedIRPBiZoneCase) SetReporterEmailAddress(v string) {
	vc.ReporterEmailAddress = v
}

// SetAnyReporterEmailAddress email адрес докладчика
func (vc *VerifiedIRPBiZoneCase) SetAnyReporterEmailAddress(a any) {
	vc.SetReporterEmailAddress(fmt.Sprint(a))
}

// GetPrimaryCategory первичная категория
func (vc *VerifiedIRPBiZoneCase) GetPrimaryCategory() string {
	return vc.PrimaryCategory
}

// SetPrimaryCategory первичная категория
func (vc *VerifiedIRPBiZoneCase) SetPrimaryCategory(v string) {
	vc.PrimaryCategory = v
}

// SetAnyPrimaryCategory первичная категория
func (vc *VerifiedIRPBiZoneCase) SetAnyPrimaryCategory(a any) {
	vc.SetPrimaryCategory(fmt.Sprint(a))
}

// GetUpdated время обновления (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetUpdated() string {
	return vc.Updated
}

// SetUpdated время обновления (преобразует в формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetUpdated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	vc.Updated = timeStr

	return nil
}

// SetAnyUpdated время обновления
func (vc *VerifiedIRPBiZoneCase) SetAnyUpdated(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetUpdated(v)
	}

	return errors.New("type conversion error")
}

// GetTimestamp временная метка (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetTimestamp() string {
	return vc.Timestamp
}

// SetTimestamp временная метка (преобразует в формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetTimestamp(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	vc.Timestamp = timeStr

	return nil
}

// SetAnyTimestamp временная метка
func (vc *VerifiedIRPBiZoneCase) SetAnyTimestamp(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetTimestamp(v)
	}

	return errors.New("type conversion error")
}

// GetResolutionDetailed подробная информация о разрешении
func (vc *VerifiedIRPBiZoneCase) GetResolutionDetailed() string {
	return vc.ResolutionDetailed
}

// SetResolutionDetailed подробная информация о разрешении
func (vc *VerifiedIRPBiZoneCase) SetResolutionDetailed(v string) {
	vc.ResolutionDetailed = v
}

// SetAnyResolutionDetailed подробная информация о разрешении
func (vc *VerifiedIRPBiZoneCase) SetAnyResolutionDetailed(a any) {
	vc.SetResolutionDetailed(fmt.Sprint(a))
}

// GetPlatformURL
func (vc *VerifiedIRPBiZoneCase) GetPlatformURL() string {
	return vc.PlatformURL
}

// SetPlatformURL
func (vc *VerifiedIRPBiZoneCase) SetPlatformURL(v string) {
	vc.PlatformURL = v
}

// SetAnyPlatformURL
func (vc *VerifiedIRPBiZoneCase) SetAnyPlatformURL(a any) {
	vc.SetPlatformURL(fmt.Sprint(a))
}

// GetCustomerStarRatingComment комментарий к звездному рейтингу клиента
func (vc *VerifiedIRPBiZoneCase) GetCustomerStarRatingComment() string {
	return vc.CustomerStarRatingComment
}

// SetCustomerStarRatingComment комментарий к звездному рейтингу клиента
func (vc *VerifiedIRPBiZoneCase) SetCustomerStarRatingComment(v string) {
	vc.CustomerStarRatingComment = v
}

// SetAnyCustomerStarRatingComment комментарий к звездному рейтингу клиента
func (vc *VerifiedIRPBiZoneCase) SetAnyCustomerStarRatingComment(a any) {
	vc.SetCustomerStarRatingComment(fmt.Sprint(a))
}

// GetRecommendations рекомендации
func (vc *VerifiedIRPBiZoneCase) GetRecommendations() string {
	return vc.Recommendations
}

// SetRecommendations рекомендации
func (vc *VerifiedIRPBiZoneCase) SetRecommendations(v string) {
	vc.Recommendations = v
}

// SetAnyRecommendations рекомендации
func (vc *VerifiedIRPBiZoneCase) SetAnyRecommendations(a any) {
	vc.SetRecommendations(fmt.Sprint(a))
}

// GetParsedActivityDetected обнаружение активности
func (vc *VerifiedIRPBiZoneCase) GetParsedActivityDetected() string {
	return vc.ParsedActivityDetected
}

// SetParsedActivityDetected обнаружение активности
func (vc *VerifiedIRPBiZoneCase) SetParsedActivityDetected(v string) {
	vc.ParsedActivityDetected = v
}

// SetAnyParsedActivityDetected обнаружение активности
func (vc *VerifiedIRPBiZoneCase) SetAnyParsedActivityDetected(a any) {
	vc.SetParsedActivityDetected(fmt.Sprint(a))
}

// GetAffectedService затронутый сервис
func (vc *VerifiedIRPBiZoneCase) GetAffectedService() string {
	return vc.AffectedService
}

// SetAffectedService затронутый сервис
func (vc *VerifiedIRPBiZoneCase) SetAffectedService(v string) {
	vc.AffectedService = v
}

// SetAnyAffectedService затронутый сервис
func (vc *VerifiedIRPBiZoneCase) SetAnyAffectedService(a any) {
	vc.SetAffectedService(fmt.Sprint(a))
}

// GetFakeAsPath подделка как путь
func (vc *VerifiedIRPBiZoneCase) GetFakeAsPath() string {
	return vc.FakeAsPath
}

// SetFakeAsPath подделка как путь
func (vc *VerifiedIRPBiZoneCase) SetFakeAsPath(v string) {
	vc.FakeAsPath = v
}

// SetAnyFakeAsPath подделка как путь
func (vc *VerifiedIRPBiZoneCase) SetAnyFakeAsPath(a any) {
	vc.SetFakeAsPath(fmt.Sprint(a))
}

// GetFakePrefix подделка как префикс
func (vc *VerifiedIRPBiZoneCase) GetFakePrefix() string {
	return vc.FakePrefix
}

// SetFakePrefix подделка как префикс
func (vc *VerifiedIRPBiZoneCase) SetFakePrefix(v string) {
	vc.FakePrefix = v
}

// SetAnyFakePrefix подделка как префикс
func (vc *VerifiedIRPBiZoneCase) SetAnyFakePrefix(a any) {
	vc.SetFakePrefix(fmt.Sprint(a))
}

// GetFpType fp тип
func (vc *VerifiedIRPBiZoneCase) GetFpType() string {
	return vc.FpType
}

// SetFpType fp тип
func (vc *VerifiedIRPBiZoneCase) SetFpType(v string) {
	vc.FpType = v
}

// SetAnyFpType fp тип
func (vc *VerifiedIRPBiZoneCase) SetAnyFpType(a any) {
	vc.SetFpType(fmt.Sprint(a))
}

// GetLookingGlass зазеркалье
func (vc *VerifiedIRPBiZoneCase) GetLookingGlass() string {
	return vc.LookingGlass
}

// SetLookingGlass зазеркалье
func (vc *VerifiedIRPBiZoneCase) SetLookingGlass(v string) {
	vc.LookingGlass = v
}

// SetAnyLookingGlass зазеркалье
func (vc *VerifiedIRPBiZoneCase) SetAnyLookingGlass(a any) {
	vc.SetLookingGlass(fmt.Sprint(a))
}

// GetSpamRecipients получатели спама
func (vc *VerifiedIRPBiZoneCase) GetSpamRecipients() string {
	return vc.SpamRecipients
}

// SetSpamRecipients получатели спама
func (vc *VerifiedIRPBiZoneCase) SetSpamRecipients(v string) {
	vc.SpamRecipients = v
}

// SetAnySpamRecipients получатели спама
func (vc *VerifiedIRPBiZoneCase) SetAnySpamRecipients(a any) {
	vc.SetSpamRecipients(fmt.Sprint(a))
}

// GetTLP
func (vc *VerifiedIRPBiZoneCase) GetTLP() string {
	return vc.TLP
}

// SetTLP
func (vc *VerifiedIRPBiZoneCase) SetTLP(v string) {
	vc.TLP = v
}

// SetAnyTLP
func (vc *VerifiedIRPBiZoneCase) SetAnyTLP(a any) {
	vc.SetTLP(fmt.Sprint(a))
}

// GetUsualPrefix обычный префикс
func (vc *VerifiedIRPBiZoneCase) GetUsualPrefix() string {
	return vc.UsualPrefix
}

// SetUsualPrefix обычный префикс
func (vc *VerifiedIRPBiZoneCase) SetUsualPrefix(v string) {
	vc.UsualPrefix = v
}

// SetAnyUsualPrefix обычный префикс
func (vc *VerifiedIRPBiZoneCase) SetAnyUsualPrefix(a any) {
	vc.SetUsualPrefix(fmt.Sprint(a))
}

// GetUsualAsPath обычный путь
func (vc *VerifiedIRPBiZoneCase) GetUsualAsPath() string {
	return vc.UsualAsPath
}

// SetUsualAsPath обычный путь
func (vc *VerifiedIRPBiZoneCase) SetUsualAsPath(v string) {
	vc.UsualAsPath = v
}

// SetAnyUsualAsPath обычный путь
func (vc *VerifiedIRPBiZoneCase) SetAnyUsualAsPath(a any) {
	vc.SetUsualAsPath(fmt.Sprint(a))
}

// GetSource источник
func (vc *VerifiedIRPBiZoneCase) GetSource() string {
	return vc.Source
}

// SetSource источник
func (vc *VerifiedIRPBiZoneCase) SetSource(v string) {
	vc.Source = v
}

// SetAnySource источник
func (vc *VerifiedIRPBiZoneCase) SetAnySource(a any) {
	vc.SetSource(fmt.Sprint(a))
}

// GetExternalId внешний идентификатор
func (vc *VerifiedIRPBiZoneCase) GetExternalId() string {
	return vc.ExternalId
}

// SetExternalId внешний идентификатор
func (vc *VerifiedIRPBiZoneCase) SetExternalId(v string) {
	vc.ExternalId = v
}

// SetAnyExternalId внешний идентификатор
func (vc *VerifiedIRPBiZoneCase) SetAnyExternalId(a any) {
	vc.SetExternalId(fmt.Sprint(a))
}

// GetUnderliningSource источник
func (vc *VerifiedIRPBiZoneCase) GetUnderliningSource() string {
	return vc.UnderliningSource
}

// SetUnderliningSource источник
func (vc *VerifiedIRPBiZoneCase) SetUnderliningSource(v string) {
	vc.UnderliningSource = v
}

// SetAnyUnderliningSource источник
func (vc *VerifiedIRPBiZoneCase) SetAnyUnderliningSource(a any) {
	vc.SetUnderliningSource(fmt.Sprint(a))
}

// GetUpdatedAll время обновления всего (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetUpdatedAll() string {
	return vc.UpdatedAll
}

// SetUpdatedAll время обновления всего (преобразует в формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetUpdatedAll(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	vc.UpdatedAll = timeStr

	return nil
}

// SetUpdatedAll время обновления всего
func (vc *VerifiedIRPBiZoneCase) SetAnyUpdatedAll(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetUpdatedAll(v)
	}

	return errors.New("type conversion error")
}

// GetGossopkaKey ключ ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) GetGossopkaKey() string {
	return vc.GossopkaKey
}

// SetGossopkaKey ключ ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) SetGossopkaKey(v string) {
	vc.GossopkaKey = v
}

// GetGossopkaKey ключ ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) SetAnyGossopkaKey(a any) {
	vc.SetGossopkaKey(fmt.Sprint(a))
}

// GetID идентификатор
func (vc *VerifiedIRPBiZoneCase) GetID() uint64 {
	return vc.ID
}

// SetID идентификатор
func (vc *VerifiedIRPBiZoneCase) SetID(v uint64) {
	vc.ID = v
}

// SetAnyID идентификатор
func (vc *VerifiedIRPBiZoneCase) SetAnyID(a any) {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return
	}

	vc.SetID(v)
}

// GetSystem идентификатор системы
func (vc *VerifiedIRPBiZoneCase) GetSystem() uint64 {
	return vc.System
}

// SetSystem идентификатор системы
func (vc *VerifiedIRPBiZoneCase) SetSystem(v uint64) {
	vc.System = v
}

// SetAnySystem идентификатор системы
func (vc *VerifiedIRPBiZoneCase) SetAnySystem(a any) {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return
	}

	vc.SetSystem(v)
}

// GetCancelGossopkaSendButton кнопка отмены ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) GetCancelGossopkaSendButton() bool {
	return vc.CancelGossopkaSendButton
}

// SetCancelGossopkaSendButton кнопка отмены ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) SetCancelGossopkaSendButton(v bool) {
	vc.CancelGossopkaSendButton = v
}

// SetAnyCancelGossopkaSendButton кнопка отмены ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) SetAnyCancelGossopkaSendButton(a any) {
	if v, ok := a.(bool); ok {
		vc.SetCancelGossopkaSendButton(v)
	}
}

// GetIsVisibleForCustomer видимость для потребителя
func (vc *VerifiedIRPBiZoneCase) GetIsVisibleForCustomer() bool {
	return vc.IsVisibleForCustomer
}

// SetIsVisibleForCustomer видимость для потребителя
func (vc *VerifiedIRPBiZoneCase) SetIsVisibleForCustomer(v bool) {
	vc.IsVisibleForCustomer = v
}

// SetAnyIsVisibleForCustomer видимость для потребителя
func (vc *VerifiedIRPBiZoneCase) SetAnyIsVisibleForCustomer(a any) {
	if v, ok := a.(bool); ok {
		vc.SetIsVisibleForCustomer(v)
	}
}

// GetShowGossopkaButton показать кнопку ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) GetShowGossopkaButton() bool {
	return vc.ShowGossopkaButton
}

// SetShowGossopkaButton показать кнопку ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) SetShowGossopkaButton(v bool) {
	vc.ShowGossopkaButton = v
}

// SetAnyShowGossopkaButton показать кнопку ГосСОПКА
func (vc *VerifiedIRPBiZoneCase) SetAnyShowGossopkaButton(a any) {
	if v, ok := a.(bool); ok {
		vc.SetShowGossopkaButton(v)
	}
}

// GetShowGtiButton показать Gti кнопку
func (vc *VerifiedIRPBiZoneCase) GetShowGtiButton() bool {
	return vc.ShowGtiButton
}

// SetShowGtiButton показать Gti кнопку
func (vc *VerifiedIRPBiZoneCase) SetShowGtiButton(v bool) {
	vc.ShowGtiButton = v
}

// SetAnyShowGtiButton показать Gti кнопку
func (vc *VerifiedIRPBiZoneCase) SetAnyShowGtiButton(a any) {
	if v, ok := a.(bool); ok {
		vc.SetShowGtiButton(v)
	}
}

// ToStringBeautiful форматированный вывод
func (vc *VerifiedIRPBiZoneCase) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'_source': '%s'\n", ws, vc.UnderliningSource)
	fmt.Fprintf(&str, "%s'id': '%d'\n", ws, vc.ID)
	fmt.Fprintf(&str, "%s'system': '%d'\n", ws, vc.System)
	fmt.Fprintf(&str, "%s'attack_type': '%s'\n", ws, vc.AttackType)
	fmt.Fprintf(&str, "%s'created': '%s'\n", ws, vc.Created)
	fmt.Fprintf(&str, "%s'creator_displayName': '%s'\n", ws, vc.CreatorDisplayName)
	fmt.Fprintf(&str, "%s'detection_date': '%s'\n", ws, vc.DetectionDate)
	fmt.Fprintf(&str, "%s'key': '%s'\n", ws, vc.Key)
	fmt.Fprintf(&str, "%s'issue_type': '%s'\n", ws, vc.IssueType)
	fmt.Fprintf(&str, "%s'priority': '%s'\n", ws, vc.Priority)
	fmt.Fprintf(&str, "%s'summary': '%s'\n", ws, vc.Summary)
	fmt.Fprintf(&str, "%s'description': '%s'\n", ws, vc.Description)
	fmt.Fprintf(&str, "%s'issue_type': '%s'\n", ws, vc.IssueType)
	fmt.Fprintf(&str, "%s'rendered_description': '%s'\n", ws, vc.RenderedDescription)
	fmt.Fprintf(&str, "%s'status': '%s'\n", ws, vc.Status)
	fmt.Fprintf(&str, "%s'status_description': '%s'\n", ws, vc.StatusDescription)
	fmt.Fprintf(&str, "%s'reporter_displayName': '%s'\n", ws, vc.ReporterDisplayName)
	fmt.Fprintf(&str, "%s'reporter_emailAddress': '%s'\n", ws, vc.ReporterEmailAddress)
	fmt.Fprintf(&str, "%s'primary_category': '%s'\n", ws, vc.PrimaryCategory)
	fmt.Fprintf(&str, "%s'updated': '%s'\n", ws, vc.Updated)
	fmt.Fprintf(&str, "%s'timestamp': '%s'\n", ws, vc.Timestamp)
	fmt.Fprintf(&str, "%s'resolution_detailed': '%s'\n", ws, vc.ResolutionDetailed)
	fmt.Fprintf(&str, "%s'platform_url': '%s'\n", ws, vc.PlatformURL)
	fmt.Fprintf(&str, "%s'customer_star_rating_comment': '%s'\n", ws, vc.CustomerStarRatingComment)
	fmt.Fprintf(&str, "%s'recommendations': '%s'\n", ws, vc.Recommendations)
	fmt.Fprintf(&str, "%s'parsed_activity_detected': '%s'\n", ws, vc.ParsedActivityDetected)
	fmt.Fprintf(&str, "%s'affected_service': '%s'\n", ws, vc.AffectedService)
	fmt.Fprintf(&str, "%s'fake_as_path': '%s'\n", ws, vc.FakeAsPath)
	fmt.Fprintf(&str, "%s'fake_prefix': '%s'\n", ws, vc.FakePrefix)
	fmt.Fprintf(&str, "%s'fp_type': '%s'\n", ws, vc.FpType)
	fmt.Fprintf(&str, "%s'looking_glass': '%s'\n", ws, vc.LookingGlass)
	fmt.Fprintf(&str, "%s'spam_recipients': '%s'\n", ws, vc.SpamRecipients)
	fmt.Fprintf(&str, "%s'tlp': '%s'\n", ws, vc.TLP)
	fmt.Fprintf(&str, "%s'usual_prefix': '%s'\n", ws, vc.UsualPrefix)
	fmt.Fprintf(&str, "%s'usual_as_path': '%s'\n", ws, vc.UsualAsPath)
	fmt.Fprintf(&str, "%s'source': '%s'\n", ws, vc.Source)
	fmt.Fprintf(&str, "%s'external_id': '%s'\n", ws, vc.ExternalId)
	fmt.Fprintf(&str, "%s'updated_all': '%s'\n", ws, vc.UpdatedAll)
	fmt.Fprintf(&str, "%s'gossopka_key': '%s'\n", ws, vc.GossopkaKey)
	fmt.Fprintf(&str, "%s'gossopka_errors': \n%s", ws, func(list map[string]any) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range list {
			fmt.Fprintf(&str, "%s'%s':'%+v'\n", supportingfunctions.GetWhitespace(num+1), k, v)
		}

		return str.String()
	}(vc.GossopkaErrors))
	fmt.Fprintf(&str, "%s'observed_iocs': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.ObservedIocs))
	fmt.Fprintf(&str, "%s'secondary_category_ref': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.SecondaryCategoryRef))
	fmt.Fprintf(&str, "%s'watchers': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.Watchers))
	fmt.Fprintf(&str, "%s'mitre_cov': \n%s", ws, vc.MITRECOV.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'detection_rules': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.DetectionRules))
	fmt.Fprintf(&str, "%s'secondary_category': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.SecondaryCategory))
	fmt.Fprintf(&str, "%s'platform_hostname': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.PlatformHostname))
	fmt.Fprintf(&str, "%s'watchers_id': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.WatchersId))
	fmt.Fprintf(&str, "%s'activity_detected,omitzero': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.ActivityDetected))
	fmt.Fprintf(&str, "%s'attachments': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.Attachments))
	fmt.Fprintf(&str, "%s'badges': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.Badges))
	fmt.Fprintf(&str, "%s'emls': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.Emls))
	fmt.Fprintf(&str, "%s'slas': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.Slas))
	fmt.Fprintf(&str, "%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.Tags))
	fmt.Fprintf(&str, "%s'keyfield': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, vc.KeyField))
	fmt.Fprintf(&str, "%s'issue_type_ref': \n%s", ws, vc.IssueTypeRef.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'primary_category_ref': \n%s", ws, vc.PrimaryCategoryRef.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'assignee': \n%v", ws, vc.Assignee)
	fmt.Fprintf(&str, "%s'assignee_displayName': \n%v", ws, vc.AssigneeDisplayName)
	fmt.Fprintf(&str, "%s'service': \n%v", ws, vc.Service)
	fmt.Fprintf(&str, "%s'resolutiondate': \n%v", ws, vc.ResolutionDate)
	fmt.Fprintf(&str, "%s'resolution_name': \n%v", ws, vc.ResolutionName)
	fmt.Fprintf(&str, "%s'resolution_name_ref': \n%v", ws, vc.ResolutionNameRef)
	fmt.Fprintf(&str, "%s'gossopka_id': \n%v", ws, vc.GossopkaId)
	fmt.Fprintf(&str, "%s'gossopka_send_time': \n%v", ws, vc.GossopkaSendTime)
	fmt.Fprintf(&str, "%s'gti_id': \n%v", ws, vc.GtiId)
	fmt.Fprintf(&str, "%s'gti_send_time': \n%v", ws, vc.GtiSendTime)
	fmt.Fprintf(&str, "%s'customer_star_rating': \n%v", ws, vc.CustomerStarRating)
	fmt.Fprintf(&str, "%s'first_sent_notification_time': \n%v", ws, vc.FirstSentNotificationTime)
	fmt.Fprintf(&str, "%s'response_team': \n%v", ws, vc.ResponseTeam)
	fmt.Fprintf(&str, "%s'gossopka_key_link': \n%v", ws, vc.GossopkaKeyLink)
	fmt.Fprintf(&str, "%s'cancel_gossopka_send_button': \n%v", ws, vc.CancelGossopkaSendButton)
	fmt.Fprintf(&str, "%s'is_visible_for_customer': \n%v", ws, vc.IsVisibleForCustomer)
	fmt.Fprintf(&str, "%s'show_gossopka_button': \n%v", ws, vc.ShowGossopkaButton)
	fmt.Fprintf(&str, "%s'show_gti_button': \n%v", ws, vc.ShowGtiButton)

	return str.String()
}
