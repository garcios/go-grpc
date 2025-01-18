package jobsearch

import (
	"encoding/json"
	"github.com/oskiegarcia/go-grpc/intro/protogen/basic"
	"github.com/oskiegarcia/go-grpc/intro/protogen/dummy"
	"github.com/oskiegarcia/go-grpc/intro/protogen/jobsearch"
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

	jsonBytes, err := json.Marshal(js)
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

	jsonBytes, err := json.Marshal(jc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", jsonBytes)
}
