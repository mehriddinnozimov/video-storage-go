package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	FullName string             `bson:"full_name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserRegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserJSON struct {
	ID       primitive.ObjectID `json:"_id"`
	FullName string             `json:"full_name"`
	Email    string             `json:"email"`
	Password string             `json:"password,omitempty"`
}

type UserResponse struct {
	User UserJSON `json:"user"`
}
