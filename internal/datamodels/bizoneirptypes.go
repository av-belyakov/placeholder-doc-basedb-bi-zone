package datamodels

// VerifiedBiZoneIRPAlert основной объект alert
type VerifiedBiZoneIRPAlert struct {
	AdditionalInformation
	Snapshots          []BiZoneIRPSnapshot `json:"snapshots"`
	Tags               []BiZoneIRPTag      `json:"tags"`
	AffectedLogSources []string            `json:"affected_log_sources"`
	Data               BiZoneIRPData       `json:"data"`
	Confidence         string              `json:"confidence"`
	CreatedTime        string              `json:"created_time"` //дата создания (формат RFC3339) или ISO 8601
	UpdatedTime        string              `json:"updated_time"` //дата обновления (формат RFC3339) или ISO 8601
	CustomerSystem     string              `json:"customer_system"`
	Description        string              `json:"description"`
	DetectionRule      string              `json:"detection_rule"`
	EventStartTime     string              `json:"event_start_time"`     //дата начала события (формат RFC3339) или ISO 8601
	EventEndTime       string              `json:"event_end_time"`       //дата конца события (формат RFC3339) или ISO 8601
	FirstDetectionTime string              `json:"first_detection_time"` //дата первого обнаружения (формат RFC3339) или ISO 8601
	LastDetectionTime  string              `json:"last_detection_time"`  //дата последнего обнаружения (формат RFC3339) или ISO 8601
	PlatformHostname   string              `json:"platform_hostname"`
	PlatformType       string              `json:"platform_type"`
	Severity           string              `json:"severity"`
	PriorityID         string              `json:"priority_id"`
	Title              string              `json:"title"`
	Recommendations    string              `json:"recommendations"`
	UUID               string              `json:"uuid"`
	ExternalID         string              `json:"external_id"`
	ID                 uint64              `json:"id"`
	ResponseTeam       uint64              `json:"response_team"`
}

// BiZoneTag структура для тегов
type BiZoneIRPTag struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	Created   string `json:"created"` //дата создания (формат RFC3339) или ISO 8601
	CreatedBy struct {
		Username string `json:"username"`
		ID       uint64 `json:"id"`
	} `json:"created_by"`
	IsVisibleForCustomer bool `json:"is_visible_for_customer"`
}

// BiZoneData структура для вложенного объекта data
type BiZoneIRPData struct {
	DataSecurity              []BiZoneIRPDataSecurity `json:"data_security"`
	Tags                      []string                `json:"tags"`
	UnmappedDstEndpointArray  []string                `json:"unmapped_dst_endpoint_array"`
	UnmappedHomeEndpointArray []string                `json:"unmapped_home_endpoint_array"`
	DetectionPattern          []uint64                `json:"detection_pattern"`
	UnmappedAgentArray        []uint64                `json:"unmapped_agent_array"`
	Desc                      string                  `json:"desc"`
	EventUID                  string                  `json:"event_uid"`
	JobTitle                  string                  `json:"job_title"`
	FirstSeenTime             string                  `json:"first_seen_time"` //дата создания (формат RFC3339) или ISO 8601
	LastSeenTime              string                  `json:"last_seen_time"`  //дата создания (формат RFC3339) или ISO 8601
	MetadataProductName       string                  `json:"metadata_product_name"`
	SourceIP                  string                  `json:"source_ip"`
	TargetIP                  string                  `json:"target_ip"`
	UnmappedHiveAlertID       string                  `json:"unmapped_hive_alert_id"`
	UnmappedSensorIP          string                  `json:"unmapped_sensor_ip"`
	UnmappedSensorName        string                  `json:"unmapped_sensor_name"`
	Agent                     uint64                  `json:"agent"`
	SeverityID                uint64                  `json:"severity_id"`
}

