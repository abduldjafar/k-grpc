package server

import (
	"k-grpc/entitypb"
	"k-grpc/service"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	talentService entitypb.TalentServiceServer = service.NewTalentService()

	logrusLogger = logrus.New()
	customFunc   = func(code codes.Code) logrus.Level {
		if code == codes.OK {
			return logrus.InfoLevel
		}
		return logrus.ErrorLevel
	}
)

func GrpcServer() *grpc.Server {
	log.Println("Server Starting")
	logrusEntry := logrus.NewEntry(logrusLogger)

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customFunc),
	}

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	// Create a server, make sure we put the grpc_ctxtags context before everything else.
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
		),
	)

	entitypb.RegisterTalentServiceServer(grpcServer, talentService)

	return grpcServer
}
