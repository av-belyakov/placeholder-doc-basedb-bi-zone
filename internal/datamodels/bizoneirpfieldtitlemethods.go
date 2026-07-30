package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPFieldTitle() *BiZoneIRPFieldTitle {
	return &BiZoneIRPFieldTitle{}
}

// GetValue
func (ft *BiZoneIRPFieldTitle) GetValue() string {
	return ft.Value
}

// SetValue
func (ft *BiZoneIRPFieldTitle) SetValue(v string) {
	ft.Value = v
}

// SetAnyValue
func (ft *BiZoneIRPFieldTitle) SetAnyValue(a any) {
	ft.SetValue(fmt.Sprint(a))
}

// GetLanguage
func (ft *BiZoneIRPFieldTitle) GetLanguage() string {
	return ft.Language
}

// SetLanguage
func (ft *BiZoneIRPFieldTitle) SetLanguage(v string) {
	ft.Language = v
}

// SetAnyLanguage
func (ft *BiZoneIRPFieldTitle) SetAnyLanguager(a any) {
	ft.SetLanguage(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (ft *BiZoneIRPFieldTitle) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'value': '%s'\n", ws, ft.Value))
	str.WriteString(fmt.Sprintf("%s'language': '%s'\n", ws, ft.Language))

	return str.String()
}
