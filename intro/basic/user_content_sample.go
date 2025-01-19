package basic

import (
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/this-is-v1",
		//	Title:         "10 Strongest People In The World",
		HtmlContent: "<p>Just a dummy content for 10 Strongest People In The World</p>",
		//	AuthorId:      99,
	}

	err := writeProtoToFile(&uc, "user_content_v1.msg")
	if err != nil {
		log.Fatal(err)
	}
}

func BasicReadUserContentV1() {
	log.Println("Read V1")

	var uc basic.UserContent
	err := readProtoFromFile("user_content_v1.msg", &uc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", &uc)

	ucJson, err := protojson.Marshal(&uc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", string(ucJson))
	log.Println()
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 1002,
		Slug:          "/this-is-v2",
		//	Title:         "10 Strongest People In The World Version 2",
		HtmlContent: "<p>Just a dummy content for 10 Strongest People In The World Version 2</p>",
		//	AuthorId:      100,
		Category: "NEWS",
	}

	err := writeProtoToFile(&uc, "user_content_v2.msg")
	if err != nil {
		log.Fatal(err)
	}
}

func BasicReadUserContentV2() {
	log.Println("Read V2")

	var uc basic.UserContent
	err := readProtoFromFile("user_content_v2.msg", &uc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", &uc)

	ucJson, err := protojson.Marshal(&uc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", string(ucJson))
	log.Println()
}

func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 1003,
		Slug:          "/this-is-v3",
		//	Title:         "10 Strongest People In The World Version 3",
		HtmlContent: "<p>Just a dummy content for 10 Strongest People In The World Version 3</p>",
		//	AuthorId:      100,
		Category: "NEWS",
		//SubCategory: "PEOPLE",
	}

	err := writeProtoToFile(&uc, "user_content_v3.msg")
	if err != nil {
		log.Fatal(err)
	}
}

func BasicReadUserContentV3() {
	log.Println("Read V3")

	var uc basic.UserContent
	err := readProtoFromFile("user_content_v3.msg", &uc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", &uc)

	ucJson, err := protojson.Marshal(&uc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", string(ucJson))
	log.Println()
}
