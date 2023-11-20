package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Me(db mongo.Database, group *gin.RouterGroup) {
	meController := *controllers.NewMeController(db)

	group.GET("/", meController.Get)
	group.PUT("/", meController.Update)
	group.PUT("/password", meController.UpdatePassword)
}
