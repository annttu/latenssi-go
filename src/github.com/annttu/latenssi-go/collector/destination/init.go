package destination

var Destinations []DestinationWriter

func AddDestination(destination DestinationType) {

	writer := DestinationWriter{Destination: destination}
	Destinations = append(Destinations, writer)
}