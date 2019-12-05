package main

import (
	"context"
	"log"
	"time"
	"cachee/client"
)

func main() {
	
	cli := client.GetETCDClient()
	defer func() { _ = cli.Close() }()

	log.Println("put key {sample_key} : {sample_value}")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	putResp, err := cli.Put(ctx, "sample_key", "sample_value")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(putResp)
	}
	cancel()

	log.Println("get key {sample_key}")
	if resp, err := cli.Get(context.TODO(), "sample_key"); err != nil {
		log.Fatal(err)
	} else {
		log.Println("resp: ", resp)
	}

	// // 这个会一直阻塞
	// ctx = context.Background()
	// rch := cli.Watch(ctx, "sample_key")
	// for wresp := range rch {
	// 	for _, ev := range wresp.Events {
	// 		log.Println("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	// 	}
	// }

	log.Println("delete {sample_key}")
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	delResp, err := cli.Delete(ctx, "sample_key")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(delResp)
	}
	cancel()
}
