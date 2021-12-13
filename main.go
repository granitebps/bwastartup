package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/granitebps/bwastartup/handler"
	"github.com/granitebps/bwastartup/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:granite97@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&user.User{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("api/v1")

	api.POST("users", userHandler.RegisterUser)

	router.Run()

	fmt.Println("Connected to Database")
}
