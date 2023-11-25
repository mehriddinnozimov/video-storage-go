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
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	database   mongo.Database
	collection *mongo.Collection
}

func NewUserService(db mongo.Database) *UserService {
	return &UserService{
		database:   db,
		collection: db.Collection(configs.Collection.User),
	}
}

func (us *UserService) Create(c context.Context, payload *types.User) error {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	payload.ID = primitive.NewObjectID()

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	payload.Password = string(encryptedPassword)

	_, err = us.collection.InsertOne(ctx, payload)
	return err
}

func (us *UserService) GetMany(c context.Context, filter types.UserFilter) ([]types.User, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	query := bson.M{}

	if utils.Has(filter, "FullName") {
		query["full_name"] = bson.M{
			"$regex": filter.FullName,
		}
	}

	if utils.Has(filter, "Email") {
		query["email"] = bson.M{
			"$regex": filter.Email,
		}
	}
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := us.collection.Find(ctx, query, opts)

	if err != nil {
		return []types.User{}, err
	}

	var users []types.User

	err = cursor.All(c, &users)
	if users == nil {
		return []types.User{}, err
	}

	return users, err
}

func (us *UserService) GetOneByEmail(c context.Context, email string) (types.User, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	var user types.User
	err := us.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (us *UserService) GetOneByEmailAndPassword(c context.Context, email string, password string) (types.User, error) {
	user, err := us.GetOneByEmail(c, email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return user, err
}

func (us *UserService) GetOneByIDAndPassword(c context.Context, id string, password string) (types.User, error) {
	user, err := us.GetOneByID(c, id)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return user, err
}

func (us *UserService) GetOneByID(c context.Context, id string) (types.User, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	var user types.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = us.collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func (us *UserService) GetOneByIDWithVideos(c context.Context, options types.UserGetOneByIDWithVideosOptions) (types.User, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	var users []types.User

	idHex, err := primitive.ObjectIDFromHex(options.Id)
	if err != nil {
		return types.User{}, err
	}

	matchPipeline := bson.M{
		"$match": bson.M{
			"_id": idHex,
		},
	}

	videoLookupPipeline := bson.M{
		"$lookup": bson.M{
			"from":         configs.Collection.Video,
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "videos",
		},
	}

	if options.VideoIsPublic != nil {
		videoLookupPipeline["pipline"] = []bson.M{
			{
				"$match": bson.M{
					"is_public": options.VideoIsPublic,
				},
			},
		}
	}

	pipeline := []bson.M{
		matchPipeline,
		videoLookupPipeline,
	}

	cursor, err := us.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return types.User{}, err
	}

	err = cursor.All(c, &users)
	if users == nil {
		return types.User{}, mongo.ErrNoDocuments
	}

	return users[0], err
}

func (us *UserService) UpdateOneByID(c context.Context, id string, payload types.UserUpdate) (int64, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	if utils.Has(payload, "Password") {
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(payload.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return 0, err
		}

		payload.Password = string(encryptedPassword)
	}

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := us.collection.UpdateOne(ctx, bson.M{"_id": idHex}, bson.M{
		"$set": payload,
	})
	if err != nil {
		return 0, err
	}
	return result.MatchedCount, err
}

func (us *UserService) RemoveOneByID(c context.Context, id string) (int64, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := us.collection.DeleteOne(ctx, bson.M{"_id": idHex})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, err
}
