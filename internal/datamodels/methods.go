package datamodels

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

var (
	fieldsFroTagsRepresentedAsList []string = []string{
		"tags.name",
		"tags.color",
		"tags.created",
		"tags.created_by.id",
		"tags.created_by.username",
		"tags.is_visible_for_customer",
	}

	fieldsForSnapshotsRepresentedAsList []string = []string{
		"os",
		"fqdn",
		"domain",
		"cmdb_id",
		"os_type",
		"hostname",
		"ip_addresses",
		"mac_addresses",
		"user_cmdb_name",
	}
)

// GetSensorsInformation объекты с информацией о сенсоре
func (ai *AdditionalInformation) GetSensorsInformation() []SensorInformation {
	return ai.Sensors
}

// GetSensorId идентификатор сенсора
func (si *SensorInformation) GetSensorId() string {
	return si.SensorId
}

// AddSensorInformation добавляет информацию о сенсоре
func (ai *AdditionalInformation) AddSensorInformation(v SensorInformation) {
	if len(ai.Sensors) == 0 || !slices.ContainsFunc(ai.Sensors, func(obj SensorInformation) bool {
		return obj.SensorId == v.SensorId
	}) {
		ai.Sensors = append(ai.Sensors, v)
	}
}

// SetSensorInformation информация о сенсорах
func (ai *AdditionalInformation) SetSensorInformation(v []SensorInformation) {
	ai.Sensors = append(ai.Sensors, v...)
}

// GetSpecialUUID специальный уникальный идентификатор
func (ai *AdditionalInformation) GetSpecialUUID() string {
	return ai.SpecialUUID
}

// SetSpecialUUID специальный уникальный идентификатор
func (ai *AdditionalInformation) SetSpecialUUID(v string) {
	ai.SpecialUUID = v
}

// GetIpAddressesInformation объекты с информацией об ip адресах
func (ai *AdditionalInformation) GetIpAddressesInformation() []IpAddressInformation {
	return ai.IpAddresses
}

// SetIpAddressesInformation объекты с информацией об ip адресах
func (ai *AdditionalInformation) SetIpAddressesInformation(v []IpAddressInformation) {
	ai.IpAddresses = append(ai.IpAddresses, v...)
}

// GetIpAddrString ip адрес в виде строки
func (i IpAddressInformation) GetIpAddrString() string {
	return i.Ip
}

// AddIpAddressInformation добавляет информацию об ip адресе
func (ai *AdditionalInformation) AddIpAddressInformation(v IpAddressInformation) {
	if len(ai.IpAddresses) == 0 || !slices.ContainsFunc(ai.IpAddresses, func(obj IpAddressInformation) bool {
		return obj.Ip == v.Ip
	}) {
		ai.IpAddresses = append(ai.IpAddresses, v)
	}
}

// ToStringBeautiful дополнительная информация по сенсорам и ip адресам
func (ai *AdditionalInformation) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	fmt.Fprintf(&str, "%s'@special_uuid': '%s'\n", supportingfunctions.GetWhitespace(num), ai.SpecialUUID)
	fmt.Fprintf(&str, "%s'@sensor_additional_information':\n", supportingfunctions.GetWhitespace(num))
	for k, v := range ai.Sensors {
		fmt.Fprintf(&str, "%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1)
		str.WriteString(v.ToStringBeautiful(num + 2))
	}
	fmt.Fprintf(&str, "%s'@ip_address_additional_information':\n", supportingfunctions.GetWhitespace(num))
	for k, v := range ai.IpAddresses {
		fmt.Fprintf(&str, "%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1)
		str.WriteString(v.ToStringBeautiful(num + 2))
	}

	return str.String()
}

// ToStringBeautiful для информации по сенсору
func (si *SensorInformation) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	fmt.Fprintf(&str, "%s'sensor_id': '%s'\n", ws, si.SensorId)
	fmt.Fprintf(&str, "%s'host_id': '%s'\n", ws, si.HostId)
	fmt.Fprintf(&str, "%s'geo_code': '%s'\n", ws, si.GeoCode)
	fmt.Fprintf(&str, "%s'object_area': '%s'\n", ws, si.ObjectArea)
	fmt.Fprintf(&str, "%s'subject_rf': '%s'\n", ws, si.SubjectRF)
	fmt.Fprintf(&str, "%s'inn': '%s'\n", ws, si.INN)
	fmt.Fprintf(&str, "%s'home_net': '%s'\n", ws, si.HomeNet)
	fmt.Fprintf(&str, "%s'org_name': '%s'\n", ws, si.OrgName)
	fmt.Fprintf(&str, "%s'full_org_name': '%s'\n", ws, si.FullOrgName)

	return str.String()
}

// ToStringBeautiful для информации по ip адресу
func (i *IpAddressInformation) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	fmt.Fprintf(&str, "%s'ip': '%s'\n", ws, i.Ip)
	fmt.Fprintf(&str, "%s'city': '%s'\n", ws, i.City)
	fmt.Fprintf(&str, "%s'country': '%s'\n", ws, i.Country)
	fmt.Fprintf(&str, "%s'country_code': '%s'\n", ws, i.CountryCode)

	return str.String()
}
