package main

import (
	"context"
	pb "github.com/gain620/realtime-dashboard-grpc/calculator/proto"
	"log"
)

func doSum(a, b int32, c pb.SumServiceClient) {
	log.Printf("doSum function invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  a,
		SecondNumber: b,
	})
	if err != nil {
		log.Fatalf("Couldn't call sum function : %v\n", err)
	}

	log.Printf("Sum : %v\n", res.Result)
}
