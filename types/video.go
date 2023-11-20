package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Time   string             `bson:"time"`
	FileId string             `bson:"file_id"`
	UserId primitive.ObjectID `bson:"user_id"`
}
