package main

import (
	"context"
	"github.com/garcios/go-grpc/grpc-client/internal/adapter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:9090", opts...)

	if err != nil {
		log.Fatalln("Can not connect to gRPC server :", err)
	}

	defer conn.Close()

	helloAdapter, err := adapter.NewHelloAdapter(conn)

	if err != nil {
		log.Fatalln("Can not create HelloAdapter :", err)
	}

	runSayHello(helloAdapter, "Bruce Wayne")

}

func runSayHello(adapter *adapter.HelloAdapter, name string) {
	resp, err := adapter.SayHello(context.Background(), name)

	if err != nil {
		log.Fatalln("Can not call SayHello :", err)
	}

	log.Println(resp.Greet)
}
