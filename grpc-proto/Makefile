# Change this to your own Go module
GO_MODULE := github.com/garcios/go-grpc/grpc-proto


.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen	
else
	rm -fR ./protogen
endif


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc: clean
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/hello/*.proto


# This command does not work at the moment, please use protoc instead.
.PHONY: generate
generate: clean
	go generate ./...

.PHONY: build
build: protoc tidy


.PHONY: run
run:
	go run main.go


.PHONY: execute
execute: build run

