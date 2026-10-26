package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForSecondaryCategoryRefType вспомогательный тип используемый для хранения объектов типа 'secondary_category_ref'
type SupportingStructureForSecondaryCategoryRefType struct {
	secondaryCategoryRefs   []datamodels.CategoryRefType
	listAcceptedFields      []string
	secondaryCategoryRefTmp datamodels.CategoryRefType
	isCompleted             bool
}

// NewSupportingStructureForSecondaryCategoryRefType формирует вспомогательный объект для обработки объектов типа 'secondary_category_ref'
func NewSupportingStructureForSecondaryCategoryRefType() *SupportingStructureForSecondaryCategoryRefType {
	return &SupportingStructureForSecondaryCategoryRefType{
		listAcceptedFields:      []string(nil),
		secondaryCategoryRefTmp: *datamodels.NewCategoryRefType(),
		secondaryCategoryRefs:   make([]datamodels.CategoryRefType, 0),
	}
}

// GetSecondaryCategoryRefs возвращает []datamodels.CategoryRefType
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из cr.secondaryCategoryRefTmp в
// cr.secondaryCategoryRefs, так как tags автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются iocs.observedIocsTmp
func (cr *SupportingStructureForSecondaryCategoryRefType) GetSecondaryCategoryRefs() []datamodels.CategoryRefType {
	if !cr.isCompleted {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&cr.secondaryCategoryRefTmp)
		cr.secondaryCategoryRefs = append(cr.secondaryCategoryRefs, cr.secondaryCategoryRefTmp)
		cr.listAcceptedFields = []string(nil)
		cr.isCompleted = true
	}

	return cr.secondaryCategoryRefs
}

// GetSecondaryCategoryRefsTmp возвращает временный объект
func (cr *SupportingStructureForSecondaryCategoryRefType) GetSecondaryCategoryRefsTmp() datamodels.CategoryRefType {
	return cr.secondaryCategoryRefTmp
}

// HandlerValue функция обработчик значений
func (cr *SupportingStructureForSecondaryCategoryRefType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if cr.isExistFieldBranch((fieldBranch)) {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&cr.secondaryCategoryRefTmp)
		cr.secondaryCategoryRefs = append(cr.secondaryCategoryRefs, cr.secondaryCategoryRefTmp)
		cr.secondaryCategoryRefTmp = *datamodels.NewCategoryRefType()
		cr.listAcceptedFields = []string(nil)
	}

	cr.listAcceptedFields = append(cr.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (cr *SupportingStructureForSecondaryCategoryRefType) isExistFieldBranch(v string) bool {
	return slices.Contains(cr.listAcceptedFields, v)
}
