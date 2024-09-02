package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	users := []string{"saurabh", "shailesh", "rushikesh", "abhi", "akash"}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
