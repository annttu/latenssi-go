package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ParseFile(path string) (config *CollectorConfig, err error) {

	config = &CollectorConfig{
		Influxdb: Influxdb{Address:"127.0.0.1:50051", Username:"", Password: ""},
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(f, config)

	return
}
