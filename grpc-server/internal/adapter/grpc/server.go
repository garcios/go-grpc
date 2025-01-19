package grpc

import (
	"fmt"
	"github.com/garcios/go-grpc/grpc-proto/protogen/go/hello"
	"github.com/garcios/go-grpc/grpc-server/internal/port"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		grpcPort:     grpcPort,
	}
}

func (a *GrpcAdapter) Run() error {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("grpc server listening on port %d", a.grpcPort)

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC on port %d: %v", a.grpcPort, err)
	}

	return nil
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
