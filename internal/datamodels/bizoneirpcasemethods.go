package datamodels

import (
	"errors"
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ******* Основная структура Case ********
// *****************************************

// NewVerifiedBiZoneIRPCase новый объект case
func NewVerifiedBiZoneIRPCase() *VerifiedBiZoneIRPCase {
	return &VerifiedBiZoneIRPCase{}
}

// GetUUID для поля 'id' в формате UUID
func (c *VerifiedBiZoneIRPCase) GetUUID() string {
	return c.ID
}

// SetUUID для поля 'id' в формате UUID
func (c *VerifiedBiZoneIRPCase) SetUUID(v string) error {
	c.ID = v

	return nil
}

// SetAnyUUID для поля 'id' в формате UUID
func (c *VerifiedBiZoneIRPCase) SetAnyUUID(a any) error {
	return c.SetUUID(fmt.Sprint(a))
}

// GetType для поля 'type'
func (c *VerifiedBiZoneIRPCase) GetType() string {
	return c.Type
}

// SetType для поля 'type'
func (c *VerifiedBiZoneIRPCase) SetType(v string) error {
	c.Type = v

	return nil
}

// SetAnyType для поля 'type'
func (c *VerifiedBiZoneIRPCase) SetAnyType(a any) error {
	return c.SetType(fmt.Sprint(a))
}

// GetSource для поля 'source'
func (c *VerifiedBiZoneIRPCase) GetSource() string {
	return c.Source
}

// SetSource для поля 'source'
func (c *VerifiedBiZoneIRPCase) SetSource(v string) error {
	c.Source = v

	return nil
}

// SetAnySource для поля 'source'
func (c *VerifiedBiZoneIRPCase) SetAnySource(a any) error {
	return c.SetSource(fmt.Sprint(a))
}

// GetSpecVersion для поля 'specversion'
func (c *VerifiedBiZoneIRPCase) GetSpecVersion() string {
	return c.SpecVersion
}

// SetSpecVersion для поля 'specversion'
func (c *VerifiedBiZoneIRPCase) SetSpecVersion(v string) error {
	c.SpecVersion = v

	return nil
}

// SetAnySpecVersion для поля 'specversion'
func (c *VerifiedBiZoneIRPCase) SetAnySpecVersion(a any) error {
	return c.SetSpecVersion(fmt.Sprint(a))
}

// GetTime для поля 'time'
func (c *VerifiedBiZoneIRPCase) GetTime() string {
	return c.Time
}

// SetTime для поля 'time'
func (c *VerifiedBiZoneIRPCase) SetTime(v string) error {
	c.Time = v

	return nil
}

// SetAnyTime для поля 'time'
func (c *VerifiedBiZoneIRPCase) SetAnyTime(a any) error {
	return c.SetTime(fmt.Sprint(a))
}

// GetSubject для поля 'subject'
func (c *VerifiedBiZoneIRPCase) GetSubject() *string {
	return c.Subject
}

// SetSubject для поля 'subject'
func (c *VerifiedBiZoneIRPCase) SetSubject(v *string) error {
	c.Subject = v

	return nil
}

// SetAnySubject для поля 'subject'
func (c *VerifiedBiZoneIRPCase) SetAnySubject(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'subject'")
	}

	return c.SetSubject(&v)
}

// GetDataSchema для поля 'dataschema'
func (c *VerifiedBiZoneIRPCase) GetDataSchema() *string {
	return c.DataSchema
}

// SetDataSchema для поля 'dataschema'
func (c *VerifiedBiZoneIRPCase) SetDataSchema(v *string) error {
	c.DataSchema = v

	return nil
}

// SetAnyDataSchema для поля 'dataschema'
func (c *VerifiedBiZoneIRPCase) SetAnyDataSchema(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'dataSchema'")
	}

	return c.SetDataSchema(&v)
}

// GetData для поля 'data'
func (c *VerifiedBiZoneIRPCase) GetData() BiZoneIRPCaseData {
	return c.Data
}

// SetData для поля 'data'
func (c *VerifiedBiZoneIRPCase) SetData(v BiZoneIRPCaseData) error {
	c.Data = v

	return nil
}

// ToStringBeautiful форматированный вывод
func (c *VerifiedBiZoneIRPCase) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, c.ID)
	fmt.Fprintf(&str, "%s'type': '%s'\n", ws, c.Type)
	fmt.Fprintf(&str, "%s'source': '%s'\n", ws, c.Source)
	fmt.Fprintf(&str, "%s'specversion': '%s'\n", ws, c.SpecVersion)
	fmt.Fprintf(&str, "%s'time': '%s'\n", ws, c.Time)
	fmt.Fprintf(&str, "%s'subject': '%s'\n", ws, *c.Subject)
	fmt.Fprintf(&str, "%s'dataschema': '%s'\n", ws, *c.DataSchema)
	fmt.Fprintf(&str, "%s'data':\n%s", ws, c.Data.ToStringBeautiful(num+1))

	return str.String()
}
