package datamodels

import (
	"errors"
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewBiZoneIRPAlertData новый объект Data родительского объекта VerifiedBiZoneIRPAler
func NewBiZoneIRPAlertData() *BiZoneIRPAlertData {
	return &BiZoneIRPAlertData{
		DataSecurity:              map[string][]BiZoneIRPDataSecurity(nil),
		Tags:                      []string(nil),
		UnmappedDstEndpointArray:  []string(nil),
		UnmappedHomeEndpointArray: []string(nil),
		DetectionPattern:          []uint64(nil),
		UnmappedAgentArray:        []uint64(nil),
	}
}

func (d *BiZoneIRPAlertData) Get() *BiZoneIRPAlertData {
	return d
}

// GetAgent для поля 'agent'
func (d *BiZoneIRPAlertData) GetAgent() uint64 {
	return d.Agent
}

// SetAgent для поля 'agent'
func (d *BiZoneIRPAlertData) SetAgent(v uint64) error {
	d.Agent = v

	return nil
}

// SetAnyAgent для поля 'agent'
func (d *BiZoneIRPAlertData) SetAnyAgent(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetAgent(v)
}

// GetSeverityID для поля 'severity_id'
func (d *BiZoneIRPAlertData) GetSeverityID() uint64 {
	return d.SeverityID
}

// SetSeverityID для поля 'severity_id'
func (d *BiZoneIRPAlertData) SetSeverityID(v uint64) error {
	d.SeverityID = v

	return nil
}

// SetAnySeverityID для поля 'severity_id'
func (d *BiZoneIRPAlertData) SetAnySeverityID(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetSeverityID(v)
}

// GetDesc для поля 'desc'
func (d *BiZoneIRPAlertData) GetDesc() string {
	return d.Desc
}

// SetDesc для поля 'desc'
func (d *BiZoneIRPAlertData) SetDesc(v string) error {
	d.Desc = v

	return nil
}

// SetAnyDesc для поля 'desc'
func (d *BiZoneIRPAlertData) SetAnyDesc(a any) error {
	return d.SetDesc(fmt.Sprint(a))
}

// GetQuery для поля 'query'
func (d *BiZoneIRPAlertData) GetQuery() string {
	return d.Query
}

// SetQuery для поля 'query'
func (d *BiZoneIRPAlertData) SetQuery(v string) error {
	d.Query = v

	return nil
}

// SetAnyQuery для поля 'query'
func (d *BiZoneIRPAlertData) SetAnyQuery(a any) error {
	return d.SetQuery(fmt.Sprint(a))
}

// GetEventUid для поля 'event_uid'
func (d *BiZoneIRPAlertData) GetEventUid() string {
	return d.EventUID
}

// SetEventUid для поля 'event_uid'
func (d *BiZoneIRPAlertData) SetEventUid(v string) error {
	d.EventUID = v

	return nil
}

// SetAnyEventUid для поля 'event_uid'
func (d *BiZoneIRPAlertData) SetAnyEventUid(a any) error {
	return d.SetEventUid(fmt.Sprint(a))
}

// GetJobTitle для поля 'job_title'
func (d *BiZoneIRPAlertData) GetJobTitle() string {
	return d.JobTitle
}

// SetJobTitle для поля 'job_title'
func (d *BiZoneIRPAlertData) SetJobTitle(v string) error {
	d.JobTitle = v

	return nil
}

// SetAnyJobTitle для поля 'job_title'
func (d *BiZoneIRPAlertData) SetAnyJobTitle(a any) error {
	return d.SetJobTitle(fmt.Sprint(a))
}

// GetMetadataProductName для поля 'metadata_product_name'
func (d *BiZoneIRPAlertData) GetMetadataProductName() string {
	return d.MetadataProductName
}

// SetMetadataProductName для поля 'metadata_product_name'
func (d *BiZoneIRPAlertData) SetMetadataProductName(v string) error {
	d.MetadataProductName = v

	return nil
}

// SetAnyMetadataProductName для поля 'metadata_product_name'
func (d *BiZoneIRPAlertData) SetAnyMetadataProductName(a any) error {
	return d.SetMetadataProductName(fmt.Sprint(a))
}

// GetSourceIP для поля 'source_ip'
func (d *BiZoneIRPAlertData) GetSourceIP() string {
	return d.SourceIP
}

// SetSourceIP для поля 'source_ip'
func (d *BiZoneIRPAlertData) SetSourceIP(v string) error {
	d.SourceIP = v

	return nil
}

// SetAnySourceIP для поля 'source_ip'
func (d *BiZoneIRPAlertData) SetAnySourceIP(a any) error {
	return d.SetSourceIP(fmt.Sprint(a))
}

// GetTargetIP для поля 'target_ip'
func (d *BiZoneIRPAlertData) GetTargetIP() string {
	return d.TargetIP
}

// SetTargetIP для поля 'target_ip'
func (d *BiZoneIRPAlertData) SetTargetIP(v string) error {
	d.TargetIP = v

	return nil
}

// SetAnyTargetIP для поля 'target_ip'
func (d *BiZoneIRPAlertData) SetAnyTargetIP(a any) error {
	return d.SetTargetIP(fmt.Sprint(a))
}

// GetUnmappedEventCard для поля 'unmapped_event_card'
func (d *BiZoneIRPAlertData) GetUnmappedEventCard() string {
	return d.UnmappedEventCard
}

// SetUnmappedEventCard для поля 'unmapped_event_card'
func (d *BiZoneIRPAlertData) SetUnmappedEventCard(v string) error {
	d.UnmappedEventCard = v

	return nil
}

// SetAnyUnmappedEventCard для поля 'unmapped_event_card'
func (d *BiZoneIRPAlertData) SetAnyUnmappedEventCard(a any) error {
	return d.SetUnmappedEventCard(fmt.Sprint(a))
}

// GetUnmappedHiveAlertID для поля 'unmapped_hive_alert_id'
func (d *BiZoneIRPAlertData) GetUnmappedHiveAlertID() string {
	return d.UnmappedHiveAlertID
}

// SetUnmappedHiveAlertID для поля 'unmapped_hive_alert_id'
func (d *BiZoneIRPAlertData) SetUnmappedHiveAlertID(v string) error {
	d.UnmappedHiveAlertID = v

	return nil
}

// SetAnyUnmappedHiveAlertID для поля 'unmapped_hive_alert_id'
func (d *BiZoneIRPAlertData) SetAnyUnmappedHiveAlertID(a any) error {
	return d.SetUnmappedHiveAlertID(fmt.Sprint(a))
}

// GetUnmappedSensorIP для поля 'unmapped_sensor_ip'
func (d *BiZoneIRPAlertData) GetUnmappedSensorIP() string {
	return d.UnmappedSensorIP
}

// SetUnmappedSensorIP для поля 'unmapped_sensor_ip'
func (d *BiZoneIRPAlertData) SetUnmappedSensorIP(v string) error {
	d.UnmappedSensorIP = v

	return nil
}

// SetAnyUnmappedSensorIP для поля 'unmapped_sensor_ip'
func (d *BiZoneIRPAlertData) SetAnyUnmappedSensorIP(a any) error {
	return d.SetUnmappedSensorIP(fmt.Sprint(a))
}

// GetUnmappedSensorName для поля 'unmapped_sensor_name'
func (d *BiZoneIRPAlertData) GetUnmappedSensorName() string {
	return d.UnmappedSensorName
}

// SetUnmappedSensorName для поля 'unmapped_sensor_name'
func (d *BiZoneIRPAlertData) SetUnmappedSensorName(v string) error {
	d.UnmappedSensorName = v

	return nil
}

// SetAnyUnmappedSensorName для поля 'unmapped_sensor_name'
func (d *BiZoneIRPAlertData) SetAnyUnmappedSensorName(a any) error {
	return d.SetUnmappedSensorName(fmt.Sprint(a))
}

// GetFirstSeenTime для поля 'first_seen_time' (формат RFC3339)
func (d *BiZoneIRPAlertData) GetFirstSeenTime() string {
	return d.FirstSeenTime
}

// SetFirstSeenTime для поля 'first_seen_time' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPAlertData) SetFirstSeenTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.FirstSeenTime = timeStr

	return nil
}

