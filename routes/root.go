package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Root(gin *gin.Engine, db mongo.Database) {
	root := gin.Group("")

	Auth(db, root.Group("/auth"))
}
