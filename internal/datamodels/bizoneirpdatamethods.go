package datamodels

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// NewBiZoneIRPData новый объект Data
func NewBiZoneIRPData() *BiZoneIRPData {
	return &BiZoneIRPData{
		DataSecurity:              []BiZoneIRPDataSecurity(nil),
		Tags:                      []string(nil),
		UnmappedDstEndpointArray:  []string(nil),
		UnmappedHomeEndpointArray: []string(nil),
		DetectionPattern:          []uint64(nil),
		UnmappedAgentArray:        []uint64(nil),
	}
}

func (d *BiZoneIRPData) Get() *BiZoneIRPData {
	return d
}

// GetAgent для поля agent
func (d *BiZoneIRPData) GetAgent() uint64 {
	return d.Agent
}

// SetAgent для поля agent
func (d *BiZoneIRPData) SetAgent(v uint64) error {
	d.Agent = v

	return nil
}

// SetAnyAgent для поля agent
func (d *BiZoneIRPData) SetAnyAgent(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetAgent(v)
}

// GetSeverityID для поля severity_id
func (d *BiZoneIRPData) GetSeverityID() uint64 {
	return d.SeverityID
}

// SetSeverityID для поля severity_id
func (d *BiZoneIRPData) SetSeverityID(v uint64) error {
	d.SeverityID = v

	return nil
}

// SetAnySeverityID для поля severity_id
func (d *BiZoneIRPData) SetAnySeverityID(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetSeverityID(v)
}

// GetDesc для поля desc
func (d *BiZoneIRPData) GetDesc() string {
	return d.Desc
}

// SetDesc для поля desc
func (d *BiZoneIRPData) SetDesc(v string) error {
	d.Desc = v

	return nil
}

// SetAnyDesc для поля desc
func (d *BiZoneIRPData) SetAnyDesc(a any) error {
	return d.SetDesc(fmt.Sprint(a))
}

// GetEventUid для поля event_uid
func (d *BiZoneIRPData) GetEventUid() string {
	return d.EventUID
}

// SetEventUid для поля event_uid
func (d *BiZoneIRPData) SetEventUid(v string) error {
	d.EventUID = v

	return nil
}

// SetAnyEventUid для поля event_uid
func (d *BiZoneIRPData) SetAnyEventUid(a any) error {
	return d.SetEventUid(fmt.Sprint(a))
}

// GetJobTitle для поля job_title
func (d *BiZoneIRPData) GetJobTitle() string {
	return d.JobTitle
}

// SetJobTitle для поля job_title
func (d *BiZoneIRPData) SetJobTitle(v string) error {
	d.JobTitle = v

	return nil
}

// SetAnyIPExter для поля job_title
func (d *BiZoneIRPData) SetAnyIPExter(a any) error {
	return d.SetAnyIPExter(fmt.Sprint(a))
}

// GetMetadataProductName для поля metadata_product_name
func (d *BiZoneIRPData) GetMetadataProductName() string {
	return d.MetadataProductName
}

// SetMetadataProductName для поля metadata_product_name
func (d *BiZoneIRPData) SetMetadataProductName(v string) error {
	d.MetadataProductName = v

	return nil
}

// SetAnyMetadataProductName для поля metadata_product_name
func (d *BiZoneIRPData) SetAnyMetadataProductName(a any) error {
	return d.SetMetadataProductName(fmt.Sprint(a))
}

// GetSourceIP для поля source_ip
func (d *BiZoneIRPData) GetSourceIP() string {
	return d.SourceIP
}

// SetSourceIP для поля source_ip
func (d *BiZoneIRPData) SetSourceIP(v string) error {
	d.SourceIP = v

	return nil
}

// SetAnySourceIP для поля source_ip
func (d *BiZoneIRPData) SetAnySourceIP(a any) error {
	return d.SetSourceIP(fmt.Sprint(a))
}

// GetTargetIP для поля target_ip
func (d *BiZoneIRPData) GetTargetIP() string {
	return d.TargetIP
}

// SetTargetIP для поля target_ip
func (d *BiZoneIRPData) SetTargetIP(v string) error {
	d.TargetIP = v

	return nil
}

// SetAnyTargetIP для поля target_ip
func (d *BiZoneIRPData) SetAnyTargetIP(a any) error {
	return d.SetTargetIP(fmt.Sprint(a))
}

// GetUnmappedHiveAlertID для поля unmapped_hive_alert_id
func (d *BiZoneIRPData) GetUnmappedHiveAlertID() string {
	return d.UnmappedHiveAlertID
}

// SetUnmappedHiveAlertID для поля unmapped_hive_alert_id
func (d *BiZoneIRPData) SetUnmappedHiveAlertID(v string) error {
	d.UnmappedHiveAlertID = v

	return nil
}

// SetAnyUnmappedHiveAlertID для поля unmapped_hive_alert_id
func (d *BiZoneIRPData) SetAnyUnmappedHiveAlertID(a any) error {
	return d.SetUnmappedHiveAlertID(fmt.Sprint(a))
}

