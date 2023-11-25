package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v3"
)

func User(db mongo.Database, storage *drive.Service, group *gin.RouterGroup) {
	userController := *controllers.NewUserController(db)
	videoController := *controllers.NewVideoController(db, storage)

	group.GET("/", userController.GetMany)
	group.GET("/:user_id", userController.GetOne)
	group.GET("/:user_id/video", videoController.GetManyByUserId)
}
