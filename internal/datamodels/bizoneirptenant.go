package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPTenant() *BiZoneIRPTenant {
	return &BiZoneIRPTenant{}
}

// GetID для поля 'id'
func (tnt *BiZoneIRPTenant) GetID() string {
	return tnt.ID
}

// SetID для поля 'id'
func (tnt *BiZoneIRPTenant) SetID(v string) error {
	tnt.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (tnt *BiZoneIRPTenant) SetAnyID(a any) error {
	return tnt.SetID(fmt.Sprint(a))
}

// GetName для поля 'name'
func (tnt *BiZoneIRPTenant) GetName() string {
	return tnt.Name
}

// SetName для поля 'name'
func (tnt *BiZoneIRPTenant) SetName(v string) error {
	tnt.Name = v

	return nil
}

// SetAnyName для поля 'name'
func (tnt *BiZoneIRPTenant) SetAnyName(a any) error {
	return tnt.SetName(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (tnt *BiZoneIRPTenant) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, tnt.ID)
	fmt.Fprintf(&str, "%s'name': '%s'\n", ws, tnt.Name)

	return str.String()
}
