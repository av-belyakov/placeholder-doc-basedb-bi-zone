package datamodels

import (
	"errors"
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPCaseFieldData() *BiZoneIRPCaseFieldData {
	return &BiZoneIRPCaseFieldData{
		ActivityDetected:  make([]any, 0),
		DetectionRules:    make([]string, 0),
		PlatformHostname:  make([]string, 0),
		SecondaryCategory: make([]BiZoneIRPFieldWithTitle, 0),
		Tags:              make([]BiZoneIRPFieldTagDescription, 0),
	}
}

// GetTags список тегов
func (cfd *BiZoneIRPCaseFieldData) GetTags() []BiZoneIRPFieldTagDescription {
	if cfd.Tags == nil {
		cfd.Tags = make([]BiZoneIRPFieldTagDescription, 0)
	}

	return cfd.Tags
}

// SetTags список тегов
func (cfd *BiZoneIRPCaseFieldData) SetTags(list []BiZoneIRPFieldTagDescription) {
	cfd.Tags = list
}

// GetSecondaryCategory список категорий
func (cfd *BiZoneIRPCaseFieldData) GetSecondaryCategory() []BiZoneIRPFieldWithTitle {
	if cfd.SecondaryCategory == nil {
		cfd.SecondaryCategory = make([]BiZoneIRPFieldWithTitle, 0)
	}

	return cfd.SecondaryCategory
}

// SetSecondaryCategory список категорий
func (cfd *BiZoneIRPCaseFieldData) SetSecondaryCategory(list []BiZoneIRPFieldWithTitle) {
	cfd.SecondaryCategory = list
}

// GetDetectionRules список правил обнаружения
func (cfd *BiZoneIRPCaseFieldData) GetDetectionRules() []string {
	if cfd.DetectionRules == nil {
		cfd.DetectionRules = make([]string, 0)
	}

	return cfd.DetectionRules
}

// SetDetectionRules список правил обнаружения
func (cfd *BiZoneIRPCaseFieldData) SetDetectionRules(list []string) {
	cfd.DetectionRules = list
}

// AddDetectionRules список правил обнаружения
func (cfd *BiZoneIRPCaseFieldData) AddDetectionRules(a string) {
	cfd.GetDetectionRules()

	cfd.DetectionRules = append(cfd.DetectionRules, a)
}

// GetPlatformHostname список платформ
func (cfd *BiZoneIRPCaseFieldData) GetPlatformHostname() []string {
	if cfd.PlatformHostname == nil {
		cfd.PlatformHostname = make([]string, 0)
	}

	return cfd.PlatformHostname
}

// SetPlatformHostname список платформ
func (cfd *BiZoneIRPCaseFieldData) SetPlatformHostname(list []string) {
	cfd.PlatformHostname = list
}

// AddPlatformHostname список платформ
func (cfd *BiZoneIRPCaseFieldData) AddPlatformHostname(a string) {
	cfd.GetPlatformHostname()

	cfd.PlatformHostname = append(cfd.PlatformHostname, a)
}

// GetActivityDetected список обнаружения активности
func (cfd *BiZoneIRPCaseFieldData) GetActivityDetected() []any {
	if cfd.ActivityDetected == nil {
		cfd.ActivityDetected = make([]any, 0)
	}

	return cfd.ActivityDetected
}

// SetActivityDetected список обнаружения активности
func (cfd *BiZoneIRPCaseFieldData) SetActivityDetected(list []any) {
	cfd.ActivityDetected = list
}

// AddActivityDetected список обнаружения активности
func (cfd *BiZoneIRPCaseFieldData) AddActivityDetected(a any) {
	cfd.GetActivityDetected()

	cfd.ActivityDetected = append(cfd.ActivityDetected, a)
}

// GetType
func (cfd *BiZoneIRPCaseFieldData) GetType() BiZoneIRPFieldWithTitle {
	return cfd.Type
}

// SetType
func (cfd *BiZoneIRPCaseFieldData) SetType(v BiZoneIRPFieldWithTitle) {
	cfd.Type = v
}

// GetStatus
func (cfd *BiZoneIRPCaseFieldData) GetStatus() BiZoneIRPFieldWithTitle {
	return cfd.Status
}

// SetStatus
func (cfd *BiZoneIRPCaseFieldData) SetStatus(v BiZoneIRPFieldWithTitle) {
	cfd.Status = v
}

// GetPriority
func (cfd *BiZoneIRPCaseFieldData) GetPriority() BiZoneIRPFieldWithTitle {
	return cfd.Priority
}

// SetPriority
func (cfd *BiZoneIRPCaseFieldData) SetPriority(v BiZoneIRPFieldWithTitle) {
	cfd.Priority = v
}

// GetPrimaryCategory
func (cfd *BiZoneIRPCaseFieldData) GetPrimaryCategory() BiZoneIRPFieldWithTitle {
	return cfd.PrimaryCategory
}

// SetPrimaryCategory
func (cfd *BiZoneIRPCaseFieldData) SetPrimaryCategory(v BiZoneIRPFieldWithTitle) {
	cfd.PrimaryCategory = v
}

// GetMITRECOV
func (cfd *BiZoneIRPCaseFieldData) GetMITRECOV() BiZoneIRPMITRECOV {
	if cfd.MITRECOV.GetPersistence() == nil {
		cfd.MITRECOV = *NewBiZoneMITRECOV()
	}

	return cfd.MITRECOV
}

// SetMITRECOV
func (cfd *BiZoneIRPCaseFieldData) SetMITRECOV(v BiZoneIRPMITRECOV) {
	cfd.MITRECOV = v
}

// GetTenantId
func (cfd *BiZoneIRPCaseFieldData) GetTenantId() string {
	return cfd.Tenant.Id
}

// SetTenantId
func (cfd *BiZoneIRPCaseFieldData) SetTenantId(v string) {
	cfd.Tenant.Id = v
}

// SetAnyTenantId
func (cfd *BiZoneIRPCaseFieldData) SetAnyTenantId(a any) {
	cfd.SetTenantId(fmt.Sprint(a))
}

// GetTenantName
func (cfd *BiZoneIRPCaseFieldData) GetTenantName() string {
	return cfd.Tenant.Name
}

// SetTenantName
func (cfd *BiZoneIRPCaseFieldData) SetTenantName(v string) {
	cfd.Tenant.Name = v
}

// SetAnyTenantName
func (cfd *BiZoneIRPCaseFieldData) SetAnyTenantName(a any) {
	cfd.SetTenantName(fmt.Sprint(a))
}

// GetCreatedById
func (cfd *BiZoneIRPCaseFieldData) GetCreatedById() string {
	return cfd.CreatedBy.Id
}

// SetCreatedById
func (cfd *BiZoneIRPCaseFieldData) SetCreatedById(v string) {
	cfd.CreatedBy.Id = v
}

// SetAnyCreatedById
func (cfd *BiZoneIRPCaseFieldData) SetAnyCreatedBy(a any) {
	cfd.SetCreatedById(fmt.Sprint(a))
}

// GetCreatedByUsername
func (cfd *BiZoneIRPCaseFieldData) GetCreatedByUsername() string {
	return cfd.CreatedBy.Username
}

// SetCreatedByUsername
func (cfd *BiZoneIRPCaseFieldData) SetCreatedByUsername(v string) {
	cfd.CreatedBy.Username = v
}

// SetAnyCreatedByUsername
func (cfd *BiZoneIRPCaseFieldData) SetAnyCreatedByUsername(a any) {
	cfd.SetCreatedByUsername(fmt.Sprint(a))
}

// GetId
func (cfd *BiZoneIRPCaseFieldData) GetId() string {
	return cfd.Id
}

// SetId
func (cfd *BiZoneIRPCaseFieldData) SetId(v string) {
	cfd.Id = v
}

// SetAnyId
func (cfd *BiZoneIRPCaseFieldData) SetAnyId(a any) {
	cfd.SetId(fmt.Sprint(a))
}

// GetTlp
func (cfd *BiZoneIRPCaseFieldData) GetTlp() string {
	return cfd.Tlp
}

// SetTlp
func (cfd *BiZoneIRPCaseFieldData) SetTlp(v string) {
	cfd.Tlp = v
}

// SetAnyTlp
func (cfd *BiZoneIRPCaseFieldData) SetAnyTlp(a any) {
	cfd.SetTlp(fmt.Sprint(a))
}

// GetFpType
func (cfd *BiZoneIRPCaseFieldData) GetFpType() string {
	return cfd.FpType
}

// SetFpType
func (cfd *BiZoneIRPCaseFieldData) SetFpType(v string) {
	cfd.FpType = v
}

// SetAnyFpType
func (cfd *BiZoneIRPCaseFieldData) SetAnyFpType(a any) {
	cfd.SetFpType(fmt.Sprint(a))
}

// GetCreated время создания (формат времени RFC3339)
func (cfd *BiZoneIRPCaseFieldData) GetCreated() string {
	return cfd.Created
}

// SetCreated время создания (преобразует в формат времени RFC3339)
func (cfd *BiZoneIRPCaseFieldData) SetCreated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	cfd.Created = timeStr

	return nil
}

