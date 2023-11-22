package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func User(db mongo.Database, group *gin.RouterGroup) {
	userController := *controllers.NewUserController(db)
	videoController := *controllers.NewVideoController(db)

	group.GET("/", userController.GetMany)
	group.GET("/:user_id", userController.GetOne)
	group.GET("/:user_id/video", videoController.GetManyByUserId)
}
