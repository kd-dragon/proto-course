package main

import (
	"context"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"log"
	"time"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg RPC: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg Response: %v\n", res.Avg)
}
