package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPType() *BiZoneIRPType {
	return &BiZoneIRPType{
		Title: []string(nil),
	}
}

// GetID для поля 'id'
func (tp *BiZoneIRPType) GetID() string {
	return tp.ID
}

// SetID для поля 'id'
func (tp *BiZoneIRPType) SetID(v string) error {
	tp.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (tp *BiZoneIRPType) SetAnyID(a any) error {
	return tp.SetID(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (tp *BiZoneIRPType) GetTitle() []string {
	return tp.Title
}

// GetTitle для поля 'title'
func (tp *BiZoneIRPType) SetTitle(v []string) error {
	tp.Title = v

	return nil
}

// SetTitleElement добавляет значение 'title' в список
func (tp *BiZoneIRPType) SetTitleElement(v string) error {
	if tp.Title == nil {
		tp.Title = []string(nil)
	}

	tp.Title = append(tp.Title, v)

	return nil
}

// SetAnyTitleElement добавляет значение 'title' в список
func (tp *BiZoneIRPType) SetAnyTitleElement(a any) error {
	return tp.SetTitleElement(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (tp *BiZoneIRPType) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, tp.Title)
	fmt.Fprintf(&str, "%s'title': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, tp.Title))
	//fmt.Fprintf(&str, "%s'title':\n", ws)
	//for k, v := range tp.Title {
	//	fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	//}

	return str.String()
}
