package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jcxue/go-web/api"
)

type Server struct {
	api.UnimplementedHelloServer
}

func NewServer() *Server {
	return &Server{}
}

func main() {
	port := flag.Int("port", 0, "server port")
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := NewServer()
	api.RegisterStoreServer(grpcServer, server)
	reflection.Register(grpcServer)

	if grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
