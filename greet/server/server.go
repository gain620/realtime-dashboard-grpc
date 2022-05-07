package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/gain620/realtime-dashboard-grpc/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello ! " + in.FirstName,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}
}
