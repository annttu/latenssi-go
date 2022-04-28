package config

type Influxdb struct {
	Address string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type CollectorConfig struct {
	Influxdb Influxdb `yaml:"influxdb,omitempty"`
}