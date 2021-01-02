package file

// FileSystem 代表一个抽象的文件系统
type FileSystem interface {
	Roots() []Path
	GetPath(path string) Path
	Separator() string
	SeparatorChar() rune
	PathSeparator() string
	PathSeparatorChar() rune
}

// Path 代表一个路径
type Path interface {
	Parent() Path
	FileSystem() FileSystem

	Path() string
	Name() string
	Exists() bool
	IsDir() bool
	IsFile() bool

	// for file
	Size() int64

	// for dir
	Mkdir() bool
	Mkdirs() bool
	List() []string // 返回短文件名
	ListChildren() []Path
}
