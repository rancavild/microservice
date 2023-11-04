## gRPC - Go - Microservice (WIP)

The main idea behind RPC (Remote Procedure Call) is to allow that clients can directly call methods on a server application on a different machine,
making it easier to create distributed applications and microservices. 

gRPC is an Open Source framework high performance Remote Procedure Call (RPC) implemented by Google. 

A service has methods that can be called remotely by clients, passing parameters and getting a response.

gRPC doesn't depend on a programming language, we could have our services written in go, and clients written in Python,
C++, Java, etc.

gRPC has been implemented over HTTP/2 which some charasteristics are:

- Multiplexing
- Stream prioritization.
- Header compression
- And binary protocol. 

These features reduce latency, increase parallelism, and optimize resource delivery. So for these reasons gRPC is a good option to implement
high-demand and distrbuted services. 

### Install

First we have to set up our environment, installing the following (this is for Ubuntu, check the documentation for your OS).

    $ apt install protobuf protoc-gen-go protoc-gen-go-grpc

Protobuf will create for us the code from the file.proto where we define our service using a descriptive and agnostic language.
For example, with the following lines (file microservice.proto) we are defining the Methods and Messages for the service that we want to develop.

    syntax = "proto3";
    option go_package = "grpc/ms/pb";

    service PaymentService {
        rpc PayTo(PaymentServiceRequest) returns (PaymentServiceReply) {}
    }

    message PaymentServiceRequest {
        string id = 1;
    }

    message PaymentServiceReply {
        string message = 1;
    }

Since, we are going to work with go, we need to install **protoc-gen-go** and **protoc-gen-go-grpc** to create the go gRPC code (as a go package).

### Creating microservice.proto

    $ protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative microservice.proto 