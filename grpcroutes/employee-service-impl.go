package grpcroutes

import (
	"context"
	"grpc/common"
	"grpc/models"
	pb "grpc/proto"

	"golang.org/x/crypto/bcrypt"
)

type EmployeeServiceImpl struct {
	pb.UnimplementedEmployeeServiceServer
}

func NewEmployeeServiceImpl() *EmployeeServiceImpl {

	return &EmployeeServiceImpl{}

}
func (*EmployeeServiceImpl) GetEmployee(ctx context.Context, req *pb.ReadEmployeeRequest) (*pb.ReadEmployeeResponse, error) {

	username := req.Employee.GetUsername()
	password := req.Employee.GetPassword()
	var storedpassword string

	var employee models.Employee
	res := common.DB.Select("password").Find(&employee, "username = ?", username)

	res.Scan(&storedpassword)
	if err := bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(password)); err != nil {
		return &pb.ReadEmployeeResponse{
			Status: "Login Faliure",
		}, nil
	}

	return &pb.ReadEmployeeResponse{
		Status: "Login Success",
	}, nil
}
