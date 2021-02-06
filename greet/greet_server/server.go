package main

import (
	"fmt"
	"log"
	"net"

	"greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("hello sosong")
	Lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(Lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}
