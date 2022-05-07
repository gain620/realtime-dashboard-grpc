package main

import (
	pb "github.com/gain620/realtime-dashboard-grpc/calculator/proto"
	"log"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Prime function invoked : %v\n", in)

	k := int32(2)
	n := in.PrimeNumber

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n /= k
		} else {
			k += 1
		}
	}

	return nil
}
