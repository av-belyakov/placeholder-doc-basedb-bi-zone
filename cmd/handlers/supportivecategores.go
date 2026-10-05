package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForCategoresType вспомогательный тип используемый для хранения объектов типа 'category'
type SupportingStructureForCategoresType struct {
	categores          []datamodels.BiZoneIRPCategory
	listAcceptedFields []string
	categoryTmp        datamodels.BiZoneIRPCategory
	isCompleted        bool
}

// NewSupportingStructureForCategoresType формирует вспомогательный объект для обработки объектов типа 'category'
func NewSupportingStructureForCategoresType() *SupportingStructureForCategoresType {
	return &SupportingStructureForCategoresType{
		listAcceptedFields: []string(nil),
		categoryTmp:        *datamodels.NewBiZoneIRPCategory(),
		categores:          make([]datamodels.BiZoneIRPCategory, 0),
	}
}

// GetCategores возвращает []datamodels.BiZoneIRPCategory
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из c.categoryTmp в
// c.Categores, так как Categores автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются c.categoryTmp
func (c *SupportingStructureForCategoresType) GetCategores() []datamodels.BiZoneIRPCategory {
	if !c.isCompleted {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&tg.titleTmp)
		c.categores = append(c.categores, c.categoryTmp)
		c.listAcceptedFields = []string(nil)
		c.isCompleted = true
	}

	return c.categores
}

// GetCategoryTmp возвращает временный объект
func (c *SupportingStructureForCategoresType) GetTitleTmp() datamodels.BiZoneIRPCategory {
	return c.categoryTmp
}

// HandlerValue функция обработчик значений
func (c *SupportingStructureForCategoresType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if c.isExistFieldBranch((fieldBranch)) {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&sc.categoryTmp)
		c.categores = append(c.categores, c.categoryTmp)
		c.categoryTmp = *datamodels.NewBiZoneIRPCategory()
		c.listAcceptedFields = []string(nil)
	}

	c.listAcceptedFields = append(c.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (c *SupportingStructureForCategoresType) isExistFieldBranch(v string) bool {
	return slices.Contains(c.listAcceptedFields, v)
}
