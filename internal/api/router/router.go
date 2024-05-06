package router

import (
	"app/internal/api/controllers"
	"app/internal/api/middlewares"
	"app/internal/repository"
	"app/internal/services"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//Env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return nil
	}
	// Middlewares
	app.Use(gin.Recovery())
	app.NoRoute(middlewares.NoRouteHandler())
	// Routes
	connection, err := repository.NewConnection()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	userRepository := repository.NewUserRepository(connection)
	userService := services.NewUserService(userRepository)
	userHandler := controllers.NewUserController(userService)

	app.GET("/user", userHandler.GET_HelloUser)
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return app
}
