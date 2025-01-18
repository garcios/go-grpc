package main

import (
	"fmt"
	"github.com/oskiegarcia/go-grpc/intro/basic"
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

	basic.BasicUser()
	//basic.ProtoToJsonUser()
	//basic.JsonToProtoUser()
	//basic.UserGroupSample()

	//jobsearch.JobSearchSofware()
	//jobsearch.JobSearchCandidate()
}
