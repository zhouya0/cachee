package util

import (
	"encoding/json"
)

type TestObject struct {
	RegistryName string `json:"registry_name"`
	Name         string `json:"name"`
}

func GetTestObject(data []byte) (TestObject, error) {
	var t TestObject
	json.Unmarshal(data, &t)
	return t, nil
}
