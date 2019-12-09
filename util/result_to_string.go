package util

import (
	"cachee/meta"
	"fmt"
)

func PrintKeyVersionObjects(objects []interface{}) {
	for _, object := range objects {
		kv, _ := object.(meta.KeyVersionObject)
		fmt.Printf("Data is: %s\n", kv.Data)
		fmt.Printf("Version is: %d\n", kv.Version)
		fmt.Printf("Key is: %s\n", kv.Key)
	}
}
