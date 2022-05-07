package main

import (
	pb "github.com/gain620/realtime-dashboard-grpc/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}
}
