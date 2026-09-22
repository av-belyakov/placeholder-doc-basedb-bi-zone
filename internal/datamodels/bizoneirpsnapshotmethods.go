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

// GetIPAddresses для поля 'ip_addresses'
func (s *BiZoneIRPSnapshot) GetIPAddresses() []string {
	return s.IPAddresses
}

// SetIPAddresses для поля 'ip_addresses'
func (s *BiZoneIRPSnapshot) SetIPAddresses(v []string) error {
	s.IPAddresses = v

	return nil
}

// SetIPAddresse добавляет IP адрес в поле 'ip_addresses'
func (s *BiZoneIRPSnapshot) SetIPAddresse(v string) error {
	if s.IPAddresses == nil {
		s.IPAddresses = []string(nil)
	}

	if slices.Contains(s.IPAddresses, v) {
		return nil
	}

	s.IPAddresses = append(s.IPAddresses, v)

	return nil
}

// SetAnyIPAddresse добавляет IP адрес в поле 'ip_addresses'
func (s *BiZoneIRPSnapshot) SetAnyIPAddresse(a any) error {
	return s.SetIPAddresse(fmt.Sprint(a))
}

// GetMACAddresses для поля 'mac_addresses'
func (s *BiZoneIRPSnapshot) GetMACAddresses() []string {
	return s.MACAddresses
}

// SetMACAddresses для поля 'mac_addresses'
func (s *BiZoneIRPSnapshot) SetMACAddresses(v []string) error {
	s.MACAddresses = v

	return nil
}

// SetMACAddresse добавляет MAC адрес в поле 'mac_addresses'
func (s *BiZoneIRPSnapshot) SetMACAddresse(v string) error {
	if s.MACAddresses == nil {
		s.MACAddresses = []string(nil)
	}

	if slices.Contains(s.MACAddresses, v) {
		return nil
	}

	s.MACAddresses = append(s.MACAddresses, v)

	return nil
}

// SetAnyMACAddresse добавляет MAC адрес в поле 'mac_addresses'
func (s *BiZoneIRPSnapshot) SetAnyMACAddresse(a any) error {
	return s.SetMACAddresse(fmt.Sprint(a))
}

// GetDomain для поля 'domain'
func (s *BiZoneIRPSnapshot) GetDomain() string {
	return s.Domain
}

// SetDomain для поля 'domain'
func (s *BiZoneIRPSnapshot) SetDomain(v string) error {
	s.Domain = v

	return nil
}

// SetAnyDomain для поля 'domain'
func (s *BiZoneIRPSnapshot) SetAnyDomain(a any) error {
	return s.SetDomain(fmt.Sprint(a))
}

// GetFqdn для поля 'fqdn'
func (s *BiZoneIRPSnapshot) GetFqdn() string {
	return s.Fqdn
}

// SetFqdn для поля 'fqdn'
func (s *BiZoneIRPSnapshot) SetFqdn(v string) error {
	s.Fqdn = v

	return nil
}

// SetAnyFqdn для поля 'fqdn'
func (s *BiZoneIRPSnapshot) SetAnyFqdn(a any) error {
	return s.SetFqdn(fmt.Sprint(a))
}

// GetOS для поля 'os'
func (s *BiZoneIRPSnapshot) GetOS() string {
	return s.OS
}

// SetOS для поля 'os'
func (s *BiZoneIRPSnapshot) SetOS(v string) error {
	s.OS = v

	return nil
}

// SetAnyOS для поля 'os'
func (s *BiZoneIRPSnapshot) SetAnyOS(a any) error {
	return s.SetOS(fmt.Sprint(a))
}

// GetUserCMDBName для поля 'user_cmdb_name'
func (s *BiZoneIRPSnapshot) GetUserCMDBName() string {
	return s.UserCmdbName
}

// SetUserCMDBName для поля 'user_cmdb_name'
func (s *BiZoneIRPSnapshot) SetUserCMDBName(v string) error {
	s.UserCmdbName = v

	return nil
}

// SetAnyUserCMDBName для поля 'user_cmdb_name'
func (s *BiZoneIRPSnapshot) SetAnyUserCMDBName(a any) error {
	return s.SetUserCMDBName(fmt.Sprint(a))
}

