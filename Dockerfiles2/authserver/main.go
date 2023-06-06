package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:5052")
	if err != nil {
		panic(err)
	}
	fmt.Println("gRPC auth server running on " + listener.Addr().String())

	s := grpc.NewServer()
	RegisterAuthServer(s, &server{})

	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
