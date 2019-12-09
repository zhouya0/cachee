package meta

type KeyVersionObject struct {
	Data    []byte
	Version int64
	Key     string
}

func NewKeyVersionObject(data []byte, version int64, key string) KeyVersionObject {
	return KeyVersionObject{
		Data:    data,
		Version: version,
		Key:     key,
	}
}
