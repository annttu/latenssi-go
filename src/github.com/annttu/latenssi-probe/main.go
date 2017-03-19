package main

import (
	"fmt"
	"github.com/annttu/latenssi-probe/probe"
	"sync"
)

func main() {
	fmt.Printf("Foo\n")

	var wg sync.WaitGroup

	p := &probe.Probe{
		Command: "/usr/local/bin/fping",
		Args: []string{"-p", "100", "-Q", "1", "-c", "10", "8.8.8.8", "8.8.4.4" ,"1.2.3.4"},
	}
	go p.Run()
	wg.Add(1)

	p2 := &probe.Probe{
		Command: "/usr/local/bin/fping",
		Args: []string{"-p", "100", "-Q", "1", "-c", "10", "annttu.fi", "www.remod.fi" ,"www.kapsi.fi"},
	}
	go p2.Run()
	wg.Add(1)

	wg.Wait()

}