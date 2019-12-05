package watch 

import (
	"go.etcd.io/etcd/clientv3"
	"context"
	"strings"
	"time"
	"log"
	"cachee/client"
)

type Interface interface {
	Stop()
	
	// ResultChan() <-chan Event
}


type watchChan struct {
	client *clientv3.Client
	key string
	initialRev int64
	recursive bool
	ctx context.Context
	cancel context.CancelFunc
}


func Watch(key string, rev int64, recursive bool) (*watchChan, error) {
	if recursive && !strings.HasSuffix(key, "/") {
		key += "/"
	}
	etcdClient := client.GetETCDClient()

	defer etcdClient.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wc := &watchChan{
		client: etcdClient,
		key: key,
		initialRev: rev,
		ctx: ctx,
		cancel: cancel,
	}
	go wc.Run()

	time.Sleep(20 * time.Second)

	return wc, nil
}

func (wc *watchChan) Run() {
	watchClosedCh := make(chan struct{})
	log.Println("Start run")
	go wc.StartWatching(watchClosedCh)

	// var resultChanWG sync.WaitGroup
	// resultChanWG.Add(1)
	time.Sleep(20 * time.Second)
	wc.cancel()
}


func (wc *watchChan) Stop() {
	wc.cancel()
}

// func (wc *watchChan) ResultChan() {
// 	return wc.resultChan
// }

func (wc *watchChan) StartWatching(watchClosedCh chan struct{}) {
	log.Println("Start watching")

	opts := []clientv3.OpOption{clientv3.WithRev(wc.initialRev)}
	if wc.recursive {
		opts = append(opts, clientv3.WithPrefix())
	}

	wch := wc.client.Watch(wc.ctx, wc.key, opts...)


	for wres := range wch {
		// if wres.Err() != nil {
		// 	err := wres.Err()
		// 	log.Fatal("watch chan error: %v", err)
		// 	wc.sendError(err)
		// 	return 
		// }

		for _, e := range wres.Events {
			log.Println(e)
			log.Println("Event received! %s executed on %q with value %q\n", e.Type, e.Kv.Key, e.Kv.Value)
		}
	}
}
