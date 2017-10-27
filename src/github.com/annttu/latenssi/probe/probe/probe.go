package probe


import (
	"context"
	"os/exec"
	"time"
	"fmt"
	"io"
	"sync"
	"bufio"
	"strings"
	"github.com/annttu/latenssi/probe/result"
	"github.com/annttu/latenssi/probe/grpc"
)


type ProbeRunner struct {
	Probe Probe
	wg sync.WaitGroup
	lineChannel chan string
	resultChannel chan *result.Result
}


type Probe interface {
	GetName() string
	GetCommand() string
	GetArgs() []string
	Parser(string)(*result.Result, error)
}


func (runner *ProbeRunner) reader(p io.ReadCloser) {
	runner.wg.Add(1)
	defer runner.wg.Done()
	//defer close(runner.stderrChannel)

	var line string
	var err error
	myreader := bufio.NewReader(p)
	for {
		line, err = myreader.ReadString('\n')

		if err != nil {
				if err == io.EOF {
						return
				}
				fmt.Printf("Got error %v", err)
				return
		}
		runner.lineChannel <- line
	}
}

func (runner *ProbeRunner) parser() {
	runner.wg.Add(1)
	defer runner.wg.Done()
	defer close(runner.resultChannel)
	var line string
	var ok bool
	var r *result.Result
	var err error
	for {
		line, ok = <- runner.lineChannel
		if !ok {
			fmt.Print("Failed to read from channel\n")
			return
		}
		//fmt.Print(line)
		line = strings.TrimRight(line, "\n")

		r, err = runner.Probe.Parser(line)

		if err != nil {
			fmt.Printf("Got error from parser: %s\n", err)
		}
		if r != nil {
			runner.resultChannel <- r
		}

		r = nil
	}
}

func (runner *ProbeRunner) reporter() {
	runner.wg.Add(1)
	defer runner.wg.Done()
	var result *result.Result
	var ok bool
	for {
		result, ok = <- runner.resultChannel
		if !ok {
			fmt.Print("Failed to read from channel\n")
			return
		}
		fmt.Printf("%s\n", result)
		result.Probe = runner.Probe.GetName()
		grpc.SendResults(result)
	}
}

func (runner *ProbeRunner) execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, runner.Probe.GetCommand(), runner.Probe.GetArgs()...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Failed to open stderr for command, %v\n", err)
		return err
	}

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Printf("Failed to open stdout for command, %v\n", err)
		return err
	}

	err = cmd.Start()

	if err != nil {
			fmt.Printf("Failed to run command %s, %v\n", cmd, err)
			return err
	}

	go runner.reader(stdout)
	go runner.reader(stderr)

	err = cmd.Wait()
	//fmt.Printf("Command exited with error %v\n", err)
	return err
}

func (runner *ProbeRunner) Run() {
	// Initialize resources
	runner.lineChannel = make(chan string, 100)
	runner.resultChannel = make(chan *result.Result, 100)
	go runner.parser()

	go runner.reporter()

	var err error
	for {
		err = runner.execute()
		if err != nil {
			fmt.Printf("Failed to execute command: %v\n", err)
			// Wait minute and retry
			<- time.After(60 * time.Second)
		}
	}
	runner.wg.Wait()
}

type ProbeInitiator func(name string, config map[string]interface{}) ProbeFunction

type ProbeFunction func(host string, interval uint64) Probe
