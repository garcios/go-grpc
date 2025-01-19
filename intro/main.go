package main

import (
	"fmt"
	"github.com/garcios/go-grpc/intro/basic"

	//"github.com/garcios/go-grpc/intro/jobsearch"
	"log"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	//basic.BasicHello()

	//basic.BasicUser()
	//basic.ProtoToJsonUser()
	//basic.JsonToProtoUser()
	//basic.UserGroupSample()

	//jobsearch.JobSearchSofware()
	//jobsearch.JobSearchCandidate()

	//basic.BasicUnMarshallAnyKnown()
	basic.BasicUnmarshallAnyNotKnowm()
}