// GetCMDBID для поля 'cmdb_id'
func (s *BiZoneIRPSnapshot) GetCMDBID() string {
	return s.CmdbId
}

// SetCMDBID для поля 'cmdb_id'
func (s *BiZoneIRPSnapshot) SetCMDBID(v string) error {
	s.CmdbId = v

	return nil
}

// SetAnyCMDBID для поля 'cmdb_id'
func (s *BiZoneIRPSnapshot) SetAnyCMDBID(a any) error {
	return s.SetCMDBID(fmt.Sprint(a))
}

// GetHostname для поля 'hostname'
func (s *BiZoneIRPSnapshot) GetHostname() string {
	return s.Hostname
}

// SetHostname для поля 'hostname'
func (s *BiZoneIRPSnapshot) SetHostname(v string) error {
	s.Hostname = v

	return nil
}

// SetAnyHostname для поля 'hostname'
func (s *BiZoneIRPSnapshot) SetAnyHostname(a any) error {
	return s.SetHostname(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (s *BiZoneIRPSnapshot) GetTitle() string {
	return s.Title
}

// SetTitle для поля 'title'
func (s *BiZoneIRPSnapshot) SetTitle(v string) error {
	s.Title = v

	return nil
}

// SetAnyTitle для поля 'title'
func (s *BiZoneIRPSnapshot) SetAnyTitle(a any) error {
	return s.SetTitle(fmt.Sprint(a))
}

// GetSeverity для поля 'severity'
func (s *BiZoneIRPSnapshot) GetSeverity() string {
	return s.Severity
}

// SetSeverity для поля 'severity'
func (s *BiZoneIRPSnapshot) SetSeverity(v string) error {
	s.Severity = v

	return nil
}

// SetAnySeverity для поля 'severity'
func (s *BiZoneIRPSnapshot) SetAnySeverity(a any) error {
	return s.SetSeverity(fmt.Sprint(a))
}

// GetOSType для поля 'os_type'
func (s *BiZoneIRPSnapshot) GetOSType() string {
	return s.OSType
}

// SetOSType для поля 'os_type'
func (s *BiZoneIRPSnapshot) SetOSType(v string) error {
	s.OSType = v

	return nil
}

// SetAnyOSType для поля 'os_type'
func (s *BiZoneIRPSnapshot) SetAnyOSType(a any) error {
	return s.SetOSType(fmt.Sprint(a))
}

// GetUserCmdbId для поля 'user_cmdb_id'
func (s *BiZoneIRPSnapshot) GetUserCmdbId() string {
	return *s.UserCmdbId
}

// SetUserCmdbId для поля 'user_cmdb_id'
func (s *BiZoneIRPSnapshot) SetUserCmdbId(v string) error {
	s.UserCmdbId = &v

	return nil
}

// SetAnyUserCmdbId для поля 'user_cmdb_id'
func (s *BiZoneIRPSnapshot) SetAnyUserCmdbId(a any) error {
	return s.SetUserCmdbId(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (s *BiZoneIRPSnapshot) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'os': '%s'\n", ws, s.OS)
	fmt.Fprintf(&str, "%s'fqdn': '%s'\n", ws, s.Fqdn)
	fmt.Fprintf(&str, "%s'title': '%s'\n", ws, s.Title)
	fmt.Fprintf(&str, "%s'domain': '%s'\n", ws, s.Domain)
	fmt.Fprintf(&str, "%s'cmdb_id': '%s'\n", ws, s.CmdbId)
	fmt.Fprintf(&str, "%s'os_type': '%s'\n", ws, s.OSType)
	fmt.Fprintf(&str, "%s'hostname': '%s'\n", ws, s.Hostname)
	fmt.Fprintf(&str, "%s'severity': '%s'\n", ws, s.Severity)
	fmt.Fprintf(&str, "%s'user_cmdb_id': '%s'\n", ws, *s.UserCmdbId)
	fmt.Fprintf(&str, "%s'user_cmdb_name': '%s'\n", ws, s.UserCmdbName)
	fmt.Fprintf(&str, "%s'ip_addressess': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, s.IPAddresses))
	fmt.Fprintf(&str, "%s'mac_addressess': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, s.MACAddresses))

	return str.String()
}
