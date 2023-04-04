package main

import (
	"fmt"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	fmt.Println("Max function was invoked")

	var max int32 = 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if res.Number > max {
			max = res.Number
			err := stream.Send(&pb.MaxResponse{
				Result: max,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
			}
		}

		fmt.Println("Current max: ", max)
	}
}
