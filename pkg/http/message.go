package http

import (
	"github.com/gin-gonic/gin"
)

func HttpErrorMessage(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": errorMessage,
	})
}
func HttpOKMessage(c *gin.Context, message string) {
	c.JSON(200, gin.H{
		"message": message,
	})
}
