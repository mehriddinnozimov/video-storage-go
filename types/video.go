package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name" binding:"required"`
	Time     string             `bson:"time" json:"time"`
	FileId   string             `bson:"file_id" json:"file_id" binding:"required"`
	UserId   primitive.ObjectID `bson:"user_id" json:"user_id"`
	IsPublic bool               `bson:"is_public" json:"is_public" binding:"required"`
}

type VideoFilter struct {
	UserId   string `json:"user_id,omitempty"`
	IsPublic bool   `bson:"is_public,omitempty"`
}

type VideoResponse struct {
	Video Video `json:"video"`
}

type VideosResponse struct {
	Videos []Video `json:"videos"`
}
