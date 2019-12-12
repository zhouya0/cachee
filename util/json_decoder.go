package util

import (
	"encoding/json"
)

type TestObject struct {
	RegistryName string      `json:"registry_name"`
	Name         interface{} `json:"name"`
}

func GetTestObject(data []byte) (TestObject, error) {
	var t TestObject
	err := json.Unmarshal(data, &t)
	return t, err
}
