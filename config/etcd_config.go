package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Config struct {
	CAPath   string `yaml:"caPath"`
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
	Server   string `yaml:"server"`
}

func (c *Config) GetConfig() *Config {
	absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := path.Join(absPath, "etcd.conf")
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.UnmarshalStrict(yamlFile, c)

	if err != nil {
		log.Fatal(err)
	}
	return c
}
