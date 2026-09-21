package datamodels

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPTag() *BiZoneIRPTag {
	return &BiZoneIRPTag{
		Created: time.Now().Format(time.RFC3339),
	}
}

// GetCreated для поля 'created' (формат RFC3339)
func (tg *BiZoneIRPTag) GetCreated() string {
	return tg.Created
}

// SetCreated для поля 'created' (преобразует в формат времени RFC3339)
func (tg *BiZoneIRPTag) SetCreated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	tg.Created = timeStr

	return nil
}

// SetAnyCreated для поля 'created'
func (tg *BiZoneIRPTag) SetAnyCreated(a any) error {
	if v, ok := a.(string); ok {
		return tg.SetCreated(v)
	}

	return errors.New("type conversion error")
}

// GetName для поля 'name'
func (tg *BiZoneIRPTag) GetName() string {
	return tg.Name
}

// SetName для поля 'name'
func (tg *BiZoneIRPTag) SetName(v string) error {
	tg.Name = v

	return nil
}

// SetAnyName для поля 'name'
func (tg *BiZoneIRPTag) SetAnyName(a any) error {
	return tg.SetName(fmt.Sprint(a))
}

// GetColor для поля 'color'
func (tg *BiZoneIRPTag) GetColor() string {
	return tg.Color
}

// SetColor для поля 'color'
func (tg *BiZoneIRPTag) SetColor(v string) error {
	tg.Color = v

	return nil
}

// SetAnyColor для поля 'color'
func (tg *BiZoneIRPTag) SetAnyColor(a any) error {
	return tg.SetColor(fmt.Sprint(a))
}

// GetCreatedByID для поля 'created_by.id'
func (tg *BiZoneIRPTag) GetCreatedByID() uint64 {
	return tg.CreatedBy.ID
}

// SetCreatedByID для поля 'created_by.id'
func (tg *BiZoneIRPTag) SetCreatedByID(v uint64) error {
	tg.CreatedBy.ID = v

	return nil
}

// SetAnyCreatedByID для поля 'created_by.id'
func (tg *BiZoneIRPTag) SetAnyCreatedByID(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return tg.SetCreatedByID(v)
}

// GetCreatedByUsername для поля 'created_by.username'
func (tg *BiZoneIRPTag) GetCreatedByUsername() string {
	return tg.CreatedBy.Username
}

// SetCreatedByUsername для поля 'created_by.username'
func (tg *BiZoneIRPTag) SetCreatedByUsername(v string) error {
	tg.CreatedBy.Username = v

	return nil
}

// SetAnyCreatedByUsername для поля 'created_by.username'
func (tg *BiZoneIRPTag) SetAnyCreatedByUsername(a any) error {
	return tg.SetCreatedByUsername(fmt.Sprint(a))
}

// GetIsVisibleForCustomer для поля 'is_visible_for_customer'
func (tg *BiZoneIRPTag) GetIsVisibleForCustomer() bool {
	return tg.IsVisibleForCustomer
}

// SetIsVisibleForCustomer для поля 'is_visible_for_customer'
func (tg *BiZoneIRPTag) SetIsVisibleForCustomer(v bool) error {
	tg.IsVisibleForCustomer = v

	return nil
}

// SetAnyIsVisibleForCustomer для поля 'is_visible_for_customer'
func (tg *BiZoneIRPTag) SetAnyIsVisibleForCustomer(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'is_visible_for_customer'")
	}

	return tg.SetIsVisibleForCustomer(v)
}

// ToStringBeautiful форматированный вывод
func (tg *BiZoneIRPTag) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'name': '%s'\n", ws, tg.Name)
	fmt.Fprintf(&str, "%s'color': '%s'\n", ws, tg.Color)
	fmt.Fprintf(&str, "%s'created': '%s'\n", ws, tg.Created)
	fmt.Fprintf(&str, "%s'created_by.id': '%d'\n", ws, tg.CreatedBy.ID)
	fmt.Fprintf(&str, "%s'created_by.username': '%s'\n", ws, tg.CreatedBy.Username)
	fmt.Fprintf(&str, "%s'is_visible_for_customer': '%t'\n", ws, tg.IsVisibleForCustomer)

	return str.String()
}
