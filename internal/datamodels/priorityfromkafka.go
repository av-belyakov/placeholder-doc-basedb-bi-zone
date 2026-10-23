package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPPriority() *BiZoneIRPPriority {
	return &BiZoneIRPPriority{
		Title: []string(nil),
		//Title: []BiZoneIRPTitle(nil),
	}
}

// Get содержимое значения 'priority'
func (p *BiZoneIRPPriority) Get() *BiZoneIRPPriority {
	return p
}

// GetID для поля 'id'
func (p *BiZoneIRPPriority) GetID() string {
	return p.ID
}

// SetID для поля 'id'
func (p *BiZoneIRPPriority) SetID(v string) error {
	p.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (p *BiZoneIRPPriority) SetAnyID(a any) error {
	return p.SetID(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (p *BiZoneIRPPriority) GetTitle() []string {
	return p.Title
}

// GetTitle для поля 'title'
func (p *BiZoneIRPPriority) SetTitle(v []string) error {
	p.Title = v

	return nil
}

// SetTitleElement добавляет значение 'title' в список
func (p *BiZoneIRPPriority) SetTitleElement(v string) error {
	if p.Title == nil {
		p.Title = []string(nil)
	}

	p.Title = append(p.Title, v)

	return nil
}

/*// GetTitle для поля 'title'
func (p *BiZoneIRPPriority) GetTitle() []BiZoneIRPTitle {
	return p.Title
}

// GetTitle для поля 'title'
func (p *BiZoneIRPPriority) SetTitle(v []BiZoneIRPTitle) error {
	p.Title = v

	return nil
}

// SetTitleElement добавляет значение 'title' в список
func (p *BiZoneIRPPriority) SetTitleElement(v BiZoneIRPTitle) error {
	if p.Title == nil {
		p.Title = []BiZoneIRPTitle(nil)
	}

	p.Title = append(p.Title, v)

	return nil
}*/

// SetAnyTitleElement добавляет значение 'title' в список
func (p *BiZoneIRPPriority) SetAnyTitleElement(a any) error {
	return p.SetTitleElement(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (p *BiZoneIRPPriority) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, p.Title)
	fmt.Fprintf(&str, "%s'title': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, p.Title))
	//fmt.Fprintf(&str, "%s'title':\n", ws)
	//for k, v := range p.Title {
	//fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))
	//}

	return str.String()
}
