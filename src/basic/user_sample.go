package basic

import (
	"github.com/oskiegarcia/go-grpc/src/protogen/basic"
	"log"
)

func BasicUser() {
	u := basic.User{
		Id:       100,
		IsActive: true,
		Username: "user1",
		Password: []byte("password1"),
	}

	log.Println(&u)
}
