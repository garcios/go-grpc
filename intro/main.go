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

	basic.BasicUser()
	//basic.ProtoToJsonUser()
	//basic.JsonToProtoUser()
	//basic.UserGroupSample()

	//jobsearch.JobSearchSofware()
	//jobsearch.JobSearchCandidate()

	// Any
	//basic.BasicUnMarshallAnyKnown()
	//basic.BasicUnmarshallAnyNotKnowm()
	//basic.BasicUnmarshallAnyIs()
	//basic.BasicOneof()

	// Read/Write
	//basic.TestReadWriteProto()
	//basic.TestReadWriteProtoJSON()

	// Schema evolution
	//basic.BasicWriteUserContentV1()
	//basic.BasicReadUserContentV1()
	//basic.BasicWriteUserContentV2()
	//basic.BasicReadUserContentV2()
	//basic.BasicWriteUserContentV3()
	//basic.BasicReadUserContentV3()
	//basic.BasicReadUserPayment()

}
