package server

import (
	"fmt"
	"lithium-test/pb"
	"lithium-test/services"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitServer() {
	var opts []grpc.ServerOption
	port := 6000

	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(opts...)
	pb.RegisterProductServiceServer(server, services.NewProductService())
	pb.RegisterSubscriptionServiceServer(server, services.NewSubscriptionService())
	reflection.Register(server)

	go func() {
		if e := server.Serve(listen); e != nil {
			log.Fatalf("Failed to serve gRPC %v", e.Error())
		}
	}()

	log.Printf("Service is running on port %d\n", port)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	<-sigint
	log.Println("Service shutting down")
}
