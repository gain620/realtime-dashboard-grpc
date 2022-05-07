package main

import (
	"context"
	pb "github.com/gain620/realtime-dashboard-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:50052"

func doSum(a, b int32, c pb.SumServiceClient) {
	log.Printf("doSum function invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: a,
		B: b,
	})
	if err != nil {
		log.Fatalf("Couldn't call sum function : %v\n", err)
	}

	log.Printf("Sum : %v\n", res.Sum)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close the connection! : %v\n", err)
		}
	}(conn)

	c := pb.NewSumServiceClient(conn)

	doSum(10, 3, c)
}
