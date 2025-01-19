# gRPC

## Creating go module
```shell
go mod init github.com/garcios/go-grpc/grpc-proto
go mod tidy
```

## Introduction
- gRPC is Remote Procedure Call (RPC) framework.
- Works over HTTP/2 protocol.
- Data transfer performance improvement by using protocol buffers as data format.
- Default behavior is encrypted communication over SSL/TLS.
- Support data streaming (on client/server/both)

## REST API or gRPC
- Focus on gRPC as service-to-service communication.
- Us gRPC-REST API gateway that able to expose gRPC as REST API
  - This means write as gRPC and release as both gRPC and REST API.

## gRPC Transport Protocol
- gRPC uses HTTP/2
  - Support multiplexing request/responses over single connection.
  - Improve performance and efficiency.