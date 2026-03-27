package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPFieldTagDescription() *BiZoneIRPFieldTagDescription {
	return &BiZoneIRPFieldTagDescription{}
}

// GetName
func (td *BiZoneIRPFieldTagDescription) GetName() string {
	return td.Name
}

// SetName
func (td *BiZoneIRPFieldTagDescription) SetName(v string) {
	td.Name = v
}

// SetAnyName
func (td *BiZoneIRPFieldTagDescription) SetAnyName(a any) {
	td.SetName(fmt.Sprint(a))
}

// GetColor
func (td *BiZoneIRPFieldTagDescription) GetColor() string {
	return td.Color
}

// SetColor
func (td *BiZoneIRPFieldTagDescription) SetColor(v string) {
	td.Color = v
}

// SetAnyColor
func (td *BiZoneIRPFieldTagDescription) SetAnyColor(a any) {
	td.SetColor(fmt.Sprint(a))
}

// GetIsVisibleForCustomer
func (td *BiZoneIRPFieldTagDescription) GetIsVisibleForCustomer() bool {
	return td.IsVisibleForCustomer
}

// SetIsVisibleForCustomer
func (td *BiZoneIRPFieldTagDescription) SetIsVisibleForCustomer(v bool) {
	td.IsVisibleForCustomer = v
}

// SetAnyIsVisibleForCustomer
func (td *BiZoneIRPFieldTagDescription) SetAnyIsVisibleForCustomer(a any) {
	if v, ok := a.(bool); ok {
		td.SetIsVisibleForCustomer(v)
	}
}

// ToStringBeautiful форматированный вывод
func (td *BiZoneIRPFieldTagDescription) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, td.Name))
	str.WriteString(fmt.Sprintf("%s'color': '%s'\n", ws, td.Color))
	str.WriteString(fmt.Sprintf("%s'is_visible_for_customer': '%s'\n", ws, td.IsVisibleForCustomer))

	return str.String()
}
