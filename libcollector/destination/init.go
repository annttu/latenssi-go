package destination

var Destinations []DestinationWriter

func AddDestination(destination DestinationType) {

	writer := DestinationWriter{Destination: destination}
	writer.Init()
	Destinations = append(Destinations, writer)
}