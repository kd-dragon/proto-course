package main

import (
	pb "github.com/kd-dragon/proto-course/grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes was invoked with %v\n", in)

	k := int64(2)
	n := in.Number

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{Result: k})
			n = n / k
		} else {
			k = k + 1
		}
	}

	return nil
}
