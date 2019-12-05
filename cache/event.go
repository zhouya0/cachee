package cache

import (
	"cachee/watch"
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
	Type EventType
	Object Object
}

func ToCacheEvent(e *watch.etcdEvent) (c *CacheEvent) {
	curObj, oldObj, _ := prepareObjs(e)
	
	switch {
	case e.isDeleted:
		c = &CacheEvent{
			Type: Deleted,
			Object: oldObj,
		}
	
	case e.isCreated:
		c = &CacheEvent{
			Type: Added,
			Object: curObj,
		}
	
	default:
		c = &CacheEvent{
			Type: Modified,
			Object: curObj,
		}

	}
	return c
}

func prepareObjs(e *watch.etcdEvent) (curObj Object, oldObj Object, err error) {
	if !e.isDeleted {
		curObj, _ := util.GetTestObject(e.value)
	}

	if len(e.preValue) >0 {
		oldObj, _ := util.GetTestObject(e.preValue)
	}

	return curObj, oldObj, nil
}