// SetAnyCreated время создания
func (cfd *BiZoneIRPCaseFieldData) SetAnyCreated(a any) error {
	if v, ok := a.(string); ok {
		return cfd.SetCreated(v)
	}

	return errors.New("type conversion error")
}

// GetUpdated время обновления (формат времени RFC3339)
func (cfd *BiZoneIRPCaseFieldData) GetUpdated() string {
	return cfd.Updated
}

// SetUpdated время обновления (преобразует в формат времени RFC3339)
func (cfd *BiZoneIRPCaseFieldData) SetUpdated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	cfd.Updated = timeStr

	return nil
}

// SetAnyUpdated время обновления
func (cfd *BiZoneIRPCaseFieldData) SetAnyUpdated(a any) error {
	if v, ok := a.(string); ok {
		return cfd.SetUpdated(v)
	}

	return errors.New("type conversion error")
}

// GetDetectionDate дата обнаружения (формат времени RFC3339)
func (cfd *BiZoneIRPCaseFieldData) GetDetectionDate() string {
	return cfd.DetectionDate
}

// SetDetectionDate дата обнаружения (преобразует в формат времени RFC3339)
func (cfd *BiZoneIRPCaseFieldData) SetDetectionDate(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	cfd.DetectionDate = timeStr

	return nil
}

