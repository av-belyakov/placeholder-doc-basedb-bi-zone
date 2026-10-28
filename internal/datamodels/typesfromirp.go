package datamodels

/***
* Структуры относящиеся к типу Incident
***/

// VerifiedBiZoneIRPIncident основной объект типа 'Incident'
type VerifiedBiZoneIRPIncident struct {
	AdditionalInformation
	Tags                      []IRPTagType       `json:"tags"`
	Watchers                  []WatchersType     `json:"watchers"`
	ObservedIocs              []ObservedIocsType `json:"observed_iocs"`
	SecondaryCategoryRef      []CategoryRefType  `json:"secondary_category_ref"`
	KeyField                  []string           `json:"keyfield"`
	ActivityDetected          []string           `json:"activity_detected"`
	PlatformHostname          []string           `json:"platform_hostname"`
	SecondaryCategory         []string           `json:"secondary_category"`
	DetectionRules            []uint64           `json:"detection_rules"`
	Emls                      []any              `json:"emls"`
	Slas                      []any              `json:"slas"`
	Badges                    []any              `json:"badges"`
	WatchersId                []any              `json:"watchers_id"`
	Attachments               []any              `json:"attachments"`
	ParsedActivityDetected    []any              `json:"parsed_activity_detected"`
	IssueTypeRef              IssueTypeRefType   `json:"issue_type_ref"`
	PrimaryCategoryRef        CategoryRefType    `json:"primary_category_ref"`
	ResolutionNameRef         CategoryRefType    `json:"resolution_name_ref"`
	Created                   string             `json:"created"`                      // дата создания (формат RFC3339) или ISO 8601
	Updated                   string             `json:"updated"`                      // дата обновления (формат RFC3339) или ISO 8601
	Timestamp                 string             `json:"timestamp"`                    // временная метка (формат RFC3339) или ISO 8601
	UpdatedAll                string             `json:"updated_all"`                  // дата обновления (формат RFC3339) или ISO 8601
	DetectionDate             string             `json:"detection_date"`               // дата обнаружения (формат RFC3339) или ISO 8601
	FirstSentNotificationTime string             `json:"first_sent_notification_time"` // первая дата отправки сообщения (формат RFC3339) или ISO 8601
	Key                       string             `json:"key"`
	Status                    string             `json:"status"`
	Summary                   string             `json:"summary"`
	TenantId                  string             `json:"tenant_id"`
	Priority                  string             `json:"priority"`
	IssueType                 string             `json:"issue_type"`
	Description               string             `json:"description"`
	RenderedDescription       string             `json:"rendered_description"`
	StatusDescription         string             `json:"status_description"`
	AssigneeDisplayName       string             `json:"assignee_displayName"`
	ReporterDisplayName       string             `json:"reporter_displayName"`
	CreatorDisplayName        string             `json:"creator_displayName"`
	ReporterEmailAddress      string             `json:"reporter_emailAddress"`
	PrimaryCategory           string             `json:"primary_category"`
	ResolutionDate            string             `json:"resolutiondate"`
	ResolutionName            string             `json:"resolution_name"`
	ResolutionDetailed        string             `json:"resolution_detailed"`
	PlatformUrl               string             `json:"platform_url"`
	AttackType                string             `json:"attack_type"`
	GossopkaSendTime          string             `json:"gossopka_send_time"`
	GtiSendTime               string             `json:"gti_send_time"`
	CustomerStarRatingComment string             `json:"customer_star_rating_comment"`
	Recommendations           string             `json:"recommendations"`
	AffectedService           string             `json:"affected_service"`
	FakeAsPath                string             `json:"fake_as_path"`
	FakePrefix                string             `json:"fake_prefix"`
	FpType                    string             `json:"fp_type"`
	LookingGlass              string             `json:"looking_glass"`
	SpamRecipients            string             `json:"spam_recipients"`
	TLP                       string             `json:"tlp"`
	UsualPrefix               string             `json:"usual_prefix"`
	UsualAsPath               string             `json:"usual_as_path"`
	Source                    string             `json:"source"`
	ExternalId                string             `json:"external_id"`
	UnderlineSource           string             `json:"_source"`
	GossopkaKey               string             `json:"gossopka_key"`
	Service                   string             `json:"service"`
	Assignee                  string             `json:"assignee"`
	GtiId                     string             `json:"gti_id"`
	CustomerStarRating        string             `json:"customer_star_rating"`
	GossopkaId                string             `json:"gossopka_id"`
	MitreCov                  string             `json:"mitre_cov"`
	ResponseTeam              string             `json:"response_team"`
	GossopkaErrors            string             `json:"gossopka_errors"`
	GossopkaKeyLink           string             `json:"gossopka_key_link"`
	ID                        uint64             `json:"id"`
	System                    uint64             `json:"system"`
	IsVisibleForCustomer      bool               `json:"is_visible_for_customer"`
	ShowGossopkaButton        bool               `json:"show_gossopka_button"`
	CancelGossopkaSendButton  bool               `json:"cancel_gossopka_send_button"`
	ShowGtiButton             bool               `json:"show_gti_button"`
}

type IssueTypeRefType struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CategoryRefType struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ObservedIocsType struct {
	Category []string `json:"category"`
	Ioc      string   `json:"ioc"`
	IocType  string   `json:"ioc_type"`
}

type WatchersType struct {
	Email      string `json:"email"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	Patronimic string `json:"patronimic"`
	Id         uint64 `json:"id"`
	IsActive   bool   `json:"is_active"`
}
