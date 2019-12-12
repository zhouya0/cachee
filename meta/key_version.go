package meta

import (
	"cachee/util"
)

type KeyVersionObject struct {
	//Data    []byte
	Data    map[string]interface{}
	Version int64
	Key     string
}

func NewKeyVersionObject(data []byte, version int64, key string) KeyVersionObject {
	mapObject, _ := util.GetObjectToMap(data)
	return KeyVersionObject{
		Data:    mapObject,
		Version: version,
		Key:     key,
	}
}
