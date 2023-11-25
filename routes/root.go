package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v3"
)

func Root(gin *gin.Engine, db mongo.Database, storage *drive.Service) {
	root := gin.Group("/")

	publicRoute := root.Group("/")
	protectedRoute := root.Group("/")

	// protectedRoute.Use(middlewares.Authorization)

	Auth(db, publicRoute.Group("/auth"))
	Me(db, protectedRoute.Group("/me"))
	User(db, storage, protectedRoute.Group("/user"))
	Video(db, storage, protectedRoute.Group("/video"))
	File(db, storage, protectedRoute.Group("/file"))
}
