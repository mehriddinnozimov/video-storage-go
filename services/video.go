package services

import (
	"context"
	"video-storage/configs"
	"video-storage/types"
	"video-storage/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoService struct {
	database   mongo.Database
	collection *mongo.Collection
}

func NewVideoService(db mongo.Database) *VideoService {
	return &VideoService{
		database:   db,
		collection: db.Collection(configs.Collection.Video),
	}
}

func (vs *VideoService) Create(c context.Context, payload *types.Video) error {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	payload.ID = primitive.NewObjectID()

	_, err := vs.collection.InsertOne(ctx, payload)
	return err
}

func (vs *VideoService) GetMany(c context.Context, filter types.VideoFilter) ([]types.Video, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	var videos []types.Video

	query := bson.M{}

	if utils.Has(filter, "UserId") {
		idHex, err := primitive.ObjectIDFromHex(filter.UserId)
		if err != nil {
			return []types.Video{}, err
		}
		query["user_id"] = idHex
	}

	if utils.Has(filter, "IsPublic") {
		query["is_public"] = filter.IsPublic
	}

	opts := options.Find()
	cursor, err := vs.collection.Find(ctx, query, opts)

	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &videos)
	if videos == nil {
		return []types.Video{}, err
	}

	return videos, err
}

func (vs *VideoService) GetOneByFileID(c context.Context, file_id string) (types.Video, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	var video types.Video
	err := vs.collection.FindOne(ctx, bson.M{"file_id": file_id}).Decode(&video)
	return video, err
}

func (vs *VideoService) GetOneByID(c context.Context, id string) (types.Video, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	var video types.Video

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return video, err
	}

	err = vs.collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&video)
	return video, err
}

func (vs *VideoService) UpdateOneByID(c context.Context, id string, payload *types.Video) (int64, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := vs.collection.UpdateOne(ctx, bson.M{"_id": idHex}, payload)
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, err
}

func (vs *VideoService) RemoveOneByID(c context.Context, id string) (int64, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := vs.collection.DeleteOne(ctx, bson.M{"_id": idHex})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, err
}

func (vs *VideoService) RemoveManyByUserID(c context.Context, user_id string) (int64, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	userIdHex, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return 0, err
	}

	result, err := vs.collection.DeleteMany(ctx, bson.M{"user_id": userIdHex})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, err
}
