package watch

import (
	"cachee/cache"
	"cachee/util"
)

type EventType string

const (
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
	Error    EventType = "ERROR"

	DefaultChanSize int32 = 100
)

type CacheEvent struct {
	Type   EventType
	Object cache.Object
}

func toCacheEvent(e *etcdEvent) (c *CacheEvent) {
	curObj, oldObj, _ := prepareObjs(e)

	switch {
	case e.isDeleted:
		c = &CacheEvent{
			Type:   Deleted,
			Object: oldObj,
		}

	case e.isCreated:
		c = &CacheEvent{
			Type:   Added,
			Object: curObj,
		}

	default:
		c = &CacheEvent{
			Type:   Modified,
			Object: curObj,
		}

	}
	return c
}

func prepareObjs(e *etcdEvent) (curObj cache.Object, oldObj cache.Object, err error) {
	if !e.isDeleted {
		curObj, _ = util.GetTestObject(e.value)
	}

	if len(e.preValue) > 0 {
		oldObj, _ = util.GetTestObject(e.preValue)
	}

	return curObj, oldObj, nil
}
