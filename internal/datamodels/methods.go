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

// SetSensorInformation добавляет информацию о сенсорах
func (ai *AdditionalInformation) SetSensorInformation(v []SensorInformation) {
	ai.Sensors = append(ai.Sensors, v...)
}

// GetIpAddressesInformation объекты с информацией об ip адресах
func (ai *AdditionalInformation) GetIpAddressesInformation() []IpAddressInformation {
	return ai.IpAddresses
}

// SetIpAddressesInformation добавляет информацию об ip адресах
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
	str.WriteString(fmt.Sprintf("%s'@sensor_additional_information':\n", supportingfunctions.GetWhitespace(num)))
	for k, v := range ai.Sensors {
		str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
		str.WriteString(v.ToStringBeautiful(num + 2))
	}
	str.WriteString(fmt.Sprintf("%s'@ip_address_additional_information':\n", supportingfunctions.GetWhitespace(num)))
	for k, v := range ai.IpAddresses {
		str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
		str.WriteString(v.ToStringBeautiful(num + 2))
	}

	return str.String()
}

// ToStringBeautiful для информации по сенсору
func (si *SensorInformation) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'sensor_id': '%s'\n", ws, si.SensorId))
	str.WriteString(fmt.Sprintf("%s'host_id': '%s'\n", ws, si.HostId))
	str.WriteString(fmt.Sprintf("%s'geo_code': '%s'\n", ws, si.GeoCode))
	str.WriteString(fmt.Sprintf("%s'object_area': '%s'\n", ws, si.ObjectArea))
	str.WriteString(fmt.Sprintf("%s'subject_rf': '%s'\n", ws, si.SubjectRF))
	str.WriteString(fmt.Sprintf("%s'inn': '%s'\n", ws, si.INN))
	str.WriteString(fmt.Sprintf("%s'home_net': '%s'\n", ws, si.HomeNet))
	str.WriteString(fmt.Sprintf("%s'org_name': '%s'\n", ws, si.OrgName))
	str.WriteString(fmt.Sprintf("%s'full_org_name': '%s'\n", ws, si.FullOrgName))

	return str.String()
}

// ToStringBeautiful для информации по ip адресу
func (i *IpAddressInformation) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'ip': '%s'\n", ws, i.Ip))
	str.WriteString(fmt.Sprintf("%s'city': '%s'\n", ws, i.City))
	str.WriteString(fmt.Sprintf("%s'country': '%s'\n", ws, i.Country))
	str.WriteString(fmt.Sprintf("%s'country_code': '%s'\n", ws, i.CountryCode))

	return str.String()
}
