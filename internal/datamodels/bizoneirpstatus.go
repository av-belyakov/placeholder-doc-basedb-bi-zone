package datamodels

import (
	"errors"
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPStatus() *BiZoneIRPStatus {
	return &BiZoneIRPStatus{
		Title: []BiZoneIRPTitle(nil),
	}
}

// GetID для поля 'id'
func (t *BiZoneIRPStatus) GetID() string {
	return t.ID
}

// SetID для поля 'id'
func (t *BiZoneIRPStatus) SetID(v string) error {
	t.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (t *BiZoneIRPStatus) SetAnyID(a any) error {
	return t.SetID(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (t *BiZoneIRPStatus) GetTitle() []BiZoneIRPTitle {
	return t.Title
}

// GetTitle для поля 'title'
func (t *BiZoneIRPStatus) SetTitle(v []BiZoneIRPTitle) error {
	t.Title = v

	return nil
}

// SetTitleElement добавляет значение 'title' в список
func (t *BiZoneIRPStatus) SetTitleElement(v BiZoneIRPTitle) error {
	if t.Title == nil {
		t.Title = []BiZoneIRPTitle(nil)
	}

	t.Title = append(t.Title, v)

	return nil
}

// SetAnyTitleElement добавляет значение 'title' в список
func (ds *BiZoneIRPStatus) SetAnyTitleElement(a any) error {
	if v, ok := a.(BiZoneIRPTitle); ok {
		return ds.SetTitleElement(v)
	}

	return errors.New("type conversion error for field 'title'")
}

// ToStringBeautiful форматированный вывод
func (t *BiZoneIRPStatus) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, t.Title)
	fmt.Fprintf(&str, "%s'title':\n", ws)
	for k, v := range t.Title {
		fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	}

	return str.String()
}
