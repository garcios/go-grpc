package basic

import (
	"fmt"
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

// This is very similar with writeProtoToFile, the only difference is we use protojson library instead.
func writeProtoToJSON(msg proto.Message, fname string) error {
	b, err := protojson.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshaling message to binary: %w", err)
	}

	err = os.WriteFile(fname, b, 0644)
	if err != nil {
		return fmt.Errorf("error writing message to file: %w", err)
	}

	return nil
}

// This is very similar with readProtoToFile, the only difference is we use protojson library instead.
func readProtoFromJSON(fname string, dest proto.Message) error {
	in, err := os.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("error reading message from file: %w", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		return fmt.Errorf("error unmarshaling message from file: %w", err)
	}

	return nil
}

func TestReadWriteProtoJSON() {
	u := basic.User{
		Id:       100,
		IsActive: true,
		Username: "user1",
		Password: []byte("password1"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"user1@gmail.com", "user1@hotmail.com"},
	}

	err := writeProtoToJSON(&u, "user.json")
	if err != nil {
		log.Fatal(err)
	}

	var dest basic.User
	err = readProtoFromJSON("user.json", &dest)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", dest)
}
