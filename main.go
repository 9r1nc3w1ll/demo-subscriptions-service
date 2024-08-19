package main

import (
	"fmt"
	"lithium-test/db"
	"lithium-test/pb"
	"lithium-test/services"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 6000
	dsn := "host=localhost user=root password=root dbname=lithium_test_db port=5490 sslmode=disable TimeZone=UTC"

	db, err := db.InitDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterProductServiceServer(server, services.NewProductService(db))
	pb.RegisterSubscriptionServiceServer(server, services.NewSubscriptionService(db))
	reflection.Register(server)

	go func() {
		listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

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
