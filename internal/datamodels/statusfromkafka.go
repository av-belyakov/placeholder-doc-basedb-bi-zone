package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPStatus() *BiZoneIRPStatus {
	return &BiZoneIRPStatus{
		Title: []string(nil),
	}
}

// Get содержимое значения 'status'
func (s *BiZoneIRPStatus) Get() *BiZoneIRPStatus {
	return s
}

// GetID для поля 'id'
func (s *BiZoneIRPStatus) GetID() string {
	return s.ID
}

// SetID для поля 'id'
func (s *BiZoneIRPStatus) SetID(v string) error {
	s.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (s *BiZoneIRPStatus) SetAnyID(a any) error {
	return s.SetID(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (s *BiZoneIRPStatus) GetTitle() []string {
	return s.Title
}

// GetTitle для поля 'title'
func (s *BiZoneIRPStatus) SetTitle(v []string) error {
	s.Title = v

	return nil
}

// SetTitleElement добавляет значение 'title' в список
func (s *BiZoneIRPStatus) SetTitleElement(v string) error {
	if s.Title == nil {
		s.Title = []string(nil)
	}

	s.Title = append(s.Title, v)

	return nil
}

// SetAnyTitleElement добавляет значение 'title' в список
func (s *BiZoneIRPStatus) SetAnyTitleElement(a any) error {
	return s.SetTitleElement(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (s *BiZoneIRPStatus) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, s.Title)
	fmt.Fprintf(&str, "%s'title': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, s.Title))
	//fmt.Fprintf(&str, "%s'title':\n", ws)
	//for k, v := range s.Title {
	//	fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	//}

	return str.String()
}
