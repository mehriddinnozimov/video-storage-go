package main

import (
	"log"
	"os"
	"path"
	"video-storage/configs"
	"video-storage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(configs.ENV.GinMode)

	app := gin.Default()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	app.Static("/public/", path.Join(wd, "public"))

	client := configs.NewClient(configs.ENV)
	defer configs.Disconnect(client)

	database := configs.NewDatabase(client, configs.ENV.DB)
	storage := configs.NewDrive()

	routes.Root(app, database, storage)

	app.Run(configs.ENV.ServerAddress)
}
