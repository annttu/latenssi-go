package destination

import (
	"github.com/annttu/latenssi-go/proto"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
	"sync"
)

type Influxdb struct {
	Address string
	Username string
	Password string
	Database string
	connection client.Client
	mx sync.Mutex
}

func (i *Influxdb) Connect() {
	var err error
	i.connection, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:    i.Address,
		Username: i.Username,
		Password: i.Password,
	})
	if err != nil {
		log.Fatal(err)
	}

}

func (i *Influxdb) WritePoints(source string, host string, probe string, timestamp time.Time, points []*proto.ResultRow) error {

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	var fields  map[string]interface{} = make(map[string]interface{})

	for _, r := range points {
		if res := r.GetIntresult(); res != nil {
			fields[res.Key] = res.Value
			fmt.Printf("Writing to influxdb %s %s %s %s %v\n", source, host, probe, res.Key, res.Value)
		} else if res := r.GetFloatresult(); res != nil {
			fields[res.Key] = res.Value
			fmt.Printf("Writing to influxdb %s %s %s %s %v\n", source, host, probe, res.Key, res.Value)
		} else {
			continue
		}

	}

	/*fields := map[string]interface{}{
		"loss":   100,
		"send": 53.3,
		"min":   46.6,
	}*/

	// Create a point and add to batch
	tags := map[string]string{"host": host, "source": source}

	pt, err := client.NewPoint(probe, tags, fields, timestamp)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	i.mx.Lock()
	defer i.mx.Unlock()
	// Write the batch
	err = i.connection.Write(bp)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}