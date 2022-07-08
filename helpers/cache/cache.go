package cache

import (
	"container/list"
	"fmt"
	"sync"
)

var MAX_SIZE int = 10

var ErrKeyNotFound error = fmt.Errorf("key not found")

type Cache struct {
	queue *list.List
	cMap  map[string][]byte
	lock  *sync.Mutex
}

func NewCache() Cache {
	cache := Cache{
		lock:  &sync.Mutex{},
		cMap:  make(map[string][]byte, MAX_SIZE),
		queue: list.New(),
	}
	return cache
}

func (c *Cache) Remove(key string) error {
	c.lock.Lock()
	if _, isInMap := c.find(key); !isInMap {
		return ErrKeyNotFound
	}

	delete(c.cMap, key)
	for e := c.queue.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			c.queue.Remove(e)
		}
	}
	c.lock.Unlock()
	return nil
}

func (c *Cache) Set(key string, value []byte) {
	if c.queue.Len() >= MAX_SIZE {
		c.Remove(c.queue.Front().Value.(string))
	}

	c.lock.Lock()

	c.cMap[key] = value
	c.queue.PushBack(key)

	c.lock.Unlock()
}

func (c *Cache) Get(key string) (bool, []byte) {
	c.lock.Lock()
	val, isInMap := c.find(key)
	if isInMap {
		for e := c.queue.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				c.queue.MoveToBack(e)
			}
		}
	}
	c.lock.Unlock()
	if !isInMap {
		return false, nil
	}
	return true, val
}

func (c *Cache) Clear() {
	c.lock.Lock()
	c.cMap = make(map[string][]byte, MAX_SIZE)
	c.queue = c.queue.Init()
	c.lock.Unlock()
}

func (c *Cache) find(key string) (val []byte, isInMap bool) {
	val, isInMap = c.cMap[key]
	return
}
