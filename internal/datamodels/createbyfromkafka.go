package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPCreatedBy() *BiZoneIRPCreatedBy {
	return &BiZoneIRPCreatedBy{}
}

// GetID для поля 'id'
func (cb *BiZoneIRPCreatedBy) GetID() uint64 {
	return cb.ID
}

// SetID для поля 'id'
func (cb *BiZoneIRPCreatedBy) SetID(v uint64) error {
	cb.ID = v

	return nil
}

// SetAnyID для поля 'id'
func (cb *BiZoneIRPCreatedBy) SetAnyID(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return cb.SetID(v)
}

// GetUsername для поля 'username'
func (cb *BiZoneIRPCreatedBy) GetUsername() string {
	return cb.Username
}

// SetUsername для поля 'username'
func (cb *BiZoneIRPCreatedBy) SetUsername(v string) error {
	cb.Username = v

	return nil
}

// SetAnyUsername для поля 'username'
func (cb *BiZoneIRPCreatedBy) SetAnyUsername(a any) error {
	return cb.SetUsername(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (cb *BiZoneIRPCreatedBy) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%d'\n", ws, cb.ID)
	fmt.Fprintf(&str, "%s'username': '%s'\n", ws, cb.Username)

	return str.String()
}
