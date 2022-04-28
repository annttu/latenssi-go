package libprobe

import (
	"time"
	"fmt"
	"sync"
	"context"
	"errors"
	"github.com/annttu/latenssi-go/libprobe/result"
	"github.com/annttu/latenssi-go/libprobe/grpc"
)

type FunctionProbeRunner struct {
	Probe FunctionProbe
	wg sync.WaitGroup
	lineChannel chan string
	resultChannel chan *result.Result

}


type FunctionProbe interface {
	GetName() string
	Test(ctx context.Context, resultChannel chan *result.Result) error
}


func (runner *FunctionProbeRunner) reporter() {
	runner.wg.Add(1)
	defer runner.wg.Done()
	var res *result.Result
	var ok bool
	for {
		select {
			case res, ok = <- runner.resultChannel:
				if !ok {
					fmt.Print("Failed to read from channel\n")
					return
				}
				fmt.Printf("%s\n", res)
				res.Probe = runner.Probe.GetName()
				grpc.SendResults(res)
			case <-time.After(time.Second * 60):

		}

	}
}

func (runner *FunctionProbeRunner) execute() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error

	var wg sync.WaitGroup

	go func () {
		err = runner.Probe.Test(ctx, runner.resultChannel)
	}()

	wg.Add(1)


	var res *result.Result
	var ok bool
	for {
		select {
		case res, ok = <- runner.resultChannel:
			if !ok {
				fmt.Print("Failed to read from channel\n")
				return errors.New("failed to read from result channel")
			}
			fmt.Printf("%s\n", res)
			res.Probe = runner.Probe.GetName()
			grpc.SendResults(res)
		case <-time.After(time.Second * 60):
			// Cancel
			fmt.Print("No results in 60 seconds, timeout\n")
			return errors.New("timeout occurred")
		}
	}

	wg.Wait()

	return err
}

func (runner *FunctionProbeRunner) Run() {
	// Initialize resources
	runner.lineChannel = make(chan string, 100)
	runner.resultChannel = make(chan *result.Result, 100)

	go runner.reporter()

	var err error
	for {
		err = runner.execute()
		if err != nil {
			fmt.Printf("Failed to execute function: %v\n", err)
			// Wait minute and retry
			<- time.After(60 * time.Second)
		}
	}
	runner.wg.Wait()
}
