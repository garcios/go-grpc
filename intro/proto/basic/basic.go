package basic

//go:generate protoc --go_opt=module=github.com/garcios/go-grpc/intro --go_out=../../  -I =.  ./hello.proto ./user.proto ./user_group.proto ./application.proto
