package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPOrganization() *BiZoneIRPOrganization {
	return &BiZoneIRPOrganization{}
}

// GetID для поля 'id'
func (o *BiZoneIRPOrganization) GetID() string {
	return o.ID
}

// SetID для поля 'id'
func (o *BiZoneIRPOrganization) SetID(v string) error {
	o.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (o *BiZoneIRPOrganization) SetAnyID(a any) error {
	return o.SetID(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (o *BiZoneIRPOrganization) GetTitle() string {
	return o.Title
}

// SetTitle для поля 'title'
func (o *BiZoneIRPOrganization) SetTitle(v string) error {
	o.Title = v

	return nil
}

// SetAnyTitle для поля 'title'
func (o *BiZoneIRPOrganization) SetAnyTitle(a any) error {
	return o.SetTitle(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (o *BiZoneIRPOrganization) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, o.ID)
	fmt.Fprintf(&str, "%s'title': '%s'\n", ws, o.Title)

	return str.String()
}
