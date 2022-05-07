package main

import (
	"context"
	pb "github.com/gain620/realtime-dashboard-grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream : %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Gain"},
		{FirstName: "Chang"},
		{FirstName: "America"},
		{FirstName: "DrStrange"},
	}

	waitChan := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request : %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second)
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
				log.Printf("Error while receiving : %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitChan)
	}()

	<-waitChan
}
