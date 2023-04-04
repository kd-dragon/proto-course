package main

import (
	"context"
	"fmt"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	fmt.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		fmt.Printf("Error while calling Max RPC: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			log.Println("Sending number: ", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			log.Printf("Received Maximun number: %v", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