// SetAnyFirstSeenTime для поля 'first_seen_time'
func (d *BiZoneIRPAlertData) SetAnyFirstSeenTime(a any) error {
	if v, ok := a.(string); ok {
		return d.SetFirstSeenTime(v)
	}

	return errors.New("type conversion error for field 'first_seen_time'")
}

// GetLastSeenTime для поля 'last_seen_time' (формат RFC3339)
func (d *BiZoneIRPAlertData) GetLastSeenTime() string {
	return d.LastSeenTime
}

// SetLastSeenTime для поля 'last_seen_time' (преобразует в формат времени RFC3339)
func (d *BiZoneIRPAlertData) SetLastSeenTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.LastSeenTime = timeStr

	return nil
}

// SetAnyLastSeenTime для поля 'last_seen_time'
func (d *BiZoneIRPAlertData) SetAnyLastSeenTime(a any) error {
	if v, ok := a.(string); ok {
		return d.SetLastSeenTime(v)
	}

	return errors.New("type conversion error for field 'last_seen_time'")
}

// GetTags для поля 'tags'
func (d *BiZoneIRPAlertData) GetTags() []string {
	return d.Tags
}

// SetTags для поля 'tags'
func (d *BiZoneIRPAlertData) SetTags(v []string) error {
	d.Tags = v

	return nil
}

