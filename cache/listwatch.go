package cache

import (
	"cachee/client"
	"cachee/watch"
	"log"

	"go.etcd.io/etcd/clientv3"
)

type listWatchCache struct {
	client *clientv3.Client
	key    string
	cache  *Cache
}

func NewListWatchCache(key string) *listWatchCache {
	etcdClient := client.GetETCDClient()
	//	defer etcdClient.Close()
	c := NewCache()

	return &listWatchCache{
		client: etcdClient,
		key:    key,
		cache:  c,
	}

}

// ListWatch will do the real basic things:
// 1. List -> Cache
// 2. Watch -> Cache
func ListWatch(key string) *Cache {
	l := NewListWatchCache(key)
	err := l.fillCacheWithList()
	if err != nil {
		return nil
	}
	go l.fillCacheWithWatch()
	return l.cache
}

func (l *listWatchCache) fillCacheWithList() error {
	items, err := watch.List(l.client, l.key, 0, true)
	for _, item := range items {
		l.cache.Add(item)
	}
	return err
}

func (l *listWatchCache) fillCacheWithWatch() error {
	watchChan, err := watch.Watch(l.client, l.key, l.cache.resourceVersion, true)
	for res := range watchChan.ResultChan() {
		log.Println(res)
		if res.Type == watch.Added {
			l.cache.Add(res)
		} else if res.Type == watch.Modified {
			l.cache.Update(res)
		} else if res.Type == watch.Deleted {
			l.cache.Delete(res)
		}
	}
	return err
}
