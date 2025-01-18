package basic

import (
	"github.com/oskiegarcia/go-grpc/src/protogen/basic"
	"log"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Oscar",
	}

	log.Println(&h)
}
