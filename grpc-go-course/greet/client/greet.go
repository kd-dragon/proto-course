package main

import (
	"context"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("goGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "D-dragon",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v\n", err)
	}

	log.Printf("Response from Greet: %v\n", res.Result)
}
