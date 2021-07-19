package main

import (
	"k-grpc/server"
	"log"
	"net"
)

func main() {
	address := "0.0.0.0:50051"

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Println(err.Error())
	}

	grpcServer := server.GrpcServer()

	grpcServer.Serve(lis)
}
