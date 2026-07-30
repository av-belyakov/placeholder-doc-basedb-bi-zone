package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPWatcher() *BiZoneIRPWatcher {
	return &BiZoneIRPWatcher{}
}

// GetId идентификатор
func (b *BiZoneIRPWatcher) GetId() string {
	return b.ID
}

// SetId идентификатор
func (b *BiZoneIRPWatcher) SetId(id string) {
	// Здесь можно добавить валидацию
	b.ID = id
}

// SetAnyId идентификатор
func (b *BiZoneIRPWatcher) SetAnyId(a any) {
	b.ID = fmt.Sprint(a)
}

// GetUsername имя пользователя
func (b *BiZoneIRPWatcher) GetUsername() string {
	return b.Username
}

// SetUsername имя пользователя
func (b *BiZoneIRPWatcher) SetUsername(username string) {
	b.Username = username
}

// SetAnyUsername имя пользователя
func (b *BiZoneIRPWatcher) SetAnyUsername(a any) {
	b.Username = fmt.Sprint(a)
}

// GetEmail электронный адрес
func (b *BiZoneIRPWatcher) GetEmail() string {
	return b.Email
}

// SetEmail электронный адрес
func (b *BiZoneIRPWatcher) SetEmail(email string) {
	b.Email = email
}

// SetAnyEmail электронный адрес
func (b *BiZoneIRPWatcher) SetAnyEmail(a any) {
	b.Email = fmt.Sprint(a)
}

// GetFirstName имя пользователя
func (b *BiZoneIRPWatcher) GetFirstName() string {
	return b.FirstName
}

// SetFirstName имя пользователя
func (b *BiZoneIRPWatcher) SetFirstName(firstName string) {
	b.FirstName = firstName
}

// SetAnyFirstName имя пользователя
func (b *BiZoneIRPWatcher) SetAnyFirstName(a any) {
	b.FirstName = fmt.Sprint(a)
}

// GetLastName фамилия пользователя
func (b *BiZoneIRPWatcher) GetLastName() string {
	return b.LastName
}

// SetLastName фамилия пользователя
func (b *BiZoneIRPWatcher) SetLastName(lastName string) {
	b.LastName = lastName
}

// SetAnyLastName фамилия пользователя
func (b *BiZoneIRPWatcher) SetAnyLastName(a any) {
	b.LastName = fmt.Sprint(a)
}

// GetPatronimic
func (b *BiZoneIRPWatcher) GetPatronimic() string {
	return b.Patronimic
}

// SetPatronimic
func (b *BiZoneIRPWatcher) SetPatronimic(patronimic string) {
	b.Patronimic = patronimic
}

// SetAnyPatronimic
func (b *BiZoneIRPWatcher) SetAnyPatronimic(a any) {
	b.Patronimic = fmt.Sprint(a)
}

// GetIsActive
func (b *BiZoneIRPWatcher) GetIsActive() bool {
	return b.IsActive
}

// SetIsActive
func (b *BiZoneIRPWatcher) SetIsActive(isActive bool) {
	b.IsActive = isActive
}

// SetIsActive
func (b *BiZoneIRPWatcher) SetAnyIsActive(a any) {
	if v, ok := a.(bool); ok {
		b.IsActive = v
	}
}

// ToStringBeautiful форматированный вывод
func (b *BiZoneIRPWatcher) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, b.ID))
	str.WriteString(fmt.Sprintf("%s'username': '%s'\n", ws, b.Username))
	str.WriteString(fmt.Sprintf("%s'firstName': '%s'\n", ws, b.FirstName))
	str.WriteString(fmt.Sprintf("%s'lastName': '%s'\n", ws, b.LastName))
	str.WriteString(fmt.Sprintf("%s'email': '%s'\n", ws, b.Email))
	str.WriteString(fmt.Sprintf("%s'patronimic': '%s'\n", ws, b.Patronimic))
	str.WriteString(fmt.Sprintf("%s'isActive': '%t'\n", ws, b.IsActive))

	return str.String()
}
