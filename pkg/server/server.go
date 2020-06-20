package server

import (
	"fmt"
	pb "gor/build/protos/runtime"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Start server.
func Start(port uint32) {
	tcp, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start server, error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRuntimeServiceServer(s, &RuntimeServer{})
	if err := s.Serve(tcp); err != nil {
		log.Fatalf("Failed to start server, error: %v", err)
	}
}
