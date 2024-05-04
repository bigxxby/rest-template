package controllers

import (
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, "Hello world!")

}
