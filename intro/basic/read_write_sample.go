package basic

import (
	"fmt"
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func writeProtoToFile(msg proto.Message, fname string) error {

	b, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshaling message to binary: %w", err)
	}

	err = os.WriteFile(fname, b, 0644)
	if err != nil {
		return fmt.Errorf("error writing message to file: %w", err)
	}

	return nil
}

func readProtoFromFile(fname string, dest proto.Message) error {
	in, err := os.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("error reading message from file: %w", err)
	}

	if err := proto.Unmarshal(in, dest); err != nil {
		return fmt.Errorf("error unmarshaling message from file: %w", err)
	}

	return nil
}

func TestReadWriteProto() {
	u := basic.User{
		Id:       100,
		IsActive: true,
		Username: "user1",
		Password: []byte("password1"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"user1@gmail.com", "user1@hotmail.com"},
	}

	err := writeProtoToFile(&u, "user.msg")
	if err != nil {
		log.Fatal(err)
	}

	var dest basic.User
	err = readProtoFromFile("user.msg", &dest)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", dest)
}
