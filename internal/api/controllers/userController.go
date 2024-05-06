package controllers

import (
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{UserService: service}
}
func (uh *UserController) GET_HelloUser(ctx *gin.Context) {
	uh.UserService.GetUserById(0)
	ctx.JSON(200, "Hello User!")
}
