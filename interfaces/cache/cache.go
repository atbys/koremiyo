package cache

type Cache interface {
	Add(interface{}) int
	Get(int) (interface{}, error)
}
