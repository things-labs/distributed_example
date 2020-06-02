#!/usr/bin/env bash

protoc --go_out=plugins=grpc:./services ./pb/*.proto