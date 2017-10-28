package destination

import "github.com/annttu/latenssi-go/proto"

type DestinationType interface {
	Connect()
	WritePoints(source string, host string, probe string, points []*proto.ResultRow) error
}

type DestinationWriter struct {
	Destination DestinationType
}

func (d *DestinationWriter) Write(result *proto.Results) error {
	return d.Destination.WritePoints(result.Source, result.Host, result.Probe, result.GetResultrows())
}