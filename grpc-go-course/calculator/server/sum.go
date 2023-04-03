package main

import (
	"context"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"log"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum was invoked with %v\n", in)
	return &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
