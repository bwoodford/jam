package path

import (
	"io/fs"
	"path/filepath"
)

type FilePath interface {
	WalkDir(string, fs.WalkDirFunc) error
}

type PathWrapper struct {
}

func (pw *PathWrapper) WalkDir(root string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(root, fn)
}
