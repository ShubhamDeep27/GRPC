package main

import (
	"context"
	"flag"
	pb "grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type Employee struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	data := &pb.Employee{
		Username: "shubham",
		Password: "123",
	}

	client := pb.NewEmployeeServiceClient(conn)

	res, err := client.GetEmployee(context.Background(), &pb.ReadEmployeeRequest{Employee: data})
	if err != nil {
		log.Fatalf("Failed to call GetEmployee: %v", err)
	}
	log.Printf("Response: %v", res)

}
