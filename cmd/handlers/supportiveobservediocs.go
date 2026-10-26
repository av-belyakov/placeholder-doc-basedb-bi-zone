package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForObservedIocsType вспомогательный тип используемый для хранения объектов типа 'observed_iocs'
type SupportingStructureForObservedIocsType struct {
	observedIocs       []datamodels.ObservedIocsType
	listAcceptedFields []string
	observedIocsTmp    datamodels.ObservedIocsType
	isCompleted        bool
}

// NewSupportingStructureForObservedIocsType формирует вспомогательный объект для обработки объектов типа 'observed_iocs'
func NewSupportingStructureForObservedIocsType() *SupportingStructureForObservedIocsType {
	return &SupportingStructureForObservedIocsType{
		listAcceptedFields: []string(nil),
		observedIocsTmp:    *datamodels.NewObservedIocsType(),
		observedIocs:       make([]datamodels.ObservedIocsType, 0),
	}
}

// GetObservedIocs возвращает []datamodels.ObservedIocsType
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из iocs.observedIocsTmp в
// iocs.observedIocs, так как tags автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются iocs.observedIocsTmp
func (iocs *SupportingStructureForObservedIocsType) GetObservedIocs() []datamodels.ObservedIocsType {
	if !iocs.isCompleted {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&iocs.observedIocsTmp)
		iocs.observedIocs = append(iocs.observedIocs, iocs.observedIocsTmp)
		iocs.listAcceptedFields = []string(nil)
		iocs.isCompleted = true
	}

	return iocs.observedIocs
}

// GetObservedIocsTmp возвращает временный объект
func (iocs *SupportingStructureForObservedIocsType) GetObservedIocsTmp() datamodels.ObservedIocsType {
	return iocs.observedIocsTmp
}

// HandlerValue функция обработчик значений
func (iocs *SupportingStructureForObservedIocsType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if iocs.isExistFieldBranch((fieldBranch)) {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&sc.tagTmp)
		iocs.observedIocs = append(iocs.observedIocs, iocs.observedIocsTmp)
		iocs.observedIocsTmp = *datamodels.NewObservedIocsType()
		iocs.listAcceptedFields = []string(nil)
	}

	iocs.listAcceptedFields = append(iocs.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (iocs *SupportingStructureForObservedIocsType) isExistFieldBranch(v string) bool {
	return slices.Contains(iocs.listAcceptedFields, v)
}
