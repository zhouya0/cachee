package main

import (
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

	defer etcdClient.Close()

	watchChan, _ := watch.Watch(etcdClient, "/registry/v1/namespaces", 0, true)

	for res := range watchChan.ResultChan() {
		log.Println(res)
	}
	time.Sleep(3 * time.Second)
}
