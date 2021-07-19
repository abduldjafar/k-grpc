package server

import (
	"k-grpc/entitypb"
	"k-grpc/service"
	"log"

	"google.golang.org/grpc"
)

var (
	talentService entitypb.TalentServiceServer = service.NewTalentService()
)

func GrpcServer() *grpc.Server {
	log.Println("Server Starting")

	grpcServer := grpc.NewServer()
	entitypb.RegisterTalentServiceServer(grpcServer, talentService)

	return grpcServer
}
