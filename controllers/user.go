package controllers

import (
	"video-storage/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type userController struct {
	userService services.UserService
}

func UserController(db mongo.Database) *userController {
	return &userController{
		userService: *services.NewUserService(db),
	}
}
