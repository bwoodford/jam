package router

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/IveGotNorto/jam/helpers/cache"
	"github.com/IveGotNorto/jam/helpers/path"
	"github.com/IveGotNorto/jam/helpers/reader"
)

type DirEntry interface {
	Name() string
	IsDir() bool
}

type FileReader interface {
	ReadFile(string) ([]byte, error)
}

type Router struct {
	paths map[string]string
	// starting path
	root   string
	cache  cache.CacheWrapper
	reader reader.FileReader
	fpath  path.FilePath
}

func NewRouter(root string) Router {
	router := Router{
		root:   root,
		cache:  cache.NewCache(),
		paths:  make(map[string]string),
		reader: &reader.ReaderWrapper{},
		fpath:  &path.PathWrapper{},
	}
	router.Init()
	return router
}

func (r *Router) Init() {
	r.fpath.WalkDir(r.root, r.mapper)
}

func (r *Router) mapper(path string, info fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() {
		fileName := info.Name()

		// Add test case for root path showing up down the tree
		routerPath := strings.Replace(path, r.root, "", 1)

		if fileName == "index.gmi" {
			// Add index.gmi variation
			r.paths[routerPath] = path
			routerPath = strings.Replace(routerPath, "/index.gmi", "", 1)
			// Add directory variation
			r.paths[routerPath] = path
			// Add directory/ variation
			routerPath += "/"
			r.paths[routerPath] = path
		} else if string(fileName[len(fileName)-4:]) == ".gmi" {
			r.paths[routerPath] = path
		}
	}
	return nil
}

func (r *Router) Load(key string) ([]byte, error) {
	var err error
	val, isInMap := r.paths[key]
	if isInMap {
		// Check if value is in the cache
		inCache, file := r.cache.Get(key)
		if !inCache {
			file, err = r.reader.ReadFile(val)
			if err != nil {
				return nil, err
			}
			r.cache.Set(key, file)
		}
		return file, nil
	}
	return nil, fmt.Errorf("value not in router map")
}
