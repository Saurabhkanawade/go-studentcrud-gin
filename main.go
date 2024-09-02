package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "saurabh",
		"email": "saurabhkanawad3@gmail.com",
	})
}

func main() {
	fmt.Println("Welcome to the gin framework.......")
	//StartServer()

	r := gin.Default()

	r.GET("/student", GetStudent)

	r.Run(":9090")
}
