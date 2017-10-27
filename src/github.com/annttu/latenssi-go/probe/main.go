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



	/*probetypes["fping"] = func(host string) probe.Probe {
		var p probe.Probe
		p = &probe.Fping{
			Command: "/usr/local/bin/fping",
			Host: host,
			Interval: 100,
		}
		return p
	}

	probetypes["fping6"] = func(host string) probe.Probe {
		var p probe.Probe
		p = &probe.Fping{
			Command: "/usr/local/bin/fping6",
			Host: host,
			Interval: 100,
		}
		return p
	}*/

	var wg sync.WaitGroup

	var runners []*probe.ProbeRunner = []*probe.ProbeRunner{}

	for _, destination := range parsedConfig.Destinations {
		if _, ok := probetypes[destination.Probe]; !ok {
			log.Printf("Skipped invalid probe type %s", destination.Probe)
			continue
		}
		p := probetypes[destination.Probe](destination.Address, 60)
		runner := &probe.ProbeRunner{Probe: p}
		go runner.Run()
		wg.Add(1)
		runners = append(runners, runner)
	}

	wg.Wait()
}
