package dicontainer

// NewDIContainer ленивая инициализация DI контейнера
func NewDIContainer(rootDir string) *DiContainer {
	return &DiContainer{
		rootDir: rootDir,
	}
}
