GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PROTO_SRC_DIR=./proto
PLUGIN_DIR=$(shell go env GOPATH)

all: protoc build

protoc:
	protoc --proto_path=$(PROTO_SRC_DIR) $(PROTO_SRC_DIR)/*proto --plugin=$(PLUGIN_DIR)/bin/protoc-gen-go --go_out=./headquarters
	protoc --proto_path=$(PROTO_SRC_DIR) $(PROTO_SRC_DIR)/*proto --plugin=$(PLUGIN_DIR)/bin/protoc-gen-go-grpc --go-grpc_out=./headquarters
	protoc --proto_path=$(PROTO_SRC_DIR) $(PROTO_SRC_DIR)/*proto --plugin=$(PLUGIN_DIR)/bin/protoc-gen-go --go_out=./ship
	protoc --proto_path=$(PROTO_SRC_DIR) $(PROTO_SRC_DIR)/*proto --plugin=$(PLUGIN_DIR)/bin/protoc-gen-go-grpc --go-grpc_out=./ship
build:
	$(GOBUILD) -o ship -v ./ship
	$(GOBUILD) -o headquarters -v ./headquarters
clean:
	rm -f ./ship/ship
	rm -f ./headquarters/headquarters