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

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
