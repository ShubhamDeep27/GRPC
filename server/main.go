package main

import (
	"log"
	"net"

	"grpc/common"
	grpcroutes "grpc/grpc-routes.go"
	"grpc/interceptors"
	pb "grpc/proto"

	"google.golang.org/grpc"
)

func init() {
	common.DatabaseConnection()
	common.InitApiLogger()
}

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	employeeServiceImpl := grpcroutes.NewEmployeeServiceImpl()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggerInterceptor))
	pb.RegisterEmployeeServiceServer(grpcServer, employeeServiceImpl)

	log.Printf("Server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
