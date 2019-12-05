package watch

import (
	"go.etcd.io/etcd/clientv3"
	"fmt"
)

type etcdEvent struct {
	key string
	value []byte
	preValue []byte
	rev int64
	isDeleted bool
	isCreated bool
}


func toETCDEvent(e *clientv3.Event) (*etcdEvent, error) {
	if !e.IsCreate() && e.PrevKv == nil {
		return nil, fmt.Errorf("etcd event received with PrevKv=nil (key=%q, modRevision=%d, type=%s)", string(e.Kv.Key), e.Kv.ModRevision, e.Type.String())
	}

	ret := &etcdEvent {
		key: string(e.Kv.Key),
		value: e.Kv.Value,
		rev: e.Kv.ModRevision,
		isDeleted: e.Type == clientv3.EventTypeDelete,
		isCreated: e.IsCreate(),
	}

	if e.PrevKv != nil {
		ret.preValue = e.PrevKv.Value
	}

	return ret, nil
} 