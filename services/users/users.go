package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getUser called"})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getUserById called id: " + id + "."})
}

func AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addUser was called"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "UpdateUser was called id: " + id + "."})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "DeleteUser was called id: " + id + "."})
}
