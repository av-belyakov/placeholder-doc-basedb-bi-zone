package handlers

import (
	"fmt"
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForDataSecurityType вспомогательный тип для обработки 'data.data_security'
type SupportingStructureForDataSecurityType struct {
	currentKey         string
	listAcceptedFields []string
	dataSecurityTmp    datamodels.BiZoneIRPDataSecurity
	dataSecurityList   map[string][]datamodels.BiZoneIRPDataSecurity
}

// NewSupportingStructureForDataSecurityType формирует вспомогательный объект для обработки объектов типа 'data.data_security'
func NewSupportingStructureForDataSecurityType() *SupportingStructureForDataSecurityType {
	return &SupportingStructureForDataSecurityType{
		listAcceptedFields: []string(nil),
		dataSecurityTmp:    *datamodels.NewBiZoneIRPDataSecurity(),
		dataSecurityList:   make(map[string][]datamodels.BiZoneIRPDataSecurity),
	}
}

// GetDataSecurity возвращает map[string][]datamodels.BiZoneIRPDataSecurity
// Ооднако, метод выполняет еще очень важное действие, перемещает содержимое из ds.dataSecurityTmp в
// ds.dataSecurityList, так как dataSecurityList автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются ds.dataSecurityTmp
func (ds *SupportingStructureForDataSecurityType) GetDataSecurity() map[string][]datamodels.BiZoneIRPDataSecurity {
	if ds.currentKey != "" {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&ds.dataSecurityTmp)
		ds.dataSecurityList[ds.currentKey] = append(ds.dataSecurityList[ds.currentKey], ds.dataSecurityTmp)
	}

	ds.currentKey = ""
	ds.listAcceptedFields = []string(nil)
	ds.dataSecurityTmp = *datamodels.NewBiZoneIRPDataSecurity()

	return ds.dataSecurityList
}

// GetDataSecurityTmp возвращает временный объект
func (ds *SupportingStructureForDataSecurityType) GetDataSecurityTmp() *datamodels.BiZoneIRPDataSecurity {
	return &ds.dataSecurityTmp
}

// HandlerValue функция обработчик значений
func (ds *SupportingStructureForDataSecurityType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	// делаем значение поля 'data.data_security.i_sid' ключём списка
	if fieldBranch == "data.data_security.i_sid" {
		str := fmt.Sprint(a)
		if _, ok := ds.dataSecurityList[str]; !ok {
			ds.dataSecurityList[str] = []datamodels.BiZoneIRPDataSecurity{}
		}

		if ds.isExistFieldBranch(fieldBranch) {
			ds.listAcceptedFields = []string(nil)
			// здесь можно выполнять постобработку некоторых пользовательский типов
			// например, изменить содержимое какого нибудь поля.
			// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
			// на основании которых было бы понятно что нужно менять.
			//_, _ = supportingfunctions.PostProcessingUserType(&ds.dataSecurityTmp)
			ds.dataSecurityList[ds.currentKey] = append(ds.dataSecurityList[ds.currentKey], ds.dataSecurityTmp)

			ds.dataSecurityTmp = *datamodels.NewBiZoneIRPDataSecurity()
		}

		ds.currentKey = str
	}

	//если поле повторяется, кроме некоторых, то считается что это уже новый объект
	if fieldBranch != "data.data_security.s_content" && ds.isExistFieldBranch((fieldBranch)) {
		ds.listAcceptedFields = []string(nil)

		if _, ok := ds.dataSecurityList[ds.currentKey]; !ok {
			ds.dataSecurityList[ds.currentKey] = []datamodels.BiZoneIRPDataSecurity(nil)
		}

		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&ds.dataSecurityTmp)
		ds.dataSecurityList[ds.currentKey] = append(ds.dataSecurityList[ds.currentKey], ds.dataSecurityTmp)

		ds.dataSecurityTmp = *datamodels.NewBiZoneIRPDataSecurity()
	}

	ds.listAcceptedFields = append(ds.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (ds *SupportingStructureForDataSecurityType) isExistFieldBranch(v string) bool {
	return slices.Contains(ds.listAcceptedFields, v)
}
