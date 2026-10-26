package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewCategoryRefType() *CategoryRefType {
	return &CategoryRefType{}
}

// Get для типа CategoryRefType
func (itr *CategoryRefType) Get() *CategoryRefType {
	return itr
}

// GetId для поля 'id'
func (itr *CategoryRefType) GetId() string {
	return itr.Id
}

// SetId для поля 'id'
func (itr *CategoryRefType) SetId(v string) error {
	itr.Id = v

	return nil
}

// SetAnyId для поля 'id'
func (itr *CategoryRefType) SetAnyId(a any) error {
	return itr.SetId(fmt.Sprint(a))
}

// GetTitle для поля 'title'
func (itr *CategoryRefType) GetTitle() string {
	return itr.Title
}

// SetTitle для поля 'title'
func (itr *CategoryRefType) SetTitle(v string) error {
	itr.Title = v

	return nil
}

// SetAnyTitle для поля 'title'
func (itr *CategoryRefType) SetAnyTitle(a any) error {
	return itr.SetTitle(fmt.Sprint(a))
}

// GetDescription для поля 'description'
func (itr *CategoryRefType) GetDescription() string {
	return itr.Id
}

// SetDescription для поля 'description'
func (itr *CategoryRefType) SetDescription(v string) error {
	itr.Id = v

	return nil
}

// SetAnyDescription для поля 'description'
func (itr *CategoryRefType) SetAnyDescription(a any) error {
	return itr.SetDescription(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (itr *CategoryRefType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, itr.Id)
	fmt.Fprintf(&str, "%s'title': '%s'\n", ws, itr.Title)
	fmt.Fprintf(&str, "%s'description': '%s'\n", ws, itr.Description)

	return str.String()
}
