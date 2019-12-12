package util

// import (
// 	"cachee/meta"
// 	"fmt"
// )

// func PrintKeyVersionObjects(objects []interface{}) error {
// 	for _, object := range objects {
// 		kv, ok := object.(meta.KeyVersionObject)
// 		if !ok {
// 			return fmt.Errorf("can't convert %T to key version object", object)
// 		}
// 		fmt.Printf("Data is: %s\n", kv.Data)
// 		fmt.Printf("Version is: %d\n", kv.Version)
// 		fmt.Printf("Key is: %s\n", kv.Key)
// 	}
// 	return nil
// }
