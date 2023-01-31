package grpcroutes

import (
	"context"
	"grpc/common"
	"grpc/models"
	pb "grpc/proto"
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
