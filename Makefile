BASE ?= $(CURDIR)
GOENV ?= GOOS=darwin
GOBUILD = $(GOENV) go

export GO111MODULE=on

.PHONY: proto
proto:
ifeq ($(shell which protoc),)
	$(info no protoc complier. install protoc)
else
	$(info gen file from .proto)
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	export PATH=$PATH:~/go/bin
	$Q cd $(BASE)/idl && \
	protoc -I. \
	--go_out=../idl \
	--go-grpc_out=require_unimplemented_servers=false:. \
	--go_opt=paths=source_relative \
	*.proto
endif

.PHONY: build
build:
	$(info build ges)
	cd $(BASE) && $(GOBUILD) build -o $(BASE)/bin/ges
