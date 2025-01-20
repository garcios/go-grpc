package adapter

import (
	"context"
	"github.com/garcios/go-grpc/grpc-client/internal/port"
	"github.com/garcios/go-grpc/grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
)

type HelloAdapter struct {
	helloClient port.HelloClientPort
}

func NewHelloAdapter(conn *grpc.ClientConn) (*HelloAdapter, error) {
	client := hello.NewHelloServiceClient(conn)

	return &HelloAdapter{
		helloClient: client,
	}, nil
}

func (a *HelloAdapter) SayHello(ctx context.Context, name string) (*hello.HelloResponse, error) {
	helloRequest := &hello.HelloRequest{
		Name: name,
	}

	resp, err := a.helloClient.SayHello(ctx, helloRequest)

	if err != nil {
		log.Fatalln("Error on SayHello :", err)
	}

	return resp, nil
}
