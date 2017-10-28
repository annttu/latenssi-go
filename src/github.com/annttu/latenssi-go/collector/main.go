package main

import (
	"github.com/annttu/latenssi-go/collector/grpc"
	"flag"
	"github.com/annttu/latenssi-go/collector/config"
	"log"
	"github.com/annttu/latenssi-go/collector/destination"
)

var (
	configFilePtr *string = flag.String("config", "collector.yaml", "Config file path")
)

func main() {

	flag.Parse()

	parsedConfig, err := config.ParseFile(*configFilePtr)

	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err.Error())
	}

	x := &destination.Influxdb{
		Address: parsedConfig.Influxdb.Address,
		Username: parsedConfig.Influxdb.Username,
		Password: parsedConfig.Influxdb.Password,
		Database: parsedConfig.Influxdb.Database,
	}
	x.Connect()
	destination.AddDestination(x)

	grpc.RunServer()
}
