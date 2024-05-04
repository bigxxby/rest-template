package router

import (
	"app/internal/api/middlewares"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.Recovery())
	app.NoRoute(middlewares.NoRouteHandler())
	// Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
