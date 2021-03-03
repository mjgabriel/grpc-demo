deps:
	go mod tidy
	go mod verify

format:
	go fmt ./...

proto:
	protoc -I ./api/proto -I ./third_party --go_out=./api --go-grpc_out=./api --grpc-gateway_out ./api ./api/proto/*.proto

build-api:
	go build -o ./bin/api ./cmd/api

build: build-api