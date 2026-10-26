package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewObservedIocsType() *ObservedIocsType {
	return &ObservedIocsType{}
}

// Get для типа ObservedIocsType
func (iocs *ObservedIocsType) Get() *ObservedIocsType {
	return iocs
}

// GetIoc для поля 'ioc'
func (iocs *ObservedIocsType) GetIoc() string {
	return iocs.Ioc
}

// SetIoc для поля 'ioc'
func (iocs *ObservedIocsType) SetIoc(v string) error {
	iocs.Ioc = v

	return nil
}

// SetAnyIoc для поля 'ioc'
func (iocs *ObservedIocsType) SetAnyIoc(a any) error {
	return iocs.SetIoc(fmt.Sprint(a))
}

// GetIocType для поля 'ioc_type'
func (iocs *ObservedIocsType) GetIocType() string {
	return iocs.IocType
}

// SetIocType для поля 'ioc_type'
func (iocs *ObservedIocsType) SetIocType(v string) error {
	iocs.IocType = v

	return nil
}

// SetAnyIocType для поля 'ioc_type'
func (iocs *ObservedIocsType) SetAnyIocType(a any) error {
	return iocs.SetIocType(fmt.Sprint(a))
}

// GetCategory для поля 'category'
func (iocs *ObservedIocsType) GetCategory() []string {
	return iocs.Category
}

// SetCategory для поля 'category'
func (iocs *ObservedIocsType) SetCategory(v []string) error {
	iocs.Category = v

	return nil
}

// SetCategoryElement добавляет значение 'category' в список
func (iocs *ObservedIocsType) SetCategoryElement(v string) error {
	if _, isExist := supportingfunctions.SliceContainsElement(v, iocs.Category); !isExist {
		iocs.Category = append(iocs.Category, v)
	}

	return nil
}

// SetAnyCategoryElement добавляет некоторое значение в список 'category'
func (iocs *ObservedIocsType) SetAnyCategoryElement(a any) error {
	return iocs.SetCategoryElement(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (iocs *ObservedIocsType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'ioc': '%s'\n", ws, iocs.Ioc)
	fmt.Fprintf(&str, "%s'ioc_type': '%s'\n", ws, iocs.IocType)
	fmt.Fprintf(&str, "%s'category': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, iocs.Category))

	return str.String()
}
