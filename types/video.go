package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty" binding:"required"`
	Time     string             `bson:"time,omitempty" json:"time,omitempty"`
	FileId   string             `bson:"file_id,omitempty" json:"file_id,omitempty" binding:"required"`
	UserId   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	IsPublic bool               `bson:"is_public,omitempty" json:"is_public,omitempty" binding:"required"`
}

type VideoUpdate struct {
	Name     string `bson:"name,omitempty" json:"name,omitempty" binding:"required"`
	Time     string `bson:"time,omitempty" json:"time,omitempty"`
	FileId   string `bson:"file_id,omitempty" json:"file_id,omitempty" binding:"required"`
	IsPublic bool   `bson:"is_public,omitempty" json:"is_public,omitempty" binding:"required"`
}

type VideoFilter struct {
	UserId   string `json:"user_id,omitempty"`
	IsPublic bool   `bson:"is_public,omitempty"`
}

type VideoParams struct {
	VideoId string `uri:"video_id"`
}

type VideoResponse struct {
	Video Video `json:"video"`
}

type VideosResponse struct {
	Videos []Video `json:"videos"`
}
