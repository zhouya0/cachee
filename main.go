package main

import (
	"log"
)

func main() {

	log.Println("Cachee started!")

	watch.Watch("sample_key", 0, false)
}
