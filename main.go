package main

import (
	"log"
	"cachee/watch"
)

// // /registry/v1/namespace/puckel

func main() {

	log.Println("Cachee started!")

	watch.Watch("/registry/v1/namespaces/test", 0, false)

}
