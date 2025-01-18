package basic

//go:generate protoc --go_opt=module=intro --go_out=../../ --go-grpc_out=. -I=. hello.proto user.proto user_group.proto
