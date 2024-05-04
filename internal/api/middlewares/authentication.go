package middlewares

import (
	"app/pkg/jwt"
	"log"

	"github.com/gin-gonic/gin"
)

func BasidAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		userId, err := jwt.ParseToken(authorizationHeader)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.Set("userId", userId)
		c.Next()

	}
}
