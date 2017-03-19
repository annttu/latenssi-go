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
	"strconv"
)


type ProbeRunner struct {
	Probe Probe
	wg sync.WaitGroup
	stderrChannel chan string
	resultChannel chan *Result
}


type Probe interface {
	GetCommand() string
	GetArgs() []string
	Parser(string)(*Result, error)
}

type Fping struct {
	Command string
	Args []string
}

func (f *Fping) GetCommand() string {
	return f.Command
}

func (f *Fping) GetArgs() []string {
	return f.Args
}

func (f *Fping) Parser(line string) (*Result, error) {

}

func (probe *ProbeRunner) reader(p io.ReadCloser) {
	probe.wg.Add(1)
        defer probe.wg.Done()
	//defer close(probe.stderrChannel)

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
                probe.stderrChannel <- line
        }
}

func (probe *ProbeRunner) parser() {
	probe.wg.Add(1)
	defer probe.wg.Done()
	defer close(probe.resultChannel)
	var line string
	var ok bool
	var result *Result
	var loss int64
	for {
		line, ok = <- probe.stderrChannel
		if !ok {
			fmt.Print("Failed to read from channel\n")
			return
		}
		fmt.Print(line)
		line = strings.TrimRight(line, "\n")
		if strings.HasPrefix(line, "[") {
			// This is timestamp, let's ignore it
			continue
		}

		parts := strings.Split(line, " ")

		if len(parts) >= 5 {

			// 1.2.3.4 : xmt/rcv/%loss = 1/0/100%
			xmtrcvloss := strings.Split(parts[4], "/")
			result = new(Result)
			result.Results = make([]ResultRow, 6)
			result.Address = parts[0]
			send, err := strconv.ParseInt(xmtrcvloss[0], 10, 64)
			if err != nil {
				fmt.Printf("Failed to parse send packets from row %s\n", line)
				continue
			}
			received, err := strconv.ParseInt(xmtrcvloss[1], 10, 64)
			if err != nil {
				fmt.Printf("Failed to parse received packets from row %s\n", line)
				continue
			}
			loss = 0
			if received > 0 {
				loss = (send * 100)  / received
			} else {
				loss = 100
			}
			result.Results[0] = &ResultRowInt64{Key: "send", Value: send}
			result.Results[1] = &ResultRowInt64{Key: "received", Value: received}
			result.Results[2] = &ResultRowInt64{Key: "loss", Value: loss}
		}
		if len(parts) == 8 {
			// 8.8.8.8 : xmt/rcv/%loss = 1/1/0%, min/avg/max = 38.6/38.6/38.6
			if parts[1] != ":" {
				fmt.Print("Invalid line from fping\n")
				continue
			}
			minmaxavg := strings.Split(parts[7], "/")
			min, err := strconv.ParseFloat(minmaxavg[0], 64)
			if err != nil {
				fmt.Printf("Failed to parse min from row %s\n", line)
				continue
			}
			max, err := strconv.ParseFloat(minmaxavg[2], 64)
			if err != nil {
				fmt.Printf("Failed to parse max from row %s\n", line)
				continue
			}
			avg, err := strconv.ParseFloat(minmaxavg[1], 64)
			if err != nil {
				fmt.Printf("Failed to parse avg from row %s\n", line)
				continue
			}
			result.Results[3] = &ResultRowFloat64{Key: "min", Value: min}
			result.Results[4] = &ResultRowFloat64{Key: "max", Value: max}
			result.Results[5] = &ResultRowFloat64{Key: "avg", Value: avg}
		}
		if result == nil {
			fmt.Printf("Got invalid amount of parts from line %s\n", line)
			continue
		}
		probe.resultChannel <- result
		result = nil
	}
}

func (probe *ProbeRunner) reporter() {
	probe.wg.Add(1)
	defer probe.wg.Done()
	var result *Result
	var ok bool
	for {
		result, ok = <- probe.resultChannel
		if !ok {
			fmt.Print("Failed to read from channel\n")
			return
		}
		fmt.Printf("%s\n", result)
	}
}

func (probe *ProbeRunner) execute() {
        ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
        defer cancel()

        cmd := exec.CommandContext(ctx, probe.Probe.GetCommand(), probe.Probe.GetArgs()...)
        stderr, err := cmd.StderrPipe()
        if err != nil {
                fmt.Errorf("Failed to open stderr for command, %v\n", err)
        }

        err = cmd.Start()

        if err != nil {
                fmt.Errorf("Failed to run command %s, %v", cmd, err)
                return
        }

        go probe.reader(stderr)

        err = cmd.Wait()
        fmt.Printf("Command exited with error %v\n", err)
}

func (probe *ProbeRunner) Run() {
	// Initialize resources
	probe.stderrChannel = make(chan string, 10)
	probe.resultChannel = make(chan *Result, 100)
	go probe.parser()

	go probe.reporter()
	for {
		probe.execute()
	}
	probe.wg.Wait()
}
