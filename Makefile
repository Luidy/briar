BASE ?= $(CURDIR)
export GO111MODULE=on

.PHONY: proto
proto:
ifeq ($(shell which protoc),)
	$(info no protoc complier. install protoc)
else
	$(info gen file from .proto)
	$Q cd $(BASE)/idl && \
	protoc -I. \
	--go_out=../model \
	--go_opt=paths=source_relative *.proto
endif