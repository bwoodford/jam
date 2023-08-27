package router

import (
	"io/fs"
	"testing"

	"github.com/IveGotNorto/jam/helpers/cache"
)

type MockDirEntry struct {
	name  string
	isDir bool
}

func (mde *MockDirEntry) Name() string {
	return mde.name
}

func (mde *MockDirEntry) IsDir() bool {
	return mde.isDir
}

// not used
func (mde *MockDirEntry) Type() fs.FileMode {
	return fs.FileMode(int(0777))
}

// not used
func (mde *MockDirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

type MockReader struct{}

func (mr *MockReader) ReadFile(path string) ([]byte, error) {
	return nil, nil
}

type MockFilePath struct{}

func (mfp *MockFilePath) WalkDir(root string, fn fs.WalkDirFunc) error {
	base := "/something/in/front/to/mimick/file/path"
	tests := []struct {
		path  string
		entry fs.DirEntry
	}{
		{path: base + "/", entry: &MockDirEntry{name: "path", isDir: true}},
		{path: base + "/index.gmi", entry: &MockDirEntry{name: "index.gmi", isDir: false}},
		{path: base + "/wee.gmi", entry: &MockDirEntry{name: "wee.gmi", isDir: false}},
		{path: base + "/woo.gmi", entry: &MockDirEntry{name: "woo.gmi", isDir: false}},
		{path: base + "/number/number.gmi", entry: &MockDirEntry{name: "number.gmi", isDir: false}},
		{path: base + "/number/", entry: &MockDirEntry{name: "number", isDir: true}},
		{path: base + "/number/index.gmi", entry: &MockDirEntry{name: "index.gmi", isDir: false}},
		{path: base + "/about/", entry: &MockDirEntry{name: "about", isDir: true}},
		{path: base + "/about/index.gmi", entry: &MockDirEntry{name: "index.gmi", isDir: false}},
		{path: base + "/about/about.gmi", entry: &MockDirEntry{name: "about.gmi", isDir: false}},
		{path: base + "/foo/", entry: &MockDirEntry{name: "foo", isDir: true}},
		{path: base + "/foo/bar/", entry: &MockDirEntry{name: "bar", isDir: true}},
		{path: base + "/foo/bar/baz.gmi", entry: &MockDirEntry{name: "baz.gmi", isDir: false}},
		{path: base + "/foo/bar/index.gmi", entry: &MockDirEntry{name: "index.gmi", isDir: false}},
	}

	for _, t := range tests {
		fn(t.path, t.entry, nil)
	}

	return nil
}

func TestInit(t *testing.T) {

	router := Router{
		root:   "/something/in/front/to/mimick/file/path",
		cache:  cache.NewCache(),
		paths:  make(map[string]string),
		reader: &MockReader{},
		fpath:  &MockFilePath{},
	}
	router.Init()

}
