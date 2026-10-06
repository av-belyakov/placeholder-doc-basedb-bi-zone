package handlers

import (
	"fmt"
	"slices"
)

// BiZoneIRPTitle локализованное значение
type TitleTemporary struct {
	ID, Value, Language string
}

func NewTitleTemporary() *TitleTemporary {
	return &TitleTemporary{
		Language: "ru",
	}
}

// GetID
func (t *TitleTemporary) GetID() string {
	return t.ID
}

// SetID
func (t *TitleTemporary) SetID(v string) error {
	t.ID = v

	return nil
}

// SetAnyID
func (t *TitleTemporary) SetAnyID(a any) error {
	return t.SetID(fmt.Sprint(a))
}

// GetValue
func (t *TitleTemporary) GetValue() string {
	return t.Value
}

// SetValue
func (t *TitleTemporary) SetValue(v string) error {
	t.Value = v

	return nil
}

// SetAnyValue
func (t *TitleTemporary) SetAnyValue(a any) error {
	return t.SetValue(fmt.Sprint(a))
}

// GetLanguage
func (t *TitleTemporary) GetLanguage() string {
	return t.Language
}

// SetLanguage
func (t *TitleTemporary) SetLanguage(v string) error {
	t.Language = v

	return nil
}

// SetAnyLanguage
func (t *TitleTemporary) SetAnyLanguage(a any) error {
	return t.SetLanguage(fmt.Sprint(a))
}

// SupportingStructureForTitlesType вспомогательный тип используемый для хранения объектов типа 'titles'
type SupportingStructureForTitlesType struct {
	titles             []TitleTemporary
	listAcceptedFields []string
	titleTmp           TitleTemporary
	isCompleted        bool
}

// NewSupportingStructureForTitlesType формирует вспомогательный объект для обработки объектов типа 'titles'
func NewSupportingStructureForTitlesType() *SupportingStructureForTitlesType {
	return &SupportingStructureForTitlesType{
		listAcceptedFields: []string(nil),
		titleTmp:           *NewTitleTemporary(),
		titles:             make([]TitleTemporary, 0),
	}
}

// GetTitles возвращает []TitleTemporary
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из tg.titleTmp в
// tg.titles, так как titles автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются tg.titleTmp
func (tl *SupportingStructureForTitlesType) GetTitles() []TitleTemporary {
	if !tl.isCompleted {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&tg.titleTmp)
		tl.titles = append(tl.titles, tl.titleTmp)
		tl.listAcceptedFields = []string(nil)
		tl.isCompleted = true
	}

	return tl.titles
}

// GetTagTmp возвращает временный объект
func (tl *SupportingStructureForTitlesType) GetTitleTmp() TitleTemporary {
	return tl.titleTmp
}

// HandlerValue функция обработчик значений
func (tl *SupportingStructureForTitlesType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if tl.isExistFieldBranch((fieldBranch)) {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&sc.tagTmp)
		tl.titles = append(tl.titles, tl.titleTmp)
		tl.titleTmp = *NewTitleTemporary()
		tl.listAcceptedFields = []string(nil)
	}

	tl.listAcceptedFields = append(tl.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (tl *SupportingStructureForTitlesType) isExistFieldBranch(v string) bool {
	return slices.Contains(tl.listAcceptedFields, v)
}
