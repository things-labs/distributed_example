#!/usr/bin/env bash

protoc --go_out=./pb --go-grpc_out=./pb ./pb/*.proto