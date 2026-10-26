package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForCategoresType вспомогательный тип используемый для хранения объектов типа 'category'
type SupportingStructureForCategoreRefsType struct {
	categoreRefs       []datamodels.CategoryRefType
	listAcceptedFields []string
	categoryRefTmp     datamodels.CategoryRefType
	isCompleted        bool
}

// NewSupportingStructureForCategoreRefsType формирует вспомогательный объект для обработки объектов типа 'category_ref'
func NewSupportingStructureForCategoreRefsType() *SupportingStructureForCategoreRefsType {
	return &SupportingStructureForCategoreRefsType{
		listAcceptedFields: []string(nil),
		categoryRefTmp:     *datamodels.NewCategoryRefType(),
		categoreRefs:       make([]datamodels.CategoryRefType, 0),
	}
}

// GetCategores возвращает []datamodels.CategoryRefType
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из c.categoryRefTmp в
// c.categoreRefs, так как Categores автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются c.categoryTmp
func (cr *SupportingStructureForCategoreRefsType) GetCategoreRefs() []datamodels.CategoryRefType {
	if !cr.isCompleted {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&tg.titleTmp)
		cr.categoreRefs = append(cr.categoreRefs, cr.categoryRefTmp)
		cr.listAcceptedFields = []string(nil)
		cr.isCompleted = true
	}

	return cr.categoreRefs
}

// GetCategoryTmp возвращает временный объект
func (cr *SupportingStructureForCategoreRefsType) GetCategoreTmp() datamodels.CategoryRefType {
	return cr.categoryRefTmp
}

// HandlerValue функция обработчик значений
func (cr *SupportingStructureForCategoreRefsType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if cr.isExistFieldBranch((fieldBranch)) {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&cr.categoryRefTmp)
		cr.categoreRefs = append(cr.categoreRefs, cr.categoryRefTmp)
		cr.categoryRefTmp = *datamodels.NewCategoryRefType()
		cr.listAcceptedFields = []string(nil)
	}

	cr.listAcceptedFields = append(cr.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (cr *SupportingStructureForCategoreRefsType) isExistFieldBranch(v string) bool {
	return slices.Contains(cr.listAcceptedFields, v)
}
