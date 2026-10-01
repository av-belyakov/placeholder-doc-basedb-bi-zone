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
func (t *BiZoneIRPTenant) GetID() string {
	return t.ID
}

// SetID для поля 'id'
func (t *BiZoneIRPTenant) SetID(v string) error {
	t.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (t *BiZoneIRPTenant) SetAnyID(a any) error {
	return t.SetID(fmt.Sprint(a))
}

// GetName для поля 'name'
func (t *BiZoneIRPTenant) GetName() string {
	return t.Name
}

// SetName для поля 'name'
func (t *BiZoneIRPTenant) SetName(v string) error {
	t.Name = v

	return nil
}

// SetAnyName для поля 'name'
func (t *BiZoneIRPTenant) SetAnyName(a any) error {
	return t.SetName(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (t *BiZoneIRPTenant) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, t.ID)
	fmt.Fprintf(&str, "%s'name': '%s'\n", ws, t.Name)

	return str.String()
}
