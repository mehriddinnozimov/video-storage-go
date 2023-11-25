package controllers

import (
	"net/http"
	"video-storage/services"
	"video-storage/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v3"
)

type VideoController struct {
	videoService   services.VideoService
	storageService services.StorageService
}

func NewVideoController(db mongo.Database, storage *drive.Service) *VideoController {
	return &VideoController{
		videoService:   *services.NewVideoService(db),
		storageService: *services.NewStorageService(storage),
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

func (vc *VideoController) Update(c *gin.Context) {
	var filter types.VideoParams

	err := c.ShouldBindUri(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	var payload types.VideoUpdate

	err = c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	_, err = vc.videoService.UpdateOneByID(c, filter.VideoId, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.CustomResponse{Ok: true})
}

func (vc *VideoController) Remove(c *gin.Context) {
	var filter types.VideoParams
	user_id := c.GetString("user_id")

	err := c.ShouldBindUri(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	video, err := vc.videoService.GetOneByID(c, filter.VideoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	err = vc.storageService.Remove(video.FileId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	_, err = vc.videoService.RemoveOneByIDAndUserId(c, filter.VideoId, user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.CustomResponse{Ok: true})
}
