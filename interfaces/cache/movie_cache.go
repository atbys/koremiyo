package cache

type MovieCache struct {
	Cache
}

func (mc *MovieCache) Store(data interface{}) int {
	id := mc.Add(data)
	return id
}

func (mc *MovieCache) Find(id int) (interface{}, error) {
	data, err := mc.Get(id)
	return data, err
}
