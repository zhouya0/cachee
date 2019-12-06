package watch 

import (
	"go.etcd.io/etcd/clientv3"
	"context"
	"strings"
	"sync"
	"log"
)

const (
	// We have set a buffer in order to reduce times of context switches.
	etcdEventChanBufSize = 100
	cacheEventChanBufSize = 100
)


type Interface interface {
	Stop()
	ResultChan() <-chan CacheEvent
}


type watchChan struct {
	client *clientv3.Client
	key string
	initialRev int64
	recursive bool
	ctx context.Context
	cancel context.CancelFunc
	etcdEventChan chan *etcdEvent
	cacheEventChan chan CacheEvent
	errChan chan error
}


func Watch(etcdClient *clientv3.Client, key string, rev int64, recursive bool) (*watchChan, error) {
	if recursive && !strings.HasSuffix(key, "/") {
		key += "/"
	}

	ctx, cancel := context.WithCancel(context.Background())
	wc := &watchChan{
		client: etcdClient,
		key: key,
		initialRev: rev,
		ctx: ctx,
		cancel: cancel,
		etcdEventChan: make(chan *etcdEvent, etcdEventChanBufSize),
		cacheEventChan: make(chan CacheEvent, cacheEventChanBufSize),
		errChan: make(chan error, 1),
	}
	go wc.Run()

	return wc, nil
}

func (wc *watchChan) Run() {
	watchClosedCh := make(chan struct{})
	go wc.StartWatching(watchClosedCh)

	var resultChanWG sync.WaitGroup
	resultChanWG.Add(1)
	go wc.processEvent(&resultChanWG)

	select {
	case err := <-wc.errChan:
		if err == context.Canceled {
			break
		}
	case <- watchClosedCh:
	case <-wc.ctx.Done():
	}

	wc.cancel()
	resultChanWG.Wait()
	close(wc.cacheEventChan)
}


func (wc *watchChan) Stop() {
	wc.cancel()
}

func (wc *watchChan) ResultChan() <-chan CacheEvent{
 	return wc.cacheEventChan
}

func (wc *watchChan) StartWatching(watchClosedCh chan struct{}) {
	log.Println("Start watching")
	// WithPrevKV will store the previous revision and value, it's must be setted.
	opts := []clientv3.OpOption{clientv3.WithRev(wc.initialRev), clientv3.WithPrevKV()}
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
			log.Printf("Event received! %s executed on %q with value %q\n", e.Type, e.Kv.Key, e.Kv.Value)
			log.Printf("The revision is %d", e.Kv.ModRevision)
			etcdEvent,_ := toETCDEvent(e)
			log.Println(etcdEvent)
			wc.sendETCDEvent(etcdEvent)
		}
	}

	close(watchClosedCh)
}


func (wc *watchChan) processEvent(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		log.Println("precocess running!")
		select {
		case e := <-wc.etcdEventChan:
			res := toCacheEvent(e)
			select {
			case wc.cacheEventChan <- *res:
			case <-wc.ctx.Done():
				return
			}
		
		case <-wc.ctx.Done():
			return 
		}
	}
}

func (wc *watchChan)sendETCDEvent(e *etcdEvent) {
	if len(wc.etcdEventChan) == etcdEventChanBufSize {
		log.Printf("etcd event chan buffer size %d is not enough!", etcdEventChanBufSize)
	}

	select {
	case wc.etcdEventChan <- e:
	case <-wc.ctx.Done():
	}
}

func (wc *watchChan)sendError(err error) {
	select {
	case wc.errChan <- err:
	case <-wc.ctx.Done():
	}
}