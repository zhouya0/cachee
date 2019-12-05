package util

import (
	"testing"
)
func TestDecode(t *testing.T) {
	testJSON := []byte(`
	{
		"registry_name": "test1", 
		"name": "puckel"
	}
	`)
	object, _ := GetTestObject(testJSON)
	if object.RegistryName != "test1" {
		t.Fatal("Decoder is wrong!")
	}
}