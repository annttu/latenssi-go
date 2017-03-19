export GOPATH = $(shell pwd)
SOURCEDIR = src

sources := $(shell find $(SOURCEDIR) -name '*.go')
BINARY := "bin/latenssi-probe"

#all: deps build
all: build

build: $(BINARY)

deps:
	go get github.com/op/go-logging
	go get google.golang.org/grpc
	go get github.com/jessevdk/go-flags


$(BINARY): $(sources)
	go install github.com/annttu/latenssi-probe

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