/*
type BiZoneIRPData struct {
+	DataSecurity              []BiZoneIRPDataSecurity `json:"data_security"`
+	Tags                      []string                `json:"tags"`
+	UnmappedDstEndpointArray  []string                `json:"unmapped_dst_endpoint_array"`
+	UnmappedHomeEndpointArray []string                `json:"unmapped_home_endpoint_array"`
+	DetectionPattern          []uint64                `json:"detection_pattern"`
+	UnmappedAgentArray        []uint64                `json:"unmapped_agent_array"`
	Desc                      string                  `json:"desc"`
	EventUID                  string                  `json:"event_uid"`
	JobTitle                  string                  `json:"job_title"`
	FirstSeenTime             string                  `json:"first_seen_time"` //дата создания (формат RFC3339) или ISO 8601
	LastSeenTime              string                  `json:"last_seen_time"`  //дата создания (формат RFC3339) или ISO 8601
	MetadataProductName       string                  `json:"metadata_product_name"`
	SourceIP                  string                  `json:"source_ip"`
	TargetIP                  string                  `json:"target_ip"`
	UnmappedHiveAlertID       string                  `json:"unmapped_hive_alert_id"`
	UnmappedSensorIP          string                  `json:"unmapped_sensor_ip"`
	UnmappedSensorName        string                  `json:"unmapped_sensor_name"`
+	Agent                     uint64                  `json:"agent"`
+	SeverityID                uint64                  `json:"severity_id"`
}
*/

// BiZoneIRPDataSecurity элемент массива data_security
type BiZoneIRPDataSecurity struct {
	SContent    []BiZoneIRPSContent `json:"s_content"`
	SRuleBody   string              `json:"s_rule_body"`
	SClasstype  string              `json:"s_classtype"`
	SMsg        string              `json:"s_msg"`
	SSourceName string              `json:"s_source_name"`
	ISid        uint64              `json:"i_sid"`
	IRev        uint64              `json:"i_rev"`
	IBRec       uint64              `json:"i_b_rec"`
	SAddDate    uint64              `json:"s_add_date"`
	IPriority   uint64              `json:"i_priority"`
}

// BiZoneIRPSContent элемент массива content
type BiZoneIRPSContent struct {
	Content        string  `json:"content"`
	Nocase         bool    `json:"nocase"`
	HTTPURI        bool    `json:"http_uri"`
	Rawbytes       bool    `json:"rawbytes"`
	HTTPCookie     bool    `json:"http_cookie"`
	HTTPHeader     bool    `json:"http_header"`
	HTTPMethod     bool    `json:"http_method"`
	HTTPRawURI     bool    `json:"http_raw_uri"`
	HTTPStatMsg    bool    `json:"http_stat_msg"`
	HTTPStatCode   bool    `json:"http_stat_code"`
	HTTPRawCookie  bool    `json:"http_raw_cookie"`
	HTTPRawHeader  bool    `json:"http_raw_header"`
	HTTPClientBody bool    `json:"http_client_body"`
	Distance       *string `json:"distance:"`    // null или строка (например "0")
	FastPattern    *bool   `json:"fast_pattern"` // null или bool
	Depth          *int    `json:"depth:"`       // null или число
	Offset         *int    `json:"offset:"`      // null или число
	Within         *int    `json:"within:"`      // null или число
}

