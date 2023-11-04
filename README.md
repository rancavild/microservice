## gRPC - Go - Microservice

### Install

    $ apt install protobuf protoc-gen-go protoc-gen-go-grpc

### Creating microservice.proto

    $ protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative microservice.proto 