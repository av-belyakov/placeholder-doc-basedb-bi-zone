package datamodels

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPSnapshot() *BiZoneIRPSnapshot {
	return &BiZoneIRPSnapshot{}
}

// GetIPAddresses для поля IPAddresses
func (s *BiZoneIRPSnapshot) GetIPAddresses() []string {
	return s.IPAddresses
}

// SetIPAddresses для поля IPAddresses
func (s *BiZoneIRPSnapshot) SetIPAddresses(ipAddress []string) {
	s.IPAddresses = ipAddress
}

// SetIPAddresse добавляет IP адрес в поле IPAddresses
func (s *BiZoneIRPSnapshot) SetIPAddresse(ipAddress string) {
	if s.IPAddresses == nil {
		s.IPAddresses = []string(nil)
	}

	if slices.Contains(s.IPAddresses, ipAddress) {
		return
	}

	s.IPAddresses = append(s.IPAddresses, ipAddress)
}

// SetAnyIPAddresse добавляет IP адрес в поле IPAddresses
func (s *BiZoneIRPSnapshot) SetAnyIPAddresse(a any) {
	s.SetIPAddresse(fmt.Sprint(a))
}

// GetMACAddresses для поля MACAddresses
func (s *BiZoneIRPSnapshot) GetMACAddresses() []string {
	return s.MACAddresses
}

// SetMACAddresses для поля MACAddresses
func (s *BiZoneIRPSnapshot) SetMACAddresses(macAddress []string) {
	s.MACAddresses = macAddress
}

// SetMACAddresse добавляет MAC адрес в поле MACAddresses
func (s *BiZoneIRPSnapshot) SetMACAddresse(macAddress string) {
	if s.MACAddresses == nil {
		s.MACAddresses = []string(nil)
	}

	if slices.Contains(s.MACAddresses, macAddress) {
		return
	}

	s.MACAddresses = append(s.MACAddresses, macAddress)
}

// SetAnyMACAddresse добавляет MAC адрес в поле MACAddresses
func (s *BiZoneIRPSnapshot) SetAnyMACAddresse(a any) {
	s.SetMACAddresse(fmt.Sprint(a))
}

// GetDomain для поля Domain
func (s *BiZoneIRPSnapshot) GetDomain() string {
	return s.Domain
}

// SetDomain для поля Domain
func (s *BiZoneIRPSnapshot) SetDomain(domain string) {
	s.Domain = domain
}

// SetAnyDomain для поля Domain
func (s *BiZoneIRPSnapshot) SetAnyDomain(a any) {
	s.Domain = fmt.Sprint(a)
}

// GetFqdn для поля Fqdn
func (s *BiZoneIRPSnapshot) GetFqdn() string {
	return s.Fqdn
}

// SetFqdn для поля Fqdn
func (s *BiZoneIRPSnapshot) SetFqdn(fqdn string) {
	s.Fqdn = fqdn
}

// SetAnyFqdn для поля Fqdn
func (s *BiZoneIRPSnapshot) SetAnyFqdn(a any) {
	s.Fqdn = fmt.Sprint(a)
}

// GetOS для поля OS
func (s *BiZoneIRPSnapshot) GetOS() string {
	return s.OS
}

// SetOS для поля OS
func (s *BiZoneIRPSnapshot) SetOS(value string) {
	s.OS = value
}

// SetAnyOS для поля OS
func (s *BiZoneIRPSnapshot) SetAnyOS(a any) {
	s.OS = fmt.Sprint(a)
}

// GetUserCMDBName для поля UserCMDBName
func (s *BiZoneIRPSnapshot) GetUserCMDBName() string {
	return s.UserCMDBName
}

// SetUserCMDBName для поля UserCMDBName
func (s *BiZoneIRPSnapshot) SetUserCMDBName(value string) {
	s.UserCMDBName = value
}

// SetAnyUserCMDBName для поля UserCMDBName
func (s *BiZoneIRPSnapshot) SetAnyUserCMDBName(a any) {
	s.UserCMDBName = fmt.Sprint(a)
}

// GetCMDBID для поля CMDBID
func (s *BiZoneIRPSnapshot) GetCMDBID() string {
	return s.CMDBID
}

// SetCMDBID для поля CMDBID
func (s *BiZoneIRPSnapshot) SetCMDBID(value string) {
	s.CMDBID = value
}

// SetAnyCMDBID для поля CMDBID
func (s *BiZoneIRPSnapshot) SetAnyCMDBID(a any) {
	s.CMDBID = fmt.Sprint(a)
}

// GetHostname для поля Hostname
func (s *BiZoneIRPSnapshot) GetHostname() string {
	return s.Hostname
}

// SetHostname для поля Hostname
func (s *BiZoneIRPSnapshot) SetHostname(hostname string) {
	s.Hostname = hostname
}

// SetAnyHostname для поля Hostname
func (s *BiZoneIRPSnapshot) SetAnyHostname(a any) {
	s.Hostname = fmt.Sprint(a)
}

// GetOSType для поля OSType
func (s *BiZoneIRPSnapshot) GetOSType() string {
	return s.OSType
}

// SetOSType для поля OSType
func (s *BiZoneIRPSnapshot) SetOSType(osType string) {
	s.OSType = osType
}

// SetAnyOSType для поля OSType
func (s *BiZoneIRPSnapshot) SetAnyOSType(a any) {
	s.OSType = fmt.Sprint(a)
}

// ToStringBeautiful форматированный вывод
func (s *BiZoneIRPSnapshot) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'os': '%s'\n", ws, s.OS))
	str.WriteString(fmt.Sprintf("%s'fqdn': '%s'\n", ws, s.Fqdn))
	str.WriteString(fmt.Sprintf("%s'domain': '%s'\n", ws, s.Domain))
	str.WriteString(fmt.Sprintf("%s'cmdb_id': '%s'\n", ws, s.CMDBID))
	str.WriteString(fmt.Sprintf("%s'os_type': '%s'\n", ws, s.OSType))
	str.WriteString(fmt.Sprintf("%s'hostname': '%s'\n", ws, s.Hostname))
	str.WriteString(fmt.Sprintf("%s'user_cmdb_name': '%s'\n", ws, s.UserCMDBName))
	str.WriteString(fmt.Sprintf("%s'ip_addressess': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, s.IPAddresses)))
	str.WriteString(fmt.Sprintf("%s'mac_addressess': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, s.MACAddresses)))

	return str.String()
}
