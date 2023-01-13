package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "grpc/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	DatabaseConnection()
}

var (
	DB  *gorm.DB
	err error
)

type Employee struct {
	ID       uint
	Name     string
	Mobile   string
	Address  string
	Salary   int
	Age      int
	Username string
	Password string
}

func getEnvVariable(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func DatabaseConnection() {
	host := getEnvVariable("DB_HOST")
	port := getEnvVariable("DB_PORT")
	user := getEnvVariable("DB_USER")
	dbname := getEnvVariable("DB_NAME")
	password := getEnvVariable("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Employee{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connected...")

}

func (b *Employee) TableName() string {
	return "Employee"
}

type server struct {
	pb.UnimplementedEmployeeServiceServer
}

func (*server) GetEmployee(ctx context.Context, req *pb.ReadEmployeeRequest) (*pb.ReadEmployeeResponse, error) {

	username := req.Employee.GetUsername()
	password := req.Employee.GetPassword()

	var employee Employee
	res := DB.Find(&employee, "username = ? AND password = ? ", username, password)

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

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
