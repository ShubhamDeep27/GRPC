package common

import (
	"fmt"
	"log"

	"grpc/models"

	"grpc/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {
	host := util.GetEnvVariable("DB_HOST")
	port := util.GetEnvVariable("DB_PORT")
	user := util.GetEnvVariable("DB_USER")
	dbname := util.GetEnvVariable("DB_NAME")
	password := util.GetEnvVariable("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(models.Employee{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connected...")

}
