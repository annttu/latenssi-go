PROBE_BINARY := "bin/latenssi-probe"
COLLECTOR_BINARY := "bin/latenssi-collector"

GO_OPTS := 

all: build

build: proto/proto.pb.go $(PROBE_BINARY) $(COLLECTOR_BINARY)
proto/proto.pb.go: proto/proto.proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --plugin=./bin/protoc-gen-go -I .  $< --go_out=plugins=grpc:.

$(PROBE_BINARY):
	go build -o $(PROBE_BINARY) probe.go


$(COLLECTOR_BINARY):
	go build -o $(COLLECTOR_BINARY) collector.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
