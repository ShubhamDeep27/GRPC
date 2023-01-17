package main

import (
	"context"
	"log"
	"net"

	"grpc/common"
	"grpc/interceptors"
	"grpc/models"
	pb "grpc/proto"

	"google.golang.org/grpc"
)

func init() {
	common.DatabaseConnection()
	common.InitApiLogger()
}

type server struct {
	pb.UnimplementedEmployeeServiceServer
}

func (*server) GetEmployee(ctx context.Context, req *pb.ReadEmployeeRequest) (*pb.ReadEmployeeResponse, error) {

	username := req.Employee.GetUsername()
	password := req.Employee.GetPassword()

	var employee models.Employee
	res := common.DB.Find(&employee, "username = ? AND password = ? ", username, password)

	if res.RowsAffected == 0 {
		return &pb.ReadEmployeeResponse{
			Status: "Login Faliure",
		}, nil
	}

	return &pb.ReadEmployeeResponse{
		Status: "Login Success",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggerInterceptor))
	pb.RegisterEmployeeServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
