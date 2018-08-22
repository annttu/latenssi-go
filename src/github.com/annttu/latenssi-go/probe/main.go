package main

import (
	"github.com/annttu/latenssi-go/probe/probe"
	"sync"
	"github.com/annttu/latenssi-go/probe/grpc"
	"github.com/annttu/latenssi-go/probe/config"
	"flag"
	"log"
	"fmt"
)

var (
	configFilePtr *string = flag.String("config", "probe.yaml", "Config file path")
)

func main() {

	var probetypes map[string]probe.ProbeFunction = make(map[string]probe.ProbeFunction)

	flag.Parse()

	parsedConfig, err := config.ParseFile(*configFilePtr)

	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err.Error())
	}

	if parsedConfig.Hostname != "" {
		grpc.Hostname = parsedConfig.Hostname
	}

	for name, p := range parsedConfig.Probes {

		probeType, ok := p["probe"]
		if !ok {
			panic(fmt.Sprintf("Probe type missing from probe %s", name))
		}

		initiator, ok := probe.Initiators[probeType.(string)]
		if !ok {
			panic(fmt.Sprintf("Probe type %s is not supported", probetypes))
		}

		probetypes[name] = initiator(name, p)

	}

	grpc.InitializeConnection(parsedConfig.Collector.Address)
	defer grpc.CloseConnection()

	var wg sync.WaitGroup

	var runners []probe.ProbeRunner = []probe.ProbeRunner{}

	for _, destination := range parsedConfig.Destinations {
		if len(destination.Probes) == 0 {
			log.Printf("No probes configured for address %s", destination.Address)
		}
		for _, probeType := range  destination.Probes {
			if _, ok := probetypes[probeType]; !ok {
				log.Printf("Skipped invalid probe type %s", probeType)
				continue
			}
			var runner probe.ProbeRunner = probetypes[probeType](destination.Address, 300)
			go runner.Run()
			wg.Add(1)
			runners = append(runners, runner)
		}
	}

	wg.Wait()
}
