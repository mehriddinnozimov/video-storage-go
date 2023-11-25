package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v3"
)

func Video(db mongo.Database, storage *drive.Service, group *gin.RouterGroup) {
	videoController := *controllers.NewVideoController(db, storage)

	group.GET("/", videoController.GetMany)
	group.POST("/", videoController.Create)
}