// SetAnyDetectionDate дата обнаружения
func (cfd *BiZoneIRPCaseFieldData) SetAnyDetectionDate(a any) error {
	if v, ok := a.(string); ok {
		return cfd.SetDetectionDate(v)
	}

	return errors.New("type conversion error")
}

// GetSummary резюме
func (cfd *BiZoneIRPCaseFieldData) GetSummary() string {
	return cfd.Summary
}

// SetSummary резюме
func (cfd *BiZoneIRPCaseFieldData) SetSummary(v string) {
	cfd.Summary = v
}

// SetAnySummary резюме
func (cfd *BiZoneIRPCaseFieldData) SetAnySummary(a any) {
	cfd.SetSummary(fmt.Sprint(a))
}

// GetDescription описание
func (cfd *BiZoneIRPCaseFieldData) GetDescription() string {
	return cfd.Description
}

// SetDescription описание
func (cfd *BiZoneIRPCaseFieldData) SetDescription(v string) {
	cfd.Description = v
}

// SetAnyDescription описание
func (cfd *BiZoneIRPCaseFieldData) SetAnyDescription(a any) {
	cfd.SetDescription(fmt.Sprint(a))
}

// GetRecommendations рекомендации
func (cfd *BiZoneIRPCaseFieldData) GetRecommendations() string {
	return cfd.Recommendations
}

// SetRecommendations рекомендации
func (cfd *BiZoneIRPCaseFieldData) SetRecommendations(v string) {
	cfd.Recommendations = v
}

// SetAnyRecommendations рекомендации
func (cfd *BiZoneIRPCaseFieldData) SetAnyRecommendations(a any) {
	cfd.SetRecommendations(fmt.Sprint(a))
}

// GetExternalId внешний идентификатор
func (cfd *BiZoneIRPCaseFieldData) GetExternalId() string {
	return cfd.ExternalId
}

// SetExternalId внешний идентификатор
func (cfd *BiZoneIRPCaseFieldData) SetExternalId(v string) {
	cfd.ExternalId = v
}

// SetAnyExternalId внешний идентификатор
func (cfd *BiZoneIRPCaseFieldData) SetAnyExternalId(a any) {
	cfd.SetExternalId(fmt.Sprint(a))
}

