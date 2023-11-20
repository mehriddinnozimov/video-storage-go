package routes

import (
	"video-storage/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Root(gin *gin.Engine, db mongo.Database) {
	root := gin.Group("")

	publicRoute := root.Group("")
	protectedRoute := root.Group("")

	protectedRoute.Use(middlewares.Authorization)

	Auth(db, publicRoute.Group("/auth"))
	Me(db, protectedRoute.Group("/me"))
	User(db, protectedRoute.Group("/user"))
}
