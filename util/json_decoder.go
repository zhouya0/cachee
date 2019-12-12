package util

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
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

func GetObjectToMap(data []byte) (map[string]interface{}, error) {
	m, ok := gjson.Parse(string(data)).Value().(map[string]interface{})
	if !ok {
		return m, fmt.Errorf("Decoder failed due to")
	}
	return m, nil
}
