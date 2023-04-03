package main

import (
	"context"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	r, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 1, SecondNumber: 1})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", r.Result)
}
