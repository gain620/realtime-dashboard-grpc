package main

import (
	"context"
	"fmt"
	pb "github.com/gain620/realtime-dashboard-grpc/calculator/proto"
	"io"
	"log"
)

func doPrime(n int, c pb.CalculatorServiceClient) {
	log.Printf("doPrime invoked")

	req := &pb.PrimeRequest{
		PrimeNumber: int32(n),
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Prime: %v\n", err)
	}

	fmt.Printf("Prime number of %v = ", req.PrimeNumber)
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream : %v\n", err)
		}

		fmt.Printf("%v * ", msg.Result)
	}

}
