package main

import (
	"context"
	pb "github.com/kd-dragon/proto-course/grpc-go-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Drake"},
		{FirstName: "Denny"},
		{FirstName: "Ben"},
		{FirstName: "Ted"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while sending data to server: %v", err)
			}
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
				log.Printf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received: %v", res.GetResult())
		}

		close(waitc)
	}()

	<-waitc
}
