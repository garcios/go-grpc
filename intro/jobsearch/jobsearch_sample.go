package jobsearch

import (
	"github.com/garcios/go-grpc/intro/protogen/basic"
	"github.com/garcios/go-grpc/intro/protogen/dummy"
	"github.com/garcios/go-grpc/intro/protogen/jobsearch"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func JobSearchSofware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 77,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "JobSearch Software",
			Platforms: []string{"Mac", "Linux", "Windows"},
		},
	}

	jsonBytes, err := protojson.Marshal(&js)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", jsonBytes)

}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 90,
		Application: &dummy.Application{
			ApplicationId:     897,
			ApplicantFullName: "Harry Potter",
			Phone:             "+65-111-222-333",
			Email:             "harry@gmail.com",
		},
	}

	jsonBytes, err := protojson.Marshal(&jc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", jsonBytes)
}
