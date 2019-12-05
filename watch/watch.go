package watch 

import (
	"go.etcd.io/etcd/clientv3"
	"context"
	"strings"
	"time"
	"log"
)

type Interface interface {
	Stop()
	
	ResultChan() <-chan Event
}


type watchChan struct {
	client *clientv3.Client
	key string
	initialRev int64
	recursive bool
	ctx context.context
	cancel context.CancelFunc
}


func Watch(key string, rev int64, recursive bool) (Interface, error) {
	if recursive && !strings.HasSuffix(key, "/") {
		key += "/"
	}
	
	wc := watchChan{}
}

func (wc *watchChan) Run() {
	watchClosedCh := make(chan struct{})
	go wc.StartWatching(StartWatching)

	time.Sleep(10 * time.Second)
	wc.cancel()
}


func (wc *watchChan) Stop() {
	wc.cancel()
}

func (wc *watchChan) ResultChan() {
	return wc.resultChan
}

func (wc *watchChan) StartWatching(watchClosedCh chan struct{}) {
	opts := []clientv3.OpOption{clientv3.WithRev(wc.initialRev + 1), clientv3.WithPreKV()}
	if wc.recursive {
		opts = append(opts, clientv3.WithPrefix())
	}

	wch := wc.client.Watch(wc.ctx, wc.key, opts...)

	for wres := range wch {
		if wres.Err() != nil {
			err := wres.Err()
			log.Fatal("watch chan error: %v", err)
			wc.sendError(err)
			return 
		}

		for _, e := range wres.Events {
			log.Println(e)
		}
	}
}
