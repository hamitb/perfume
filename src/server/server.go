package server

import (
	"net"
	"perfumepb"
	"storage"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

type server struct {
	cnf     *Config
	storage *storage.MongoDBDriver
}

func NewServer(config *Config) (*server, error) {
	md, err := storage.NewMongoDriver(config.DBAddr)
	if err != nil {
		logrus.WithError(err).Errorf("Server: can not create new mongodb driver with given dbAddress")
		return nil, err
	}

	server := &server{
		cnf:     config,
		storage: md,
	}

	return server, nil
}

func (s *server) Healthy() error {
	return nil
}

func (s *server) Listen() error {
	grpcPort := s.cnf.GetGrpcPortString()
	// Listen
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logrus.Fatalf("server.go: Failed to listen %v for grpc", err)
	}

	grpcServer := grpc.NewServer()

	// Register services
	perfumepb.RegisterPerfumeServiceServer(grpcServer, s)

	logrus.Infof("server.go: Binding %s for grpc", grpcPort)

	// Save
	return grpcServer.Serve(lis)
}
