package destination

import (
	"github.com/annttu/latenssi-go/proto"
	"time"
)

type DestinationType interface {
	Connect()
	WritePoints(source string, host string, probe string, timestamp time.Time, points []*proto.ResultRow) error
}

type DestinationWriter struct {
	Destination DestinationType
	queue chan *proto.Results
}

func (d *DestinationWriter) Init() {
	d.queue = make(chan *proto.Results, 1000)
	go d.writeWorker()
}

func (d *DestinationWriter) writeWorker() {
	var result *proto.Results
	for {
		result = <- d.queue
		if result == nil {
			return
		}
		d.Destination.WritePoints(result.Source, result.Host, result.Probe, time.Unix(0, int64(result.Time)), result.GetResultrows())
	}
}

func (d *DestinationWriter) Write(result *proto.Results) error {
	d.queue <- result
	return nil
}
