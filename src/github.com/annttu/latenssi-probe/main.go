package main

import (
	"fmt"
	"github.com/annttu/latenssi-probe/probe"
)

func main() {
	fmt.Printf("Foo\n")

	p := &probe.Probe{
		Command: "/usr/local/bin/fping",
		Args: []string{"-Q", "1", "-c", "10", "8.8.8.8", "8.8.4.4" ,"1.2.3.4"},
	}
	for {

		p.Execute()
	}

}