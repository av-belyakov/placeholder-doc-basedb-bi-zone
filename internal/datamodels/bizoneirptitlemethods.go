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
func (t *BiZoneIRPTitle) GetValue() string {
	return t.Value
}

// SetValue для поля 'value'
func (t *BiZoneIRPTitle) SetValue(v string) error {
	t.Value = v

	return nil
}

// SetAnyValue для поля 'value'
func (t *BiZoneIRPTitle) SetAnyValue(a any) error {
	return t.SetValue(fmt.Sprint(a))
}

// GetLanguage для поля 'language'
func (t *BiZoneIRPTitle) GetLanguage() string {
	return t.Language
}

// SetLanguage для поля 'language'
func (t *BiZoneIRPTitle) SetLanguage(v string) error {
	t.Language = v

	return nil
}

// SetAnyLanguage для поля 'language'
func (t *BiZoneIRPTitle) SetAnyLanguage(a any) error {
	return t.SetLanguage(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (t *BiZoneIRPTitle) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'value': '%s'\n", ws, t.Value)
	fmt.Fprintf(&str, "%s'language': '%s'\n", ws, t.Language)

	return str.String()
}
