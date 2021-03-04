deps:
	go mod tidy
	go mod verify

format:
	go fmt ./...

docker-proto:
	docker run --rm -v $$(pwd):/opt/src grpc-demo-protogen:latest
	
proto:
	protoc -I ./api/proto -I ./third_party --go_out=./api --go-grpc_out=./api --grpc-gateway_out ./api ./api/proto/*.proto

build-api:
	go build -o ./bin/api ./cmd/api
	
build: build-api

docker-builder-protogen:
	docker build -t "grpc-demo-protogen:Dockerfile.protogen" .