// GetUnmappedSensorIP для поля unmapped_sensor_ip
func (d *BiZoneIRPData) GetUnmappedSensorIP() string {
	return d.UnmappedSensorIP
}

// SetUnmappedSensorIP для поля unmapped_sensor_ip
func (d *BiZoneIRPData) SetUnmappedSensorIP(v string) error {
	d.UnmappedSensorIP = v

	return nil
}

// SetAnyUnmappedSensorIP для поля unmapped_sensor_ip
func (d *BiZoneIRPData) SetAnyUnmappedSensorIP(a any) error {
	return d.SetUnmappedSensorIP(fmt.Sprint(a))
}

// GetUnmappedSensorName для поля unmapped_sensor_name
func (d *BiZoneIRPData) GetUnmappedSensorName() string {
	return d.UnmappedSensorName
}

// SetUnmappedSensorName для поля unmapped_sensor_name
func (d *BiZoneIRPData) SetUnmappedSensorName(v string) error {
	d.UnmappedSensorName = v

	return nil
}

// SetAnyUnmappedSensorName для поля unmapped_sensor_name
func (d *BiZoneIRPData) SetAnyUnmappedSensorName(a any) error {
	return d.SetUnmappedSensorName(fmt.Sprint(a))
}

// GetFirstSeenTime для поля first_seen_time (формат RFC3339)
func (d *BiZoneIRPData) GetFirstSeenTime() string {
	return d.FirstSeenTime
}

// SetFirstSeenTime для поля first_seen_time (преобразует в формат времени RFC3339)
func (d *BiZoneIRPData) SetFirstSeenTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.FirstSeenTime = timeStr

	return nil
}

// SetAnyFirstSeenTime для поля first_seen_time
func (d *BiZoneIRPData) SetAnyFirstSeenTime(a any) error {
	if v, ok := a.(string); ok {
		return d.SetFirstSeenTime(v)
	}

	return errors.New("type conversion error for field 'first_seen_time'")
}

// GetLastSeenTime для поля last_seen_time (формат RFC3339)
func (d *BiZoneIRPData) GetLastSeenTime() string {
	return d.LastSeenTime
}

// SetLastSeenTime для поля last_seen_time (преобразует в формат времени RFC3339)
func (d *BiZoneIRPData) SetLastSeenTime(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	d.LastSeenTime = timeStr

	return nil
}

// SetAnyLastSeenTime для поля last_seen_time
func (d *BiZoneIRPData) SetAnyLastSeenTime(a any) error {
	if v, ok := a.(string); ok {
		return d.SetLastSeenTime(v)
	}

	return errors.New("type conversion error for field 'last_seen_time'")
}

// GetTags для поля tags
func (d *BiZoneIRPData) GetTags() []string {
	return d.Tags
}

// SetTags для поля tags
func (d *BiZoneIRPData) SetTags(v []string) error {
	d.Tags = v

	return nil
}

// SetTag добавляет значение tag в список
func (d *BiZoneIRPData) SetTag(v string) error {
	if d.Tags == nil {
		d.Tags = []string(nil)
	}

	if slices.Contains(d.Tags, v) {
		return nil
	}

	d.Tags = append(d.Tags, v)

	return nil
}

// SetAnyTag добавляет некоторое значение в список tags
func (d *BiZoneIRPData) SetAnyTag(a any) error {
	return d.SetTag(fmt.Sprint(a))
}

// GetUnmappedDstEndpointArray для поля unmapped_dst_endpoint_array
func (d *BiZoneIRPData) GetUnmappedDstEndpointArray() []string {
	return d.UnmappedDstEndpointArray
}

// SetUnmappedDstEndpointArray для поля unmapped_dst_endpoint_array
func (d *BiZoneIRPData) SetUnmappedDstEndpointArray(v []string) error {
	d.UnmappedDstEndpointArray = v

	return nil
}

// SetUnmappedDstEndpointArrayElement добавляет значение unmapped_dst_endpoint_array в список
func (d *BiZoneIRPData) SetUnmappedDstEndpointArrayElement(v string) error {
	if d.UnmappedDstEndpointArray == nil {
		d.UnmappedDstEndpointArray = []string(nil)
	}

	if slices.Contains(d.UnmappedDstEndpointArray, v) {
		return nil
	}

	d.UnmappedDstEndpointArray = append(d.UnmappedDstEndpointArray, v)

	return nil
}

// SetAnyUnmappedDstEndpointArray добавляет некоторое значение в список unmapped_dst_endpoint_array
func (d *BiZoneIRPData) SetAnyUnmappedDstEndpointArray(a any) error {
	return d.SetUnmappedDstEndpointArrayElement(fmt.Sprint(a))
}

// GetUnmappedHomeEndpointArray для поля unmapped_home_endpoint_array
func (d *BiZoneIRPData) GetUnmappedHomeEndpointArray() []string {
	return d.UnmappedHomeEndpointArray
}

// SetUnmappedHomeEndpointArray для поля unmapped_home_endpoint_array
func (d *BiZoneIRPData) SetUnmappedHomeEndpointArray(v []string) error {
	d.UnmappedHomeEndpointArray = v

	return nil
}

