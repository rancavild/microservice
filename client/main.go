package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc/ms/pb"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:9999", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewPaymentServiceClient(cc)
	request := &pb.PaymentServiceRequest{Id: "2036"}

	resp, err := client.PayTo(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received :", resp.Message)
}
