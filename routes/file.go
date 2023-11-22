package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v3"
)

func File(db mongo.Database, drive *drive.Service, group *gin.RouterGroup) {
	fileController := controllers.NewFileController(db, drive)

	group.POST("/", fileController.Create)
	group.GET("/:file_id", fileController.GetOne)
}
