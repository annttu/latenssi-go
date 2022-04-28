package grpc

import (
	"golang.org/x/net/context"
	pb "github.com/annttu/latenssi-go/proto"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/annttu/latenssi-go/libcollector/destination"
)

const (
	listenAddress = ":50051"
)

type collector struct{}

func (s *collector) SendResults(ctx context.Context, res *pb.Results) (*pb.ResultResponse, error) {
	log.Printf("Got result from client")
	for _, r := range res.Resultrows {
		log.Printf("%s: %s: %s: %v", res.Source, res.Host, res.Probe, r.GetResult())
	}

	for _, d := range destination.Destinations {
		err := d.Write(res)
		if err != nil {
			log.Printf("Error: Failed to write points to database: %v", err.Error())
		}
	}

	return &pb.ResultResponse{Status: pb.ResultStatus_RESULT_OK, Message: ""}, nil
}


func RunServer() error {
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", listenAddress)
	s := grpc.NewServer()
	pb.RegisterResultServiceServer(s, &collector{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
