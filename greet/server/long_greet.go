package main

import (
	"fmt"
	pb "github.com/gain620/realtime-dashboard-grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			err := stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
			if err != nil {
				return err
			}
		}

		if err != nil {
			log.Fatalf("Error while reading client stream : %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s \n", req.FirstName)
	}

	return nil
}