// SetUnmappedHomeEndpointArrayElement добавляет значение unmapped_home_endpoint_array в список
func (d *BiZoneIRPData) SetUnmappedHomeEndpointArrayElement(v string) error {
	if d.UnmappedHomeEndpointArray == nil {
		d.UnmappedHomeEndpointArray = []string(nil)
	}

	if slices.Contains(d.UnmappedHomeEndpointArray, v) {
		return nil
	}

	d.UnmappedHomeEndpointArray = append(d.UnmappedHomeEndpointArray, v)

	return nil
}

// SetAnyUnmappedHomeEndpointArray добавляет некоторое значение в список unmapped_home_endpoint_array
func (d *BiZoneIRPData) SetAnyUnmappedHomeEndpointArray(a any) error {
	return d.SetUnmappedDstEndpointArrayElement(fmt.Sprint(a))
}

// GetDetectionPattern для поля detection_pattern
func (d *BiZoneIRPData) GetDetectionPattern() []uint64 {
	return d.DetectionPattern
}

// SetDetectionPattern для поля detection_pattern
func (d *BiZoneIRPData) SetDetectionPattern(v []uint64) error {
	d.DetectionPattern = v

	return nil
}

// SetDetectionPatternElement добавляет значение detection_pattern в список
func (d *BiZoneIRPData) SetDetectionPatternElement(v uint64) error {
	if d.DetectionPattern == nil {
		d.DetectionPattern = []uint64(nil)
	}

	if slices.Contains(d.DetectionPattern, v) {
		return nil
	}

	d.DetectionPattern = append(d.DetectionPattern, v)

	return nil
}

// SetAnyDetectionPatternElement добавляет некоторое значение в список detection_pattern
func (d *BiZoneIRPData) SetAnyDetectionPatternElement(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetDetectionPatternElement(v)
}

// GetUnmappedAgentArray для поля unmapped_agent_array
func (d *BiZoneIRPData) GetUnmappedAgentArray() []uint64 {
	return d.UnmappedAgentArray
}

// SetUnmappedAgentArrayn для поля unmapped_agent_array
func (d *BiZoneIRPData) SetUnmappedAgentArrayn(v []uint64) error {
	d.UnmappedAgentArray = v

	return nil
}

// SetUnmappedAgentArrayElement добавляет значение unmapped_agent_array в список
func (d *BiZoneIRPData) SetUnmappedAgentArrayElement(v uint64) error {
	if d.UnmappedAgentArray == nil {
		d.UnmappedAgentArray = []uint64(nil)
	}

	if slices.Contains(d.UnmappedAgentArray, v) {
		return nil
	}

	d.UnmappedAgentArray = append(d.UnmappedAgentArray, v)

	return nil
}

// SetAnyUnmappedAgentArrayElement добавляет некоторое значение в список unmapped_agent_array
func (d *BiZoneIRPData) SetAnyUnmappedAgentArrayElement(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return d.SetDetectionPatternElement(v)
}

// GetDataSecurity для поля data_security
func (d *BiZoneIRPData) GetDataSecurity() []BiZoneIRPDataSecurity {
	return d.DataSecurity
}

// SetDataSecurity для поля data_security
func (d *BiZoneIRPData) SetDataSecurity(v []BiZoneIRPDataSecurity) error {
	d.DataSecurity = v

	return nil
}

// SetDataSecurityElement добавляет значение data_security в список
func (d *BiZoneIRPData) SetDataSecurityElement(v BiZoneIRPDataSecurity) error {
	if d.DataSecurity == nil {
		d.DataSecurity = []BiZoneIRPDataSecurity(nil)
	}

	d.DataSecurity = append(d.DataSecurity, v)

	return nil
}

// ToStringBeautiful форматированный вывод
func (d *BiZoneIRPData) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	fmt.Fprintf(&str, "%s'desc': '%s'\n", ws, d.Desc)
	fmt.Fprintf(&str, "%s'event_uid': '%s'\n", ws, d.EventUID)
	fmt.Fprintf(&str, "%s'job_title': '%s'\n", ws, d.JobTitle)
	fmt.Fprintf(&str, "%s'first_seen_time': '%s'\n", ws, d.FirstSeenTime)
	fmt.Fprintf(&str, "%s'last_seen_time': '%s'\n", ws, d.LastSeenTime)
	fmt.Fprintf(&str, "%s'metadata_product_name': '%s'\n", ws, d.MetadataProductName)
	fmt.Fprintf(&str, "%s'source_ip': '%s'\n", ws, d.SourceIP)
	fmt.Fprintf(&str, "%s'target_ip': '%s'\n", ws, d.TargetIP)
	fmt.Fprintf(&str, "%s'unmapped_hive_alert_id': '%s'\n", ws, d.GetUnmappedHiveAlertID())
	fmt.Fprintf(&str, "%s'unmapped_sensor_ip': '%s'\n", ws, d.UnmappedSensorIP)
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
		fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	}

	return str.String()
}
