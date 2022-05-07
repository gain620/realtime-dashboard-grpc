package main

import (
	"context"
	pb "github.com/gain620/realtime-dashboard-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:50051"

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet function invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Gain",
	})
	if err != nil {
		log.Fatalf("Couldn't Greet : %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
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

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	//	doGreetManyTimes(c)
	doLongGreet(c)
	//doGreetEveryone(c)
}
