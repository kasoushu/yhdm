package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Testget(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "hello",
		"number": "123456",
	})
}

