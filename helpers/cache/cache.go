package cache

import (
	"container/list"
	"fmt"
	"sync"
)

const MAX_SIZE int = 10

var ErrKeyNotFound error = fmt.Errorf("key not found")

type Cache struct {
	order    list.List
	cMap     map[string][]byte
	cMapLock *sync.Mutex
}

func NewCache() Cache {
	cache := Cache{
		cMapLock: &sync.Mutex{},
		cMap:     make(map[string][]byte, MAX_SIZE),
		order:    *list.New(),
	}
	return cache
}

func (c *Cache) Remove(key string) error {
	c.cMapLock.Lock()
	if _, isInMap := c.find(key); !isInMap {
		return ErrKeyNotFound
	}

	delete(c.cMap, key)
	// start at the back - more likely to find the element
	for e := c.order.Back(); e != nil; e = e.Prev() {
		if e.Value == key {
			c.order.Remove(e)
		}
	}
	c.cMapLock.Unlock()
	return nil
}

func (c *Cache) Set(key string, value []byte) {
	var keyRemoval string

	c.cMapLock.Lock()

	c.cMap[key] = value
	c.order.PushFront(key)

	if len(c.cMap) > MAX_SIZE {
		keyRemoval = c.order.Back().Value.(string)
	}

	c.cMapLock.Unlock()

	if keyRemoval != "" {
		c.Remove(keyRemoval)
	}
}

func (c *Cache) Get(key string) (bool, []byte) {
	c.cMapLock.Lock()
	val, isInMap := c.find(key)
	if isInMap {
		for e := c.order.Back(); e != nil; e = e.Prev() {
			if e.Value == key {
				c.order.MoveToFront(e)
			}
		}
	}
	c.cMapLock.Unlock()
	if !isInMap {
		return false, nil
	}
	return true, val
}

func (c *Cache) Clear() {
	c.cMapLock.Lock()
	c.cMap = make(map[string][]byte, MAX_SIZE)
	c.order = *c.order.Init()
	c.cMapLock.Unlock()
}

func (c *Cache) find(key string) (val []byte, isInMap bool) {
	val, isInMap = c.cMap[key]
	return
}