// SetTag добавляет значение 'tag' в список
func (d *BiZoneIRPAlertData) SetTag(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.Tags); !isExist {
		d.Tags = append(d.Tags, v)
	}

	return nil
}

// SetAnyTag добавляет некоторое значение в список 'tags'
func (d *BiZoneIRPAlertData) SetAnyTag(a any) error {
	return d.SetTag(fmt.Sprint(a))
}

// GetUnmappedDstEndpointArray для поля 'unmapped_dst_endpoint_array'
func (d *BiZoneIRPAlertData) GetUnmappedDstEndpointArray() []string {
	return d.UnmappedDstEndpointArray
}

// SetUnmappedDstEndpointArray для поля 'unmapped_dst_endpoint_array'
func (d *BiZoneIRPAlertData) SetUnmappedDstEndpointArray(v []string) error {
	d.UnmappedDstEndpointArray = v

	return nil
}

// SetUnmappedDstEndpointArrayElement добавляет значение 'unmapped_dst_endpoint_array' в список
func (d *BiZoneIRPAlertData) SetUnmappedDstEndpointArrayElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.UnmappedDstEndpointArray); !isExist {
		d.UnmappedDstEndpointArray = append(d.UnmappedDstEndpointArray, v)
	}

	return nil
}

// SetAnyUnmappedDstEndpointArray добавляет некоторое значение в список 'unmapped_dst_endpoint_array'
func (d *BiZoneIRPAlertData) SetAnyUnmappedDstEndpointArray(a any) error {
	return d.SetUnmappedDstEndpointArrayElement(fmt.Sprint(a))
}

// GetUnmappedHomeEndpointArray для поля 'unmapped_home_endpoint_array'
func (d *BiZoneIRPAlertData) GetUnmappedHomeEndpointArray() []string {
	return d.UnmappedHomeEndpointArray
}

// SetUnmappedHomeEndpointArray для поля 'unmapped_home_endpoint_array'
func (d *BiZoneIRPAlertData) SetUnmappedHomeEndpointArray(v []string) error {
	d.UnmappedHomeEndpointArray = v

	return nil
}

// SetUnmappedHomeEndpointArrayElement добавляет значение 'unmapped_home_endpoint_array' в список
func (d *BiZoneIRPAlertData) SetUnmappedHomeEndpointArrayElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.UnmappedHomeEndpointArray); !isExist {
		d.UnmappedHomeEndpointArray = append(d.UnmappedHomeEndpointArray, v)
	}

	return nil
}

// SetAnyUnmappedHomeEndpointArray добавляет некоторое значение в список 'unmapped_home_endpoint_array'
func (d *BiZoneIRPAlertData) SetAnyUnmappedHomeEndpointArray(a any) error {
	return d.SetUnmappedHomeEndpointArrayElement(fmt.Sprint(a))
}

// GetDetectionPattern для поля 'detection_pattern'
func (d *BiZoneIRPAlertData) GetDetectionPattern() []uint64 {
	return d.DetectionPattern
}

// SetDetectionPattern для поля 'detection_pattern'
func (d *BiZoneIRPAlertData) SetDetectionPattern(v []uint64) error {
	d.DetectionPattern = v

	return nil
}

// SetDetectionPatternElement добавляет значение 'detection_pattern' в список
func (d *BiZoneIRPAlertData) SetDetectionPatternElement(v uint64) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.DetectionPattern); !isExist {
		d.DetectionPattern = append(d.DetectionPattern, v)
	}

	return nil
}

// SetAnyDetectionPatternElement добавляет некоторое значение в список 'detection_pattern'
func (d *BiZoneIRPAlertData) SetAnyDetectionPatternElement(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetDetectionPatternElement(v)
}

