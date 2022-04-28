package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ParseFile(path string) (config *ProbeConfig, err error) {

	config = &ProbeConfig{
		Collector: Collector{Address:"127.0.0.1:50051"},
		Destinations: []Destination{},
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(f, config)

	return
}

