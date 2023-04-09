package main

import (
	"github.com/snowmerak/guide/src/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	service.RegisterServiceServer(server, ServiceServer{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
