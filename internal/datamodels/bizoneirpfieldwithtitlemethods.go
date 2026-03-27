package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPFieldWithTitle() *BiZoneIRPFieldWithTitle {
	return &BiZoneIRPFieldWithTitle{
		Title: make([]BiZoneIRPFieldTitle, 0),
	}
}

// GetId
func (fwt *BiZoneIRPFieldWithTitle) GetId() string {
	return fwt.Id
}

// SetId
func (fwt *BiZoneIRPFieldWithTitle) SetId(v string) {
	fwt.Id = v
}

// SetAnyId
func (fwt *BiZoneIRPFieldWithTitle) SetAnyId(a any) {
	fwt.SetId(fmt.Sprint(a))
}

// GetTitle
func (fwt *BiZoneIRPFieldWithTitle) GetTitle() []BiZoneIRPFieldTitle {
	return fwt.Title
}

// SetTitle
func (fwt *BiZoneIRPFieldWithTitle) SetTitle(v []BiZoneIRPFieldTitle) {
	fwt.Title = v
}

// ToStringBeautiful форматированный вывод
func (fwt *BiZoneIRPFieldWithTitle) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, fwt.Id))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws))
	for k, v := range fwt.Title {
		str.WriteString(fmt.Sprintf("%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2)))
	}

	return str.String()
}
