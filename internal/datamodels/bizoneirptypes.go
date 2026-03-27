package datamodels

// VerifiedBiZoneAlert основная структура объекта Alert
type VerifiedBiZoneIRPAlert struct {
	AdditionalInformation
	Snapshots          []BiZoneIRPSnapshot `json:"snapshots"`
	Tags               []BiZoneIRPTag      `json:"tags"`
	SpecialUUID        string              `json:"@special_uuid"`
	UUID               string              `json:"uuid"`
	Title              string              `json:"title"`
	Severity           string              `json:"severity"`
	ExternalID         string              `json:"external_id"`
	Confidence         string              `json:"confidence"`
	Description        string              `json:"description"`
	CreatedTime        string              `json:"created_time"` //дата создания (формат RFC3339)
	UpdatedTime        string              `json:"updated_time"` //дата обновления (формат RFC3339)
	PlatformType       string              `json:"platform_type"`
	EventStartTime     string              `json:"event_start_time"` //дата начала события (формат RFC3339)
	EventEndTime       string              `json:"event_end_time"`   //дата конца события (формат RFC3339)
	DetectionRule      string              `json:"detection_rule"`
	CustomerSystem     string              `json:"customer_system"`
	Recommendations    string              `json:"recommendations"`
	PlatformHostname   string              `json:"platform_hostname"`
	FirstDetectionTime string              `json:"first_detection_time"` //дата первого обнаружения (формат RFC3339)
	LastDetectionTime  string              `json:"last_detection_time"`  //дата последнего обнаружения (формат RFC3339)
	AffectedLogSources string              `json:"affected_log_sources"`
	IDNum              uint64              `json:"id"`
	Data               BiZoneIRPData       `json:"data"`
}

// BiZoneData структура для вложенного объекта data
type BiZoneIRPData struct {
	AllIPHome    []string `json:"all__ip_home"`
	SnortSid     []uint64 `json:"snort_sid"`
	AllSensors   []uint64 `json:"all_sensors"`
	ID           string   `json:"_id"`
	Title        string   `json:"title"`
	IPHome       string   `json:"ip_home"`
	URLFTP       string   `json:"url____ftp"`
	IPExter      string   `json:"ip_exter"`
	URLHTTP      string   `json:"url___http"`
	EventType    string   `json:"event_type"`
	URLArkime    string   `json:"url_arkime"`
	Sensor       uint64   `json:"sensor"`
	AllIPExt     uint64   `json:"all____ip_ext"`
	ResponseTeam uint64   `json:"response_team"`
}

// BiZoneIRPTypeCase основная структура объекта Case (новая)
/* основной объект состоит из структуры типа
"id": "socd-soar-prod-issue-v10267578",
    "headers": [
        {
            "key": "content-type",
            "value": "[97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110]"
        }
    ],
    "offset": 267578,
    "timestamp": 1774346255756,
    "timestampType": {
        "id": "CreateTime",
        "name": "CreateTime"
    },
    "topic": "socd-soar-prod-issue-v1",
    "partition": 0,
    "value": {
        "source": "soar.core",
        "id": "516d7e1b-8e10-4d3b-960a-87099a45bb41",
        "type": "soar.core.incident.changed.v1",
        "specversion": "1.0",
        "time": "2026-03-24T09:57:28.791931+00:00",
        "subject": null,
        "dataschema": null,
        "data": {

		}
    },
    "key": null,
    "keyMeta": null,
    "valueMeta": {},
    "valueDecoder": "Json",
    "keyDecoder": null
	где основные поля, нобходимые для анализа кейса находятся в value.data
*/
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

// BiZoneTag структура для тегов
type BiZoneIRPTag struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	Created   string `json:"created"` //дата создания (формат RFC3339)
	CreatedBy struct {
		Username string `json:"username"`
		ID       uint64 `json:"id"`
	} `json:"created_by"`
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
