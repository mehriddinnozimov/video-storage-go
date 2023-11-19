package services

import (
	"context"
	"video-storage/configs"
	"video-storage/types"

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

func (vs *VideoService) GetMany(c context.Context) ([]types.Video, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := vs.collection.Find(ctx, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var videos []types.Video

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
	return result.DeletedCount, err
}
