package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewIssueTypeRefType() *IssueTypeRefType {
	return &IssueTypeRefType{}
}

// Get для типа IssueTypeRefType
func (itr *IssueTypeRefType) Get() *IssueTypeRefType {
	return itr
}

// GetId для поля 'id'
func (itr *IssueTypeRefType) GetId() string {
	return itr.Id
}

// SetId для поля 'id'
func (itr *IssueTypeRefType) SetId(v string) error {
	itr.Id = v

	return nil
}

// SetAnyId для поля 'id'
func (itr *IssueTypeRefType) SetAnyId(a any) error {
	return itr.SetId(fmt.Sprint(a))
}

// GetTitle для поля 'ioc_type'
func (itr *IssueTypeRefType) GetTitle() string {
	return itr.Title
}

// SetTitle для поля 'ioc_type'
func (itr *IssueTypeRefType) SetTitle(v string) error {
	itr.Title = v

	return nil
}

// SetAnyTitle для поля 'ioc_type'
func (itr *IssueTypeRefType) SetAnyTitle(a any) error {
	return itr.SetTitle(fmt.Sprint(a))
}

// GetDescription для поля 'description'
func (itr *IssueTypeRefType) GetDescription() string {
	return itr.Id
}

// SetDescription для поля 'description'
func (itr *IssueTypeRefType) SetDescription(v string) error {
	itr.Id = v

	return nil
}

// SetAnyDescription для поля 'description'
func (itr *IssueTypeRefType) SetAnyDescription(a any) error {
	return itr.SetDescription(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (itr *IssueTypeRefType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%s'\n", ws, itr.Id)
	fmt.Fprintf(&str, "%s'ioc_type': '%s'\n", ws, itr.Title)
	fmt.Fprintf(&str, "%s'description': '%s'\n", ws, itr.Description)

	return str.String()
}
