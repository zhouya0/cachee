package cache

import (
	"cachee/util"
	"cachee/watch"
	"fmt"
	"k8s.io/client-go/tools/cache"
)

type Cache struct {
	store           cache.Store
	resourceVersion uint64
}

func (c *Cache) NewCache() *Cache {
	cache := &Cache{
		store:           cache.NewStore(getCacheObjectKey),
		resourceVersion: 0,
	}
	return cache
}

func (c *Cache) Add(event watch.CacheEvent) error {
	return c.store.Add(event.Object)
}

// 接下来的任务，搞懂这个到底能不能行？

func (c *Cache) List() []interface{} {
	return c.store.List()
}

func main() {
	indexers := cache.Indexers{
		"bla": func(obj interface{}) (strings []string, e error) {
			indexes := []string{obj.(string)}
			return indexes, nil
		},
	}
	indices := cache.Indices{}
	store := cache.NewThreadSafeStore(indexers, indices)
	for {
		key := "test"
		store.Add(key, key)
		store.Delete(key)
	}
}

func getCacheObjectKey(obj interface{}) (string, error) {
	cacheEvent, ok := obj.(*util.KeyVersionObject)
	if !ok {
		return "", fmt.Errorf("not a cacheEvent!%v", obj)
	}
	return cacheEvent.Key, nil
}
