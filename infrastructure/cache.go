package infrastructure

import (
	"errors"
)

type Cacher struct {
	InMemoryCache  map[int]interface{}
	LastCacheIndex int
}

func NewCacher() *Cacher {
	return &Cacher{
		InMemoryCache:  make(map[int]interface{}),
		LastCacheIndex: 0,
	}
}

const maxCache = 1024

func (c *Cacher) Add(data interface{}) int {
	idx := c.LastCacheIndex
	idx += 1
	if idx >= maxCache {
		idx = 0
	}
	c.InMemoryCache[idx] = data
	return idx
}

func (c *Cacher) Get(id int) (interface{}, error) {
	res, ok := c.InMemoryCache[id]
	if !ok {
		return nil, errors.New("Invalid ID")
	}
	return res, nil
}
