package main

import (
	"log"
	"net"

	"github.com/kudo07/go-grpc/module2/internal/hello"
	"github.com/kudo07/go-grpc/module2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	helloService := hello.Service{}

	proto.RegisterHelloServiceServer(grpcServer, &helloService)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting grpc server on address :%s", ":50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