// GetStatusDescription описание статуса
func (cfd *BiZoneIRPCaseFieldData) GetStatusDescription() string {
	return cfd.StatusDescription
}

// SetStatusDescription описание статуса
func (cfd *BiZoneIRPCaseFieldData) SetStatusDescription(v string) {
	cfd.StatusDescription = v
}

// SetAnyStatusDescription описание статуса
func (cfd *BiZoneIRPCaseFieldData) SetAnyStatusDescription(a any) {
	cfd.SetStatusDescription(fmt.Sprint(a))
}

// GetResolutionDetailed подробная информация о разрешении
func (cfd *BiZoneIRPCaseFieldData) GetResolutionDetailed() string {
	return cfd.ResolutionDetailed
}

// SetResolutionDetailed подробная информация о разрешении
func (cfd *BiZoneIRPCaseFieldData) SetResolutionDetailed(v string) {
	cfd.ResolutionDetailed = v
}

// SetAnyResolutionDetailed подробная информация о разрешении
func (cfd *BiZoneIRPCaseFieldData) SetAnyResolutionDetailed(a any) {
	cfd.SetResolutionDetailed(fmt.Sprint(a))
}

// GetCustomerStarRatingComment комментарий к звездному рейтингу клиента
func (cfd *BiZoneIRPCaseFieldData) GetCustomerStarRatingComment() string {
	return cfd.CustomerStarRatingComment
}

// SetCustomerStarRatingComment комментарий к звездному рейтингу клиента
func (cfd *BiZoneIRPCaseFieldData) SetCustomerStarRatingComment(v string) {
	cfd.CustomerStarRatingComment = v
}

// SetAnyCustomerStarRatingComment комментарий к звездному рейтингу клиента
func (cfd *BiZoneIRPCaseFieldData) SetAnyCustomerStarRatingComment(a any) {
	cfd.SetCustomerStarRatingComment(fmt.Sprint(a))
}

// GetAssignee правопреемник
func (cfd *BiZoneIRPCaseFieldData) GetAssignee() any {
	return cfd.Assignee
}

// SetAssignee правопреемник
func (cfd *BiZoneIRPCaseFieldData) SetAssignee(a any) {
	cfd.Assignee = a
}

// GetCustomerAssignee правопреемник клиента
func (cfd *BiZoneIRPCaseFieldData) GetCustomerAssignee() any {
	return cfd.CustomerAssignee
}

// SetCustomerAssignee правопреемник клиента
func (cfd *BiZoneIRPCaseFieldData) SetCustomerAssignee(a any) {
	cfd.CustomerAssignee = a
}

// GetCustomerStarRating звездный рейтинг клиента
func (cfd *BiZoneIRPCaseFieldData) GetCustomerStarRating() any {
	return cfd.CustomerStarRating
}

// SetCustomerStarRating звездный рейтинг клиента
func (cfd *BiZoneIRPCaseFieldData) SetCustomerStarRating(a any) {
	cfd.CustomerStarRating = a
}

// GetResolutionDate дата разрешения
func (cfd *BiZoneIRPCaseFieldData) GetResolutionDate() any {
	return cfd.ResolutionDate
}

// SetResolutionDate дата разрешения
func (cfd *BiZoneIRPCaseFieldData) SetResolutionDate(a any) {
	cfd.ResolutionDate = a
}

// GetResolution разрешение
func (cfd *BiZoneIRPCaseFieldData) GetResolution() any {
	return cfd.Resolution
}

// SetResolution разрешение
func (cfd *BiZoneIRPCaseFieldData) SetResolution(a any) {
	cfd.Resolution = a
}

// GetResponseTeam группа разрешения
func (cfd *BiZoneIRPCaseFieldData) GetResponseTeam() any {
	return cfd.ResponseTeam
}

// SetResponseTeam группа разрешения
func (cfd *BiZoneIRPCaseFieldData) SetResponseTeam(a any) {
	cfd.ResponseTeam = a
}

// GetIsPublic видимость
func (cfd *BiZoneIRPCaseFieldData) GetIsPublic() bool {
	return cfd.IsPublic
}

// SetIsPublic видимость
func (cfd *BiZoneIRPCaseFieldData) SetIsPublic(v bool) {
	cfd.IsPublic = v
}

