package datamodels

import (
	"errors"
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewWatcherType() *WatchersType {
	return &WatchersType{}
}

// Get для типа WatchersType
func (w *WatchersType) Get() *WatchersType {
	return w
}

// GetId для поля 'id'
func (w *WatchersType) GetId() uint64 {
	return w.Id
}

// SetId для поля 'id'
func (w *WatchersType) SetId(v uint64) error {
	w.Id = v

	return nil
}

// SetAnyId для поля 'id'
func (w *WatchersType) SetAnyId(a any) error {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return err
	}

	return w.SetId(v)
}

// GetIsActive для поля 'is_active'
func (w *WatchersType) GetIsActive() bool {
	return w.IsActive
}

// SetIsActive для поля 'is_active'
func (w *WatchersType) SetIsActive(v bool) error {
	w.IsActive = v

	return nil
}

// SetAnyIsActive для поля 'is_active'
func (w *WatchersType) SetAnyIsActive(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'is_active'")
	}

	return w.SetIsActive(v)
}

// GetEmail для поля 'email'
func (w *WatchersType) GetEmail() string {
	return w.Email
}

// SetEmail для поля 'email'
func (w *WatchersType) SetEmail(v string) error {
	w.Email = v

	return nil
}

// SetAnyEmail для поля 'email'
func (w *WatchersType) SetAnyEmail(a any) error {
	return w.SetEmail(fmt.Sprint(a))
}

// GetLastName для поля 'last_name'
func (w *WatchersType) GetLastName() string {
	return w.LastName
}

// SetLastName для поля 'last_name'
func (w *WatchersType) SetLastName(v string) error {
	w.LastName = v

	return nil
}

// SetAnyLastName для поля 'last_name'
func (w *WatchersType) SetAnyLastName(a any) error {
	return w.SetLastName(fmt.Sprint(a))
}

// GetUsername для поля 'username'
func (w *WatchersType) GetUsername() string {
	return w.Username
}

// SetUsername для поля 'username'
func (w *WatchersType) SetUsername(v string) error {
	w.Username = v

	return nil
}

// SetAnyUsername для поля 'username'
func (w *WatchersType) SetAnyUsername(a any) error {
	return w.SetUsername(fmt.Sprint(a))
}

// GetFirstName для поля 'first_name'
func (w *WatchersType) GetFirstName() string {
	return w.FirstName
}

// SetFirstName для поля 'first_name'
func (w *WatchersType) SetFirstName(v string) error {
	w.FirstName = v

	return nil
}

// SetAnyFirstName для поля 'first_name'
func (w *WatchersType) SetAnyFirstName(a any) error {
	return w.SetFirstName(fmt.Sprint(a))
}

// GetPatronimic для поля 'patronimic'
func (w *WatchersType) GetPatronimic() string {
	return w.Patronimic
}

// SetPatronimic для поля 'patronimic'
func (w *WatchersType) SetPatronimic(v string) error {
	w.Patronimic = v

	return nil
}

// SetAnyPatronimic для поля 'patronimic'
func (w *WatchersType) SetAnyPatronimic(a any) error {
	return w.SetPatronimic(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (w *WatchersType) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'id': '%d'\n", ws, w.Id)
	fmt.Fprintf(&str, "%s'is_active': '%t'\n", ws, w.IsActive)
	fmt.Fprintf(&str, "%s'email': '%s'\n", ws, w.Email)
	fmt.Fprintf(&str, "%s'last_name': '%s'\n", ws, w.LastName)
	fmt.Fprintf(&str, "%s'username': '%s'\n", ws, w.Username)
	fmt.Fprintf(&str, "%s'first_name': '%s'\n", ws, w.FirstName)
	fmt.Fprintf(&str, "%s'patronimic': '%s'\n", ws, w.Patronimic)

	return str.String()
}
