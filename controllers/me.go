package controllers

import (
	"fmt"
	"net/http"
	"video-storage/services"
	"video-storage/types"
	"video-storage/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MeController struct {
	userService services.UserService
}

func NewMeController(db mongo.Database) *MeController {
	return &MeController{
		userService: *services.NewUserService(db),
	}
}

func (mc *MeController) Get(c *gin.Context) {
	user_id := c.GetString("user_id")

	user, err := mc.userService.GetOneByID(c, user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.UserResponse{User: user})
}

func (mc *MeController) Update(c *gin.Context) {
	user_id := c.GetString("user_id")
	var requestPayload types.UserUpdate

	err := c.ShouldBind(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	payload := types.UserUpdate{
		FullName: requestPayload.FullName,
		Email:    requestPayload.Email,
	}

	if utils.Has(payload, "FullName") || utils.Has(payload, "Email") {
		fmt.Println(payload)
		_, err = mc.userService.UpdateOneByID(c, user_id, payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, types.CustomResponse{Ok: true})
}

func (mc *MeController) UpdatePassword(c *gin.Context) {
	user_id := c.GetString("user_id")
	var requestPayload types.UserUpdatePassword

	err := c.ShouldBind(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	_, err = mc.userService.GetOneByIDAndPassword(c, user_id, requestPayload.OldPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	payload := types.UserUpdate{
		Password: requestPayload.Password,
	}

	_, err = mc.userService.UpdateOneByID(c, user_id, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.CustomResponse{Ok: true})
}

func (mc *MeController) Remove(c *gin.Context) {
	user_id := c.GetString("user_id")
	var requestPayload types.UserUpdatePassword

	err := c.ShouldBind(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	_, err = mc.userService.GetOneByIDAndPassword(c, user_id, requestPayload.OldPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	payload := types.UserUpdate{
		Password: requestPayload.Password,
	}

	_, err = mc.userService.UpdateOneByID(c, user_id, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.CustomResponse{Ok: true})
}
