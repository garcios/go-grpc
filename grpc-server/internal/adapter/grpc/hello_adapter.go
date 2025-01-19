package grpc

import (
	"context"
	"github.com/garcios/go-grpc/grpc-proto/protogen/go/hello"
	"log"
)

func (a *GrpcAdapter) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	log.Println("request:", req)

	greet := a.helloService.Greet(req.Name)

	return &hello.HelloResponse{
		Greet: greet,
	}, nil
}
