# Introduction
This is an introduction to protobuf.

## Requirements

- Protobuf Compiler
```shell
brew install protobuf
```

- Protobuf Go Plugin
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

- Open API
```shell
go get github.com/google/gnostic
go install github.com/google/gnostic/cmd/protoc-gen-openapi
```

- Make
```shell
brew install make
```

## Proto compiler command
```shell
cd src
protoc --go_out=. ./proto/basic/*.proto
```

or using go generate command
```shell
cd <project root directory>
go generate ./...
```

or using make command
```shell
make generate
```

## Update go modules to satisfy golang imports for the generated code
```shell
cd ..
go mod tidy
```

or using make command
```shell
make tidy
```

## Run the code
```shell
cd src
go run .
```

or using make command
```shell
make run
```

## Using protojson
When you're using Protobuf as your data structure, you might want to use protojson for the following reasons:
1. __Native Compatibility__: protojson is built to work with Protobuf. It directly understands Protobuf message 
definitions, enumerations, nested messages, etc. It's more seamless and natural to marshal/unmarshal using protojson 
if your primary data structure is Protobuf.
2. __Efficiency__: protojson is designed to be fast and efficient with Protobuf data structures. The standard 
encoding/json library isn't optimized specifically for Protobuf.
3. __Additional Features__: protojson also provides additional options that are not available in encoding/json, 
such as emitting fields with default values when marshaling, treating null values as empty messages, etc. For example, 
this code block adapts the JSON marshaller settings:

```go
import "google.golang.org/protobuf/encoding/protojson"

marshaler := protojson.MarshalOptions{
	EmitUnpopulated: true,
}
jsonBytes, err := marshaler.Marshal(protoMessage)
if err != nil {
	// handle error
}
```
Remember, if you're not working with Protobuf data, then encoding/json is perfectly suitable for most json 
serialization/deserialization tasks. But when working with Protobuf, protojson could provide a more native and efficient
option.


## References
- https://github.com/uber/prototool/blob/dev/style/README.md
- https://cloud.google.com/apis/design
- https://github.com/uber-go/guide/blob/master/style.md