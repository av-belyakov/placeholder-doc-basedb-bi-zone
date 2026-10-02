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

// GetActivityDetected для поля 'activity_detected'
func (d *BiZoneIRPCaseData) GetActivityDetected() []any {
	return d.ActivityDetected
}

// SetActivityDetected для поля 'activity_detected'
func (d *BiZoneIRPCaseData) SetActivityDetected(v []any) error {
	d.ActivityDetected = v

	return nil
}

// SetActivityDetected добавляет значение 'activity_detected' в список
func (d *BiZoneIRPCaseData) SetActivityDetectedElement(v string) error {
	//if _, isExist := supportingfunctions.SliceContainsElement(v, d.ActivityDetected); !isExist {
	d.ActivityDetected = append(d.ActivityDetected, v)
	//}

	return nil
}

// SetAnyActivityDetected добавляет некоторое значение в список 'activity_detected'
func (d *BiZoneIRPCaseData) SetAnyActivityDetected(a any) error {
	return d.SetActivityDetectedElement(fmt.Sprint(a))
}

// GetPlatformHostname для поля 'platform_hostname'
func (d *BiZoneIRPCaseData) GetPlatformHostname() []string {
	return d.PlatformHostname
}

// SetPlatformHostname для поля 'platform_hostname'
func (d *BiZoneIRPCaseData) SetPlatformHostname(v []string) error {
	d.PlatformHostname = v

	return nil
}

// SetPlatformHostname добавляет значение 'platform_hostname' в список
func (d *BiZoneIRPCaseData) SetPlatformHostnameElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.PlatformHostname); !isExist {
		d.PlatformHostname = append(d.PlatformHostname, v)
	}

	return nil
}

// SetAnyPlatformHostname добавляет некоторое значение в список 'detection_rules'
func (d *BiZoneIRPCaseData) SetAnyPlatformHostname(a any) error {
	return d.SetPlatformHostnameElement(fmt.Sprint(a))
}

// GetDetectionRules для поля 'detection_rules'
func (d *BiZoneIRPCaseData) GetDetectionRules() []uint64 {
	return d.DetectionRules
}

// SetDetectionRules для поля 'detection_rules'
func (d *BiZoneIRPCaseData) SetDetectionRules(v []uint64) error {
	d.DetectionRules = v

	return nil
}

// SetDetectionRules добавляет значение 'detection_rules' в список
func (d *BiZoneIRPCaseData) SetDetectionRulesElement(v uint64) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.DetectionRules); !isExist {
		d.DetectionRules = append(d.DetectionRules, v)
	}

	return nil
}

// SetAnyDetectionRules добавляет некоторое значение в список 'detection_rules'
func (d *BiZoneIRPCaseData) SetAnyDetectionRules(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetDetectionRulesElement(v)
}

// GetType для поля 'type'
func (d *BiZoneIRPCaseData) GetType() *BiZoneIRPType {
	return &d.Type
}

// SetType для поля 'type'
func (d *BiZoneIRPCaseData) SetType(v BiZoneIRPType) error {
	d.Type = v

	return nil
}

// GetPriority для поля 'priority'
func (d *BiZoneIRPCaseData) GetPriority() *BiZoneIRPPriority {
	return &d.Priority
}

// SetPriority для поля 'priority'
func (d *BiZoneIRPCaseData) SetPriority(v BiZoneIRPPriority) error {
	d.Priority = v

	return nil
}

// GetStatus для поля 'status'
func (d *BiZoneIRPCaseData) GetStatus() *BiZoneIRPStatus {
	return &d.Status
}

// SetStatus для поля 'status'
func (d *BiZoneIRPCaseData) SetStatus(v BiZoneIRPStatus) error {
	d.Status = v

	return nil
}

// GetPrimaryCategory для поля 'primary_category'
func (d *BiZoneIRPCaseData) GetPrimaryCategory() *BiZoneIRPCategory {
	return &d.PrimaryCategory
}

// SetPrimaryCategory для поля 'primary_category'
func (d *BiZoneIRPCaseData) SetPrimaryCategory(v BiZoneIRPCategory) error {
	d.PrimaryCategory = v

	return nil
}

// GetTenant для поля 'tenant'
func (d *BiZoneIRPCaseData) GetTenant() *BiZoneIRPTenant {
	return &d.Tenant
}

