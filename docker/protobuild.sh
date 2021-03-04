#!/bin/bash

protoc -I./api/proto -I./third_party --go_out=./api --go-grpc_out=./api ./api/proto/*.proto