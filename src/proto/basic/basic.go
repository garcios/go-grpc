package basic

//go:generate protoc --go_opt=module=go-grpc --go_out=../../ --go-grpc_out=. -I=. ./hello.proto ./user.proto
