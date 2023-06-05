// package main
// import "fmt"
// func main() {
//     fmt.Println("hello world")
// }

package main

import (
	"fmt"
	"google.golang.org/grpc"
	gen "main/gen/go"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:5062")
	if err != nil {
		panic(err)
	}
	fmt.Println("gRPC biz server running on " + listener.Addr().String())

	s := grpc.NewServer()
	gen.RegisterBizServer(s, &server{})

	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
