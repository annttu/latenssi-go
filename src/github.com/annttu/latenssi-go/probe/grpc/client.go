package grpc

import (
	"log"
	"google.golang.org/grpc"
	pb "github.com/annttu/latenssi-go/proto"
	"golang.org/x/net/context"
	"github.com/annttu/latenssi-go/probe/result"
	"fmt"
	"os"
)


var conn *grpc.ClientConn
var Hostname string

func InitializeConnection(serverAddress string) (err error) {



	if conn != nil {
		return fmt.Errorf("Connection is already initialized")
	}

	if Hostname == "" {
		Hostname, err = os.Hostname()
	}

	if err != nil {
		return err
	}

	conn, err = grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return nil
}

func GetConnection() *grpc.ClientConn {
	return conn
}

func CloseConnection() {
	conn.Close()
}

func SendResults(r *result.Result) error {
	c := pb.NewResultServiceClient(GetConnection())

	var results []*pb.ResultRow = make([]*pb.ResultRow, len(r.Results))

	for idx, res := range r.Results {
		if res == nil {
			// Skip nil results
			continue
		}
		k, v := res.Get()
		switch res.Type() {
		case result.ResultTypeInt64:
			results[idx] = &pb.ResultRow{
				Result: &pb.ResultRow_Intresult{
					Intresult: &pb.ResultInt64{Key: k, Value: v.(int64)},
				},
			}
		case result.ResultTypeFloat:
			results[idx] = &pb.ResultRow{
				Result: &pb.ResultRow_Floatresult{
					Floatresult: &pb.ResultFloat{Key: k, Value: v.(float64)},
				},
			}
		}

	}

	response, err := c.SendResults(context.Background(), &pb.Results{
		Source: Hostname,
		Host: r.Address,
		Probe: r.Probe,
		Resultrows: results,
		Time: uint64(r.Time.UnixNano()),
	})
	if err != nil {
		log.Printf("Failed to send Results to server: %v", err)
		return fmt.Errorf("Failed to send Results to server: %v", err)
	}

	if response.Status != pb.ResultStatus_RESULT_OK {
		log.Printf("Result save failed with message: %s", response.Message)
		return fmt.Errorf("Result save failed with message: %s", response.Message)
	}
	return nil
}