// SetTenant для поля 'tenant'
func (d *BiZoneIRPCaseData) SetTenant(v BiZoneIRPTenant) error {
	d.Tenant = v

	return nil
}

// GetCreatedBy для поля 'created_by'
func (d *BiZoneIRPCaseData) GetCreatedBy() *BiZoneIRPCreatedBy {
	return &d.CreatedBy
}

// SetCreatedBy для поля 'created_by'
func (d *BiZoneIRPCaseData) SetCreatedBy(v BiZoneIRPCreatedBy) error {
	d.CreatedBy = v

	return nil
}

// ToStringBeautiful форматированный вывод
func (d *BiZoneIRPCaseData) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	//wsStr := supportingfunctions.GetWhitespace(num + 1)
	wsInt := supportingfunctions.GetWhitespace(num + 2)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, d.ID)
	fmt.Fprintf(&str, "%s'external_id': '%s'\n", ws, d.ExternalID)
	fmt.Fprintf(&str, "%s'summary': '%s'\n", ws, d.Summary)
	fmt.Fprintf(&str, "%s'description': '%s'\n", ws, d.Description)
	fmt.Fprintf(&str, "%s'recommendations': '%s'\n", ws, d.Recommendations)
	fmt.Fprintf(&str, "%s'created': '%s'\n", ws, d.Created)
	fmt.Fprintf(&str, "%s'updated': '%s'\n", ws, d.Updated)
	fmt.Fprintf(&str, "%s'detection_date': '%s'\n", ws, d.DetectionDate)
	fmt.Fprintf(&str, "%s'resolution_date': '%s'\n", ws, d.ResolutionDate)
	fmt.Fprintf(&str, "%s'resolution_detailed': '%s'\n", ws, d.ResolutionDetailed)
	fmt.Fprintf(&str, "%s'customer_star_rating_comment': '%s'\n", ws, d.CustomerStarRatingComment)
	fmt.Fprintf(&str, "%s'status_description': '%s'\n", ws, d.StatusDescription)
	fmt.Fprintf(&str, "%s'fp_type': '%s'\n", ws, d.FpType)
	fmt.Fprintf(&str, "%s'tlp': '%s'\n", ws, d.TLP)
	fmt.Fprintf(&str, "%s'customer_assignee': '%s'\n", ws, *d.CustomerAssignee)
	fmt.Fprintf(&str, "%s'assignee': '%s'\n", ws, *d.Assignee)
	fmt.Fprintf(&str, "%s'mitre_cov': '%s'\n", ws, *d.MitreCov)
	fmt.Fprintf(&str, "%s'resolution': '%s'\n", ws, *d.Resolution)
	fmt.Fprintf(&str, "%s'response_team': '%s'\n", ws, *d.ResponseTeam)
	fmt.Fprintf(&str, "%s'customer_star_rating': '%d'\n", ws, *d.CustomerStarRating)
	fmt.Fprintf(&str, "%s'is_public': '%t'\n", ws, d.IsPublic)
	fmt.Fprintf(&str, "%s'activity_detected':\n", ws)
	for k, v := range d.ActivityDetected {
		fmt.Fprintf(&str, "%s%d.\n%s%s%+v\n", wsInt, k, wsInt, wsInt, v)
	}
	fmt.Fprintf(&str, "%s'tags':\n", ws)
	for k, v := range d.Tags {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInt, k, v.ToStringBeautiful(num+2))
	}
	fmt.Fprintf(&str, "%s'secondary_category':\n", ws)
	for k, v := range d.SecondaryCategory {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInt, k, v.ToStringBeautiful(num+2))
	}
	fmt.Fprintf(&str, "%s'platform_hostname': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.PlatformHostname))
	fmt.Fprintf(&str, "%s'detection_rules': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.DetectionRules))
	fmt.Fprintf(&str, "%s'type':\n%s", ws, d.Type.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'priority':\n%s", ws, d.Priority.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'status':\n%s", ws, d.Status.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'primary_category':\n%s", ws, d.PrimaryCategory.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'tenant':\n%s", ws, d.Tenant.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'created_by':\n%s", ws, d.CreatedBy.ToStringBeautiful(num+1))

	return str.String()
}
