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

// GetCreated для поля Created (формат RFC3339)
func (t *BiZoneIRPTag) GetCreated() string {
	return t.Created
}

// SetCreated для поля Created (преобразует в формат времени RFC3339)
func (t *BiZoneIRPTag) SetCreated(v string) error {
	timeStr, err := supportingfunctions.SmartConvertToRFC3339(v)
	if err != nil {
		return err
	}

	t.Created = timeStr

	return nil
}

// SetAnyCreated для поля Created
func (t *BiZoneIRPTag) SetAnyCreated(a any) error {
	if v, ok := a.(string); ok {
		return t.SetCreated(v)
	}

	return errors.New("type conversion error")
}

// GetName для поля Name
func (t *BiZoneIRPTag) GetName() string {
	return t.Name
}

// SetName для поля Name
func (t *BiZoneIRPTag) SetName(name string) {
	t.Name = name
}

// SetAnyName для поля Name
func (t *BiZoneIRPTag) SetAnyName(a any) {
	t.Name = fmt.Sprint(a)
}

// GetColor для поля Color
func (t *BiZoneIRPTag) GetColor() string {
	return t.Color
}

// SetColor для поля Color
func (t *BiZoneIRPTag) SetColor(color string) {
	t.Color = color
}

// SetAnyColor для поля Color
func (t *BiZoneIRPTag) SetAnyColor(a any) {
	t.Color = fmt.Sprint(a)
}

// GetCreatedByID для поля CreatedBy.ID
func (t *BiZoneIRPTag) GetCreatedByID() uint64 {
	return t.CreatedBy.ID
}

// SetCreatedByID для поля CreatedBy.ID
func (t *BiZoneIRPTag) SetCreatedByID(id uint64) {
	t.CreatedBy.ID = id
}

// SetAnyCreatedByID для поля CreatedBy.ID
func (t *BiZoneIRPTag) SetAnyCreatedByID(a any) {
	if v, ok := a.(float32); ok {
		t.CreatedBy.ID = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		t.CreatedBy.ID = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		t.CreatedBy.ID = v
	}
}

// GetCreatedByUsername для поля CreatedBy.Username
func (t *BiZoneIRPTag) GetCreatedByUsername() string {
	return t.CreatedBy.Username
}

// SetCreatedByUsername для поля CreatedBy.Username
func (t *BiZoneIRPTag) SetCreatedByUsername(username string) {
	t.CreatedBy.Username = username
}

// SetAnyCreatedByUsername для поля CreatedBy.Username
func (t *BiZoneIRPTag) SetAnyCreatedByUsername(a any) {
	t.CreatedBy.Username = fmt.Sprint(a)
}

// ToStringBeautiful форматированный вывод
func (t *BiZoneIRPTag) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, t.Name))
	str.WriteString(fmt.Sprintf("%s'color': '%s'\n", ws, t.Color))
	str.WriteString(fmt.Sprintf("%s'created': '%s'\n", ws, t.Created))
	str.WriteString(fmt.Sprintf("%s'created_by.id': '%d'\n", ws, t.CreatedBy.ID))
	str.WriteString(fmt.Sprintf("%s'created_by.username': '%s'\n", ws, t.CreatedBy.Username))

	return str.String()
}
