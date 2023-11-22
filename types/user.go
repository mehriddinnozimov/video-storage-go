package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName string             `bson:"full_name,omitempty" json:"full_name,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"-"`

	Videos []Video `json:"videos"`
}

type UserUpdate struct {
	FullName string `bson:"full_name,omitempty" json:"full_name,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"-"`
}

type UserRegister struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserFilter struct {
	FullName string `form:"full_name,omitempty"`
	Email    string `form:"email,omitempty"`
}

type UserParams struct {
	UserId string `uri:"user_id"`
}

type UserUpdatePassword struct {
	Password    string `json:"password" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserResponse struct {
	User User `json:"user"`
}

type UsersResponse struct {
	Users []User `json:"users"`
}
