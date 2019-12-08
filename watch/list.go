package watch

import (
	"cachee/meta"
	"context"
	"fmt"
	"log"
	"strings"

	"go.etcd.io/etcd/clientv3"
)

func List(etcdClient *clientv3.Client, key string, rev int64, recursive bool) []meta.Object {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if recursive && !strings.HasSuffix(key, "/") {
		key += "/"
	}
	opts := []clientv3.OpOption{clientv3.WithRev(rev)}
	if recursive {
		opts = append(opts, clientv3.WithPrefix())
	}
	items, err := etcdClient.Get(ctx, key, opts...)
	if err != nil {
		log.Fatal(err)
	}
	version := items.Header.Revision
	var objects []meta.Object
	for _, ev := range items.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		object := meta.NewKeyVersionObject(ev.Value, version, string(ev.Key))
		objects = append(objects, object)
	}
	return objects
}
