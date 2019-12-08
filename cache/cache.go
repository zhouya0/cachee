package cache

import (
	"cachee/meta"
	"cachee/watch"
	"fmt"
	"log"

	"k8s.io/client-go/tools/cache"
)

type Cache struct {
	store           cache.Store
	resourceVersion int64
}

func NewCache() *Cache {
	cache := &Cache{
		store:           cache.NewStore(getCacheObjectKey),
		resourceVersion: 0,
	}
	return cache
}

func (c *Cache) Add(event watch.CacheEvent) error {
	version, err := getCacheObjectVersion(event.Object)
	log.Printf("Add called %d", version)
	if err != nil {
		log.Fatal(err)
	}
	c.resourceVersion = version
	return c.store.Add(event.Object)
}

func (c *Cache) Update(event watch.CacheEvent) error {
	version, _ := getCacheObjectVersion(event.Object)
	log.Printf("Updated called %d", version)
	c.resourceVersion = version
	return c.store.Update(event.Object)
}

func (c *Cache) Delete(event watch.CacheEvent) error {
	version, _ := getCacheObjectVersion(event.Object)
	log.Printf("Deleted called %d", version)
	c.resourceVersion = version
	return c.store.Delete(event.Object)
}

func (c *Cache) List() []interface{} {
	return c.store.List()
}

func (c *Cache) ListKeys() []string {
	return c.store.ListKeys()
}

func getCacheObjectKey(obj interface{}) (string, error) {
	cacheEvent, ok := obj.(meta.KeyVersionObject)
	if !ok {
		return "", fmt.Errorf("not a cacheEvent!%v", obj)
	}
	return cacheEvent.Key, nil
}

func getCacheObjectVersion(obj interface{}) (int64, error) {
	cacheEvent, ok := obj.(meta.KeyVersionObject)
	if !ok {
		return 0, fmt.Errorf("not a cacheEßvent!%v", obj)
	}
	return cacheEvent.Version, nil
}
