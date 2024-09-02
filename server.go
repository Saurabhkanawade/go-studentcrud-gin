package main

import (
	"github.com/gin-gonic/gin"
	"saurabhkanawade/gin/internal/controller"
)

func StartServer() {
	router := gin.Default()

	controller.StudentController(router)
	router.Run(":8080")
}
