package router

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/IveGotNorto/jam/helpers/cache"
)

type Router struct {
	// mapped out path structure
	paths map[string]string
	// starting path
	root  string
	cache cache.Cache
}

func NewRouter(root string) Router {
	cache := cache.NewCache()
	router := Router{
		root:  root,
		cache: cache,
		paths: make(map[string]string),
	}
	router.Init()
	return router
}

func (r *Router) Init() {
	r.traverse()
	fmt.Printf("%v", r.paths)
}

func (r *Router) traverse() {
	filepath.WalkDir(r.root, r.mapper)
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
			file, err = ioutil.ReadFile(val)
			if err != nil {
				return nil, err
			}
			r.cache.Set(key, file)
		}
		return file, nil
	}
	return nil, fmt.Errorf("value not in router map")
}
