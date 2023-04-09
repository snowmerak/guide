package cache

import (
	"context"
	"github.com/snowmerak/guide/src/cache/tuple"
)

type Cache[K comparable, V any] struct {
	m map[K]V

	get chan tuple.Tuple[K, func(V)]
	set chan tuple.Tuple[K, V]
	del chan K
}

func New[K comparable, V any](ctx context.Context) *Cache[K, V] {
	m := make(map[K]V)

	const size = 32
	get := make(chan tuple.Tuple[K, func(V)], size)
	set := make(chan tuple.Tuple[K, V], size)
	del := make(chan K, size)

	go func() {
		prevCap := cap(get)
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-get:
				key, callback := t.Unpack()
				callback(m[key])
			case t := <-set:
				key, value := t.Unpack()
				m[key] = value
				prevCap = len(m)
			case k := <-del:
				delete(m, k)
				if prevCap > len(m)*2 {
					newM := make(map[K]V, len(m))
					for k, v := range m {
						newM[k] = v
					}
					m = newM
					prevCap = len(m)
				}
			}
		}
	}()

	return &Cache[K, V]{m, get, set, del}
}

func (c *Cache[K, V]) Get(key K) V {
	var value V
	ch := make(chan struct{})
	c.get <- tuple.New(key, func(v V) {
		value = v
		close(ch)
	})
	<-ch
	return value
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.set <- tuple.New(key, value)
}

func (c *Cache[K, V]) Del(key K) {
	c.del <- key
}
