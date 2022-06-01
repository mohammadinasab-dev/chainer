package grpc

import (
	"chainer/cmd/di"
	"chainer/config"
	"chainer/pkg/pbs"
	"fmt"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

var (
	grpcServer *grpc.Server
)

type server struct{}

type grpcConf struct {
	ConnectionType string
	Address        string
}

func GetInstance() *server {
	return &server{}
}

func (s *server) loadConfigs() *grpcConf {
	connectionType := config.GetGrpcConnectionTypeString()
	address := fmt.Sprintf("%s:%s", config.GetGrpcAddressString(), config.GetGrpcPortString())
	return &grpcConf{ConnectionType: connectionType, Address: address}
}

func (s *server) Start() error {
	fmt.Println("GRPC SERVER STARTING...")

	conf := s.loadConfigs()

	lis, err := s.makeListener(conf.ConnectionType, conf.Address)
	if err != nil {
		log.Println("cant make listener for GRPC server")
		return err
	}

	grpcServer = s.makeServer()

	s.registerServers(grpcServer)

	s.serve(grpcServer, lis)
	fmt.Printf("gRPC server started on: %s \n", conf.Address)

	return nil
}

func (s *server) Stop() {
	if grpcServer != nil {
		fmt.Println("STOPPING GRPC SERVER")
		grpcServer.GracefulStop()
	}
}

func (s *server) makeListener(connctionType, address string) (net.Listener, error) {
	lis, err := net.Listen(connctionType, address)
	if err != nil {
		return nil, err
	}
	return lis, nil
}

func (s *server) makeServer() *grpc.Server {
	return grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcValidator.UnaryServerInterceptor(),
			),
		),
	)
}

func (s *server) registerServers(server *grpc.Server) {
	pbs.RegisterAddressGeneratorServer(server, di.AddressGenerator())
}

func (s *server) serve(server *grpc.Server, lis net.Listener) {
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Println("faild to serve GRPC server")
		}
	}()
}
