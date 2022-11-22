# chat-app-grpc

A twitter like solution using golang GRPC

__Compile Proto__

``protoc -I . --go_out ./app --go-grpc_out ./app proto/homefeed.proto``

## Run Server

``go run main.go``

## Test GRPC using postman

- new->gRPC Request
- server : grpc://127.0.0.1:4444
- import proto
- select api
