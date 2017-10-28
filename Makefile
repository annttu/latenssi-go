export GOPATH = $(shell pwd)
SOURCEDIR = src

sources := $(shell find $(SOURCEDIR) -name '*.go')
PROBE_BINARY := "bin/latenssi-probe"
COLLECTOR_BINARY := "bin/latenssi-collector"

GO_OPTS := 

#all: deps build
all: build

build: src/github.com/annttu/latenssi-go/proto/proto.pb.go $(PROBE_BINARY) $(COLLECTOR_BINARY)

deps:
	go get ${GO_OPTS} github.com/op/go-logging
	go get ${GO_OPTS} google.golang.org/grpc
	go get ${GO_OPTS} github.com/jessevdk/go-flags
	go get ${GO_OPTS} github.com/golang/protobuf/proto
	go get ${GO_OPTS} github.com/golang/protobuf/protoc-gen-go
	go get ${GO_OPTS} gopkg.in/yaml.v2
	go get ${GO_OPTS} github.com/influxdata/influxdb/client/v2

src/github.com/annttu/latenssi-go/proto/proto.pb.go: src/github.com/annttu/latenssi-go/proto/proto.proto
	protoc --plugin=$(GOPATH)/bin/protoc-gen-go -I .  $< --go_out=plugins=grpc:.

$(PROBE_BINARY): $(sources)
	go install github.com/annttu/latenssi-go/probe


$(COLLECTOR_BINARY): $(sources)
	go install github.com/annttu/latenssi-go/collector

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
