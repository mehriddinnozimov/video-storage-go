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
	var payload types.UserRegister

	err := c.ShouldBind(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	_, err = ac.userService.GetOneByEmail(c, payload.Email)
	if err == nil {
		c.JSON(http.StatusConflict, types.CustomResponse{Ok: false, Message: "User already exists with the given email"})
		return
	}

	user := types.User{
		FullName: payload.FullName,
		Password: payload.Password,
		Email:    payload.Email,
	}

	err = ac.userService.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	accress_token, err := utils.CreateToken(user.ID.Hex(), configs.ENV.JwtSecret, configs.ACCESS_TOKEN_EXPIRE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(user.ID.Hex(), configs.ENV.JwtSecret, configs.ACCESS_TOKEN_EXPIRE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.SetCookie("refresh_token", refresh_token, int(configs.REFRESH_TOKEN_EXPIRE_BY_SECOND), "/", configs.ENV.Domain, configs.ENV.IsSecure, true)
	c.SetCookie("access_token", accress_token, int(configs.ACCESS_TOKEN_EXPIRE_BY_SECOND), "/", configs.ENV.Domain, configs.ENV.IsSecure, true)

	c.JSON(http.StatusOK, types.UserResponse{User: user})
}

func (ac *AuthController) Login(c *gin.Context) {
	var payload types.UserLogin

	err := c.ShouldBind(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	user, err := ac.userService.GetOneByEmailAndPassword(c, payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	accress_token, err := utils.CreateToken(user.ID.Hex(), configs.ENV.JwtSecret, configs.ACCESS_TOKEN_EXPIRE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(user.ID.Hex(), configs.ENV.JwtSecret, configs.REFRESH_TOKEN_EXPIRE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.SetCookie("refresh_token", refresh_token, int(configs.REFRESH_TOKEN_EXPIRE_BY_SECOND), "/", configs.ENV.Domain, configs.ENV.IsSecure, true)
	c.SetCookie("access_token", accress_token, int(configs.ACCESS_TOKEN_EXPIRE_BY_SECOND), "/", configs.ENV.Domain, configs.ENV.IsSecure, true)

	c.JSON(http.StatusOK, types.UserResponse{User: user})
}
