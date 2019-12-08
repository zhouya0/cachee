package main

import (
	"cachee/cache"
	"cachee/client"
	"cachee/watch"
	"log"
	"time"
)

// /registry/v1/namespace/puckel
// ssh 10.6.170.191

func main() {
	log.Println("Cachee started!")

	etcdClient := client.GetETCDClient()

	c := cache.NewCache()

	defer etcdClient.Close()

	watchChan, _ := watch.Watch(etcdClient, "/registry/v1/namespaces", 0, true)

	for res := range watchChan.ResultChan() {
		log.Println(res)
		if res.Type == watch.Added {
			c.Add(res)
		} else if res.Type == watch.Modified {
			c.Update(res)
		} else if res.Type == watch.Deleted {
			c.Delete(res)
		}
		log.Println(c.ListKeys())
		log.Println(c.List())
	}
	time.Sleep(3 * time.Second)
}