// GetUnmappedAgentArray для поля 'unmapped_agent_array'
func (d *BiZoneIRPAlertData) GetUnmappedAgentArray() []uint64 {
	return d.UnmappedAgentArray
}

// SetUnmappedAgentArrayn для поля 'unmapped_agent_array'
func (d *BiZoneIRPAlertData) SetUnmappedAgentArrayn(v []uint64) error {
	d.UnmappedAgentArray = v

	return nil
}

// SetUnmappedAgentArrayElement добавляет значение 'unmapped_agent_array' в список
func (d *BiZoneIRPAlertData) SetUnmappedAgentArrayElement(v uint64) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, d.UnmappedAgentArray); !isExist {
		d.UnmappedAgentArray = append(d.UnmappedAgentArray, v)
	}

	return nil
}

// SetAnyUnmappedAgentArrayElement добавляет некоторое значение в список 'unmapped_agent_array'
func (d *BiZoneIRPAlertData) SetAnyUnmappedAgentArrayElement(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetUnmappedAgentArrayElement(v)
}

// GetDataSecurity для поля 'data_security'
func (d *BiZoneIRPAlertData) GetDataSecurity() map[string][]BiZoneIRPDataSecurity {
	return d.DataSecurity
}

// SetDataSecurity для поля 'data_security'
func (d *BiZoneIRPAlertData) SetDataSecurity(v map[string][]BiZoneIRPDataSecurity) error {
	d.DataSecurity = v

	return nil
}

// SetDataSecurityElement добавляет значение 'data_security' в список
func (d *BiZoneIRPAlertData) SetDataSecurityElement(k string, v []BiZoneIRPDataSecurity) error {
	d.DataSecurity[k] = v

	return nil
}

// ToStringBeautiful форматированный вывод
func (d *BiZoneIRPAlertData) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsStr := supportingfunctions.GetWhitespace(num + 1)
	wsInt := supportingfunctions.GetWhitespace(num + 2)

	fmt.Fprintf(&str, "%s'desc': '%s'\n", ws, d.Desc)
	fmt.Fprintf(&str, "%s'query': '%s'\n", ws, d.Query)
	fmt.Fprintf(&str, "%s'source_ip': '%s'\n", ws, d.SourceIP)
	fmt.Fprintf(&str, "%s'target_ip': '%s'\n", ws, d.TargetIP)
	fmt.Fprintf(&str, "%s'event_uid': '%s'\n", ws, d.EventUID)
	fmt.Fprintf(&str, "%s'job_title': '%s'\n", ws, d.JobTitle)
	fmt.Fprintf(&str, "%s'first_seen_time': '%s'\n", ws, d.FirstSeenTime)
	fmt.Fprintf(&str, "%s'last_seen_time': '%s'\n", ws, d.LastSeenTime)
	fmt.Fprintf(&str, "%s'unmapped_sensor_ip': '%s'\n", ws, d.UnmappedSensorIP)
	fmt.Fprintf(&str, "%s'unmapped_event_card': '%s'\n", ws, d.UnmappedEventCard)
	fmt.Fprintf(&str, "%s'metadata_product_name': '%s'\n", ws, d.MetadataProductName)
	fmt.Fprintf(&str, "%s'unmapped_hive_alert_id': '%s'\n", ws, d.UnmappedHiveAlertID)
	fmt.Fprintf(&str, "%s'unmapped_sensor_name': '%s'\n", ws, d.UnmappedSensorName)
	fmt.Fprintf(&str, "%s'agent': '%d'\n", ws, d.Agent)
	fmt.Fprintf(&str, "%s'severity_id': '%d'\n", ws, d.SeverityID)
	fmt.Fprintf(&str, "%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.Tags))
	fmt.Fprintf(&str, "%s'unmapped_dst_endpoint_array': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.UnmappedDstEndpointArray))
	fmt.Fprintf(&str, "%s'unmapped_home_endpoint_array': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.UnmappedHomeEndpointArray))
	fmt.Fprintf(&str, "%s'detection_pattern': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.DetectionPattern))
	fmt.Fprintf(&str, "%s'unmapped_agent_array': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.UnmappedAgentArray))
	fmt.Fprintf(&str, "%s'data_security':\n", ws)
	for k, v := range d.DataSecurity {
		fmt.Fprintf(&str, "%s%s:\n", wsStr, k)

		for item, value := range v {
			fmt.Fprintf(&str, "%s%d.\n%s\n", wsInt, item, value.ToStringBeautiful(num+3))
		}
	}

	return str.String()
}
