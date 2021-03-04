# gRPC Demo Service
### Generate gRPC, gRPC-Gateway

With Docker running, run the following command to build the `grpc-demo-protogen` image.

```
docker build -t grpc-demo-protogen -f ./docker/Dockerfile.protogen ./docker
```

This will only need to be done once or whenever the `Dockerfile.protogen` is updated. At least, until the image is available in a container registry for download.

Once the image is built, to generate the protobuf files from the  `.proto` files defined for the service run the following:

```
make proto
```

Which simply runs the following `docker run`

```
docker run --rm -v $(pwd):/opt/src grpc-demo-protogen:latest
```

The reason for using the docker image is that the `protoc` compiler and plugins used are version locked so that the service can be consistently rebuilt and remain functional despite changes new versions of the `protoc` compiler introduce going forward. Tool dependency version locking.