package file

import (
	"path/filepath"
	"runtime"
)

type apiPlatform interface {
	Roots() []string
	PathSeparatorChar() rune
	SeparatorChar() rune
}

type innerFSCore struct {
	fs       FileSystem
	platform apiPlatform

	separator         string
	separatorChar     rune
	pathSeparator     string
	pathSeparatorChar rune
}

type innerPath struct {
	// impl Path
	core *innerFSCore
	path string
}

type innerFileSystem struct {
	// impl FileSystem
	core *innerFSCore
}

// impl innerFileSystem

// Default 创建一个默认的 FileSystem 实例
func Default() FileSystem {

	sys := runtime.GOOS
	var platform apiPlatform

	if sys == "windows" {
		platform = &innerWindowsPlatform{}
	} else {
		platform = &innerPosixPlatform{}
	}

	// create
	core := &innerFSCore{}
	fs := &innerFileSystem{}

	// binding
	core.fs = fs
	core.platform = platform
	core.pathSeparatorChar = platform.PathSeparatorChar()
	core.pathSeparator = string(platform.PathSeparatorChar())
	core.separatorChar = platform.SeparatorChar()
	core.separator = string(platform.SeparatorChar())

	fs.core = core

	return fs
}

// impl innerFileSystem

func (inst *innerFileSystem) GetPath(path string) Path {
	path, _ = filepath.Abs(path)
	return &innerPath{
		core: inst.core,
		path: path,
	}
}

func (inst *innerFileSystem) Roots() []Path {
	roots := inst.core.platform.Roots()
	list := make([]Path, len(roots))
	for index := range list {
		path := roots[index]
		list[index] = inst.GetPath(path)
	}
	return list
}

func (inst *innerFileSystem) Separator() string {
	return inst.core.separator
}

func (inst *innerFileSystem) SeparatorChar() rune {
	return inst.core.separatorChar
}

func (inst *innerFileSystem) PathSeparator() string {
	return inst.core.pathSeparator
}

func (inst *innerFileSystem) PathSeparatorChar() rune {
	return inst.core.pathSeparatorChar
}

// impl innerPath

func (inst *innerPath) Name() string {
	return ""
}

func (inst *innerPath) Path() string {
	return ""
}

func (inst *innerPath) Parent() Path {
	return nil
}

func (inst *innerPath) Exists() bool {
	return false
}

func (inst *innerPath) IsDir() bool {
	return false
}

func (inst *innerPath) IsFile() bool {
	return false
}

func (inst *innerPath) Mkdir() bool {
	return false
}

func (inst *innerPath) Mkdirs() bool {
	return false
}

func (inst *innerPath) Size() int64 {
	return 0
}

func (inst *innerPath) FileSystem() FileSystem {
	return inst.core.fs
}

func (inst *innerPath) List() []string {
	return nil
}

func (inst *innerPath) ListChildren() []Path {
	return nil
}
