package main

import (
	"video-storage/configs"
	"video-storage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(configs.ENV.GinMode)

	app := gin.Default()
	client := configs.NewClient(configs.ENV)
	defer configs.Disconnect(client)

	database := configs.NewDatabase(client, configs.ENV.DB)

	routes.Root(app, database)

	app.Run(configs.ENV.ServerAddress)
}
