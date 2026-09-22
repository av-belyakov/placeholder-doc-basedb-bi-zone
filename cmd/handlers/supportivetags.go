package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForTagsType вспомогательный тип используемый для хранения объектов типа 'tags'
type SupportingStructureForTagsType struct {
	listAcceptedFields []string
	tagTmp             *datamodels.BiZoneIRPTag
	tags               []datamodels.BiZoneIRPTag
}

// NewSupportingStructureForTagsType формирует вспомогательный объект для обработки объектов типа 'tags'
func NewSupportingStructureForTagsType() *SupportingStructureForTagsType {
	return &SupportingStructureForTagsType{
		listAcceptedFields: []string(nil),
		tagTmp:             datamodels.NewBiZoneIRPTag(),
		tags:               make([]datamodels.BiZoneIRPTag, 0),
	}
}

// GetTags возвращает []datamodels.BiZoneIRPTag
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из tg.tagTmp в
// tg.tags, так как tags автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются tg.tagTmp
func (tg *SupportingStructureForTagsType) GetTags() []datamodels.BiZoneIRPTag {
	if tg.tagTmp != nil {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&tg.tagTmp)
		tg.tags = append(tg.tags, *tg.tagTmp)

		tg.tagTmp = nil
		tg.listAcceptedFields = []string(nil)
	}

	return tg.tags
}

// GetTagTmp возвращает временный объект
func (tg *SupportingStructureForTagsType) GetTagTmp() *datamodels.BiZoneIRPTag {
	return tg.tagTmp
}

// HandlerValue функция обработчик значений
func (tg *SupportingStructureForTagsType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if tg.isExistFieldBranch((fieldBranch)) {
		tg.listAcceptedFields = []string(nil)

		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&sc.tagTmp)
		tg.tags = append(tg.tags, *tg.tagTmp)
		tg.tagTmp = datamodels.NewBiZoneIRPTag()
	}

	tg.listAcceptedFields = append(tg.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (tg *SupportingStructureForTagsType) isExistFieldBranch(v string) bool {
	return slices.Contains(tg.listAcceptedFields, v)
}
