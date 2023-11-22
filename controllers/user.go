package controllers

import (
	"net/http"
	"video-storage/services"
	"video-storage/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(db mongo.Database) *UserController {
	return &UserController{
		userService: *services.NewUserService(db),
	}
}

func (uc *UserController) GetMany(c *gin.Context) {
	var filter types.UserFilter

	err := c.ShouldBindQuery(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	users, err := uc.userService.GetMany(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.UsersResponse{Users: users})
}

func (uc *UserController) GetOne(c *gin.Context) {
	var filter types.UserParams

	err := c.ShouldBindUri(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	user, err := uc.userService.GetOneByIDWithVideos(c, filter.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.UserResponse{User: user})
}
