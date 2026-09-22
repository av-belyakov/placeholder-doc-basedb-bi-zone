package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForSnapshotsType вспомогательный тип используемый для хранения объектов типа 'snapshots'
type SupportingStructureForSnapshotsType struct {
	listAcceptedFields []string
	snapshotTmp        *datamodels.BiZoneIRPSnapshot
	snapshots          []datamodels.BiZoneIRPSnapshot
}

// NewSupportingStructureForSnapshotsTypeType формирует вспомогательный объект для обработки объектов типа 'snapshots'
func NewSupportingStructureForSnapshotsType() *SupportingStructureForSnapshotsType {
	return &SupportingStructureForSnapshotsType{
		listAcceptedFields: []string(nil),
		snapshotTmp:        datamodels.NewBiZoneIRPSnapshot(),
		snapshots:          make([]datamodels.BiZoneIRPSnapshot, 0),
	}
}

// GetSnapshots возвращает []datamodels.BiZoneIRPSnapshot
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из s.snapshotTmp в
// s.snapshots, так как snapshots автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются s.snapshotTmp
func (s *SupportingStructureForSnapshotsType) GetSnapshots() []datamodels.BiZoneIRPSnapshot {
	if s.snapshotTmp != nil {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&s.snapshotTmp)
		s.snapshots = append(s.snapshots, *s.snapshotTmp)

		s.snapshotTmp = nil
		s.listAcceptedFields = []string(nil)
	}

	return s.snapshots
}

// GetSnapshotTmp возвращает временный объект
func (s *SupportingStructureForSnapshotsType) GetSnapshotTmp() *datamodels.BiZoneIRPSnapshot {
	return s.snapshotTmp
}

// HandlerValue функция обработчик значений
func (s *SupportingStructureForSnapshotsType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if s.isExistFieldBranch((fieldBranch)) {
		s.listAcceptedFields = []string(nil)

		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&s.snapshotTmp)
		s.snapshots = append(s.snapshots, *s.snapshotTmp)
		s.snapshotTmp = datamodels.NewBiZoneIRPSnapshot()
	}

	s.listAcceptedFields = append(s.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (s *SupportingStructureForSnapshotsType) isExistFieldBranch(v string) bool {
	return slices.Contains(s.listAcceptedFields, v)
}
