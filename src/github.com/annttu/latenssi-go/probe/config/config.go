package config

type Destination struct {
	Address string `yaml:"address"`
	Probe string `yaml:"probe"`
}

type Collector struct {
	Address string `yaml:"address"`
}

type ProbeConfig struct {
	Hostname string `yaml:"hostname"`
	Destinations []Destination `yaml:"destinations"`
	Collector Collector `yaml:"collector"`
	Probes map[string]map[string]interface{}
}
