package controller

import (
	"github.com/gin-gonic/gin"
	"saurabhkanawade/gin/internal/service"
)

func StudentController(r *gin.Engine) {
	r.GET("/student", service.GetUser)
}
