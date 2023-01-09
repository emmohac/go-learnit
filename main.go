package main

import (
	"log"

	"github.com/emmohac/go-learnit/databases"
	UserService "github.com/emmohac/go-learnit/services/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	err := databases.ConnectDatabase()
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	{
		v1.GET("users", UserService.GetUsers)
		// v1.GET("users/:id", UserService)
		// v1.POST("users", services.AddUser)
		// v1.PUT("users/:id", services.UpdateUser)
		// v1.DELETE("users/:id", services.DeleteUser)
	}

	router.Run()
}
