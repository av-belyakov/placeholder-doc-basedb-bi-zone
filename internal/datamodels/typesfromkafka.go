package datamodels

/***
* Структуры относящиеся к типу Alert
***/

// VerifiedBiZoneIRPAlert основной объект alert
type VerifiedBiZoneIRPAlert struct {
	AdditionalInformation
	Snapshots          []BiZoneIRPSnapshot `json:"snapshots"`
	Tags               []IRPTagType        `json:"tags"`
	AffectedLogSources []string            `json:"affected_log_sources"`
	Data               BiZoneIRPAlertData  `json:"data"`
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

// BiZoneIRPAlertData структура для вложенного объекта data
type BiZoneIRPAlertData struct {
	DataSecurity              map[string][]BiZoneIRPDataSecurity `json:"data_security"`
	Tags                      []string                           `json:"tags"`
	UnmappedDstEndpointArray  []string                           `json:"unmapped_dst_endpoint_array"`
	UnmappedHomeEndpointArray []string                           `json:"unmapped_home_endpoint_array"`
	DetectionPattern          []uint64                           `json:"detection_pattern"`
	UnmappedAgentArray        []uint64                           `json:"unmapped_agent_array"`
	Desc                      string                             `json:"desc"`
	Query                     string                             `json:"query"`
	EventUID                  string                             `json:"event_uid"`
	JobTitle                  string                             `json:"job_title"`
	FirstSeenTime             string                             `json:"first_seen_time"` //дата создания (формат RFC3339) или ISO 8601
	LastSeenTime              string                             `json:"last_seen_time"`  //дата создания (формат RFC3339) или ISO 8601
	MetadataProductName       string                             `json:"metadata_product_name"`
	SourceIP                  string                             `json:"source_ip"`
	TargetIP                  string                             `json:"target_ip"`
	UnmappedEventCard         string                             `json:"unmapped_event_card"`
	UnmappedHiveAlertID       string                             `json:"unmapped_hive_alert_id"`
	UnmappedSensorIP          string                             `json:"unmapped_sensor_ip"`
	UnmappedSensorName        string                             `json:"unmapped_sensor_name"`
	Agent                     uint64                             `json:"agent"`
	SeverityID                uint64                             `json:"severity_id"`
}

// BiZoneIRPDataSecurity элемент массива data_security
type BiZoneIRPDataSecurity struct {
	SContent    []BiZoneIRPSContent `json:"s_content"`
	SMsg        string              `json:"s_msg"`
	SRuleBody   string              `json:"s_rule_body"`
	SClasstype  string              `json:"s_classtype"`
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
	Depth          *string `json:"depth:"`       // null или строка (например "0")
	Within         *string `json:"within:"`      // null или строка (например "0")
	Distance       *string `json:"distance:"`    // null или строка (например "0")
	FastPattern    *string `json:"fast_pattern"` // null или строка (например "0")
	Offset         *string `json:"offset:"`      // null или строка (например "0")
}

// BiZoneSnapshot структура для снимков
type BiZoneIRPSnapshot struct {
	IPAddresses  []string `json:"ip_addresses"`
	MACAddresses []string `json:"mac_addresses"`
	OS           string   `json:"os"`
	Fqdn         string   `json:"fqdn"`
	Title        string   `json:"title"`
	Domain       string   `json:"domain"`
	CmdbId       string   `json:"cmdb_id"`
	OSType       string   `json:"os_type"`
	Hostname     string   `json:"hostname"`
	Severity     string   `json:"severity"`
	UserCmdbName string   `json:"user_cmdb_name"`
	UserCmdbId   *string  `json:"user_cmdb_id"`
}

/***
* Структуры относящиеся к типу Case
***/

// VerifiedBiZoneIRPCase основной объект case
type VerifiedBiZoneIRPCase struct {
	AdditionalInformation
	Data        BiZoneIRPCaseData `json:"data"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Time        string            `json:"time"` //дата создания (формат RFC3339) или ISO 8601
	Source      string            `json:"source"`
	SpecVersion string            `json:"specversion"`
	Subject     *string           `json:"subject"`
	DataSchema  *string           `json:"dataschema"`
}

// BiZoneIRPCaseData полезная нагрузка события
type BiZoneIRPCaseData struct {
	Tags                      []IRPTagType          `json:"tags"`
	SecondaryCategory         []BiZoneIRPCategory   `json:"secondary_category"`
	ActivityDetected          []any                 `json:"activity_detected"`
	PlatformHostname          []string              `json:"platform_hostname"`
	DetectionRules            []uint64              `json:"detection_rules"`
	Type                      BiZoneIRPType         `json:"type"`
	Status                    BiZoneIRPStatus       `json:"status"`
	Priority                  BiZoneIRPPriority     `json:"priority"`
	PrimaryCategory           BiZoneIRPCategory     `json:"primary_category"`
	Tenant                    BiZoneIRPTenant       `json:"tenant"`
	CreatedBy                 BiZoneIRPCreatedBy    `json:"created_by"`
	Organization              BiZoneIRPOrganization `json:"organization,omitzero"`
	Created                   string                `json:"created"`        // дата создания (формат RFC3339) или ISO 8601
	Updated                   string                `json:"updated"`        // дата обновления (формат RFC3339) или ISO 8601
	DetectionDate             string                `json:"detection_date"` // дата обнаружения (формат RFC3339) или ISO 8601
	ResolutionDate            string                `json:"resolutiondate"` // дата принятия решения (формат RFC3339) или ISO 8601
	ID                        string                `json:"id"`
	TLP                       string                `json:"tlp"`
	Title                     string                `json:"title,omitzero"`
	FpType                    string                `json:"fp_type"`
	Summary                   string                `json:"summary"`
	ExternalID                string                `json:"external_id"`
	Description               string                `json:"description"`
	Recommendations           string                `json:"recommendations"`
	StatusDescription         string                `json:"status_description"`
	ResolutionDetailed        string                `json:"resolution_detailed"`
	CustomerStarRatingComment string                `json:"customer_star_rating_comment"`
	Assignee                  *string               `json:"assignee"`
	MitreCov                  *string               `json:"mitre_cov"`
	Resolution                *string               `json:"resolution"`
	ResponseTeam              *string               `json:"response_team"`
	CustomerAssignee          *string               `json:"customer_assignee"`
	CustomerStarRating        *uint64               `json:"customer_star_rating"`
	IsPublic                  bool                  `json:"is_public"`
}

// BiZoneIRPType тип инцидента
type BiZoneIRPType struct {
	ID    string   `json:"id"`
	Title []string `json:"title"`
	//Title []BiZoneIRPTitle `json:"title"`
}

// BiZoneIRPPriority приоритет
type BiZoneIRPPriority struct {
	ID    string   `json:"id"`
	Title []string `json:"title"`
	//Title []BiZoneIRPTitle `json:"title"`
}

// BiZoneIRPStatus статус
type BiZoneIRPStatus struct {
	ID    string   `json:"id"`
	Title []string `json:"title"`
	//Title []BiZoneIRPTitle `json:"title"`
}

// BiZoneIRPCategory категория (первичная/вторичная)
type BiZoneIRPCategory struct {
	ID    string   `json:"id"`
	Title []string `json:"title"`
	//Title []BiZoneIRPTitle `json:"title"`
}

// BiZoneIRPTenant арендатор (СОА)
type BiZoneIRPTenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BiZoneIRPCreatedBy автор инцидента
type BiZoneIRPCreatedBy struct {
	Username string `json:"username"`
	ID       uint64 `json:"id"`
}

// BiZoneIRPTitle локализованное значение
type BiZoneIRPTitle struct {
	Value    string `json:"value"`
	Language string `json:"language"`
}

type BiZoneIRPOrganization struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