type BiZoneIRPCaseFieldData struct {
	Tags              []BiZoneIRPFieldTagDescription `json:"tags"`
	SecondaryCategory []BiZoneIRPFieldWithTitle      `json:"secondary_category"`
	DetectionRules    []string                       `json:"detection_rules"`
	PlatformHostname  []string                       `json:"platform_hostname"`
	ActivityDetected  []any                          `json:"activity_detected"` // пока тип не ясен
	Type              BiZoneIRPFieldWithTitle        `json:"type"`
	Status            BiZoneIRPFieldWithTitle        `json:"status"`
	Priority          BiZoneIRPFieldWithTitle        `json:"priority"`
	PrimaryCategory   BiZoneIRPFieldWithTitle        `json:"primary_category"`
	MITRECOV          BiZoneIRPMITRECOV              `json:"mitre_cov"`
	//MitreCov                  any    `json:"mitre_cov"`                     // пока тип не ясен
	Tenant struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"tenant"`
	CreatedBy struct {
		Id       string `json:"id"`
		Username string `json:"username"`
	} `json:"created_by"`
	Id                        string `json:"id"`
	Tlp                       string `json:"tlp"`
	FpType                    string `json:"fp_type"`
	Created                   string `json:"created"`        //"2026-03-17T12:33:03.716622Z"
	Updated                   string `json:"updated"`        //"2026-03-24T09:57:28.724958Z"
	DetectionDate             string `json:"detection_date"` // "2025-02-21T17:22:46.031534Z"
	Summary                   string `json:"summary"`
	Description               string `json:"description"`
	Recommendations           string `json:"recommendations"`
	ExternalId                string `json:"external_id"`
	StatusDescription         string `json:"status_description"`
	ResolutionDetailed        string `json:"resolution_detailed"`
	CustomerStarRatingComment string `json:"customer_star_rating_comment"`
	Assignee                  any    `json:"assignee,omitzero"` // пока тип не ясен
	CustomerAssignee          any    `json:"customer_assignee"`
	CustomerStarRating        any    `json:"customer_star_rating,omitzero"` // пока тип не ясен
	ResolutionDate            any    `json:"resolutiondate,omitzero"`       // пока тип не ясен
	Resolution                any    `json:"resolution"`                    // пока тип не ясен
	ResponseTeam              any    `json:"response_team"`                 // пока тип не ясен
	IsPublic                  bool   `json:"is_public"`
}

type BiZoneIRPFieldWithTitle struct {
	Title []BiZoneIRPFieldTitle `json:"title"`
	Id    string                `json:"id"`
}

type BiZoneIRPFieldTitle struct {
	Value    string `json:"value"`
	Language string `json:"language"`
}

type BiZoneIRPFieldTagDescription struct {
	Name                 string `json:"name"`
	Color                string `json:"color"`
	IsVisibleForCustomer bool   `json:"is_visible_for_customer"`
}

// VerifiedIRPBiZoneCase основная структура объекта Case
type VerifiedIRPBiZoneCase struct {
	GossopkaErrors            map[string]any          `json:"gossopka_errors,omitzero"` // пока тип в мапе не ясен
	ObservedIocs              []BiZoneIRPObservedIocs `json:"observed_iocs"`
	SecondaryCategoryRef      []BiZoneIRPTypeRef      `json:"secondary_category_ref"`
	Watchers                  []BiZoneIRPWatcher      `json:"watchers"`
	MITRECOV                  BiZoneIRPMITRECOV       `json:"mitre_cov"`
	DetectionRules            []string                `json:"detection_rules"`
	SecondaryCategory         []string                `json:"secondary_category"`
	PlatformHostname          []string                `json:"platform_hostname"`
	WatchersId                []uint64                `json:"watchers_id"`
	ActivityDetected          []any                   `json:"activity_detected,omitzero"` // пока тип в срезе не ясен
	Attachments               []any                   `json:"attachments,omitzero"`       // пока тип в срезе не ясен
	Badges                    []any                   `json:"badges,omitzero"`            // пока тип в срезе не ясен
	Emls                      []any                   `json:"emls,omitzero"`              // пока тип в срезе не ясен
	Slas                      []any                   `json:"slas,omitzero"`              // пока тип в срезе не ясен
	Tags                      []any                   `json:"tags,omitzero"`              // пока тип в срезе не ясен
	KeyField                  []any                   `json:"keyfield,omitzero"`          // пока тип в срезе не ясен
	Assignee                  any                     `json:"assignee,omitzero"`          // пока тип не ясен
	AssigneeDisplayName       any                     `json:"assignee_displayName,omitzero"`
	Service                   any                     `json:"service,omitzero"`                      // пока тип не ясен
	ResolutionDate            any                     `json:"resolutiondate,omitzero"`               // пока тип не ясен
	ResolutionName            any                     `json:"resolution_name,omitzero"`              // пока тип не ясен
	ResolutionNameRef         any                     `json:"resolution_name_ref,omitzero"`          // пока тип не ясен
	GossopkaId                any                     `json:"gossopka_id,omitzero"`                  // пока тип не ясен
	GossopkaSendTime          any                     `json:"gossopka_send_time,omitzero"`           // пока тип не ясен (возможно время)
	GtiId                     any                     `json:"gti_id,omitzero"`                       // пока тип не ясен
	GtiSendTime               any                     `json:"gti_send_time,omitzero"`                // пока тип не ясен (возможно время)
	CustomerStarRating        any                     `json:"customer_star_rating,omitzero"`         // пока тип не ясен
	FirstSentNotificationTime any                     `json:"first_sent_notification_time,omitzero"` // пока тип не ясен (возможно время)
	ResponseTeam              any                     `json:"response_team,omitzero"`                // пока тип не ясен
	GossopkaKeyLink           any                     `json:"gossopka_key_link,omitzero"`            // пока тип не ясен
	IssueTypeRef              BiZoneIRPTypeRef        `json:"issue_type_ref"`
	PrimaryCategoryRef        BiZoneIRPTypeRef        `json:"primary_category_ref"`
	AttackType                string                  `json:"attack_type"`
	Created                   string                  `json:"created"`                  // дата нужно сделать формат RFC3339
	Updated                   string                  `json:"updated"`                  // дата нужно сделать формат RFC3339
	UpdatedAll                string                  `json:"updated_all"`              // дата нужно сделать формат RFC3339
	Timestamp                 string                  `json:"timestamp"`                // дата нужно сделать формат RFC3339
	DetectionDate             string                  `json:"detection_date"`           // дата нужно сделать формат RFC3339
	ParsedActivityDetected    string                  `json:"parsed_activity_detected"` // пока тип в срезе не ясен
	CreatorDisplayName        string                  `json:"creator_displayName"`
	Key                       string                  `json:"key"`
	IssueType                 string                  `json:"issue_type"`
	Priority                  string                  `json:"priority"`
	Summary                   string                  `json:"summary"`
	Description               string                  `json:"description"`
	RenderedDescription       string                  `json:"rendered_description"`
	Status                    string                  `json:"status"`
	StatusDescription         string                  `json:"status_description"`
	ReporterDisplayName       string                  `json:"reporter_displayName"`
	ReporterEmailAddress      string                  `json:"reporter_emailAddress"`
	PrimaryCategory           string                  `json:"primary_category"`
	ResolutionDetailed        string                  `json:"resolution_detailed"`
	PlatformURL               string                  `json:"platform_url"`
	CustomerStarRatingComment string                  `json:"customer_star_rating_comment"`
	Recommendations           string                  `json:"recommendations"`
	AffectedService           string                  `json:"affected_service"`
	FakeAsPath                string                  `json:"fake_as_path"`
	FakePrefix                string                  `json:"fake_prefix"`
	FpType                    string                  `json:"fp_type"`
	LookingGlass              string                  `json:"looking_glass"`
	SpamRecipients            string                  `json:"spam_recipients"`
	TLP                       string                  `json:"tlp"`
	UsualPrefix               string                  `json:"usual_prefix"`
	UsualAsPath               string                  `json:"usual_as_path"`
	Source                    string                  `json:"source"`
	ExternalId                string                  `json:"external_id"`
	UnderliningSource         string                  `json:"_source"`
	GossopkaKey               string                  `json:"gossopka_key"`
	ID                        uint64                  `json:"id"`
	System                    uint64                  `json:"system"`
	CancelGossopkaSendButton  bool                    `json:"cancel_gossopka_send_button"`
	IsVisibleForCustomer      bool                    `json:"is_visible_for_customer"`
	ShowGossopkaButton        bool                    `json:"show_gossopka_button"`
	ShowGtiButton             bool                    `json:"show_gti_button"`
}

type BiZoneIRPTypeRef struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BiZoneIRPObservedIocs struct {
	Category []string `json:"category"`
	IocType  string   `json:"ioc_type"`
	Ioc      string   `json:"ioc"`
}

type BiZoneIRPMITRECOV struct {
	Persistence []string `json:"persistence"`
}

type BiZoneIRPWatcher struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Patronimic string `json:"patronimic"`
	IsActive   bool   `json:"is_active"`
}

// BiZoneSnapshot структура для снимков
type BiZoneIRPSnapshot struct {
	IPAddresses  []string `json:"ip_addresses"`
	MACAddresses []string `json:"mac_addresses"`
	OS           string   `json:"os"`
	Fqdn         string   `json:"fqdn"`
	Domain       string   `json:"domain"`
	CMDBID       string   `json:"cmdb_id"`
	OSType       string   `json:"os_type"`
	Hostname     string   `json:"hostname"`
	UserCMDBName string   `json:"user_cmdb_name"`
}
