package basic

import (
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"log"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Oscar",
	}

	log.Println(&h)
}
