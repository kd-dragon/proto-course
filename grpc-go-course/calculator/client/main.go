package main

import (
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	//doPrimes(c)
	//doAvg(c)
	doMax(c)

}
