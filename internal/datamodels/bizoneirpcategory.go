package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPCategory() *BiZoneIRPCategory {
	return &BiZoneIRPCategory{
		//Title: []BiZoneIRPTitle(nil),
		Title: []string(nil),
	}
}

// GetID для поля 'id'
func (c *BiZoneIRPCategory) GetID() string {
	return c.ID
}

// SetID для поля 'id'
func (c *BiZoneIRPCategory) SetID(v string) error {
	c.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (c *BiZoneIRPCategory) SetAnyID(a any) error {
	return c.SetID(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (c *BiZoneIRPCategory) GetTitle() []string {
	return c.Title
}

// GetTitle для поля 'title'
func (c *BiZoneIRPCategory) SetTitle(v []string) error {
	c.Title = v

	return nil
}

// SetTitleElement добавляет значение 'title' в список
func (c *BiZoneIRPCategory) SetTitleElement(v string) error {
	if c.Title == nil {
		c.Title = []string(nil)
	}

	c.Title = append(c.Title, v)

	return nil
}

// SetAnyTitleElement добавляет значение 'title' в список
func (c *BiZoneIRPCategory) SetAnyTitleElement(a any) error {
	return c.SetTitleElement(fmt.Sprint(a))
}

/*// GetTitle для поля 'title'
func (c *BiZoneIRPCategory) GetTitle() []BiZoneIRPTitle {
	return c.Title
}

// GetTitle для поля 'title'
func (c *BiZoneIRPCategory) SetTitle(v []BiZoneIRPTitle) error {
	c.Title = v

	return nil
}

// SetTitleElement добавляет значение 'title' в список
func (c *BiZoneIRPCategory) SetTitleElement(v BiZoneIRPTitle) error {
	if c.Title == nil {
		c.Title = []BiZoneIRPTitle(nil)
	}

	c.Title = append(c.Title, v)

	return nil
}

// SetAnyTitleElement добавляет значение 'title' в список
func (ds *BiZoneIRPCategory) SetAnyTitleElement(a any) error {
	if v, ok := a.(BiZoneIRPTitle); ok {
		return ds.SetTitleElement(v)
	}

	return errors.New("type conversion error for field 'title'")
}*/

// ToStringBeautiful форматированный вывод
func (c *BiZoneIRPCategory) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, c.Title)
	fmt.Fprintf(&str, "%s'title': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, c.Title))
	//fmt.Fprintf(&str, "%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2))

	return str.String()
}
