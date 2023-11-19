package controllers

import (
	"net/http"
	"video-storage/configs"
	"video-storage/services"
	"video-storage/types"
	"video-storage/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(db mongo.Database) *AuthController {
	return &AuthController{
		userService: *services.NewUserService(db),
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var payload types.UserRegisterRequest

	err := c.ShouldBind(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = ac.userService.GetOneByEmail(c, payload.Email)
	if err == nil {
		c.JSON(http.StatusConflict, types.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	user := types.User{
		FullName: payload.FullName,
		Password: payload.Password,
		Email:    payload.Email,
	}

	err = ac.userService.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Message: err.Error()})
	}

	userJSON := types.UserJSON{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}

	accress_token, err := utils.CreateToken(userJSON, configs.ENV.JwtSecret, configs.ACCESS_TOKEN_EXPIRE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Message: err.Error()})
	}

	refresh_token, err := utils.CreateToken(userJSON, configs.ENV.JwtSecret, configs.ACCESS_TOKEN_EXPIRE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Message: err.Error()})
	}

	c.SetCookie("refresh_token", refresh_token, configs.REFRESH_TOKEN_EXPIRE, "/", configs.ENV.Domain, configs.ENV.IsSecure == "true", true)
	c.SetCookie("access_token", accress_token, configs.ACCESS_TOKEN_EXPIRE, "/", configs.ENV.Domain, configs.ENV.IsSecure == "true", true)

	c.JSON(http.StatusOK, types.UserResponse{User: userJSON})
}