// SetAnyIsPublic видимость
func (cfd *BiZoneIRPCaseFieldData) SetAnyIsPublic(a any) {
	if v, ok := a.(bool); ok {
		cfd.SetIsPublic(v)
	}
}

// ToStringBeautiful форматированный вывод
func (cfd *BiZoneIRPCaseFieldData) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%d'\n", ws, cfd.Id)
	fmt.Fprintf(&str, "%s'system': '%d'\n", ws, cfd.System)
	fmt.Fprintf(&str, "%s'attack_type': '%s'\n", ws, cfd.AttackType)
	fmt.Fprintf(&str, "%s'created': '%s'\n", ws, cfd.Created)
	fmt.Fprintf(&str, "%s'creator_displayName': '%s'\n", ws, cfd.CreatorDisplayName)
	fmt.Fprintf(&str, "%s'detection_date': '%s'\n", ws, cfd.DetectionDate)
	fmt.Fprintf(&str, "%s'key': '%s'\n", ws, cfd.Key)
	fmt.Fprintf(&str, "%s'issue_type': '%s'\n", ws, cfd.IssueType)
	fmt.Fprintf(&str, "%s'priority': '%s'\n", ws, cfd.Priority)
	fmt.Fprintf(&str, "%s'summary': '%s'\n", ws, cfd.Summary)
	fmt.Fprintf(&str, "%s'description': '%s'\n", ws, cfd.Description)
	fmt.Fprintf(&str, "%s'issue_type': '%s'\n", ws, cfd.IssueType)
	fmt.Fprintf(&str, "%s'rendered_description': '%s'\n", ws, cfd.RenderedDescription)
	fmt.Fprintf(&str, "%s'status': '%s'\n", ws, cfd.Status)
	fmt.Fprintf(&str, "%s'status_description': '%s'\n", ws, cfd.StatusDescription)
	fmt.Fprintf(&str, "%s'reporter_displayName': '%s'\n", ws, cfd.ReporterDisplayName)
	fmt.Fprintf(&str, "%s'reporter_emailAddress': '%s'\n", ws, cfd.ReporterEmailAddress)
	fmt.Fprintf(&str, "%s'primary_category': '%s'\n", ws, cfd.PrimaryCategory)
	fmt.Fprintf(&str, "%s'updated': '%s'\n", ws, cfd.Updated)
	fmt.Fprintf(&str, "%s'timestamp': '%s'\n", ws, cfd.Timestamp)
	fmt.Fprintf(&str, "%s'resolution_detailed': '%s'\n", ws, cfd.ResolutionDetailed)
	fmt.Fprintf(&str, "%s'platform_url': '%s'\n", ws, cfd.PlatformURL)
	fmt.Fprintf(&str, "%s'customer_star_rating_comment': '%s'\n", ws, cfd.CustomerStarRatingComment)
	fmt.Fprintf(&str, "%s'recommendations': '%s'\n", ws, cfd.Recommendations)
	fmt.Fprintf(&str, "%s'parsed_activity_detected': '%s'\n", ws, cfd.ParsedActivityDetected)
	fmt.Fprintf(&str, "%s'affected_service': '%s'\n", ws, cfd.AffectedService)
	fmt.Fprintf(&str, "%s'fake_as_path': '%s'\n", ws, cfd.FakeAsPath)
	fmt.Fprintf(&str, "%s'fake_prefix': '%s'\n", ws, cfd.FakePrefix)
	fmt.Fprintf(&str, "%s'fp_type': '%s'\n", ws, cfd.FpType)
	fmt.Fprintf(&str, "%s'looking_glass': '%s'\n", ws, cfd.LookingGlass)
	fmt.Fprintf(&str, "%s'spam_recipients': '%s'\n", ws, cfd.SpamRecipients)
	fmt.Fprintf(&str, "%s'tlp': '%s'\n", ws, cfd.TLP)
	fmt.Fprintf(&str, "%s'usual_prefix': '%s'\n", ws, cfd.UsualPrefix)
	fmt.Fprintf(&str, "%s'usual_as_path': '%s'\n", ws, cfd.UsualAsPath)
	fmt.Fprintf(&str, "%s'source': '%s'\n", ws, cfd.Source)
	fmt.Fprintf(&str, "%s'external_id': '%s'\n", ws, cfd.ExternalId)
	fmt.Fprintf(&str, "%s'updated_all': '%s'\n", ws, cfd.UpdatedAll)
	fmt.Fprintf(&str, "%s'gossopka_key': '%s'\n", ws, cfd.GossopkaKey)
	fmt.Fprintf(&str, "%s'gossopka_errors': \n%s", ws, func(list map[string]any) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range list {
			fmt.Fprintf(&str, "%s'%s':'%+v'\n", supportingfunctions.GetWhitespace(num+1), k, v)
		}

		return str.String()
	}(cfd.GossopkaErrors))
	fmt.Fprintf(&str, "%s'observed_iocs': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.ObservedIocs))
	fmt.Fprintf(&str, "%s'secondary_category_ref': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.SecondaryCategoryRef))
	fmt.Fprintf(&str, "%s'watchers': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.Watchers))
	fmt.Fprintf(&str, "%s'mitre_cov': \n%s", ws, cfd.MITRECOV.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'detection_rules': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.DetectionRules))
	fmt.Fprintf(&str, "%s'secondary_category': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.SecondaryCategory))
	fmt.Fprintf(&str, "%s'platform_hostname': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.PlatformHostname))
	fmt.Fprintf(&str, "%s'watchers_id': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.WatchersId))
	fmt.Fprintf(&str, "%s'activity_detected,omitzero': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.ActivityDetected))
	fmt.Fprintf(&str, "%s'attachments': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.Attachments))
	fmt.Fprintf(&str, "%s'badges': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.Badges))
	fmt.Fprintf(&str, "%s'emls': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.Emls))
	fmt.Fprintf(&str, "%s'slas': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.Slas))
	fmt.Fprintf(&str, "%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.Tags))
	fmt.Fprintf(&str, "%s'keyfield': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num+1, cfd.KeyField))
	fmt.Fprintf(&str, "%s'issue_type_ref': \n%s", ws, cfd.IssueTypeRef.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'primary_category_ref': \n%s", ws, cfd.PrimaryCategoryRef.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'assignee': \n%v", ws, cfd.Assignee)
	fmt.Fprintf(&str, "%s'assignee_displayName': \n%v", ws, cfd.AssigneeDisplayName)
	fmt.Fprintf(&str, "%s'service': \n%v", ws, cfd.Service)
	fmt.Fprintf(&str, "%s'resolutiondate': \n%v", ws, cfd.ResolutionDate)
	fmt.Fprintf(&str, "%s'resolution_name': \n%v", ws, cfd.ResolutionName)
	fmt.Fprintf(&str, "%s'resolution_name_ref': \n%v", ws, cfd.ResolutionNameRef)
	fmt.Fprintf(&str, "%s'gossopka_id': \n%v", ws, cfd.GossopkaId)
	fmt.Fprintf(&str, "%s'gossopka_send_time': \n%v", ws, cfd.GossopkaSendTime)
	fmt.Fprintf(&str, "%s'gti_id': \n%v", ws, cfd.GtiId)
	fmt.Fprintf(&str, "%s'gti_send_time': \n%v", ws, cfd.GtiSendTime)
	fmt.Fprintf(&str, "%s'customer_star_rating': \n%v", ws, cfd.CustomerStarRating)
	fmt.Fprintf(&str, "%s'first_sent_notification_time': \n%v", ws, cfd.FirstSentNotificationTime)
	fmt.Fprintf(&str, "%s'response_team': \n%v", ws, cfd.ResponseTeam)
	fmt.Fprintf(&str, "%s'gossopka_key_link': \n%v", ws, cfd.GossopkaKeyLink)
	fmt.Fprintf(&str, "%s'cancel_gossopka_send_button': \n%v", ws, cfd.CancelGossopkaSendButton)
	fmt.Fprintf(&str, "%s'is_visible_for_customer': \n%v", ws, cfd.IsVisibleForCustomer)
	fmt.Fprintf(&str, "%s'show_gossopka_button': \n%v", ws, cfd.ShowGossopkaButton)
	fmt.Fprintf(&str, "%s'show_gti_button': \n%v", ws, cfd.ShowGtiButton)

	return str.String()
}
