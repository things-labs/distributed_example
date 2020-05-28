#!/usr/bin/env bash

protoc --proto_path=./pb --go_out=plugins=grpc:./services ./pb/*.proto
protoc --proto_path=./pb --grpc-gateway_out=logtostderr=true:./services ./pb/*.proto