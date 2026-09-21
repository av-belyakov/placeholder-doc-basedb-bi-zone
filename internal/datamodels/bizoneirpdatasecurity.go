package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPDataSecurity() *BiZoneIRPDataSecurity {
	return &BiZoneIRPDataSecurity{
		SContent: []BiZoneIRPSContent(nil),
	}
}

func (ds *BiZoneIRPDataSecurity) Get() *BiZoneIRPDataSecurity {
	return ds
}

// GetSRuleBody для поля s_rule_body
func (ds *BiZoneIRPDataSecurity) GetSRuleBody() string {
	return ds.SRuleBody
}

// SetSRuleBody для поля s_rule_body
func (ds *BiZoneIRPDataSecurity) SetSRuleBody(v string) error {
	ds.SRuleBody = v

	return nil
}

// SetAnySRuleBody для поля s_rule_body
func (ds *BiZoneIRPDataSecurity) SetAnySRuleBody(a any) error {
	return ds.SetSRuleBody(fmt.Sprint(a))
}

// GetSClasstype для поля s_classtype
func (ds *BiZoneIRPDataSecurity) GetSClasstype() string {
	return ds.SClasstype
}

// SetSClasstype для поля s_classtype
func (ds *BiZoneIRPDataSecurity) SetSClasstype(v string) error {
	ds.SClasstype = v

	return nil
}

// SetAnySClasstype для поля s_classtype
func (ds *BiZoneIRPDataSecurity) SetAnySClasstype(a any) error {
	return ds.SetSClasstype(fmt.Sprint(a))
}

// GetSMsg для поля s_msg
func (ds *BiZoneIRPDataSecurity) GetSMsg() string {
	return ds.SMsg
}

// SetSMsg для поля s_msg
func (ds *BiZoneIRPDataSecurity) SetSMsg(v string) error {
	ds.SMsg = v

	return nil
}

// SetAnySMsg для поля s_msg
func (ds *BiZoneIRPDataSecurity) SetAnySMsg(a any) error {
	return ds.SetSMsg(fmt.Sprint(a))
}

// GetSSourceName для поля s_source_name
func (ds *BiZoneIRPDataSecurity) GetSSourceName() string {
	return ds.SSourceName
}

// SetSSourceName для поля s_source_name
func (ds *BiZoneIRPDataSecurity) SetSSourceName(v string) error {
	ds.SSourceName = v

	return nil
}

// SetAnySSourceName для поля s_source_name
func (ds *BiZoneIRPDataSecurity) SetAnySSourceName(a any) error {
	return ds.SetSSourceName(fmt.Sprint(a))
}

// GetISid для поля i_sid
func (ds *BiZoneIRPDataSecurity) GetISid() uint64 {
	return ds.ISid
}

// SetISid для поля i_sid
func (ds *BiZoneIRPDataSecurity) SetISid(v uint64) error {
	ds.ISid = v

	return nil
}

// SetAnyISid для поля i_sid
func (ds *BiZoneIRPDataSecurity) SetAnyISid(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return ds.SetISid(v)
}

// GetIRev для поля i_rev
func (ds *BiZoneIRPDataSecurity) GetIRev() uint64 {
	return ds.IRev
}

// SetIRev для поля i_rev
func (ds *BiZoneIRPDataSecurity) SetIRev(v uint64) error {
	ds.IRev = v

	return nil
}

// SetAnyIRev для поля i_rev
func (ds *BiZoneIRPDataSecurity) SetAnyIRev(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return ds.SetIRev(v)
}

// GetIBRec для поля i_b_rec
func (ds *BiZoneIRPDataSecurity) GetIBRec() uint64 {
	return ds.IBRec
}

// SetIBRec для поля i_b_rec
func (ds *BiZoneIRPDataSecurity) SetIBRec(v uint64) error {
	ds.IBRec = v

	return nil
}

// SetAnyIBRec для поля i_b_rec
func (ds *BiZoneIRPDataSecurity) SetAnyIBRec(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return ds.SetIBRec(v)
}

// GetSAddDate для поля s_add_date
func (ds *BiZoneIRPDataSecurity) GetSAddDate() uint64 {
	return ds.SAddDate
}

// SetSAddDate для поля s_add_date
func (ds *BiZoneIRPDataSecurity) SetSAddDate(v uint64) error {
	ds.SAddDate = v

	return nil
}

// SetAnySAddDate для поля s_add_date
func (ds *BiZoneIRPDataSecurity) SetAnySAddDate(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return ds.SetSAddDate(v)
}

// GetIPriority для поля i_priority
func (ds *BiZoneIRPDataSecurity) GetIPriority() uint64 {
	return ds.IPriority
}

// SetIPriority для поля ai_prioritygent
func (ds *BiZoneIRPDataSecurity) SetIPriority(v uint64) error {
	ds.IPriority = v

	return nil
}

// SetAnyIPriority для поля i_priority
func (ds *BiZoneIRPDataSecurity) SetAnyIPriority(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return ds.SetIPriority(v)
}

// ToStringBeautiful форматированный вывод
func (ds *BiZoneIRPDataSecurity) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	fmt.Fprintf(&str, "%s's_rule_body': '%s'\n", ws, ds.SRuleBody)
	fmt.Fprintf(&str, "%s's_classtype': '%s'\n", ws, ds.SClasstype)
	fmt.Fprintf(&str, "%s's_msg': '%s'\n", ws, ds.SMsg)
	fmt.Fprintf(&str, "%s's_source_name': '%s'\n", ws, ds.SSourceName)
	fmt.Fprintf(&str, "%s'i_sid': '%d'\n", ws, ds.ISid)
	fmt.Fprintf(&str, "%s'i_rev': '%d'\n", ws, ds.IRev)
	fmt.Fprintf(&str, "%s'i_b_rec': '%d'\n", ws, ds.IBRec)
	fmt.Fprintf(&str, "%s's_add_date': '%d'\n", ws, ds.SAddDate)
	fmt.Fprintf(&str, "%s'i_priority': '%d'\n", ws, ds.IPriority)
	fmt.Fprintf(&str, "%s's_content':\n", ws)
	for k, v := range ds.SContent {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	}
	return str.String()
}
