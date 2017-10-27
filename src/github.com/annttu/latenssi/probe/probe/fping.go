package probe

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/annttu/latenssi/probe/result"
	"time"
)


type Fping struct {
	Name string
	Command string
	Interval uint64
	Host string
	previousLineWasTimestamp bool
}

func (f *Fping) GetName() string {
	return f.Name
}

func (f *Fping) GetCommand() string {
	return f.Command
}

func (f *Fping) GetArgs() []string {
	return []string{"-p", fmt.Sprintf("%d", f.Interval), "-Q", "5", "-c", fmt.Sprintf("%d", 60000 / f.Interval), f.Host}
}

func (f *Fping) Parser(line string) (r *result.Result, err error) {
	if strings.HasPrefix(line, "[") {
		// This is timestamp, let's ignore it
		f.previousLineWasTimestamp = true
		return nil, nil
	}

	// Skip summary lines
	if !f.previousLineWasTimestamp {
		fmt.Printf("Skipping summary line %s", line)
		return nil, nil
	}

	parts := strings.Split(line, " ")
	var loss int64

	if len(parts) >= 5 {

		// 1.2.3.4 : xmt/rcv/%loss = 1/0/100%
		xmtrcvloss := strings.Split(parts[4], "/")
		r = new(result.Result)
		r.Probe = "fping"
		r.Time = time.Now()
		if len(parts) == 8 {
			r.Results = make([]result.ResultRow, 6)
		} else {
			r.Results = make([]result.ResultRow, 3)
		}
		r.Address = parts[0]
		send, err := strconv.ParseInt(xmtrcvloss[0], 10, 64)
		if send == 100 {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to parse send packets from row %s, %v", line, err)
		}
		received, err := strconv.ParseInt(xmtrcvloss[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse received packets from row %s, %v", line, err)
		}
		loss = 0
		if received > send {
			received = send
		}
		if received > 0 {
			loss = ((send - received) * 100)  / send
		} else {
			loss = 100
		}

		r.Results[0] = &result.ResultRowInt64{Key: "send", Value: send}
		r.Results[1] = &result.ResultRowInt64{Key: "received", Value: received}
		r.Results[2] = &result.ResultRowInt64{Key: "loss", Value: loss}
	}
	if len(parts) == 8 {
		// 8.8.8.8 : xmt/rcv/%loss = 1/1/0%, min/avg/max = 38.6/38.6/38.6
		if parts[1] != ":" {
			return nil, fmt.Errorf("Invalid line from fping: %s", line)
		}
		minmaxavg := strings.Split(parts[7], "/")
		min, err := strconv.ParseFloat(minmaxavg[0], 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse min from row %s", line)
		}
		max, err := strconv.ParseFloat(minmaxavg[2], 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse max from row %s", line)
		}
		avg, err := strconv.ParseFloat(minmaxavg[1], 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse avg from row %s", line)
		}
		r.Results[3] = &result.ResultRowFloat64{Key: "min", Value: min}
		r.Results[4] = &result.ResultRowFloat64{Key: "max", Value: max}
		r.Results[5] = &result.ResultRowFloat64{Key: "avg", Value: avg}
	}
	f.previousLineWasTimestamp = false

	if r == nil {
		return nil, fmt.Errorf("Got invalid amount of parts from line %s\n", line)
	}
	return r, nil
}

func init() {
	if Initiators == nil {
		Initiators = make(map[string]ProbeInitiator)
	}
	Initiators["fping"] = func(name string, config map[string]interface{}) ProbeFunction {
		var command string
		commandInterface, ok := config["command"]
		if !ok {
			command="/usr/bin/fping"
		} else {
			command = commandInterface.(string)
		}
		return func(host string, interval uint64) Probe {
			var p Probe = &Fping{
				Name:     name,
				Command:  command,
				Host:     host,
				Interval: interval,
			}
			return p
		}
	}
}