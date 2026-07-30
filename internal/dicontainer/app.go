package dicontainer

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"

// NewDIContainer ленивая инициализация DI контейнера
func NewDIContainer(rootDir string, ch chan interfaces.Messager) *DiContainer {
	return &DiContainer{
		rootDir: rootDir,
		ch:      ch,
	}
}
