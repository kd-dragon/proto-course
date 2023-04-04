package main

import (
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg was invoked")

	var sum int32
	var count int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			res := float64(sum) / float64(count)
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Received %v\n", req)
		sum += req.Number
		count++
	}

}
