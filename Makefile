BASE ?= $(CURDIR)

.PHONY: proto
proto:
ifeq ($(shell which protoc),)
	$(info no protoc complier. install protoc)
else
	$(info gen file from .proto)
	cd $(BASE)/idl && \
	protoc -I. \
	--go_out=plugins=grpc:$(BASE)/model 
endif