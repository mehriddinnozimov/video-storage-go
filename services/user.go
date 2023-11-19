package services

import (
	"context"
	"video-storage/configs"
	"video-storage/types"

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

func (us *UserService) GetMany(c context.Context) ([]types.User, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := us.collection.Find(ctx, bson.D{}, opts)

	if err != nil {
		return nil, err
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

func (us *UserService) UpdateOneByID(c context.Context, id string, payload *types.User) (int64, error) {
	ctx, cancel := context.WithTimeout(c, configs.CONTEXT_TIMEOUT)
	defer cancel()

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := us.collection.UpdateOne(ctx, bson.M{"_id": idHex}, payload)
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
	return result.DeletedCount, err
}
