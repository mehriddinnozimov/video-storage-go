package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Video(db mongo.Database, group *gin.RouterGroup) {
	videoController := *controllers.NewVideoController(db)

	group.GET("/", videoController.GetMany)
	group.POST("/", videoController.Create)
}
