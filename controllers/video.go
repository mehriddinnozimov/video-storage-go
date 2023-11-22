package controllers

import (
	"net/http"
	"video-storage/services"
	"video-storage/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoController struct {
	videoService services.VideoService
}

func NewVideoController(db mongo.Database) *VideoController {
	return &VideoController{
		videoService: *services.NewVideoService(db),
	}
}

func (vc *VideoController) GetMany(c *gin.Context) {

}

func (vc *VideoController) GetManyByUserId(c *gin.Context) {

}

func (vc *VideoController) Create(c *gin.Context) {
	var payload types.Video

	err := c.ShouldBind(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	user_id := c.GetString("user_id")

	payload.UserId, err = primitive.ObjectIDFromHex(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	err = vc.videoService.Create(c, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.VideoResponse{Video: payload})
}
