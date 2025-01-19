package basic

import (
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func BasicReadUserPayment() {
	log.Println("Read User Payment")

	var up basic.UserPayment

	err := readProtoFromFile("user_content_v1.msg", &up)
	if err != nil {
		log.Fatal(err)
	}

	upJson, err := protojson.Marshal(&up)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("User Payment JSON: %s\n", upJson)
}
