# Go gRPC
This comprises key insights and crucial programs that have contributed to my understanding of the gRPC using the 
Go language.

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

## References
- https://github.com/uber/prototool/blob/dev/style/README.md
- https://cloud.google.com/apis/design
- https://github.com/uber-go/guide/blob/master/style.md