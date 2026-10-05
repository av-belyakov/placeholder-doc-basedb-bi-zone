package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForTitlesType вспомогательный тип используемый для хранения объектов типа 'titles'
type SupportingStructureForTitlesType struct {
	titles             []datamodels.BiZoneIRPTitle
	listAcceptedFields []string
	titleTmp           datamodels.BiZoneIRPTitle
	isCompleted        bool
}

// NewSupportingStructureForTitlesType формирует вспомогательный объект для обработки объектов типа 'titles'
func NewSupportingStructureForTitlesType() *SupportingStructureForTitlesType {
	return &SupportingStructureForTitlesType{
		listAcceptedFields: []string(nil),
		titleTmp:           *datamodels.NewBiZoneIRPTitle(),
		titles:             make([]datamodels.BiZoneIRPTitle, 0),
	}
}

// GetTitles возвращает []datamodels.BiZoneIRPTitle
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из tg.titleTmp в
// tg.titles, так как titles автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются tg.titleTmp
func (tl *SupportingStructureForTitlesType) GetTitles() []datamodels.BiZoneIRPTitle {
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
func (tl *SupportingStructureForTitlesType) GetTitleTmp() datamodels.BiZoneIRPTitle {
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
		tl.titleTmp = *datamodels.NewBiZoneIRPTitle()
		tl.listAcceptedFields = []string(nil)
	}

	tl.listAcceptedFields = append(tl.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (tl *SupportingStructureForTitlesType) isExistFieldBranch(v string) bool {
	return slices.Contains(tl.listAcceptedFields, v)
}
