package main

import (
	UserService "github.com/emmohac/go-learnit/services/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	{
		v1.GET("users", UserService.GetUsers)
		v1.GET("users/:id", UserService.GetUserById)
		v1.POST("users", UserService.AddUser)
		v1.PUT("users/:id", UserService.UpdateUser)
		v1.DELETE("users/:id", UserService.DeleteUser)
	}

	router.Run()
}
