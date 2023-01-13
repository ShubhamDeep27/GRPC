package main

import (
	"flag"
	pb "grpc/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
	"net/http"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type Employee struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewEmployeeServiceClient(conn)

	r := gin.Default()
	r.GET("/Employee/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		log.Print(username, password)
		data := &pb.Employee{
			Username: username,
			Password: password,
		}
		res, err := client.GetEmployee(c, &pb.ReadEmployeeRequest{Employee: data})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": res,
		})
	})

	r.Run()
}
