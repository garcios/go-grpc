package basic

import (
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"math/rand"
	"time"
)

func BasicUser() {

	comm := randomCommunicationChannel()

	u := basic.User{
		Id:       100,
		IsActive: true,
		Username: "user1",
		Password: []byte("password1"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"user1@gmail.com", "user1@hotmail.com"},
		Address: &basic.Address{
			Street:  "Southern Cross",
			City:    "Melbourne",
			Country: "Australia",
			Coordinate: &basic.Address_Coordinate{
				Latitude:  -37.8176783959,
				Longitude: 144.951446194,
			},
		},
		CommunicationChannel: &comm,
	}

	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatalln("Got an error:", err)
	}

	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	u := basic.User{
		Id:       100,
		IsActive: true,
		Username: "user1",
		Password: []byte("password1"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"user1@gmail.com", "user1@hotmail.com"},
	}

	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	jsonBytes, err := marshaler.Marshal(&u)
	if err != nil {
		log.Fatalln("Got an error:", err)
	}

	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
    {
       "id":101,
        "username":"user2",
        "is_active":true,
         "password":"cGFzc3dvcmQx",
         "emails":["user2@gmail.com","user2@hotmail.com"],
         "gender":"FEMALE"
    }`

	var u basic.User
	err := protojson.Unmarshal([]byte(json), &u)
	if err != nil {
		log.Fatalln("Got an error:", err)
	}

	log.Println(&u)
}

func randomCommunicationChannel() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PapelMail{
		PaperMailAddress: "123 Main St New York",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "Facebook",
		SocialMediaUsername: "byteme",
	}

	instantMessging := basic.InstantMessaging{
		InstantMessagingProduct:  "WhatsApp",
		InstantMessagingUsername: "batman123",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	case 2:
		anypb.MarshalFrom(&a, &instantMessging, proto.MarshalOptions{})
	}

	return a
}

// BasicUnMarshallAnyKnown is a method to demonstrate unmarshalling when type is known.
func BasicUnMarshallAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "Facebook",
		SocialMediaUsername: "byteme123",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (Social Media)
	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		log.Fatalln("Got an error:", err)
	}

	jsonBytes, err := protojson.Marshal(&sm)
	if err != nil {
		log.Fatalf("Error marshaling json: %v", err)
	}
	log.Printf("JSON representation: %s\n", jsonBytes)
}

// BasicUnmarshallAnyNotKnowm is a method to demonstrate unmarshalling when type is NOT known.
func BasicUnmarshallAnyNotKnowm() {
	a := randomCommunicationChannel()

	var unmarchalled protoreflect.ProtoMessage

	unmarchalled, err := a.UnmarshalNew()
	if err != nil {
		log.Fatalln("Got an error:", err)
	}

	log.Println("Unmashall as", unmarchalled.ProtoReflect().Descriptor().Name())

	jsonBytes, err := protojson.Marshal(unmarchalled)
	if err != nil {
		log.Fatalf("Error marshaling json: %v", err)
	}

	log.Printf("JSON representation: %s\n", jsonBytes)
}

func BasicUnmarshallAnyIs() {

}
