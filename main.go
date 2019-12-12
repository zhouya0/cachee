package main

import (
	"cachee/cache"
	"fmt"
	"time"
)

// /registry/v1/namespace/puckel
// ssh 10.6.170.191

func main() {
	c := cache.ListWatch("/registry/v1/namespaces")
	for {
		fmt.Println("============================Cache DATA")
		objects := c.List()
		//util.PrintKeyVersionObjects(objects)
		fmt.Println(objects)
		fmt.Println("============================")
		time.Sleep(3 * time.Second)
	}
}
