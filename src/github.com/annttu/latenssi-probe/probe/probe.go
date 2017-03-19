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


type Probe struct {
        Command string
        Args []string
	wg sync.WaitGroup
}

type Result struct {
	Address string
	Min float64
	Max float64
	Avg float64
	Send int64
	Received int64
	Loss int64
}


func (r *Result) String () (string) {
	return fmt.Sprintf("%s min/max/avg: %f/%f/%f send/rcv/loss%%: %d/%d/%d", r.Address, r.Min, r.Max, r.Avg, r.Send,
		           r.Received, r.Loss)
}

func (probe *Probe) reader(p io.ReadCloser, out chan string) {
	probe.wg.Add(1)
        defer probe.wg.Done()
	defer close(out)

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
                // fmt.Printf("%s", line)
                out <- line
        }
}

func (probe *Probe)parser(in chan string, out chan *Result) {
	probe.wg.Add(1)
	defer probe.wg.Done()
	defer close(out)
	var line string
	var ok bool
	var err error
	var result *Result
	for {
		line, ok = <- in
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
			result.Address = parts[0]
			result.Send, err = strconv.ParseInt(xmtrcvloss[0], 10, 64)
			if err != nil {
				fmt.Printf("Failed to parse send packets from row %s\n", line)
				continue
			}
			result.Received, err = strconv.ParseInt(xmtrcvloss[1], 10, 64)
			if err != nil {
				fmt.Printf("Failed to parse received packets from row %s\n", line)
				continue
			}
			if result.Received > 0 {
				result.Loss = (result.Send * 100)  / result.Received
			} else {
				result.Loss = 100
			}

		}
		if len(parts) == 8 {
			// 8.8.8.8 : xmt/rcv/%loss = 1/1/0%, min/avg/max = 38.6/38.6/38.6
			if parts[1] != ":" {
				fmt.Print("Invalid line from fping\n")
				continue
			}
			minmaxavg := strings.Split(parts[7], "/")
			result.Min, err = strconv.ParseFloat(minmaxavg[0], 64)
			if err != nil {
				fmt.Printf("Failed to parse min from row %s\n", line)
				continue
			}
			result.Max, err = strconv.ParseFloat(minmaxavg[2], 64)
			if err != nil {
				fmt.Printf("Failed to parse max from row %s\n", line)
				continue
			}
			result.Avg, err = strconv.ParseFloat(minmaxavg[1], 64)
			if err != nil {
				fmt.Printf("Failed to parse avg from row %s\n", line)
				continue
			}
		}
		if result == nil {
			fmt.Printf("Got invalid amount of parts from line %s\n", line)
			continue
		}
		out <- result
		result = nil
	}
}

func (probe *Probe) reporter(in chan *Result) {
	probe.wg.Add(1)
	defer probe.wg.Done()
	var result *Result
	var ok bool
	for {
		result, ok = <- in
		if !ok {
			fmt.Print("Failed to read from channel\n")
			return
		}
		fmt.Printf("%s\n", result)
	}
}

func (probe *Probe) Execute() {
        ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
        defer cancel()

        cmd := exec.CommandContext(ctx, probe.Command, probe.Args...)
        stderr, err := cmd.StderrPipe()
        if err != nil {
                fmt.Errorf("Failed to open stderr for command, %v\n", err)
        }

        err = cmd.Start()

        if err != nil {
                fmt.Errorf("Failed to run command %s, %v", cmd, err)
                return
        }

        //stdoutChannel := make(chan string, 10)
        stderrChannel := make(chan string, 10)
	resultChannel := make(chan *Result, 100)


        //go reader(stdout, &wg, stdoutChannel)
        go probe.reader(stderr, stderrChannel)

	//go parser(stdoutChannel, &wg)
	go probe.parser(stderrChannel, resultChannel)

	go probe.reporter(resultChannel)

        err = cmd.Wait()
        fmt.Printf("Command exited with error %v\n", err)
        probe.wg.Wait()
}

