package services

import (
	"log"
	"net/http"

	"github.com/emmohac/go-learnit/models"
	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetUsers(c *gin.Context) {
	users, err := models.GetPersons(10)
	checkErr(err)

	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
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
