package handlers

import (
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

// SupportingStructureForWatchersType вспомогательный тип используемый для хранения объектов типа 'watchers'
type SupportingStructureForWatchersType struct {
	watchers           []datamodels.WatchersType
	listAcceptedFields []string
	watchersTmp        datamodels.WatchersType
	isCompleted        bool
}

// NewSupportingStructureForWatchersType формирует вспомогательный объект для обработки объектов типа 'watchers'
func NewSupportingStructureForWatchersType() *SupportingStructureForWatchersType {
	return &SupportingStructureForWatchersType{
		listAcceptedFields: []string(nil),
		watchersTmp:        *datamodels.NewWatcherType(),
		watchers:           make([]datamodels.WatchersType, 0),
	}
}

// GetWatchers возвращает []datamodels.WatchersType
// Однако, метод выполняет еще очень важное действие, перемещает содержимое из w.watchersTmp в
// w.watchers, так как tags автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются w.watchersTmp
func (w *SupportingStructureForWatchersType) GetWatchers() []datamodels.WatchersType {
	if !w.isCompleted {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&w.watchersTmp)
		w.watchers = append(w.watchers, w.watchersTmp)
		w.listAcceptedFields = []string(nil)
		w.isCompleted = true
	}

	return w.watchers
}

// GetWatchersTmp возвращает временный объект
func (w *SupportingStructureForWatchersType) GetWatchersTmp() datamodels.WatchersType {
	return w.watchersTmp
}

// HandlerValue функция обработчик значений
func (w *SupportingStructureForWatchersType) HandlerValue(fieldBranch string, a any, f func(any) error) error {
	//если поле повторяется то считается что это уже новый объект
	if w.isExistFieldBranch((fieldBranch)) {
		// здесь можно выполнять постобработку некоторых пользовательский типов
		// например, изменить содержимое какого нибудь поля.
		// Однако, пока данная функция не чего полезного не выполяет так как нет вводных
		// на основании которых было бы понятно что нужно менять.
		//_, _ = supportingfunctions.PostProcessingUserType(&sc.tagTmp)
		w.watchers = append(w.watchers, w.watchersTmp)
		w.watchersTmp = *datamodels.NewWatcherType()
		w.listAcceptedFields = []string(nil)
	}

	w.listAcceptedFields = append(w.listAcceptedFields, fieldBranch)

	return f(a)
}

// isExistFieldBranch проверка существования ветки
func (w *SupportingStructureForWatchersType) isExistFieldBranch(v string) bool {
	return slices.Contains(w.listAcceptedFields, v)
}
