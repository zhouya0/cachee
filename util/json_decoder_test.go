package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	testJSON := []byte(`
	{
		"registry_name": "test1", 
		"name": "12"
	}
	`)
	object, err := GetTestObject(testJSON)
	if err != nil {
		t.Errorf("Decoder error occurs %v", err)
	}
	fmt.Println(reflect.TypeOf(object.Name))
	if object.RegistryName != "test1" {
		t.Fatal("Decoder is wrong!")
	}
}
