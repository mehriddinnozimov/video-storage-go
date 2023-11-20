package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func User(db mongo.Database, group *gin.RouterGroup) {
	ucController := *controllers.NewUserController(db)

	group.GET("/", ucController.GetMany)
}
