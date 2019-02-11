package cache

type Cache interface {
	Get(key int) int
	Put(key int, value int)
}
