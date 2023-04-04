package main

import (
	"fmt"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %v, number %v", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
