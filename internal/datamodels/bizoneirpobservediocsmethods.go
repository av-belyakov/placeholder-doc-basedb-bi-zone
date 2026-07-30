package datamodels

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPObservedIocs() *BiZoneIRPObservedIocs {
	return &BiZoneIRPObservedIocs{
		Category: make([]string, 0),
	}
}

// GetIoc сам Ioc
func (oiocs *BiZoneIRPObservedIocs) GetIoc() string {
	return oiocs.Ioc
}

// SetIoc сам Ioc
func (oiocs *BiZoneIRPObservedIocs) SetIoc(ioc string) {
	oiocs.Ioc = ioc
}

// SetAnyIoc сам Ioc
func (oiocs *BiZoneIRPObservedIocs) SetAnyIoc(a any) {
	oiocs.Ioc = fmt.Sprint(a)
}

// GetIocType тип Ioc
func (oiocs *BiZoneIRPObservedIocs) GetIocType() string {
	return oiocs.IocType
}

// SetIocType тип Ioc
func (oiocs *BiZoneIRPObservedIocs) SetIocType(iocType string) {
	oiocs.IocType = iocType
}

// SetAnyIocType тип Ioc
func (oiocs *BiZoneIRPObservedIocs) SetAnyIocType(a any) {
	oiocs.IocType = fmt.Sprint(a)
}

// GetCategories категории
func (oiocs *BiZoneIRPObservedIocs) GetCategories() []string {
	if oiocs.Category == nil {
		oiocs.Category = make([]string, 0)
	}

	return oiocs.Category
}

// SetCategories категории
func (oiocs *BiZoneIRPObservedIocs) SetCategories(categories []string) {
	oiocs.Category = categories
}

// AddCategory добавляет значение в список
func (oiocs *BiZoneIRPObservedIocs) AddCategory(category string) {
	oiocs.GetCategories()

	if slices.Contains(oiocs.Category, category) {
		return
	}

	oiocs.Category = append(oiocs.Category, category)
}

// AddAnyCategory добавляет значение в список
func (oiocs *BiZoneIRPObservedIocs) AddAnyCategory(a any) {
	oiocs.AddCategory(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (oiocs *BiZoneIRPObservedIocs) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'ioc': '%s'\n", ws, oiocs.Ioc))
	str.WriteString(fmt.Sprintf("%s'iocType': '%s'\n", ws, oiocs.IocType))
	str.WriteString(fmt.Sprintf("%s'categores': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, oiocs.Category)))

	return str.String()
}
