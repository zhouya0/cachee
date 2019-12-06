package main

import (
	"cachee/client"
	"cachee/watch"
	"log"
	"time"
)

// // /registry/v1/namespace/puckel

func main() {

	log.Println("Cachee started!")

	etcdClient := client.GetETCDClient()

	defer etcdClient.Close()

	watchChan, _ := watch.Watch(etcdClient, "/registry/v1/namespaces/test", 0, false)

	for res := range watchChan.ResultChan() {
		log.Println(res)
	}
	time.Sleep(3 * time.Second)
}
