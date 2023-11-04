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

In the proto file we define the service **PaymentService** with the method **PayTo()** and the messages **PaymentServiceRequest**, and **PaymentServiceReply**. 

### Setting up

Create a directory called **microservice**.

    $ mkdir microservice

Initialize the go module.

    $ cd microservice
    $ go mod init grpc/ms

### Creating microservice.proto

Create the file microservice.go with the following lines.

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

And create the **pb** directory inside the directory **microservice**.

    $ mkdir pb

And finally execute the command. 

    $ protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative microservice.proto 

This will create the files **microservice_grpc.pb.go** and **microservice_grpc.pb.go** with the code needed by the service.

### Service code

Create the directory service inside microservice directory and go into.

    $ mkdir server
    $ cd server

Create the file **main.go** with the following code:

    package main

    import (
	    "context"
	    "fmt"
	    pb "grpc/ms/pb"
	    "log"
	    "net"

	    "google.golang.org/grpc"
    )

    type server struct {
	    pb.PaymentServiceServer
    }

    func (s *server) PayTo(ctx context.Context, req *pb.PaymentServiceRequest) (*pb.PaymentServiceReply, error) {
	    log.Println("Paying Id :", req.Id)
	    return &pb.PaymentServiceReply{
		    Message: fmt.Sprintf("Processing payment id : %v", req.Id),
	    }, nil
    }

    func main() {
	    listener, err := net.Listen("tcp", ":9999")
	    if err != nil {
		    panic(err)
	    }
	    s := grpc.NewServer()
	    pb.RegisterPaymentServiceServer(s, &server{})

        log.Println("Service running")
	    
        if err := s.Serve(listener); err != nil {
		    log.Fatalf("Failed to serve : %v", err)
	    }
    }

Notice that we are using the "grpc/ms/pb" package created from microservice.proto.

Before running our service, execute go mod tidy to get all module needed by the service (in particular "google.golang.org/grpc").

    $ go mod tidy

### Client code

Create the directory client inside microservice directory and go into. And create the file main.go inside.

    $ mkdir client
    $ cd client

File main.go:

    package main

    import (
	    "context"
	    "fmt"
	    "log"

	    pb "grpc/ms/pb"

	    "google.golang.org/grpc"
	    "google.golang.org/grpc/credentials/insecure"
    )

    func main() {
	    opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	    cc, err := grpc.Dial("localhost:9999", opts)
	    if err != nil {
		    log.Fatal(err)
	    }
	    defer cc.Close()

	    client := pb.NewPaymentServiceClient(cc)
	    request := &pb.PaymentServiceRequest{Id: "21111036"}

	    resp, err := client.PayTo(context.Background(), request)

	    if err != nil {
		    log.Fatal(err)
	    }
	    fmt.Println("Received :", resp.Message)
    }

### Running server and client

Verify you are in microservice directory.

    $ cd microservice

Then run in a terminal:

    $ go run server/main.go
    2023/11/01 12:24:48 Service running

In another terminal run the client.

    $ go run client/main.go

We should receive a response like this:

    Received : Processing payment id : 21111036

And the terminal server we should see:

    2023/11/01 12:28:01 Paying Id : 21111036

...And that's for now, we implemented a very simple gRPC microservice.