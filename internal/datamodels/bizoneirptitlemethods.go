package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPTitle() *BiZoneIRPTitle {
	return &BiZoneIRPTitle{
		Language: "ru",
	}
}

// GetValue для поля 'value'
func (tt *BiZoneIRPTitle) GetValue() string {
	return tt.Value
}

// SetValue для поля 'value'
func (tt *BiZoneIRPTitle) SetValue(v string) error {
	tt.Value = v

	return nil
}

// SetAnyValue для поля 'value'
func (tt *BiZoneIRPTitle) SetAnyValue(a any) error {
	return tt.SetValue(fmt.Sprint(a))
}

// GetLanguage для поля 'language'
func (tt *BiZoneIRPTitle) GetLanguage() string {
	return tt.Language
}

// SetLanguage для поля 'language'
func (tt *BiZoneIRPTitle) SetLanguage(v string) error {
	tt.Language = v

	return nil
}

// SetAnyLanguage для поля 'language'
func (tt *BiZoneIRPTitle) SetAnyLanguage(a any) error {
	return tt.SetLanguage(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (tt *BiZoneIRPTitle) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'value': '%s'\n", ws, tt.Value)
	fmt.Fprintf(&str, "%s'language': '%s'\n", ws, tt.Language)

	return str.String()
}
