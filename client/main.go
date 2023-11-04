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
