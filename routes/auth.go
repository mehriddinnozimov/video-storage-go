package routes

import (
	"video-storage/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Auth(db mongo.Database, group *gin.RouterGroup) {
	authController := *controllers.NewAuthController(db)

	group.POST("/register", authController.Register)
	group.POST("/login", authController.Login)
}
