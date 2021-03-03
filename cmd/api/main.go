package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	greetings "github.com/mjgabriel/grpc-demo/api"
	"google.golang.org/grpc"
)

const port = ":9000"

type server struct {
	greetings.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, r *greetings.HelloRequest) (*greetings.HelloResponse, error) {
	return &greetings.HelloResponse{Message: fmt.Sprintf("Greetings %s!", r.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// attach Greetings service to the server
	greetings.RegisterHelloServiceServer(s, &server{})

	log.Println("Serving gRPC on 0.0.0.0" + port)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+port,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	gwmux := runtime.NewServeMux()

	// register Greetings service
	err = greetings.RegisterHelloServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":9090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:9090")
	log.Fatalln(gwServer.ListenAndServe())
}
