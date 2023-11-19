package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(env *Envs) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s", env.DBHost, env.DBPort)
	option := options.Client().ApplyURI(uri)

	if env.DBUser != "" && env.DBPassword != "" {
		credential := options.Credential{
			Username: env.DBUser,
			Password: env.DBPassword,
		}
		option.SetAuth(credential)
	}

	client, err := mongo.Connect(ctx, option)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return *client
}

func NewDatabase(client mongo.Client, name string) mongo.Database {
	database := client.Database(name)

	return *database
}

func Disconnect(client mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
