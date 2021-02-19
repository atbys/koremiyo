package cache

import "github.com/atbys/koremiyo/usecase"

type MovieCache struct {
	Cache
}

type Cache interface {
	Add(interface{}) int
	Get(int) (interface{}, error)
}

type CacheData struct {
	MovieList usecase.List
	Index     int
}

func (mc *MovieCache) Store(l usecase.List, index int) int {
	ret_id := mc.Add(CacheData{MovieList: l, Index: index})
	return ret_id
}

func (mc *MovieCache) FindById(id int) (usecase.List, int, error) {
	data, err := mc.Get(id)
	movieList := data.(CacheData).MovieList
	index := data.(CacheData).Index
	return movieList, index, err
}
