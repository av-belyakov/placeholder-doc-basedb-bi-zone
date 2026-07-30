package datamodels

import (
	"fmt"
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneMITRECOV() *BiZoneIRPMITRECOV {
	return &BiZoneIRPMITRECOV{
		Persistence: make([]string, 0),
	}
}

// GetPersistence
func (mcov *BiZoneIRPMITRECOV) GetPersistence() []string {
	if mcov.Persistence == nil {
		mcov.Persistence = make([]string, 0)
	}

	return mcov.Persistence
}

// SetPersistences список значений
func (mcov *BiZoneIRPMITRECOV) SetPersistencies(persistences []string) {
	mcov.Persistence = persistences
}

// AddPersistence добавляет значение в список
func (mcov *BiZoneIRPMITRECOV) AddPersistence(persistence string) {
	mcov.GetPersistence()

	if slices.Contains(mcov.Persistence, persistence) {
		return
	}

	mcov.Persistence = append(mcov.Persistence, persistence)
}

// AddAnyPersistence добавляет значение в список
func (mcov *BiZoneIRPMITRECOV) AddAnyPersistence(a any) {
	mcov.AddPersistence(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (mcov *BiZoneIRPMITRECOV) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	return fmt.Sprintf("%s'categores': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, mcov.Persistence))
}
