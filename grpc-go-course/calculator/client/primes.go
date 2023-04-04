package main

import (
	"context"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) error {

	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{Number: 12390392840}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		log.Printf("Response from Primes: %v\n", res.Result)
	}

	return nil
}
