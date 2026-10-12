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

func (vc *VerifiedBiZoneIRPCase) Get() *VerifiedBiZoneIRPCase {
	return vc
}

// GetID для поля 'id' в формате UUID
func (vc *VerifiedBiZoneIRPCase) GetID() string {
	return vc.ID
}

// SetID для поля 'id' в формате UUID
func (vc *VerifiedBiZoneIRPCase) SetID(v string) error {
	vc.ID = v

	return nil
}

// SetAnyID для поля 'id' в формате ID
func (vc *VerifiedBiZoneIRPCase) SetAnyUUID(a any) error {
	return vc.SetID(fmt.Sprint(a))
}

// GetType для поля 'type'
func (vc *VerifiedBiZoneIRPCase) GetType() string {
	return vc.Type
}

// SetType для поля 'type'
func (vc *VerifiedBiZoneIRPCase) SetType(v string) error {
	vc.Type = v

	return nil
}

// SetAnyType для поля 'type'
func (vc *VerifiedBiZoneIRPCase) SetAnyType(a any) error {
	return vc.SetType(fmt.Sprint(a))
}

// GetSource для поля 'source'
func (vc *VerifiedBiZoneIRPCase) GetSource() string {
	return vc.Source
}

// SetSource для поля 'source'
func (vc *VerifiedBiZoneIRPCase) SetSource(v string) error {
	vc.Source = v

	return nil
}

// SetAnySource для поля 'source'
func (vc *VerifiedBiZoneIRPCase) SetAnySource(a any) error {
	return vc.SetSource(fmt.Sprint(a))
}

// GetSpecVersion для поля 'specversion'
func (vc *VerifiedBiZoneIRPCase) GetSpecVersion() string {
	return vc.SpecVersion
}

// SetSpecVersion для поля 'specversion'
func (vc *VerifiedBiZoneIRPCase) SetSpecVersion(v string) error {
	vc.SpecVersion = v

	return nil
}

// SetAnySpecVersion для поля 'specversion'
func (vc *VerifiedBiZoneIRPCase) SetAnySpecVersion(a any) error {
	return vc.SetSpecVersion(fmt.Sprint(a))
}

// GetTime для поля 'time'
func (vc *VerifiedBiZoneIRPCase) GetTime() string {
	return vc.Time
}

// SetTime для поля 'time'
func (vc *VerifiedBiZoneIRPCase) SetTime(v string) error {
	vc.Time = v

	return nil
}

// SetAnyTime для поля 'time'
func (vc *VerifiedBiZoneIRPCase) SetAnyTime(a any) error {
	return vc.SetTime(fmt.Sprint(a))
}

// GetSubject для поля 'subject'
func (vc *VerifiedBiZoneIRPCase) GetSubject() *string {
	return vc.Subject
}

// SetSubject для поля 'subject'
func (vc *VerifiedBiZoneIRPCase) SetSubject(v *string) error {
	vc.Subject = v

	return nil
}

// SetAnySubject для поля 'subject'
func (vc *VerifiedBiZoneIRPCase) SetAnySubject(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'subject'")
	}

	return vc.SetSubject(&v)
}

// GetDataSchema для поля 'dataschema'
func (vc *VerifiedBiZoneIRPCase) GetDataSchema() *string {
	return vc.DataSchema
}

// SetDataSchema для поля 'dataschema'
func (vc *VerifiedBiZoneIRPCase) SetDataSchema(v *string) error {
	vc.DataSchema = v

	return nil
}

// SetAnyDataSchema для поля 'dataschema'
func (vc *VerifiedBiZoneIRPCase) SetAnyDataSchema(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'dataSchema'")
	}

	return vc.SetDataSchema(&v)
}

// GetData для поля 'data'
func (vc *VerifiedBiZoneIRPCase) GetData() BiZoneIRPCaseData {
	return vc.Data
}

// SetData для поля 'data'
func (vc *VerifiedBiZoneIRPCase) SetData(v BiZoneIRPCaseData) error {
	vc.Data = v

	return nil
}

// GetAdditionalInformation поле дополнительной информации
func (vc *VerifiedBiZoneIRPCase) GetAdditionalInformation() *AdditionalInformation {
	return &vc.AdditionalInformation
}

// SetAdditionalInformation для поля дополнительной информации
func (vc *VerifiedBiZoneIRPCase) SetAdditionalInformation(v AdditionalInformation) error {
	vc.AdditionalInformation = v

	return nil
}

// ToStringBeautiful форматированный вывод
func (vc *VerifiedBiZoneIRPCase) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, vc.ID)
	fmt.Fprintf(&str, "%s'type': '%s'\n", ws, vc.Type)
	fmt.Fprintf(&str, "%s'source': '%s'\n", ws, vc.Source)
	fmt.Fprintf(&str, "%s'specversion': '%s'\n", ws, vc.SpecVersion)
	fmt.Fprintf(&str, "%s'time': '%s'\n", ws, vc.Time)
	fmt.Fprintf(&str, "%s'subject': '%s'\n", ws, *vc.Subject)
	fmt.Fprintf(&str, "%s'dataschema': '%s'\n", ws, *vc.DataSchema)
	fmt.Fprintf(&str, "%s'data':\n%s", ws, vc.Data.ToStringBeautiful(num+1))
	fmt.Fprintf(&str, "%s'@additional_information':\n%s", ws, vc.AdditionalInformation.ToStringBeautiful(num+1))

	return str.String()
}
