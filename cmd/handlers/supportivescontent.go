package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForSContentType вспомогательный тип для обработки 'data.data_security.scontent'
type SupportingStructureForSContentType struct {
	listAcceptedFields []string
	sContentTmp        *datamodels.BiZoneIRPSContent
	sContentList       []datamodels.BiZoneIRPSContent
}

// NewSupportingStructureForSContentTypeType формирует вспомогательный объект для обработки объектов типа 'data.data_security.scontent'
func NewSupportingStructureForSContentType() *SupportingStructureForSContentType {
	return &SupportingStructureForSContentType{
		listAcceptedFields: []string(nil),
		sContentTmp:        datamodels.NewBiZoneIRPSContent(),
		sContentList:       make([]datamodels.BiZoneIRPSContent, 0),
	}
}

// GetSContent возвращает []datamodels.BiZoneIRPSContent
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из sc.sContentTmp в
// sc.sContentList, так как sContentList автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются sc.sContentTmp
func (sc *SupportingStructureForSContentType) GetSContent() []datamodels.BiZoneIRPSContent {
	if sc.sContentTmp != nil {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&sc.sContentTmp)
		sc.sContentList = append(sc.sContentList, *sc.sContentTmp)

		sc.sContentTmp = nil
		sc.listAcceptedFields = []string(nil)
	}

	return sc.sContentList
}

// GetSContentTmp возвращает временный объект
func (sc *SupportingStructureForSContentType) GetSContentTmp() *datamodels.BiZoneIRPSContent {
	return sc.sContentTmp
}

// HandlerValue функция обработчик значений
func (sc *SupportingStructureForSContentType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if sc.isExistFieldBranch((fieldBranch)) {
		sc.listAcceptedFields = []string(nil)

		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&sc.sContentTmp)
		sc.sContentList = append(sc.sContentList, *sc.sContentTmp)
		sc.sContentTmp = datamodels.NewBiZoneIRPSContent()
	}

	sc.listAcceptedFields = append(sc.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (sc *SupportingStructureForSContentType) isExistFieldBranch(v string) bool {
	return slices.Contains(sc.listAcceptedFields, v)
}
