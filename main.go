package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		port     int
		grpcPort int
	)

	flag.IntVar(&port, "port", 8072, "The port number for ServerTwo")
	flag.IntVar(&grpcPort, "grpcPort", 8071, "The port number for ServerOne")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial gRPC: %v", err)
	}
	defer conn.Close()

	handler := NewHandler(conn)

	http.HandleFunc("/get", handler.getPerson)
	http.HandleFunc("/set", handler.setPerson)

	log.Println("serve at port ", port)

	err = http.Serve(lis, nil)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
