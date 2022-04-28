package libprobe

import (
	"github.com/miekg/dns"
	"context"
	"time"
	"github.com/annttu/latenssi-go/libprobe/result"
)

type DNSProbe struct {
	Name string
	Interval uint64
	Host string
	Query string
	QueryType string
}

func (f *DNSProbe) GetName() string {
	return f.Name
}

func (f *DNSProbe) Test(ctx context.Context, resultChannel chan *result.Result) error {
	c := dns.Client{}
	m := dns.Msg{}

	var queryType uint16 = dns.TypeA
	if f.QueryType == "AAAA" {
		queryType = dns.TypeAAAA
	}

	m.SetQuestion(f.Query+".", queryType)

	var responseReceived int64 = 0
	var numberOfResponses int64 = 0

	for {

		r, t, err := c.Exchange(&m, f.Host+":53")
		if err != nil {
			responseReceived = 0
		} else {
			responseReceived = 1
		}

		if r != nil {
			numberOfResponses = int64(len(r.Answer))
		} else {
			numberOfResponses = 0
		}


		resultChannel <- &result.Result{
			Address: f.Host,
			Time: time.Now(),
			Results: []result.ResultRow{
				&result.ResultRowInt64{Key: "result", Value: responseReceived},
				&result.ResultRowInt64{Key: "rtt", Value: t.Nanoseconds() / 1000000},
				&result.ResultRowInt64{Key: "num_responses", Value: numberOfResponses},
			},
		}

		time.Sleep(1 * time.Second)

	}
}

func init() {
	if Initiators == nil {
		Initiators = make(map[string]ProbeInitiator)
	}
	Initiators["dns"] = func(name string, config map[string]interface{}) ProbeFunction {
		return func(host string, interval uint64) ProbeRunner {
			var queryDomain, queryType string
			cInterface, ok := config["query"]
			if !ok {
				queryDomain = "a.root-servers.net"
			} else {
				queryDomain = cInterface.(string)
			}
			cInterface, ok = config["query"]
			if !ok {
				queryType = "A"
			} else {
				queryType = cInterface.(string)
			}
			var p = &DNSProbe{
				Name:     name,
				Host:     host,
				Query:    queryDomain,
				QueryType: queryType,
				Interval: interval,
			}
			var runner ProbeRunner = &FunctionProbeRunner{Probe: p}
			return runner
		}
	}
}
