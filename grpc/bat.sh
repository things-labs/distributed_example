#!/usr/bin/env bash

protoc --go_out=./pb --go-grpc_out=./pb ./pb/*.proto
# protoc --proto_path=./pb --grpc-gateway_out=logtostderr=true:./pb ./pb/*.proto