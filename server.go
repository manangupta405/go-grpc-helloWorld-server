package main

import (
	"grpc/hello"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := hello.Server{}

	grpcServer := grpc.NewServer()

	hello.RegisterHelloServer(grpcServer, &s)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
