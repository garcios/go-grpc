package basic

import (
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func UserGroupSample() {
	batman := &basic.User{
		Id:       100,
		IsActive: true,
		Username: "batman",
		Password: []byte("password1"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"batman@gmail.com", "batman@hotmail.com"},
		Address: &basic.Address{
			Street:  "Southern Cross",
			City:    "Melbourne",
			Country: "Australia",
		},
	}

	robin := &basic.User{
		Id:       100,
		IsActive: true,
		Username: "robin",
		Password: []byte("password1"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"robin@gmail.com", "robin@hotmail.com"},
		Address: &basic.Address{
			Street:  "Southern Cross",
			City:    "Melbourne",
			Country: "Australia",
		},
	}

	ug := basic.UserGroup{
		GroupId:     123,
		GroupName:   "Bat Family",
		Users:       []*basic.User{batman, robin},
		Roles:       []string{"admin", "superuser"},
		Description: "Testing only",
	}

	jsonBytes, err := protojson.Marshal(&ug)
	if err != nil {
		log.Fatalln("Got an error:", err)
	}

	log.Println(string(jsonBytes))

}
